package usecase

import "github.com/muhammadkhon-abdulloev/url-shortener/internal/service"

type serviceUC struct {

}

func NewServiceUC() service.UseCase {
	return &serviceUC{}
}