package db

import (
	"context"
	"log"
	model "sharing_vision/models"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDB(conf model.ServerConfig) *Database {
	var err error
	var db *gorm.DB

	var host, port, user, pass, name string

	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in newDB", r)
		}
	}()

	host = conf.DBConfig.Host
	port = conf.DBConfig.Port
	user = conf.DBConfig.User
	pass = conf.DBConfig.Password
	name = conf.DBConfig.Name

	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// check if database exists	 then use it
	if err := db.Exec("USE " + name).Error; err != nil {
		log.Println(err, "NewDB: recover from contract db init")
		// create database
		if err := db.Exec("CREATE DATABASE " + name).Error; err != nil {
			log.Fatal(err)
		}
	}

	var ctx context.Context
	db = db.WithContext(ctx)

	dbSql, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	// db conneection pool
	dbSql.SetMaxIdleConns(10)
	dbSql.SetMaxOpenConns(100)
	dbSql.SetConnMaxLifetime(0)

	err = dbSql.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		go doEvery(10*time.Minute, pingDb, db)
		return &Database{DB: db}
	}

	return &Database{DB: db}
}

func doEvery(d time.Duration, f func(*gorm.DB), x *gorm.DB) {
	for range time.Tick(d) {
		f(x)
	}
}

func pingDb(db *gorm.DB) {
	dbSQL, err := db.DB()

	if err != nil {
		log.Println(err, "pingDB: recover from contract db init")
	}

	err = dbSQL.Ping()
	if err != nil {
		log.Println(err, "pingDB: recover from contract db init")
	}
}

func (d *Database) AutoMigrate(schemas ...interface{}) {
	for _, schema := range schemas {
		if err := d.DB.AutoMigrate(schema); err != nil {
			log.Println(err, "AutoMigrate: recover from contract db init")
		}
	}
}

func (db *Database) DropTable(schemas ...interface{}) error {
	for _, schema := range schemas {

		if err := db.DB.Migrator().DropTable(schema); err != nil {
			log.Println(err, "DropTable: recover from contract db init")
			return err
		}
	}
	return nil
}
