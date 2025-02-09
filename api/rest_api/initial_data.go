package main

import (
	"github.com/google/uuid"
	"github.com/mgpaja8/pavs-relic/internal/domain/models"
	"github.com/mgpaja8/pavs-relic/internal/domain/valueobjects"
)

var (
	companyGoogle, _   = models.NewCompany(valueobjects.CompanyName("Google"))
	companyNewRelic, _ = models.NewCompany(valueobjects.CompanyName("New Relic"))

	companiesMap = map[uuid.UUID]models.Company{
		companyGoogle.ID():   companyGoogle,
		companyNewRelic.ID(): companyNewRelic,
	}

	customer1, _  = models.NewCustomer(companyGoogle, valueobjects.FirstName("Jane"), valueobjects.LastName("Doe"))
	customer2, _  = models.NewCustomer(companyGoogle, valueobjects.FirstName("John"), valueobjects.LastName("Smith"))
	customer3, _  = models.NewCustomer(companyGoogle, valueobjects.FirstName("Alice"), valueobjects.LastName("Johnson"))
	customer4, _  = models.NewCustomer(companyNewRelic, valueobjects.FirstName("Bob"), valueobjects.LastName("Brown"))
	customer5, _  = models.NewCustomer(companyNewRelic, valueobjects.FirstName("Eve"), valueobjects.LastName("Lee"))
	customer6, _  = models.NewCustomer(companyNewRelic, valueobjects.FirstName("Michael"), valueobjects.LastName("Wilson"))
	customer7, _  = models.NewCustomer(companyNewRelic, valueobjects.FirstName("Sarah"), valueobjects.LastName("Miller"))
	customer8, _  = models.NewCustomer(companyNewRelic, valueobjects.FirstName("David"), valueobjects.LastName("Taylor"))
	customer9, _  = models.NewCustomer(companyNewRelic, valueobjects.FirstName("Emily"), valueobjects.LastName("Anderson"))
	customer10, _ = models.NewCustomer(companyNewRelic, valueobjects.FirstName("Jane"), valueobjects.LastName("Parker"))

	customersMap = map[uuid.UUID]models.Customer{
		customer1.ID():  customer1,
		customer2.ID():  customer2,
		customer3.ID():  customer3,
		customer4.ID():  customer4,
		customer5.ID():  customer5,
		customer6.ID():  customer6,
		customer7.ID():  customer7,
		customer8.ID():  customer8,
		customer9.ID():  customer9,
		customer10.ID(): customer10,
	}
)
