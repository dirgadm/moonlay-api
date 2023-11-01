package domain

import (
	"context"
	"project-version3/moonlay-api/service/domain/dto"
	"time"
)

type Lists struct {
	Id          int `gorm:"primaryKey;autoIncrement:true"`
	Title       string
	Description string
	Priority    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (m *Lists) TableName() string {
	return "lists"
}

type ListsUsecase interface {
	GetList(ctx context.Context, offset int, limit int, search string) (res []dto.ListsResponse, total int64, err error)
	GetDetail(ctx context.Context, id int) (res dto.ListsResponse, err error)
}

type ListsRepository interface {
	GetList(ctx context.Context, offset int, limit int, search string) (lists []Lists, count int64, err error)
	GetDetail(ctx context.Context, id int) (list Lists, err error)
}
