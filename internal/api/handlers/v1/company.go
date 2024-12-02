package v1

import (
	"fmt"

	"github.com/baxromumarov/CompanyService/config"
	"github.com/baxromumarov/CompanyService/models"
	"github.com/baxromumarov/CompanyService/pkg/helper"
	"github.com/gin-gonic/gin"
)

func (h *handlerV1) CreateCompany(c *gin.Context) {
	var company *models.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if company.Name == "" || company.AmountOfEmployees < 0 || company.Type == "" {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	if _, ok := config.CompanyType[company.Type]; !ok {
		c.JSON(400, gin.H{"error": "Invalid company type"})
		return
	}

	fmt.Println("company: ", company)
	if err := h.storagePostgres.Company().Create(c, company); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, company)
}

func (h *handlerV1) UpdateCompany(c *gin.Context) {
	var (
		company *models.Company
	)

	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	company.ID = c.Param("id")
	if !helper.IsValidUUID(company.ID) {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	if err := h.storagePostgres.Company().Update(c, company); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, company)
}

func (h *handlerV1) DeleteCompany(c *gin.Context) {
	id := c.Param("id")
	if !helper.IsValidUUID(id) {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	if err := h.storagePostgres.Company().Delete(c, id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Company deleted successfully"})
}

func (h *handlerV1) GetCompany(c *gin.Context) {
	id := c.Param("id")
	if !helper.IsValidUUID(id) {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	company, err := h.storagePostgres.Company().Get(c, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, company)
}
