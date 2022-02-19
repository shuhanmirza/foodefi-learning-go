package fd_event_listener

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (server *Server) checkValidBlockchain(ctx *gin.Context, blockchainId int64) (found bool, err error) {
	_, err = server.store.Queries.GetBlockchain(ctx, blockchainId)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("blockchain id not found")
			fmt.Println(err)
			return false, nil
		}
		return false, err
	}
	return true, nil
}
