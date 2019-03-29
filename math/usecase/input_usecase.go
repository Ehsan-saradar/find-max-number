package usecase

import (
	"github.com/Ehsan-saradar/find-max-number/math/repository"
)

type InputUseCase interface {
	GetNextNumber()(int,error)
}

//Delivers numbers directly from repository layer to delivery layer
func NewSimpleInputUseCase(mathRepository repository.MathRepository) InputUseCase {
	return &simpleInputUseCase{mathRepository:mathRepository}
}
type simpleInputUseCase struct {
	maxNumber int32
	isvalidMax bool
	mathRepository repository.MathRepository
}

func (inputUseCase *simpleInputUseCase) GetNextNumber()(int,error){
	return inputUseCase.mathRepository.GetNextNumber()
}