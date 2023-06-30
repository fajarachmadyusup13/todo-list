package repository

import (
	"log"
	"os"
	"strconv"

	"github.com/DATA-DOG/go-sqlmock"
	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/fajarachmadyusup13/todo-list/internal/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initializeTest() {
	config.GetConf()
	setupLogger()
}

func setupLogger() {
	formatter := runtime.Formatter{
		ChildFormatter: &logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		},
		Line: true,
		File: true,
	}

	logrus.SetFormatter(&formatter)
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.WarnLevel)

	verbose, _ := strconv.ParseBool(os.Getenv("VERBOSE"))
	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func initializeCockroachMockConn() (db *gorm.DB, mock sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}

	db, err = gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return
}
