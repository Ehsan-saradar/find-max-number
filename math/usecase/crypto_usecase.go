package usecase

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/gob"
	"encoding/pem"
	"github.com/Ehsan-saradar/find-max-number/math"
	"io/ioutil"
	"os"
	RSA "crypto/rsa"
	ed "golang.org/x/crypto/ed25519"
	"path"
)
const (
	publicKeyFileName = "public.key"
	privateKeyFileName = "private.key"
	RSAKeyType="RSA"
	ED25519KeyType="ed25519"
	RsaKeySize=1024

)
type CryptoUseCase interface {
	GenerateKey()(Key,error)
	Sign(message string) ([]byte,error)
	Verify(message string,sign []byte) (bool)
}

type Key struct {
	KeyType string
	PublicKey []byte
	PrivateKey []byte
}
//Save public/private keys in separate files
func (key *Key)Save(homeDir string)error{
	var err error
	if _, err := os.Stat(homeDir); os.IsNotExist(err) {
		err=os.MkdirAll(homeDir,os.ModePerm)
	}
	if err!=nil{
		return err
	}
	var buffer bytes.Buffer
	encoder:=gob.NewEncoder(&buffer)
	err=encoder.Encode(Key{
		KeyType:key.KeyType,
		PublicKey:key.PublicKey,
	})
	if err!=nil{
		return err
	}
	err = ioutil.WriteFile(path.Join(homeDir,publicKeyFileName), buffer.Bytes(), 0644)
	if err!=nil{
		return err
	}
	buffer.Reset()
	encoder=gob.NewEncoder(&buffer)
	err=encoder.Encode(Key{
		KeyType:key.KeyType,
		PrivateKey:key.PrivateKey,
	})
	err = ioutil.WriteFile(path.Join(homeDir,privateKeyFileName), buffer.Bytes(), 0644)
	return err
}

//Read keys from files in home directory
func (key *Key)Load(homeDir string,includePrivateKey bool)(CryptoUseCase,error){
	var cryptoUseCase CryptoUseCase
	var err error
	err=key.LoadPublicKey(homeDir)
	if err!=nil{
		return cryptoUseCase,err
	}
	if includePrivateKey {
		err = key.LoadPrivateKey(homeDir)
		if err != nil {
			return cryptoUseCase, err
		}
	}
	switch key.KeyType {
	case RSAKeyType:
		if includePrivateKey {
			block, _ := pem.Decode(key.PrivateKey)
			privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
			if err != nil {
				return cryptoUseCase, math.InvalidPrivateKey
			}
			cryptoUseCase = &Rsa{
				PrivateKey: privateKey,
			}
		} else {
			block, _ := pem.Decode(key.PublicKey)
			publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
			if err != nil {
				return cryptoUseCase, math.InvalidPrivateKey
			}
			cryptoUseCase = &Rsa{
				PublicKey: publicKey,
			}
		}
		break
	case ED25519KeyType:
		if includePrivateKey {
			cryptoUseCase = &Ed25519{
				PublicKey:  key.PublicKey,
				PrivateKey: key.PrivateKey,
			}
		} else {
			cryptoUseCase = &Ed25519{
				PublicKey: key.PublicKey,
			}
		}
		break
	default:
		return cryptoUseCase, math.InvalidPublicKey
	}
	return cryptoUseCase,nil
}
//Decode public key fie
func (key *Key)LoadPublicKey(homeDir string)(error){
	var err error
	var tempKey Key
	var fileBytes []byte
	fileExits,err:=exists(path.Join(homeDir,publicKeyFileName))
	if !fileExits{
		return math.PublicKeyNotFound
	}
	if err!=nil{
		return math.InvalidPublicKey
	}
	fileBytes,err=ioutil.ReadFile(path.Join(homeDir,publicKeyFileName))
	if err!=nil{
		return math.InvalidPublicKey
	}
	var buffer bytes.Buffer
	buffer.Write(fileBytes)
	err=gob.NewDecoder(&buffer).Decode(&tempKey)
	if err!=nil{
		return math.InvalidPublicKey
	}
	if len(key.KeyType)>0 && key.KeyType!=tempKey.KeyType{
		return math.InvalidPublicKey
	}
	key.KeyType=tempKey.KeyType
	key.PublicKey=tempKey.PublicKey
	return nil
}
//Decode private key fie
func (key *Key)LoadPrivateKey(homeDir string)(error){
	var err error
	var tempKey Key
	var fileBytes []byte
	fileExits,err:=exists(path.Join(homeDir,privateKeyFileName))
	if !fileExits{
		return math.PrivateKeyNotFound
	}
	if err!=nil{
		return math.InvalidPrivateKey
	}
	fileBytes,err=ioutil.ReadFile(path.Join(homeDir,privateKeyFileName))
	if err!=nil{
		return math.InvalidPrivateKey
	}
	var buffer bytes.Buffer
	buffer.Write(fileBytes)
	err=gob.NewDecoder(&buffer).Decode(&tempKey)
	if err!=nil{
		return math.InvalidPrivateKey
	}
	if len(key.KeyType)>0 && key.KeyType!=tempKey.KeyType{
		return math.InvalidPrivateKey
	}
	key.KeyType=tempKey.KeyType
	key.PrivateKey=tempKey.PrivateKey
	return nil
}

