package fd_event_listener

import (
	"fmt"
	db "foodefi-go/sqlc"
	"foodefi-go/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type submitEventRequest struct {
	BlockchainId int64  `json:"blockchain_id" binding:"required"`
	BlockNumber  int64  `json:"block_number" binding:"required"`
	EventName    string `json:"event_name" binding:"required"`
	Fields       []struct {
		Type  string      `json:"type" binding:"required"`
		Name  string      `json:"name" binding:"required"`
		Value interface{} `json:"value" binding:"required"`
	} `json:"fields" binding:"required"`
}

type submitEventResponse struct {
	Ok string `json:"ok"`
}

func (server *Server) submitEvent(ctx *gin.Context) {
	var request submitEventRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		fmt.Println("validation error")
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// TODO: retrieve user's information from principal

	found, err := server.checkValidBlockchain(ctx, request.BlockchainId)
	if err != nil {
		fmt.Println("could not validate blockchain")
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if found == false {
		ctx.JSON(http.StatusBadRequest, errorResponse(&util.InvalidBlockchainId{}))
		return
	}

	eventFieldTxList, err := server.parseEventFieldTxFromSubmitEventRequest(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	txResult, err := server.store.SubmitEventTx(ctx, db.SubmitEventsTxParams{
		BlockchainId: request.BlockchainId,
		BlockNumber:  request.BlockNumber,
		EventName:    request.EventName,
		Recorder:     "shuhan_scraper", //TODO: implement jwt token
		Fields:       eventFieldTxList,
	})
	if err != nil {
		fmt.Println("could not commit transaction")
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println(txResult)
	ctx.JSON(http.StatusCreated, submitEventResponse{
		Ok: "Okay",
	})

}

func (server *Server) parseEventFieldTxFromSubmitEventRequest(request submitEventRequest) ([]db.EventFieldTx, error) {
	var eventFieldTxList []db.EventFieldTx
	for _, field := range request.Fields {
		value, err := util.ConvertFieldValueInterfaceToString(field.Value, field.Type)
		if err != nil {
			return eventFieldTxList, err
		}
		eventFieldTxList = append(eventFieldTxList, db.EventFieldTx{
			Type:  field.Type,
			Name:  field.Name,
			Value: value,
		})
	}

	return eventFieldTxList, nil
}
