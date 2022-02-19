package fd_auth

import (
	"database/sql"
	"foodefi-go/util"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error

	globalConfig, err := util.LoadGlobalConfig("../")
	if err != nil {
		log.Fatal("can not load global config", err)
	}

	testDb, err = sql.Open(globalConfig.DBDriver, globalConfig.DBSource)
	if err != nil {
		log.Fatal("can not connect to Db", err)
	}

	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
