package data

import (
	"chat_with_me/common/model/entity"
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
func (u *UserRepo) UpdateAttr(ctx context.Context, id int64, key string, val any) error {
	/*var result string
	if err := u.db.Model(&entity.User{}).WithContext(ctx).Select("attr->'profile'->>'name' as key").Where("id = ?", id).Scan(&result).Error; err != nil {
		return err
	}
	fmt.Println(result)
	return nil*/
	return u.db.Model(&entity.User{}).WithContext(ctx).Where("id = ?", id).Update("attr", datatypes.JSONSet("attr").Set("{"+key+"}", val)).Error
}
