package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

func GetConf() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("./..")
	viper.AddConfigPath("./../..")
	viper.AddConfigPath("./../../..")
	viper.SetConfigName("config")
	viper.SetEnvPrefix("svc")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Warnf("%v", err)
	}
}

func CockroachHost() string {
	return viper.GetString("cockroach.host")
}

func CockroachDatabase() string {
	return ""
}

func CockroachUsername() string {
	return ""
}

func CockroachPassword() string {
	return ""
}

func CockroachSSLMode() string {
	return ""
}

func CockroachMaxIdleConns() int {
	return 0
}

func CockroachConnMaxLifetime() time.Duration {
	return 1 * time.Minute
}

func CockroachMaxOpenConns() int {
	return 1
}

func CockroachPingInterval() time.Duration {
	return 1 * time.Second
}

func DatabaseDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		CockroachUsername(),
		CockroachPassword(),
		CockroachHost(),
		CockroachDatabase(),
		CockroachSSLMode())
}
