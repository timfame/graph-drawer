package service

type Service interface {
	Draw(input, output string) error
}
