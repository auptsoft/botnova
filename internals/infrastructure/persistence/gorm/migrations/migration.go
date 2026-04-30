package migrations

import (
	"log"

	"auptex.com/botnova/internals/infrastructure/persistence/gorm/entities"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "001_create_models",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(
					&entities.Project{},
					&entities.User{},
					&entities.Robot{},
					&entities.RobotModel{},
					&entities.Script{},
					&entities.RobotEndpoint{},
					&entities.RobotGroup{},
					&entities.RobotGroupMember{},
					&entities.CalibrationEntity{},
				)
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("model_entities")
			},
		},
	})

	if err := m.Migrate(); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migration successful")
}
