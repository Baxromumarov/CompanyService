package db

import (
	"context"
	"testing"

	"github.com/baxromumarov/CompanyService/config"
	"github.com/baxromumarov/CompanyService/models"
	"github.com/baxromumarov/CompanyService/pkg/helper"
	"github.com/baxromumarov/CompanyService/pkg/logger"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var db, _ = InitPostgres(config.Load())

var repo = NewCompanyRepo(db, logger.New("debug", "test"))

func TestCreate(t *testing.T) {
	var mockData = &models.Company{
		ID:                uuid.NewString(),
		Name:              helper.GenerateRandomString(15),
		Description:       "Test Description",
		AmountOfEmployees: 100,
		Registered:        false,
		Type:              "NonProfit",
	}
	repo := NewCompanyRepo(db, logger.New(logger.LevelDebug, "test"))

	err := repo.Create(context.Background(), mockData)
	if err != nil {
		t.Fatalf("error creating company: %v", err)
	}
}

func TestGet(t *testing.T) {
	companyId := uuid.NewString()
	company := models.Company{
		ID:                companyId,
		Name:              helper.GenerateRandomString(15),
		Description:       "mock Description",
		AmountOfEmployees: 341,
		Registered:        true,
		Type:              "NonProfit",
	}

	err := repo.Create(context.Background(), &company)
	if err != nil {
		t.Fatalf("error creating company: %v", err)

	}

	// Create the company
	result, err := repo.Get(context.Background(), companyId)
	if err != nil {
		t.Fatalf("error creating company: %v", err)
	}

	assert.Equal(t, result.Name, company.Name, "company name should match")
	assert.Equal(t, result.Description, company.Description, "company description should match")
	assert.Equal(t, result.AmountOfEmployees, company.AmountOfEmployees, "company AmountOfEmployees should match")
	assert.Equal(t, result.Registered, company.Registered, "company Registered should match")
	assert.Equal(t, result.Type, company.Type, "company Registered should match")
}

func TestUpdate(t *testing.T) {
	companyId := uuid.NewString()
	company := models.Company{
		ID:                companyId,
		Name:              helper.GenerateRandomString(15),
		Description:       "mock Description",
		AmountOfEmployees: 341,
		Registered:        true,
		Type:              "NonProfit",
	}

	err := repo.Create(context.Background(), &company)
	if err != nil {
		t.Fatalf("error creating company: %v", err)

	}

	company.Name = helper.GenerateRandomString(15)

	err = repo.Update(context.Background(), &company)
	if err != nil {
		t.Fatalf("error updating company: %v", err)
	}

	// Fetch the updated company
	updatedCompany, err := repo.Get(context.Background(), companyId)
	if err != nil {
		t.Fatalf("failed to fetch updated company: %v", err)
	}
	assert.Equal(t, company.Name, updatedCompany.Name, "company name should be updated")
}

func TestDelete(t *testing.T) {
	companyId := uuid.NewString()
	company := models.Company{
		ID:                companyId,
		Name:              helper.GenerateRandomString(15),
		Description:       "mock Description",
		AmountOfEmployees: 341,
		Registered:        true,
		Type:              "NonProfit",
	}

	err := repo.Create(context.Background(), &company)
	if err != nil {
		t.Fatalf("error creating company: %v", err)

	}

	err = repo.Delete(context.Background(), companyId)
	if err != nil {
		t.Fatalf("failed to delete company: %v", err)
	}

	deletedCompany, err := repo.Get(context.Background(), companyId)
	if err == nil {
		t.Fatalf("expected error when fetching deleted company, but got: %v", deletedCompany)
	}
}