//Check if file exits or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err
}


type Rsa struct {
	PrivateKey *RSA.PrivateKey
	PublicKey  *RSA.PublicKey
}
func (rsa *Rsa)GenerateKey() (Key,error) {
	var err error
	rsa.PrivateKey,err= RSA.GenerateKey(rand.Reader,RsaKeySize)
	rsa.PublicKey=&rsa.PrivateKey.PublicKey
	if err!=nil{
		return Key{},err
	}
	return Key{
		KeyType:RSAKeyType,
		PrivateKey:pem.EncodeToMemory(
			&pem.Block{
				Type: "RSA PRIVATE KEY",
				Bytes: x509.MarshalPKCS1PrivateKey(rsa.PrivateKey),
			},
		),
		PublicKey:pem.EncodeToMemory(
			&pem.Block{
				Type: "RSA PUBLIC KEY",
				Bytes: x509.MarshalPKCS1PublicKey(&rsa.PrivateKey.PublicKey),
			},
		),
	},nil
}
func (rsa *Rsa) Sign(message string)([]byte,error) {
	h := sha256.New()
	h.Write([]byte(message))
	hashedMessage := h.Sum(nil)
	sign,err:=RSA.SignPKCS1v15(rand.Reader, rsa.PrivateKey, crypto.SHA256, hashedMessage)
	return sign,err
}
func (rsa *Rsa) Verify(message string,sign []byte)(bool) {
	h := sha256.New()
	h.Write([]byte(message))
	hashedMessage := h.Sum(nil)
	return RSA.VerifyPKCS1v15(rsa.PublicKey,crypto.SHA256,hashedMessage,sign)==nil
}


type Ed25519 struct {
	PublicKey ed.PublicKey
	PrivateKey ed.PrivateKey
}

func (ed25519 *Ed25519) GenerateKey()(Key,error)  {
	var err error
	ed25519.PublicKey,ed25519.PrivateKey,err=ed.GenerateKey(rand.Reader)
	if err!=nil{
		return Key{},err
	}
	return Key{
		KeyType:ED25519KeyType,
		PublicKey:ed25519.PublicKey,
		PrivateKey:ed25519.PrivateKey,
	},nil
}
func (ed25519 *Ed25519) Sign(message string)([]byte,error) {
	return ed.Sign(ed25519.PrivateKey,[]byte(message)),nil
}
func (ed25519 *Ed25519) Verify(message string,sign []byte)(bool) {
	return ed.Verify(ed25519.PublicKey,[]byte(message),sign)
}



