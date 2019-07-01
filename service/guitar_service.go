package service

import (
	"log"
	"personal/guitar_collection/domain"
	"personal/guitar_collection/repository"
)

type GuitarService interface {
	Process(viewGuitar domain.ViewGuitar) error
}

type guitarService struct {
	guitarRepository repository.GuitarRepository
}

func (self guitarService) Process(viewGuitar domain.ViewGuitar) error {
	guitar := &domain.Guitar{
		Brand: viewGuitar.Brand,
		Type:  viewGuitar.Type,
		Price: viewGuitar.Price,
	}
	err := self.guitarRepository.CreateGuitar(guitar)
	log.Println(err)
	return err
}

func NewGuitarService() GuitarService {
	return guitarService{
		guitarRepository: repository.NewGuitarRepository(),
	}
}
