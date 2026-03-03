package main

import (
	"database/database"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func validate_payload(){

}

func assign_id(){
	id := uuid.New()
	fmt.Print(id)
}

func store_job(){

}

func track_status() {

}
func main() {
	db, err = database.NewConnection()
	r := gin.Default()
	r.POST("/jobs", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message":"pong",
			"id": id,
		})
	})

	r.GET("/jobs/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message":"pong",
			"id": id,
		})
	})

  	r.GET("/metrics", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":"pong",
		})
	})

	r.DELETE("/jobs/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message":"pong",
			"id": id,
		})
	})
	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
  	}


}
