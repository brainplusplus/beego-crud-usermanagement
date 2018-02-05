package repositories

import (
	_ "github.com/lib/pq"
	"time"
	"github.com/jinzhu/gorm"
)

type DataSource struct {
	*gorm.DB
	maxIdleConns     int
	maxOpenConns     int
	maxConnsLifetime time.Duration
}

func NewDatabaseConnectionWithConnectionPool(dbDriver string, dbUrl string, maxIdle int, maxOpenConnection int) (*DataSource, error) {
	db, err := gorm.Open(dbDriver, dbUrl)
	//defer db.Close()
	if err != nil {
		return nil, err
	}
	if err = db.DB().Ping(); err != nil {
		return nil, err
	}

	maxLifeTime := 1 * time.Second

	db.DB().SetMaxIdleConns(maxIdle)
	db.DB().SetMaxOpenConns(maxOpenConnection)
	db.DB().SetConnMaxLifetime(maxLifeTime)
	return &DataSource{db, maxIdle, maxOpenConnection, maxLifeTime}, nil
}

func NewDatabaseConnection(dbDriver string, dbUrl string) (*DataSource, error) {
	return NewDatabaseConnectionWithConnectionPool(dbDriver, dbUrl, 10, 100);
}
