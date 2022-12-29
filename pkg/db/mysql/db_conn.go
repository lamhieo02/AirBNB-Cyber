package mysql

import (
	"fmt"
	"go01-airbnb/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQL(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Mysql.User, cfg.Mysql.Password,
		cfg.Mysql.Host, cfg.Mysql.Port, cfg.Mysql.DBName)
	//dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	//	cfg.Mysql.User, cfg.Mysql.Password, cfg.Mysql.ContainerName, cfg.Mysql.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
