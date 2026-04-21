package main

import (
	"auptex.com/botnova/internals/infrastructure/persistence/gorm"
	"auptex.com/botnova/internals/infrastructure/persistence/gorm/migrations"
)

func main() {
	dbConfig := gorm.Config{
		Driver: "sqlite",
		DSN:    "db.sqlite",
	}

	db := gorm.NewDB(dbConfig)
	migrations.Migrate(db)
}
