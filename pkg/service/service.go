package service

import (
	"NameService/pkg/model"
	"NameService/pkg/repository"
	"context"
)

type ImageImplementation interface {
	InsertData(ctx context.Context, fio *model.NameModel) (int64, error)
	UpdateInfo(ctx context.Context, fio *model.NameModel) error
	GetPeopleByAge(ctx context.Context) ([]*model.NameModel, error)
	DeleteImageByID(ctx context.Context, id int) (int64, error)
	//GetAllImages(ctx context.Context) ([]*model.NameModel, error)
}

type Service struct {
	repository repository.Implementation
}

type Implementation interface {
	ImageImplementation
}

func NewService(repository repository.Implementation) Implementation {
	return &Service{repository: repository}
}
