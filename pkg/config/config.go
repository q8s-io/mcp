package config

import (
	"io/ioutil"
	"os"
	"strconv"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v2"
	"k8s.io/klog"

	"github.com/q8s-io/mcp/pkg/validator"
)

type Config struct {
	MysqlConfig MysqlConfig `yaml:"mysql"`
}

type MysqlConfig struct {
	Host     string `yaml:"host" validate:"required"`
	Port     int    `yaml:"port" validate:"required"`
	Database string `yaml:"database" validate:"required"`
	Username string `yaml:"username" validate:"required"`
	Password string `yaml:"password" validate:"required"`

	Debug bool `yaml:"debug" default:"false"`

	MaxOpenConns     int `yaml:"max_open_conns" default:"10"`
	MaxIdleConns     int `yaml:"max_idle_conns" default:"1"`
	ConnMaxTimeInSec int `yaml:"conn_max_time_in_sec" default:"30"`
}

// LoadConfig loads specified config file, generate to Config struct.
func LoadConfig(configFile string) (*Config, error) {
	var config Config

	if configFile == "" {
		klog.Info("set config from env")
		if err := setConfigFromEnv(&config); err != nil {
			return nil, err
		}
	} else {
		klog.Infof("set config from config file: %s", configFile)
		if err := setConfigFromFile(configFile, &config); err != nil {
			return nil, err
		}
	}

	// set default value
	if err := defaults.Set(&config); err != nil {
		return nil, err
	}

	// validate value
	if err := validator.GetValidate().Struct(&config); err != nil {
		return nil, err
	}

	klog.Infof("config loaded: %+v", config)
	return &config, nil
}

func setConfigFromEnv(config *Config) error {
	port, err := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	if err != nil {
		return err
	}

	isDebug, err := strconv.ParseBool(os.Getenv("MYSQL_DEBUG"))
	if err != nil {
		return err
	}

	config.MysqlConfig.Host = os.Getenv("MYSQL_HOST")
	config.MysqlConfig.Port = port
	config.MysqlConfig.Username = os.Getenv("MYSQL_USER")
	config.MysqlConfig.Password = os.Getenv("MYSQL_PASSWORD")
	config.MysqlConfig.Database = os.Getenv("MYSQL_DATABASE")
	config.MysqlConfig.Debug = isDebug

	return nil
}

func setConfigFromFile(configFile string, config *Config) error {
	bytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(bytes, &config); err != nil {
		return err
	}

	return nil
}
