package service

type Services struct {
	Guitar GuitarService
}

func NewServices() *Services {
	return &Services{
		Guitar: NewGuitarService(),
	}
}
