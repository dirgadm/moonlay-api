package usecase

import (
	"context"
	"time"

	"project-version3/moonlay-api/pkg/ehttp"
	"project-version3/moonlay-api/service/domain"
	"project-version3/moonlay-api/service/domain/dto"

	"github.com/labstack/gommon/log"
)

type listsUsecase struct {
	listsRepo    domain.ListsRepository
	subListsRepo domain.SubListsRepository
	uploadRepo   domain.UploadRepository
}

// NewListsUsecase will create new an articleUsecase object representation of domain.ListsUsecase interface
func NewListsUsecase(u domain.ListsRepository, s domain.SubListsRepository, uf domain.UploadRepository, timeout time.Duration) domain.ListsUsecase {
	return &listsUsecase{
		subListsRepo: s,
		uploadRepo:   uf,
		listsRepo:    u,
	}
}

func (s *listsUsecase) GetList(ctx context.Context, offset int, limit int, search string) (res []dto.ListsResponse, total int64, err error) {
	var lists []domain.Lists
	lists, total, err = s.listsRepo.GetList(ctx, offset, limit, search)
	if err != nil {
		log.Error(err)
		return
	}

	for _, l := range lists {
		res = append(res, dto.ListsResponse{
			Id:          l.Id,
			Title:       l.Title,
			Description: l.Description,
			Priority:    l.Priority,
			CreatedAt:   l.CreatedAt,
			UpdatedAt:   l.UpdatedAt,
		})
	}

	return
}

func (s *listsUsecase) GetDetail(ctx context.Context, id int) (res dto.ListsResponse, err error) {
	var list domain.Lists
	list, err = s.listsRepo.GetDetail(ctx, id)
	if err != nil {
		log.Error(err)
		return
	}

	res = dto.ListsResponse{
		Id:          list.Id,
		Title:       list.Title,
		Description: list.Description,
		Priority:    list.Priority,
		CreatedAt:   list.CreatedAt,
		UpdatedAt:   list.UpdatedAt,
	}

	return
}

func (s *listsUsecase) Update(ctx context.Context, req dto.ListsRequest) (err error) {

	// validate list
	var list domain.Lists
	list, err = s.listsRepo.GetDetail(ctx, req.Id)
	if err != nil {
		log.Error(err)
		err = ehttp.ErrorOutput("id", "The list is invalid")
		return
	}

	if len(req.Files) > 0 {
		if err = s.uploadRepo.DeleteByListId(ctx, list.Id); err != nil {
			log.Error(err)
			err = ehttp.ErrorOutput("id", "Failed to deleted file")
			return
		}

		for _, v := range req.Files {
			upload := &domain.UploadedFile{
				ListId:    list.Id,
				SubListId: 0,
				FileName:  v,
				UpdatedAt: time.Now(),
			}
			err = s.uploadRepo.Create(ctx, upload)
			if err != nil {
				log.Error(err)
				return
			}
		}

	}
	list.Title = req.Title
	list.Description = req.Description
	err = s.listsRepo.Update(ctx, &list)
	if err != nil {
		log.Error(err)
		return
	}

	return
}

func (s *listsUsecase) Create(ctx context.Context, req dto.ListsRequest) (err error) {

	// var list domain.Lists
	list := &domain.Lists{
		Title:       req.Title,
		Description: req.Description,
		Priority:    1,
		CreatedAt:   time.Now(),
	}

	err = s.listsRepo.Create(ctx, list)
	if err != nil {
		log.Error(err)
		return
	}

	if len(req.Files) > 0 {
		for _, v := range req.Files {
			upload := &domain.UploadedFile{
				ListId:    list.Id,
				SubListId: 0,
				FileName:  v,
				UpdatedAt: time.Now(),
			}
			err = s.uploadRepo.Create(ctx, upload)
			if err != nil {
				log.Error(err)
				return
			}
		}

	}

	return
}

func (s *listsUsecase) Delete(ctx context.Context, id int) (err error) {

	var list domain.Lists
	list, err = s.listsRepo.GetDetail(ctx, id)
	if err != nil {
		log.Error(err)
		return
	}

	// delete file based on list id
	if err = s.uploadRepo.DeleteByListId(ctx, list.Id); err != nil {
		log.Error(err)
		err = ehttp.ErrorOutput("id", "Failed to deleted file")
		return
	}

	// delete file based on sublist id and sublist itself
	var subLists []domain.SubLists
	subLists, _, err = s.subListsRepo.GetList(ctx, 0, 0, "", list.Id)
	for _, v := range subLists {
		if err = s.uploadRepo.DeleteBySubListId(ctx, v.Id); err != nil {
			log.Error(err)
			err = ehttp.ErrorOutput("id", "Failed to deleted file")
			return
		}

		if err = s.subListsRepo.Delete(ctx, &v); err != nil {
			log.Error(err)
			err = ehttp.ErrorOutput("id", "Failed to deleted sublist")
			return
		}
	}

	// delete list based id
	if err = s.listsRepo.Delete(ctx, &list); err != nil {
		log.Error(err)
		err = ehttp.ErrorOutput("id", "Failed to deleted list")
		return
	}

	return
}
