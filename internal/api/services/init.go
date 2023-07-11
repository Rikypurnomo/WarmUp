package services

import (
	"github.com/Rikypurnomo/warmup/internal/api/repository"
	"github.com/Rikypurnomo/warmup/pkg/database"
)

type (
	ServicessInit struct {
		RepositoryAuth        repository.UsersRepository
		RepositoryProduct     repository.ProductsRepository
		RepositoryCategori    repository.CategoriesRepository
		RepositoryCart        repository.CartsRepository
		RepositoryTransaction repository.TransactionsRepository
		// RepositoryHistory     repository.HistoryRepository
	}
)

func InitiateServicessInterface() *ServicessInit {
	return &ServicessInit{
		RepositoryAuth:        repository.NewUsersRepository(database.GetDbCon()),
		RepositoryProduct:     repository.NewProductsRepository(database.GetDbCon()),
		RepositoryCategori:    repository.NewCategoriesRepository(database.GetDbCon()),
		RepositoryCart:        repository.NewCartRepository(database.GetDbCon()),
		RepositoryTransaction: repository.NewTransactionRepository(database.GetDbCon()),
		// RepositoryHistory:     repository.NewHistoryRepository(database.GetDbCon()),
	}
}
