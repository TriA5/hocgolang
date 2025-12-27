package migrations

import (
	"log"

	"be-ep/db"
	"be-ep/model"
)

func RunMigration() {
	log.Println("🚀 Running database migrations...")

	err := db.DB.AutoMigrate(
		&model.Role{},
		&model.User{},
	)
	if err != nil {
		log.Fatal("❌ Migration failed:", err)
	}

	log.Println("✅ Migration completed")
}
