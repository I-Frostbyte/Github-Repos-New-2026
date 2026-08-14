package main

import (
	"log"
	"net/http"

	http_methods "github.com/I-Frostbyte/Github-Repos-New-2026/Tutorials/4-Simple-Gin-Server/1-Using-HTTP-Method"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		// http.StatusOK sends back a positive status.
		c.JSON(http.StatusOK, gin.H{ 
			"message": "pong",
		})
	})
	// router.Run() // listens on 0.0.0.0:8080 by default

	newRouter := http_methods.HTTPMethods()

	// newRouter.Run(":5050")
	// router.Run(":3000") for a hard coded port

	// Trying to run both servers using go routines.
	var g errgroup.Group

	// Define Server 1
	server1 := &http.Server{
		Addr: ":8080",
		Handler: router,
	}

	// Define Server 2
	server2 := &http.Server{
		Addr: ":5050",
		Handler: newRouter,
	}

	// Launch Server 1
	g.Go(func() error {
		log.Println("Server 1 listening on :8080")
		return server1.ListenAndServe()
	})

	// Launch Server 2
	g.Go(func() error {
		log.Println("Server 2 listening on :5050")
		return server2.ListenAndServe()
	})

	// Wait for both blocks to complete or for one to return an error
	if err := g.Wait(); err != nil {
		log.Fatalf("Server cluster execution stopped with error: %v", err)
	}
}
