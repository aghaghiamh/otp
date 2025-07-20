package adaptor

import (
	"fmt"
	"otp/src/model"
	"otp/src/pkg/config"
	"otp/src/pkg/log"

	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func CreatePostgresqlDbClient() *gorm.DB {
	db, err := gorm.Open(postgres.Open(generatePostgresConnectionString()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	})
	if err != nil {
		fmt.Println(err)
		// TODO: Add Logger.Fatal
	}

	if config.GetAppConfigInstance().AutoMigrationEnable {
		err := db.AutoMigrate(&model.User{})
		if err != nil {
			log.GetLoggerInstance().Fatal(err)
			return nil
		}
	}

	return db
}

func generatePostgresConnectionString() string {
	cnf := config.GetAppConfigInstance()
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cnf.Postgres.Host, cnf.Postgres.Port, cnf.Postgres.Username, cnf.Postgres.Password, cnf.Postgres.DB)
}
