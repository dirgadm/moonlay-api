package domain

import (
	"context"
	"net/http"
	"project-version3/moonlay-api/service/domain/dto"
)

// type Lists struct {
// 	Id          int `gorm:"primaryKey;autoIncrement:true"`
// 	Title       string
// 	Description string
// 	Priority    int
// 	CreatedAt   time.Time
// 	UpdatedAt   time.Time
// }

// func (m *Lists) TableName() string {
// 	return "lists"
// }

type UploadUsecase interface {
	// Upload(ctx context.Context, w http.ResponseWriter, r *http.Request) (res dto.UploadResponse, err error)
	UploadFile(ctx context.Context, w http.ResponseWriter, r *http.Request) (res []dto.UploadResponse, err error)
}

// type ListsRepository interface {
// 	GetList(ctx context.Context, offset int, limit int, search string) (carts []Lists, count int64, err error)
// }
