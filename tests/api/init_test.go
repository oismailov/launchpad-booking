package api_test

import (
	"github.com/gin-gonic/gin"
	"github.com/oismailov/launchpad-booking/config"
	"github.com/oismailov/launchpad-booking/persistance"
	"github.com/oismailov/launchpad-booking/pkg/migration"
	r "github.com/oismailov/launchpad-booking/router"
	"gopkg.in/testfixtures.v2"
	"log"
)

type Response struct {
	Message string `json:"message"`
}

var (
	fixtures *testfixtures.Context
	err      error
	router   *gin.Engine
)

func init() {
	config.LoadTestConfig()
	migration.Migrate()

	router = r.GetRouter()
	db, err := persistance.GetInstance().DB()
	if err != nil {
		log.Fatal(err)
	}

	fixtures, err = testfixtures.NewFolder(db, &testfixtures.PostgreSQL{}, "../fixtures")

	if err != nil {
		log.Fatal(err)
	}

	if err := fixtures.Load(); err != nil {
		log.Fatal(err)
	}
}

func RestoreTestDatabaseState() {
	if err := fixtures.Load(); err != nil {
		log.Fatal(err)
	}
}
