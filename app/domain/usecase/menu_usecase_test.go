package usecase_test

import (
	"log"
	"testing"

	"github.com/YutoMizutani/gohome/app/domain/usecase"
)

func TestMenuUseCaseFetch(t *testing.T) {
	MenuUseCase := usecase.MenuUseCase{}

	models, err := MenuUseCase.Get()
	if err != nil {
		log.Fatal("Error entity returns nil")
	}

	firstModel := (*models)[0]
	if firstModel.Title != "Lives" {
		log.Fatal("Error Title value")
	}
}
