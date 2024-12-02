package db

import (
	"context"

	"github.com/baxromumarov/CompanyService/models"
	"github.com/baxromumarov/CompanyService/pkg/logger"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type company struct {
	db  *sqlx.DB
	log logger.Logger
}

func NewCompanyRepo(db *sqlx.DB, log logger.Logger) CompanyRepo {
	return &company{
		db:  db,
		log: log,
	}
}

func (c *company) Create(ctx context.Context, req *models.Company) error {
	req.ID = uuid.New().String()
	_, err := c.db.Exec(`
    		INSERT INTO companies 
    		(id, name, description, amount_of_employees, registered, type)
    		VALUES ($1, $2, $3, $4, $5, $6)`,
		req.ID,
		req.Name,
		req.Description,
		req.AmountOfEmployees,
		req.Registered,
		req.Type,
	)
	if err != nil {
		c.log.Error("error while creating company: " + err.Error())
		return err
	}
	return nil
}

func (c *company) Get(ctx context.Context, id string) (*models.Company, error) {
	var company = models.Company{
		ID: id,
	}

	err := c.db.QueryRow(`
			SELECT 
                name,
                description,
                amount_of_employees,
                registered,
                type
            FROM 
                companies 
            WHERE id = $1`, id,
	).Scan(
		&company.Name,
		&company.Description,
		&company.AmountOfEmployees,
		&company.Registered,
		&company.Type,
	)
	if err != nil {
		c.log.Error("error while getting company: " + err.Error())
		return nil, err
	}

	return &company, nil
}

func (c *company) Update(ctx context.Context, req *models.Company) error {
	_, err := c.db.Exec("UPDATE companies SET name = $1 WHERE id = $2", req.Name, req.ID)
	if err != nil {
		c.log.Error("error while updating company: " + err.Error())
		return err
	}
	return nil
}

func (c *company) Delete(ctx context.Context, id string) error {
	_, err := c.db.Exec("DELETE FROM companies WHERE id = $1", id)
	if err != nil {
		c.log.Error("error while deleting company: " + err.Error())
		return err
	}
	return nil
}
