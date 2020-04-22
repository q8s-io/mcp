package persistence

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" //nolint
	"github.com/jinzhu/gorm"
	"k8s.io/klog"

	"github.com/q8s-io/mcp/pkg/config"
	"github.com/q8s-io/mcp/pkg/domain/repository"
)

var repositories *Repositories

type Repositories struct {
	DB *gorm.DB

	ClusterRepo    repository.ClusterRepository
	KubeconfigRepo repository.KubeconfigRepository
}

func GetRepositories() *Repositories {
	return repositories
}

func NewRepositories(db *gorm.DB) {
	repositories = &Repositories{
		DB: db,

		ClusterRepo:    newClusterRepository(db),
		KubeconfigRepo: newKubeconfigRepository(db),
	}
}

func InitDB(config *config.MysqlConfig) (*Repositories, bool) {
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
			return nil, false
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

	NewRepositories(db)
	return repositories, true
}

func CloseDB() {
	if repositories != nil {
		repositories.DB.Close()
	}
}
