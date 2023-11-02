package mysql_test

import (
	"context"
	"project-version3/moonlay-api/service/domain"
	_uploadRepo "project-version3/moonlay-api/service/upload/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() (*gorm.DB, error) {
	// Set up an in-memory SQLite database for testing
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate your model to create the table
	err = db.AutoMigrate(&domain.UploadedFile{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestGetListByListId(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Error setting up test database: %v", err)
	}
	defer db.Migrator().DropTable(&domain.UploadedFile{}) // Clean up: Drop the table after the tests

	repo := _uploadRepo.NewlUploadRepository(db)

	// Test Create
	upload := &domain.UploadedFile{
		ListId:    1,
		FileName:  "input test1",
		CreatedAt: time.Now(),
	}

	err = repo.Create(context.Background(), upload)
	if err != nil {
		t.Fatalf("Error creating upload: %v", err)
	}

	retrievedList, _, err := repo.GetListByListId(context.Background(), 0, 10, "", 1)
	if err != nil {
		t.Fatalf("Error getting upload detail: %v", err)
	}

	if retrievedList[0].FileName != upload.FileName {
		t.Errorf("Expected title: %s, Got: %s", upload.FileName, retrievedList[0].FileName)
	}

	assert.NoError(t, err, "Error getting upload detail")
	assert.Len(t, retrievedList, 1)
	assert.Equal(t, upload.FileName, retrievedList[0].FileName, "Test Title")
}

func TestGetListBySubListId(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Error setting up test database: %v", err)
	}
	defer db.Migrator().DropTable(&domain.UploadedFile{}) // Clean up: Drop the table after the tests

	repo := _uploadRepo.NewlUploadRepository(db)

	// Test Create
	upload := &domain.UploadedFile{
		SubListId: 1,
		FileName:  "input test1",
		CreatedAt: time.Now(),
	}

	err = repo.Create(context.Background(), upload)
	if err != nil {
		t.Fatalf("Error creating upload: %v", err)
	}

	retrievedList, _, err := repo.GetListBySubListId(context.Background(), 0, 10, "", 1)
	if err != nil {
		t.Fatalf("Error getting upload detail: %v", err)
	}

	if retrievedList[0].FileName != upload.FileName {
		t.Errorf("Expected title: %s, Got: %s", upload.FileName, retrievedList[0].FileName)
	}

	assert.NoError(t, err, "Error getting upload detail")
	assert.Len(t, retrievedList, 1)
	assert.Equal(t, upload.FileName, retrievedList[0].FileName, "Test Title")
}

func TestCreate(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Error setting up test database: %v", err)
	}
	defer db.Migrator().DropTable(&domain.UploadedFile{}) // Clean up: Drop the table after the tests

	repo := _uploadRepo.NewlUploadRepository(db)

	// Test Create
	var uploads []*domain.UploadedFile
	upload := &domain.UploadedFile{
		SubListId: 1,
		ListId:    1,
		FileName:  "input test1",
		CreatedAt: time.Now(),
	}
	upload2 := &domain.UploadedFile{
		SubListId: 2,
		ListId:    1,
		FileName:  "input test1",
		CreatedAt: time.Now(),
	}
	uploads = append(uploads, upload, upload2)
	var ID []int
	for _, v := range uploads {
		err = repo.Create(context.Background(), v)
		if err != nil {
			t.Fatalf("Error creating upload: %v", err)
		}
		ID = append(ID, v.Id)
	}

	assert.NoError(t, err)
	assert.Len(t, ID, 2)
	assert.Equal(t, int(1), ID[0])
	assert.Equal(t, int(2), ID[1])
}

func TestDeleteByListId(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Error setting up test database: %v", err)
	}
	defer db.Migrator().DropTable(&domain.UploadedFile{}) // Clean up: Drop the table after the tests

	repo := _uploadRepo.NewlUploadRepository(db)

	// Test Create
	upload := &domain.UploadedFile{
		ListId:    1,
		FileName:  "input test1",
		CreatedAt: time.Now(),
	}

	err = repo.Create(context.Background(), upload)
	if err != nil {
		t.Fatalf("Error creating upload: %v", err)
	}

	err = repo.DeleteByListId(context.Background(), upload.ListId)
	if err != nil {
		t.Fatalf("Error creating upload: %v", err)
	}

	assert.NoError(t, err)
}

func TestDeleteBySubListId(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Error setting up test database: %v", err)
	}
	defer db.Migrator().DropTable(&domain.UploadedFile{}) // Clean up: Drop the table after the tests

	repo := _uploadRepo.NewlUploadRepository(db)

	// Test Create
	upload := &domain.UploadedFile{
		SubListId: 1,
		FileName:  "input test1",
		CreatedAt: time.Now(),
	}

	err = repo.Create(context.Background(), upload)
	if err != nil {
		t.Fatalf("Error creating upload: %v", err)
	}

	err = repo.DeleteByListId(context.Background(), upload.SubListId)
	if err != nil {
		t.Fatalf("Error creating upload: %v", err)
	}

	assert.NoError(t, err)
}
