package repository

import (
	"NameService/pkg/model"
	"context"
)

type ImageImplementation interface {
	InsertData(ctx context.Context, fio *model.NameModel) (int64, error)
	UpdateInfo(ctx context.Context, fio *model.NameModel) error
	GetPeopleByAge(ctx context.Context) ([]*model.NameModel, error)
	DeleteImageByID(ctx context.Context, id int) (int64, error)
	//GetAllImages(ctx context.Context) ([]*model.NameModel, error)
}

type Implementation interface {
	ImageImplementation
}
