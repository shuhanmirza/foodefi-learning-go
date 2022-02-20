package fd_event_listener

import (
	"fmt"
	db "foodefi-go/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListEvents struct {
	PageId          int32  `json:"page_id" binding:"gt=0"`
	PageSize        int32  `json:"page_size" binding:"gt=0"`
	BlockchainId    int64  `json:"blockchain_id"`
	EventName       string `json:"event_name"`
	FromBlockNumber int64  `json:"from_block_number"`
	ToBlockNumber   int64  `json:"to_block_number"`
}

func (server *Server) listEvent(ctx *gin.Context) {
	var err error
	var request ListEvents

	// TODO: retrieve user's information from principal
	// TODO: restructure to eliminate duplicate code

	if err = ctx.ShouldBindJSON(&request); err == nil {
		if len(request.EventName) > 0 && request.BlockchainId > 0 && request.FromBlockNumber > 0 && request.ToBlockNumber > 0 {
			listEvent, err := server.store.Queries.ListEventsByBlockchainIdEventNameBlockNumberRange(ctx, db.ListEventsByBlockchainIdEventNameBlockNumberRangeParams{
				Limit:        request.PageSize,
				Offset:       request.PageSize * (request.PageId - 1),
				BlockchainID: request.BlockchainId,
				EventName:    request.EventName,
				FromBlock:    request.FromBlockNumber,
				ToBlock:      request.ToBlockNumber,
			})
			if err != nil {
				fmt.Println("could not fetch events list")
				fmt.Println(err)
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			}

			fmt.Printf("%v", listEvent)
			ctx.JSON(http.StatusOK, listEvent)
			return
		} else if request.BlockchainId > 0 && request.FromBlockNumber > 0 && request.ToBlockNumber > 0 {
			listEvent, err := server.store.Queries.ListEventsByBlockchainIdBlockNumberRange(ctx, db.ListEventsByBlockchainIdBlockNumberRangeParams{
				Limit:        request.PageSize,
				Offset:       request.PageSize * (request.PageId - 1),
				BlockchainID: request.BlockchainId,
				FromBlock:    request.FromBlockNumber,
				ToBlock:      request.ToBlockNumber,
			})
			if err != nil {
				fmt.Println("could not fetch events list")
				fmt.Println(err)
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			}

			fmt.Printf("%v", listEvent)
			ctx.JSON(http.StatusOK, listEvent)
			return
		} else if len(request.EventName) > 0 && request.BlockchainId > 0 {
			listEvent, err := server.store.Queries.ListEventsByBlockchainIdEventName(ctx, db.ListEventsByBlockchainIdEventNameParams{
				Limit:        request.PageSize,
				Offset:       request.PageSize * (request.PageId - 1),
				BlockchainID: request.BlockchainId,
				EventName:    request.EventName,
			})
			if err != nil {
				fmt.Println("could not fetch events list")
				fmt.Println(err)
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			}

			fmt.Printf("%v", listEvent)
			ctx.JSON(http.StatusOK, listEvent)
			return
		} else if request.BlockchainId > 0 {
			listEvent, err := server.store.Queries.ListEventsByBlockchainId(ctx, db.ListEventsByBlockchainIdParams{
				Limit:        request.PageSize,
				Offset:       request.PageSize * (request.PageId - 1),
				BlockchainID: request.BlockchainId,
			})
			if err != nil {
				fmt.Println("could not fetch events list")
				fmt.Println(err)
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			}

			fmt.Printf("%v", listEvent)
			ctx.JSON(http.StatusOK, listEvent)
			return
		} else {
			listEvent, err := server.store.Queries.ListEvents(ctx, db.ListEventsParams{
				Limit:  request.PageSize,
				Offset: request.PageSize * (request.PageId - 1),
			})
			if err != nil {
				fmt.Println("could not fetch events list")
				fmt.Println(err)
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			}

			fmt.Printf("%v", listEvent)
			ctx.JSON(http.StatusOK, listEvent)
			return
		}
	} else {
		fmt.Println("validation error")
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
}
