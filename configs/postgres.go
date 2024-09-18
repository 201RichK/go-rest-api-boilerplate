package configs

import (
	"strconv"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Driver   string
	Host     string
	Username string
	Password string
	Port     int
	Database string
}

type Db struct {
	*gorm.DB
}

func New(config *DatabaseConfig) (*Db, error) {
	var db *gorm.DB
	var err error

	switch strings.ToLower(config.Driver) {
	case "mysql":
		dsn := config.Username + ":" + config.Password + "@tcp(" + config.Host + ":" + strconv.Itoa(config.Port) + ")/" + config.Database + "?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=True&loc=UTC"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		break
	case "postgresql", "postgres":
		dsn := "user=" + config.Username + " password=" + config.Password + " dbname=" + config.Database + " host=" + config.Host + " port=" + strconv.Itoa(config.Port) + " TimeZone=UTC"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		break
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(DatabaseSetting.Name+".db"), &gorm.Config{})
		break
	}

	return &Db{db}, err
}

func GetDbConfig() *DatabaseConfig {
	port, err := strconv.Atoi(DatabaseSetting.Port)
	if err != nil {
		panic(err)
	}

	return &DatabaseConfig{
		Driver:   DatabaseSetting.Type,
		Host:     DatabaseSetting.Host,
		Username: DatabaseSetting.User,
		Password: DatabaseSetting.Password,
		Port:     port,
		Database: DatabaseSetting.Name,
	}
}
