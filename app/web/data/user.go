package data

import (
	"chat_with_me/common/model/entity"
	"chat_with_me/common/utils"
	"context"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) Create(ctx context.Context, user *entity.User) error {
	return u.db.WithContext(ctx).Create(user).Error
}
func (u *UserRepo) Update(ctx context.Context, id int64, user map[string]any) error {
	return u.db.Model(&entity.User{}).WithContext(ctx).Where("id = ?", id).Updates(user).Error
}
func (u *UserRepo) UpdateAttr(ctx context.Context, id int64, dbColumn string, objectPath []string, val any) error {
	return u.db.Model(&entity.User{}).WithContext(ctx).Where("id = ?", id).Transaction(func(db *gorm.DB) error {
		if err := utils.BuildPostgresJsonbMissObject(db, dbColumn, objectPath); err != nil {
			return err
		}
		path := utils.JoinPostgresJsonbPath(objectPath)
		return db.Update(dbColumn, datatypes.JSONSet(dbColumn).Set(path, val)).Error
	})
}
