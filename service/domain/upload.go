package domain

import (
	"context"
	"net/http"
	"project-version3/moonlay-api/service/domain/dto"
	"time"
)

type UploadedFile struct {
	Id        int `gorm:"primaryKey;autoIncrement:true"`
	ListId    int
	SubListId int
	FileName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *UploadedFile) TableName() string {
	return "uploaded_files"
}

type UploadUsecase interface {
	// Upload(ctx context.Context, w http.ResponseWriter, r *http.Request) (res dto.UploadResponse, err error)
	UploadFile(ctx context.Context, w http.ResponseWriter, r *http.Request) (res []dto.UploadResponse, err error)
}

type UploadRepository interface {
	GetListByListId(ctx context.Context, offset int, limit int, search string, listId int) (uf []UploadedFile, count int64, err error)
	GetListBySubListId(ctx context.Context, offset int, limit int, search string, subListId int) (uf []UploadedFile, count int64, err error)
	DeleteByListId(ctx context.Context, listId int) (err error)
	DeleteBySubListId(ctx context.Context, subListId int) (err error)
	Create(ctx context.Context, upload *UploadedFile) (err error)
}
