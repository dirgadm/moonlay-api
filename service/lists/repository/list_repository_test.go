package mysql_test

import (
	"context"
	"project-version3/moonlay-api/service/domain"
	_listRepo "project-version3/moonlay-api/service/lists/repository"
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
	err = db.AutoMigrate(&domain.Lists{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestGetList(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Error setting up test database: %v", err)
	}
	defer db.Migrator().DropTable(&domain.Lists{}) // Clean up: Drop the table after the tests

	repo := _listRepo.NewlListsRepository(db)

	// Test Create
	list := &domain.Lists{
		Title:       "Test Title",
		Description: "Test Description",
		Priority:    1,
		CreatedAt:   time.Now(),
	}

	err = repo.Create(context.Background(), list)
	if err != nil {
		t.Fatalf("Error creating list: %v", err)
	}

	retrievedList, _, err := repo.GetList(context.Background(), 0, 10, "")
	if err != nil {
		t.Fatalf("Error getting list detail: %v", err)
	}

	if retrievedList[0].Title != list.Title {
		t.Errorf("Expected title: %s, Got: %s", list.Title, retrievedList[0].Title)
	}

	assert.NoError(t, err, "Error getting list detail")
	assert.Len(t, retrievedList, 1)
	assert.Equal(t, list.Title, retrievedList[0].Title, "Test Title")
	assert.Equal(t, list.Description, retrievedList[0].Description, "Test Description")
}

func TestGetDetail(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Error setting up test database: %v", err)
	}
	defer db.Migrator().DropTable(&domain.Lists{}) // Clean up: Drop the table after the tests

	repo := _listRepo.NewlListsRepository(db)

	// Test Create
	list := &domain.Lists{
		Title:       "Test Title 2",
		Description: "Test Description 2",
		Priority:    1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = repo.Create(context.Background(), list)
	if err != nil {
		t.Fatalf("Error creating list: %v", err)
	}

	// Test GetDetail
	retrievedList, err := repo.GetDetail(context.Background(), list.Id)
	if err != nil {
		t.Fatalf("Error getting list detail: %v", err)
	}

	assert.NoError(t, err)
	assert.NotNil(t, retrievedList)
}

func TestCreate(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Error setting up test database: %v", err)
	}
	defer db.Migrator().DropTable(&domain.Lists{}) // Clean up: Drop the table after the tests

	repo := _listRepo.NewlListsRepository(db)

	// Test Create
	var lists []*domain.Lists
	list := &domain.Lists{
		Title:       "Test Title 3",
		Description: "Test Description 3",
		Priority:    1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	list2 := &domain.Lists{
		Title:       "Test Title 1",
		Description: "Test Description 1",
		Priority:    1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	lists = append(lists, list, list2)
	var ID []int
	for _, v := range lists {
		err = repo.Create(context.Background(), v)
		if err != nil {
			t.Fatalf("Error creating list: %v", err)
		}
		ID = append(ID, v.Id)
	}

	assert.NoError(t, err)
	assert.Len(t, ID, 2)
	assert.Equal(t, int(1), ID[0])
	assert.Equal(t, int(2), ID[1])
}

func TestUpdate(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Error setting up test database: %v", err)
	}
	defer db.Migrator().DropTable(&domain.Lists{}) // Clean up: Drop the table after the tests

	repo := _listRepo.NewlListsRepository(db)

	// Test Create
	list := &domain.Lists{
		Title:       "Test Title 1",
		Description: "Test Description 1",
		Priority:    1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = repo.Create(context.Background(), list)
	if err != nil {
		t.Fatalf("Error creating list: %v", err)
	}

	list.Title = "Updated Title"

	err = repo.Update(context.Background(), list)
	if err != nil {
		t.Fatalf("Error creating list: %v", err)
	}

	// Test GetDetail
	retrievedList, err := repo.GetDetail(context.Background(), list.Id)
	if err != nil {
		t.Fatalf("Error getting list detail: %v", err)
	}

	assert.NoError(t, err)
	assert.NotNil(t, retrievedList)
	assert.Equal(t, "Updated Title", list.Title)
}

func TestDelete(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Error setting up test database: %v", err)
	}
	defer db.Migrator().DropTable(&domain.Lists{}) // Clean up: Drop the table after the tests

	repo := _listRepo.NewlListsRepository(db)

	// Test Create
	list := &domain.Lists{
		Title:       "Test Title 1",
		Description: "Test Description 1",
		Priority:    1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = repo.Create(context.Background(), list)
	if err != nil {
		t.Fatalf("Error creating list: %v", err)
	}

	err = repo.Delete(context.Background(), list)
	if err != nil {
		t.Fatalf("Error creating list: %v", err)
	}

	assert.NoError(t, err)
}
