package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gophersiesta/gophersiesta/server/handlers"
	"github.com/gophersiesta/gophersiesta/server/storage"
)

var db storage.Storage

// Server is responsible for holding the config storage and the http engine to route the request
type Server struct {
	Storage storage.Storage
	*gin.Engine
}

// StartServer creates the storage and configures the routes
func StartServer() *Server {

	db = &storage.BoltDb{} // RAM

	db.Init()

	storage.CreateSampleData(db)

	router := gin.Default()

	server := &Server{db, router}

	server.GET("/", handlers.GetHome)

	// This handler will match /conf/appname but will not match neither /conf/ or /conf
	server.GET("/conf/:appname", handlers.GetConfig)

	// Return list of placeholders
	server.GET("/conf/:appname/placeholders", handlers.GetPlaceHolders)

	// Return list of set values
	server.GET("/conf/:appname/values", handlers.GetValues(db))

	// Set values
	server.POST("/conf/:appname/values", handlers.SetValues(db))

	// Return the rendered template
	server.GET("/conf/:appname/render/:format", handlers.ReplacePlaceholders(db))

	// Return list of set labels
	server.GET("/conf/:appname/labels", handlers.GetLabels(db))

	// Return list of apps on the server
	server.GET("/apps", handlers.GetApps(db))

	server.Run(getPort())

	return server
}

func getPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4747"
		log.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}

	return ":" + port
}
