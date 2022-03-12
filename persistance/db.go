package persistance

import (
	"fmt"
	"github.com/oismailov/launchpad-booking/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func GetInstance() *gorm.DB {
	if dbInstance != nil {
		return dbInstance
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.Cfg.DatabaseSettings.Host,
		config.Cfg.DatabaseSettings.Username,
		config.Cfg.DatabaseSettings.Password,
		config.Cfg.DatabaseSettings.Name,
		config.Cfg.DatabaseSettings.Port,
		config.Cfg.DatabaseSettings.SSLMode,
		config.Cfg.DatabaseSettings.Timezone,
	)
	dbInstance, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("db connection error: %v", err)
	}

	return dbInstance
}
