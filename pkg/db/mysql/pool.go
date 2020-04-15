package mysql

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" //nolint
	"github.com/jinzhu/gorm"
	"k8s.io/klog"

	"github.com/q8s-io/mcp/pkg/config"
)

var dbPool *gorm.DB

func GetDB() *gorm.DB {
	return dbPool
}

func InitDB(config *config.MysqlConfig) error {
	maxTimes := 5
	// in order to ignore difference of time zone, use UTC as mysql server time_zone
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&time_zone='UTC'",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	// connect MySQL, with max 5 times
try:
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		maxTimes--
		if maxTimes <= 0 {
			return err
		}
		if maxTimes > 0 {
			klog.Infof("error to connect to mysql server, sleep 5 sec and retry, remain retry times: %d, %v", maxTimes, err)
			time.Sleep(5 * time.Second)
			goto try
		}
	}

	db.SingularTable(true)

	db.DB().SetMaxOpenConns(config.MaxOpenConns)
	db.DB().SetMaxIdleConns(config.MaxIdleConns)
	db.DB().SetConnMaxLifetime(time.Duration(config.ConnMaxTimeInSec) * time.Second)

	if config.Debug {
		db.LogMode(true)
	}

	dbPool = db
	return nil
}

func CloseDB() {
	if dbPool != nil {
		dbPool.Close()
	}
}
