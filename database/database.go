package database

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

const (
	Host         = "127.0.0.1"
	Port         = "3306"
	DatabaseName = "testdb"
	Username     = "root"
	Password     = "root"
)

func GetDB(c *fiber.Ctx) *gorm.DB {

	return DB
}

func getHost() string {
	host := Host

	return host
}

func getPort() string {
	return Port
}

func getDatabase() string {
	dbName := DatabaseName

	return dbName
}

func getUsername() string {
	userName := Username

	return userName
}

func getPassword() string {
	password := Password

	return password
}

func Start() {
	dbName := getDatabase()
	dbUser := getUsername()
	dbPass := getPassword()
	dbHost := getHost()
	dbPort := getPort()
	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
	dsn = dsn + "?charset=utf8&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name,
		DefaultStringSize:         256,   // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision:  true,  // disable datetime precision support, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // use change when rename column, rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // smart configure based on used version
	}), &gorm.Config{
		//Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		panic(err)
	}
}
