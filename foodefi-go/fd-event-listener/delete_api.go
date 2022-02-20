package fd_event_listener

import (
	"database/sql"
	"fmt"
	db "foodefi-go/sqlc"
	"foodefi-go/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type deleteEventRequest struct {
	BlockchainId int64  `json:"blockchain_id" binding:"required"`
	BlockNumber  int64  `json:"block_number" binding:"required"`
	EventName    string `json:"event_name" binding:"required"`
}

type deleteEventResponse struct {
	Ok string `json:"ok"`
}

func (server *Server) delete(ctx *gin.Context) {
	var request deleteEventRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		fmt.Println("validation error")
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// TODO: retrieve user's information from principal

	//retrieve event
	event, err := server.store.Queries.GetEventByAllData(ctx, db.GetEventByAllDataParams{
		BlockchainID: request.BlockchainId,
		BlockNumber:  request.BlockNumber,
		EventName:    request.EventName,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("blockchain event not found")
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, errorResponse(&util.InvalidEvent{}))
			return
		}
		fmt.Println("could not get event from database")
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	//delete event fields
	err = server.store.Queries.DeleteEventFieldByEventId(ctx, event.ID)
	if err != nil {
		fmt.Println("could not delete event field from database")
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	//delete event
	err = server.store.Queries.DeleteEvent(ctx, event.ID)
	if err != nil {
		fmt.Println("could not delete event from database")
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	fmt.Println("event deleted successfully")
	ctx.JSON(http.StatusOK, deleteEventResponse{
		Ok: "okay",
	})
	return

}
