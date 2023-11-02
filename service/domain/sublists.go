package domain

import (
	"context"
	"project-version3/moonlay-api/service/domain/dto"
	"time"
)

type SubLists struct {
	Id          int `gorm:"primaryKey;autoIncrement:true"`
	ListId      int
	Title       string
	Description string
	Priority    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (m *SubLists) TableName() string {
	return "sub_lists"
}

type SubListsUsecase interface {
	GetList(ctx context.Context, offset int, limit int, search string, listId int) (res []dto.SubListsResponse, total int64, err error)
	GetDetail(ctx context.Context, id int) (res dto.SubListsResponse, err error)
	Create(ctx context.Context, req dto.SubListsRequest) (err error)
	Update(ctx context.Context, req dto.SubListsRequest) (err error)
	Delete(ctx context.Context, id int) (err error)
}

type SubListsRepository interface {
	GetList(ctx context.Context, offset int, limit int, search string, listId int) (subLists []SubLists, count int64, err error)
	GetDetail(ctx context.Context, id int) (sublist SubLists, err error)
	Create(ctx context.Context, sublists *SubLists) (err error)
	Update(ctx context.Context, sublists *SubLists) (err error)
	Delete(ctx context.Context, sublists *SubLists) (err error)
}
