package fd_event_listener

import (
	db "foodefi-go/sqlc"
	"github.com/gin-gonic/gin"
)

const PREFIX_PATH = "/event"

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST(PREFIX_PATH+"/submit", server.submitEvent)
	router.POST(PREFIX_PATH+"/update", server.updateEvent)
	router.POST(PREFIX_PATH+"/delete", server.deleteEvent)
	router.POST(PREFIX_PATH+"/list", server.listEvent)

	server.router = router

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
