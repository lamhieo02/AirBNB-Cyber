package utils

import (
	"fmt"
	"go01-airbnb/config"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunDBMigration(cfg *config.Config) {
	dsn := fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s",
		cfg.Mysql.User, cfg.Mysql.Password, cfg.Mysql.Host, cfg.Mysql.Port, cfg.Mysql.DBName)
	// dsn := fmt.Sprintf("mysql://%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	cfg.Mysql.User, cfg.Mysql.Password, cfg.Mysql.ContainerName, cfg.Mysql.DBName)
	m, err := migrate.New(cfg.App.MigrationURL, dsn)
	if err != nil {
		log.Fatalln("Cannot run migrate db a", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalln("Cannot run migrate db b", err)
	}
}
