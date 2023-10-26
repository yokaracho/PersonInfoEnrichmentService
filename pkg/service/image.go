package service

import (
	"NameService/pkg/model"
	"context"
)

func (s *Service) InsertData(ctx context.Context, fio *model.NameModel) (int64, error) {
	return s.repository.InsertData(ctx, fio)
}

func (s *Service) UpdateInfo(ctx context.Context, fio *model.NameModel) error {
	return s.repository.UpdateInfo(ctx, fio)
}

func (s *Service) GetPeopleByAge(ctx context.Context) ([]*model.NameModel, error) {
	return s.repository.GetPeopleByAge(ctx)
}

func (s *Service) DeleteImageByID(ctx context.Context, id int) (int64, error) {
	return s.repository.DeleteImageByID(ctx, id)
}

//
//func (s *Service) GetAllImages(ctx context.Context) ([]*model.NameModel, error) {
//	return s.repository.GetAllImages(ctx)
//}
