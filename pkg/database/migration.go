package database

import "github.com/Rikypurnomo/warmup/internal/api/models"

type Migrates []interface{}

var MigrateList Migrates = Migrates{
	// fill your model here
	&models.User{},
	&models.Product{},
	&models.Category{},
	&models.Cart{},
	&models.Transaction{},
	&models.History{},
}
