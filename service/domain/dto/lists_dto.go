package dto

import "time"

type ListsResponse struct {
	Id          int                `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Priority    int                `json:"priority"`
	CreatedAt   time.Time          `json:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at,omitempty"`
	SubList     []SubListsResponse `json:"sublist,omitempty"`
	Files       []UploadResponse   `json:"files,omitempty"`
}

type ListsRequest struct {
	Id          int      `json:"id"`
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Files       []string `json:"files"`
}
