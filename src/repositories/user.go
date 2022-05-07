package repositories

import (
	"context"

	userGrpc "github.com/hanherb/go-playground/grpc-gen"
	"github.com/hanherb/go-playground/src/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Get(ctx context.Context, props *userGrpc.UserGetOneRequest) error
	Data() *models.User
}

type UserRepositories interface {
	Get(ctx context.Context, props *userGrpc.UserGetListRequest) (*userGrpc.UserGetListResponse, error)
	Data() []*models.User
}

//==================================================================

type UserRepositoryImp struct {
	query *gorm.DB
	data  models.User
}

type UserRepositoriesImp struct {
	query *gorm.DB
	data  []*models.User
}

//==================================================================

func NewUserRepository(db *gorm.DB) UserRepository {
	result := new(UserRepositoryImp)
	result.query = db
	return result
}

func (r *UserRepositoryImp) Get(ctx context.Context, props *userGrpc.UserGetOneRequest) error {
	query := r.query

	if &props.Id != nil {
		query = query.Where("id = ?", props.Id)
	}

	if err := query.WithContext(ctx).First(&r.data).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepositoryImp) Data() *models.User {
	return &r.data
}

//==================================================================

func NewUserRepositories(db *gorm.DB) UserRepositories {
	result := new(UserRepositoriesImp)
	result.query = db
	return result
}

func (r *UserRepositoriesImp) Get(ctx context.Context, props *userGrpc.UserGetListRequest) (*userGrpc.UserGetListResponse, error) {
	response := &userGrpc.UserGetListResponse{}

	query := r.query

	if props.City != nil {
		query = query.Where("city = ?", props.City)
	}

	//select query
	if err := query.WithContext(ctx).Find(&r.data).Error; err != nil {
		return nil, err
	}

	//count query
	count := int64(0)
	if err := query.Model(r.data).Count(&count).Error; err != nil {
		return nil, err
	}

	response.Count = int32(count)

	return response, nil
}

func (r *UserRepositoriesImp) Data() []*models.User {
	return r.data
}
