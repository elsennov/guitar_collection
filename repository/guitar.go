package repository

import (
	"github.com/jmoiron/sqlx"
	"log"
	"personal/guitar_collection/appcontext"
	guitarDomain "personal/guitar_collection/domain"
)

const (
	createGuitarQuery = `INSERT INTO guitar (brand, type, price) VALUES (?, ?, ?)`
)

type GuitarRepository interface {
	CreateGuitar(*guitarDomain.Guitar) error
}

type guitarRepository struct {
	db *sqlx.DB
}

func (self guitarRepository) CreateGuitar(guitar *guitarDomain.Guitar) error {
	_, err := self.db.Exec(createGuitarQuery, guitar.Brand, guitar.Type, guitar.Price)
	log.Println(err)
	return err
}

func NewGuitarRepository() GuitarRepository {
	return guitarRepository{
		db: appcontext.GetDB(),
	}
}
