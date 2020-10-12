package store

import (
	"context"

	"github.com/binkkatal/proto_user/src/model"
	"github.com/binkkatal/proto_user/src/service"
	"gorm.io/gorm"
)

type UserStore struct {
	service.UnimplementedUserServiceServer
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{db: db}
}

func (st *UserStore) Add(ctx context.Context, in *model.User) (*service.UserResponse, error) {
	err := st.db.Create(in).Error
	if err != nil {
		return nil, err
	}
	return &service.UserResponse{User: in}, nil
}

func (st *UserStore) Get(ctx context.Context, in *service.GetUserRequest) (*service.UserResponse, error) {
	user, err := st.get(in.Id)
	return &service.UserResponse{User: user}, err
}

func (st *UserStore) get(uid int32) (*model.User, error) {
	user := &model.User{Id: uid}
	err := st.db.First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
