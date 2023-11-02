package mysql

import (
	"context"
	"project-version3/moonlay-api/service/domain"

	"gorm.io/gorm"
)

type uploadRepository struct {
	Conn *gorm.DB
}

func NewlUploadRepository(conn *gorm.DB) domain.UploadRepository {
	return &uploadRepository{conn}
}

func (m *uploadRepository) GetListByListId(ctx context.Context, offset int, limit int, search string, listId int) (upload []domain.UploadedFile, count int64, err error) {
	gorm := m.Conn.Model(domain.UploadedFile{})

	if listId != 0 {
		gorm = gorm.Where("list_id = ?", listId)
	}

	err = gorm.Count(&count).Error
	if err != nil {
		return
	}

	err = gorm.Offset(offset).Limit(limit).Find(&upload).Error

	return
}

func (m *uploadRepository) GetListBySubListId(ctx context.Context, offset int, limit int, search string, subListId int) (upload []domain.UploadedFile, count int64, err error) {
	gorm := m.Conn.Model(domain.UploadedFile{})

	if subListId != 0 {
		gorm = gorm.Where("sub_list_id = ?", subListId)
	}

	err = gorm.Count(&count).Error
	if err != nil {
		return
	}

	err = gorm.Offset(offset).Limit(limit).Find(&upload).Error

	return
}

func (m *uploadRepository) DeleteByListId(ctx context.Context, listId int) (err error) {
	err = m.Conn.Where("list_id = ?", listId).Delete(domain.UploadedFile{}).Error
	return
}

func (m *uploadRepository) DeleteBySubListId(ctx context.Context, listId int) (err error) {
	err = m.Conn.Where("sub_list_id = ?", listId).Delete(domain.UploadedFile{}).Error
	return
}

func (m *uploadRepository) Create(ctx context.Context, upload *domain.UploadedFile) (err error) {
	err = m.Conn.Create(upload).Error
	return
}
