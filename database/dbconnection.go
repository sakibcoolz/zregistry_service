package database

import (
	"fmt"
	"log"

	"zregistry_service/config"
	"zregistry_service/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBInstance struct {
	DB *gorm.DB
}

func Connection(configdb *config.Database) *DBInstance {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configdb.Username,
		configdb.Password,
		configdb.Host,
		configdb.Port,
		configdb.Schema)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return &DBInstance{DB: db}
}

func Migrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.TenantMaster{},
		&model.UserMaster{},
		&model.Contact{},
		&model.DeviceInfo{})
}
