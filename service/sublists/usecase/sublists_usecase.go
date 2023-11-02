package usecase

import (
	"context"
	"errors"
	"time"

	"project-version3/moonlay-api/pkg/ehttp"
	"project-version3/moonlay-api/service/domain"
	"project-version3/moonlay-api/service/domain/dto"

	"github.com/labstack/gommon/log"
)

type subListsUsecase struct {
	subListsRepo domain.SubListsRepository
	listsRepo    domain.ListsRepository
	uploadRepo   domain.UploadRepository
}

// NewListsUsecase will create new an articleUsecase object representation of domain.ListsUsecase interface
func NewSubListsUsecase(u domain.SubListsRepository, l domain.ListsRepository, uf domain.UploadRepository, timeout time.Duration) domain.SubListsUsecase {
	return &subListsUsecase{
		subListsRepo: u,
		listsRepo:    l,
		uploadRepo:   uf,
	}
}

func (s *subListsUsecase) GetList(ctx context.Context, offset int, limit int, search string, listId int) (res []dto.SubListsResponse, total int64, err error) {
	var sublists []domain.SubLists
	sublists, total, err = s.subListsRepo.GetList(ctx, offset, limit, search, listId)
	if err != nil {
		log.Error(err)
		return
	}

	for _, sl := range sublists {
		var list domain.Lists
		list, err = s.listsRepo.GetDetail(ctx, sl.ListId)
		if err != nil {
			log.Error(err)
			return
		}

		res = append(res, dto.SubListsResponse{
			Id:          sl.Id,
			Title:       sl.Title,
			ListId:      sl.ListId,
			Description: sl.Description,
			Priority:    sl.Priority,
			CreatedAt:   sl.CreatedAt,
			UpdatedAt:   sl.UpdatedAt,
			List: dto.ListsResponse{
				Id:          list.Id,
				Title:       list.Title,
				Description: list.Description,
				Priority:    list.Priority,
				CreatedAt:   list.CreatedAt,
				UpdatedAt:   list.UpdatedAt,
			},
		})
	}

	return
}

func (s *subListsUsecase) GetDetail(ctx context.Context, id int) (res dto.SubListsResponse, err error) {
	var subList domain.SubLists
	subList, err = s.subListsRepo.GetDetail(ctx, id)
	if err != nil {
		log.Error(err)
		return
	}

	var list domain.Lists
	list, err = s.listsRepo.GetDetail(ctx, subList.ListId)
	if err != nil {
		log.Error(err)
		return
	}

	var files []domain.UploadedFile
	files, _, err = s.uploadRepo.GetListBySubListId(ctx, 0, 100, "", subList.Id)
	if err != nil {
		log.Error(err)
		return
	}

	res = dto.SubListsResponse{
		Id:          subList.Id,
		Title:       subList.Title,
		ListId:      subList.ListId,
		Description: subList.Description,
		Priority:    subList.Priority,
		CreatedAt:   subList.CreatedAt,
		UpdatedAt:   subList.UpdatedAt,
		List: dto.ListsResponse{
			Id:          list.Id,
			Title:       list.Title,
			Description: list.Description,
			Priority:    list.Priority,
			CreatedAt:   list.CreatedAt,
			UpdatedAt:   list.UpdatedAt,
		},
	}

	var filesArrRes []dto.UploadResponse
	var filesRes dto.UploadResponse
	if len(files) > 0 {
		for _, v := range files {
			filesRes = dto.UploadResponse{
				FileName: v.FileName,
			}
			filesArrRes = append(filesArrRes, filesRes)
		}
		res.Files = filesArrRes
	}

	return
}

func (s *subListsUsecase) Update(ctx context.Context, req dto.SubListsRequest) (err error) {

	var list domain.Lists
	list, err = s.listsRepo.GetDetail(ctx, req.ListId)
	if err != nil {
		log.Error(err)
		return
	}
	// validate sublist
	var sublist domain.SubLists
	sublist, err = s.subListsRepo.GetDetail(ctx, req.Id)
	if err != nil {
		log.Error(err)
		err = ehttp.ErrorOutput("id", "The list is invalid")
		return
	}

	if req.ListId != sublist.ListId {
		log.Error(err)
		err = ehttp.ErrorOutput("id", "This sublist not belongs to the list. Check list Id")
		return
	}

	if len(req.Files) > 0 {
		if err = s.uploadRepo.DeleteBySubListId(ctx, sublist.Id); err != nil {
			log.Error(err)
			err = ehttp.ErrorOutput("id", "Failed to deleted file")
			return
		}

		for _, v := range req.Files {
			upload := &domain.UploadedFile{
				ListId:    0,
				SubListId: sublist.Id,
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
	sublist.Title = req.Title
	sublist.ListId = list.Id
	sublist.Description = req.Description
	err = s.subListsRepo.Update(ctx, &sublist)
	if err != nil {
		log.Error(err)
		return
	}

	return
}

func (s *subListsUsecase) Create(ctx context.Context, req dto.SubListsRequest) (err error) {

	if req.ListId == 0 {
		err = errors.New("List id is required ")
		log.Error(err)
		return
	}
	var list domain.Lists
	list, err = s.listsRepo.GetDetail(ctx, req.ListId)
	if err != nil {
		log.Error(err)
		return
	}
	// var list domain.Lists
	sublist := &domain.SubLists{
		Title:       req.Title,
		ListId:      list.Id,
		Description: req.Description,
		Priority:    1,
		CreatedAt:   time.Now(),
	}
	err = s.subListsRepo.Create(ctx, sublist)
	if err != nil {
		log.Error(err)
		return
	}

	if len(req.Files) > 0 {
		for _, v := range req.Files {
			upload := &domain.UploadedFile{
				ListId:    0,
				SubListId: sublist.Id,
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

func (s *subListsUsecase) Delete(ctx context.Context, id int) (err error) {

	var sublist domain.SubLists
	sublist, err = s.subListsRepo.GetDetail(ctx, id)
	if err != nil {
		log.Error(err)
		return
	}

	// delete file based on list id
	if err = s.uploadRepo.DeleteBySubListId(ctx, sublist.Id); err != nil {
		log.Error(err)
		err = ehttp.ErrorOutput("id", "Failed to deleted file")
		return
	}

	// delete list based id
	if err = s.subListsRepo.Delete(ctx, &sublist); err != nil {
		log.Error(err)
		err = ehttp.ErrorOutput("id", "Failed to deleted list")
		return
	}

	return
}
