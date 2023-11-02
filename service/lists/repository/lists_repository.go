package mysql

import (
	"context"
	"project-version3/moonlay-api/service/domain"

	"gorm.io/gorm"
)

type listsRepository struct {
	Conn *gorm.DB
}

func NewlListsRepository(conn *gorm.DB) domain.ListsRepository {
	return &listsRepository{conn}
}

func (m *listsRepository) GetList(ctx context.Context, offset int, limit int, search string) (lists []domain.Lists, count int64, err error) {
	gorm := m.Conn.Model(domain.Lists{})

	if search != "" {
		gorm = gorm.Where("title like ? or description like ?", "%"+search+"%", "%"+search+"%").Order("priority ASC")
	}

	err = gorm.Count(&count).Error
	if err != nil {
		return
	}

	err = gorm.Offset(offset).Limit(limit).Order("priority ASC").Find(&lists).Error

	return
}

func (m *listsRepository) GetDetail(ctx context.Context, id int) (lists domain.Lists, err error) {
	err = m.Conn.Where("id = ?", id).First(&lists).Error
	return
}

func (m *listsRepository) Create(ctx context.Context, lists *domain.Lists) (err error) {
	err = m.Conn.Create(&lists).Error
	return
}

func (m *listsRepository) Update(ctx context.Context, lists *domain.Lists) (err error) {
	err = m.Conn.Save(&lists).Error
	return
}

func (m *listsRepository) Delete(ctx context.Context, lists *domain.Lists) (err error) {
	err = m.Conn.Delete(&lists).Error
	return
}
