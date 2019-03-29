package repository

import (
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)
const (
	InputFileName = "input.txt"
	RandomNumbers=20
)
type MathRepository interface {
	// Extract numbers one by one
	GetNextNumber()(int,error)
}


func NewMathConsoleRepository() MathRepository {
	return &mathConsoleRepository{}
}
//Get numbers from console input
type mathConsoleRepository struct {
}
func (*mathConsoleRepository) GetNextNumber()(int,error){
	var number int
	_,err:=fmt.Scanln(&number)
	return number,err
}

// Open input file and read numbers
func NewMathFileRepository(homeDir string) (MathRepository,error) {
	handle, err := os.Open(path.Join(homeDir,InputFileName))
	if err != nil {
		return &mathFileRepository{},err
	}
	defer handle.Close()
	return ReadFileRepository(handle)


}
// Read numbers from file and extract numbers
func ReadFileRepository(reader io.Reader)(MathRepository,error){
	var mathRepository mathFileRepository
	fileBytes,err:=ioutil.ReadAll(reader)
	if err!=nil{
		return &mathRepository,err
	}
	mathRepository.numbers=make([]int, len(strings.Split(string(fileBytes), "\r\n")))
	for i,substr:=range strings.Split(string(fileBytes),"\r\n"){
		mathRepository.numbers[i],err=strconv.Atoi(substr)
		if err!=nil{
			return &mathRepository,err
		}
	}
	return &mathRepository,nil
}
//Generate file and save random numbers
func GenerateFileRepository(homeDir string) (error) {
	var err error
	if _, err := os.Stat(homeDir); os.IsNotExist(err) {
		err=os.MkdirAll(homeDir,os.ModePerm)
	}
	if err!=nil{
		return err
	}
	handle, err := os.OpenFile(path.Join(homeDir,InputFileName),os.O_CREATE,0600)
	if err != nil {
		return err
	}
	defer handle.Close()
	return generateRandom(handle)
}

//Generate numbers and writes each one in a line in writer
func generateRandom(writer io.Writer)error{
	var data string
	rnd:=rand.New(rand.NewSource(time.Now().Unix()))
	for i:=0;i<RandomNumbers;i++{
		if len(data)>0{
			data+="\r\n"
		}
		data+=strconv.Itoa(rnd.Intn(10000))
	}
	_,err:=writer.Write([]byte(data))
	return err
}

//Get number from file (each line has one number)
type mathFileRepository struct {
	index int
	numbers []int
}
func (mathFileRepository *mathFileRepository)GetNextNumber()(int,error) {
	if mathFileRepository.index<len(mathFileRepository.numbers) {
		mathFileRepository.index = mathFileRepository.index + 1
		return mathFileRepository.numbers[mathFileRepository.index-1], nil
	}
	return -1,io.EOF
}
