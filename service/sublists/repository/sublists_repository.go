package mysql

import (
	"context"
	"project-version3/moonlay-api/service/domain"

	"gorm.io/gorm"
)

type subListsRepository struct {
	Conn *gorm.DB
}

func NewlSubListsRepository(conn *gorm.DB) domain.SubListsRepository {
	return &subListsRepository{conn}
}

func (m *subListsRepository) GetList(ctx context.Context, offset int, limit int, search string, listId int) (sublists []domain.SubLists, count int64, err error) {
	gorm := m.Conn.Model(domain.SubLists{})

	if search != "" {
		gorm = gorm.Where("title like ? or description like ?", "%"+search+"%", "%"+search+"%").Order("priority ASC")
	}
	if listId != 0 {
		gorm = gorm.Where("list_id = ?", listId)
	}

	err = gorm.Count(&count).Error
	if err != nil {
		return
	}

	err = gorm.Offset(offset).Limit(limit).Order("priority ASC").Find(&sublists).Error

	return
}

func (m *subListsRepository) GetDetail(ctx context.Context, id int) (sublist domain.SubLists, err error) {
	err = m.Conn.Where("id = ?", id).First(&sublist).Error
	return
}

func (m *subListsRepository) Create(ctx context.Context, sublists *domain.SubLists) (err error) {
	err = m.Conn.Create(&sublists).Error
	return
}

func (m *subListsRepository) Update(ctx context.Context, sublists *domain.SubLists) (err error) {
	err = m.Conn.Save(&sublists).Error
	return
}

func (m *subListsRepository) Delete(ctx context.Context, sublists *domain.SubLists) (err error) {
	err = m.Conn.Delete(&sublists).Error
	return
}
