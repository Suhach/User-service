package user

import (
	"context"
	"gorm.io/gorm"
)

type Repository interface {
	GetAll(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id int) (*User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, id int, user *User) error
	Delete(ctx context.Context, id int) error
}

type UserREPO struct {
	db *gorm.DB
}

func NewUserREPO(db *gorm.DB) *UserREPO {
	return &UserREPO{
		db: db,
	}
}

func (r *UserREPO) GetAll(ctx context.Context) ([]User, error) {
	var users []User
	err := r.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserREPO) GetByID(ctx context.Context, id int) (user *User, err error) {
	var usr User
	err = r.db.WithContext(ctx).First(&usr, id).Error
	if err != nil {
		return nil, err
	}
	return &usr, nil
}

// TODO implement Create/Update/Delete
func (r *UserREPO) Create(ctx context.Context, user *User) error {

}

func (r *UserREPO) Update(ctx context.Context, id int, user *User) error {}

func (r *UserREPO) Delete(ctx context.Context, id int) error {}
