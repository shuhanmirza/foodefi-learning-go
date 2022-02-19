package fd_event_listener

import (
	"database/sql"
	"fmt"
	db "foodefi-go/sqlc"
	"foodefi-go/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type updateEventRequest struct {
	BlockchainId int64  `json:"blockchain_id" binding:"required"`
	BlockNumber  int64  `json:"block_number" binding:"required"`
	EventName    string `json:"event_name" binding:"required"`
	Field        struct {
		Type  string      `json:"type" binding:"required"`
		Name  string      `json:"name" binding:"required"`
		Value interface{} `json:"value" binding:"required"`
	} `json:"field" binding:"required"`
}

type updateEventResponse struct {
	Ok string `json:"ok"`
}

func (server *Server) update(ctx *gin.Context) {
	var request updateEventRequest

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

	//retrieve event field
	eventField, err := server.store.Queries.GetEventFieldByEventIdFieldName(ctx, db.GetEventFieldByEventIdFieldNameParams{
		EventID: event.ID,
		Name:    request.Field.Name,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("event field not found")
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, errorResponse(&util.InvalidEvent{}))
			return
		}
		fmt.Println("could not get event field from database")
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if request.Field.Type == eventField.Type {
		//process value
		value, err := util.ConvertFieldValueInterfaceToString(request.Field.Value, eventField.Type)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		// update event field
		err = server.store.Queries.UpdateEventField(ctx, db.UpdateEventFieldParams{
			EventID: event.ID,
			Name:    eventField.Name,
			Value:   value,
		})
		if err != nil {
			fmt.Println("could not update event field")
			fmt.Println(err)
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		fmt.Println("event updated successfully")
		ctx.JSON(http.StatusOK, updateEventResponse{
			Ok: "okay",
		})
		return
	} else {
		fmt.Println("event field type mismatch")
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(&util.InvalidEventFieldType{}))
		return
	}
}
