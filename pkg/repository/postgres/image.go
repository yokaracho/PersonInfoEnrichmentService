package db

import (
	"NameService/pkg/model"
	"context"
)

func (r *Repository) InsertData(ctx context.Context, fio *model.NameModel) (int64, error) {
	result, err := r.pool.Exec(ctx, insertImageQuery, fio.Name, fio.Surname, fio.Patronymic, fio.Age, fio.Gender, fio.Nationality)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), err
}

//func (r *Repository) GetPeopleByAge(ctx context.Context) ([]model.NameModel, error) {
//	var fios []*model.NameModel
//	err := r.pool.QueryRow(ctx, getPeopleSort).Scan(&fio.ID, fio.Name, fio.Surname, fio.Patronymic, fio.Age, fio.Gender, fio.Nationality)
//	if err != nil {
//		return nil, err
//	}
//	return &fio, nil
//}

func (r *Repository) GetPeopleByAge(ctx context.Context) ([]*model.NameModel, error) {
	var fios []*model.NameModel
	rows, err := r.pool.Query(ctx, getPeopleSort)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var fio model.NameModel

		if err := rows.Scan(&fio.ID, &fio.Name, &fio.Surname, &fio.Patronymic, &fio.Age, &fio.Gender, &fio.Nationality); err != nil {
			return nil, err
		}
		fios = append(fios, &fio)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return fios, nil
}

func (r *Repository) UpdateInfo(ctx context.Context, fio *model.NameModel) error {
	_, err := r.pool.Exec(ctx, updateInfoPeople, fio.ID, fio.Name, fio.Surname, fio.Patronymic, fio.Age, fio.Gender, fio.Nationality)
	if err != nil {
		return err
	}
	return err
}

func (r *Repository) DeleteImageByID(ctx context.Context, id int) (int64, error) {
	result, err := r.pool.Exec(ctx, deleteImageByDateQuery, id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), err
}
