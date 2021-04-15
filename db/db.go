package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"bitbucket.org/liamstask/goose/lib/goose"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_user_name = "mysql_user_name"
	mysql_password  = "mysql_password"
	mysql_host      = "mysql_host"
	mysql_schema    = "mysql_schema"
)

type UserDbService struct{}

var (
	db *gorm.DB

	// userName = os.Getenv(mysql_user_name)
	// password = os.Getenv(mysql_password)
	// host     = os.Getenv(mysql_host)
	// schema   = os.Getenv(mysql_schema)
	userName = "root"
	password = "root1"
	host     = "127.0.0.1:3306"
	schema   = "users_db"
)

func init() {
	var err error

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		userName, password, host, schema,
	)
	fmt.Println(dataSourceName)
	db, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("Connected to DB successfully")
	workingDir, err := os.Getwd()
	if err != nil {
		log.Println("Unable to fetch the working directory")
		os.Exit(1)
	}
	workingDir = workingDir + "/db/migrations"
	migrateConf := &goose.DBConf{
		MigrationsDir: workingDir,
		Driver: goose.DBDriver{
			Name:    "mysql",
			OpenStr: dataSourceName,
			Import:  "github.com/go-sql-driver/mysql",
			Dialect: &goose.MySqlDialect{},
		},
	}
	latest, err := goose.GetMostRecentDBVersion(migrateConf.MigrationsDir)
	if err != nil {
		log.Println("Unable to fetch the most recent DB version")
	}
	sqlDB, _ := db.DB()
	err = goose.RunMigrationsOnDb(migrateConf, migrateConf.MigrationsDir, latest, sqlDB)
	if err != nil {
		log.Println("Error while running migrations")
		os.Exit(1)
	}

}

func (d UserDbService) GetDb() *gorm.DB {
	return db
}
