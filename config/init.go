package config

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitMasterConfig df
func InitMasterConfig() DBConfig {
	var cfg DBConfig
	f, err := os.Open("config.yaml")
	if err != nil {
		panic(err)
	}
	decoder := yaml.NewDecoder(f)
	defer f.Close()
	err = decoder.Decode(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

// InitDB df
func InitDB(credentials DBConfig) *gorm.DB {
	dsn := credentials.Username + ":" + credentials.Password + "@(" + credentials.Host + ":" + credentials.Port + ")/" + credentials.DBName + "?charset=utf8&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // Disable color
		},
	)
	// logger := zapgorm2.New(log.GetLogger().Named("gorm"))
	// logger.SetAsDefault()
	options := &gorm.Config{
		Logger: newLogger,
	}
	db, err := gorm.Open(mysql.Open(dsn), options)
	db.Set("gorm:table_options", "ENGINE=InnoDB").Debug().AutoMigrate()
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(credentials.MaxIdleConns)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(credentials.MaxOpenConns)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(credentials.ConnMaxLifetime))
	if err != nil {
		log.Fatalf("Mysql Connection Error", err)
	}
	return db
}

func GetDB() *gorm.DB {
	cfg := InitMasterConfig()
	return InitDB(cfg)
}
