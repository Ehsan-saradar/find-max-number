package usecase


type MathUseCase interface {
	IsMax(number int32)(bool)
	Reset()
}
func NewSimpleMathUseCase() MathUseCase {
	return &simpleMathUseCase{}
}
type simpleMathUseCase struct {
	maxNumber int32
	isvalidMax bool
}

//Check if number is the maximum till now
func (math *simpleMathUseCase) IsMax(number int32)(bool){
	if !math.isvalidMax || number>math.maxNumber{
		math.maxNumber=number
		math.isvalidMax=true
		return true
	}
	return false
}

//Clear maximum number
func (math *simpleMathUseCase)Reset(){
	math.isvalidMax=false
}