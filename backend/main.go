package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gopalabhamidipati/backend/database"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Job struct {
	ID uuid.UUID `json:"id"`
	Status string `json:"status"`
    Payload sql.NullString `json:"payload"`
    Retries int `json:"retires"`
	CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type CreateJobRequest struct {
	Status string `json:"status" binding:"required"`
	Payload any    `json:"payload"`
}

type UpdateJobRequest struct {
	Status  *string `json:"status"`
	Payload any     `json:"payload"`
	Retries *int    `json:"retries"`
}

func create_job(db *sql.DB, status string, payload any) (Job, error){
	id := uuid.New()
	payloadJSON, err := jsonMarshal(payload)

	if err != nil {
		return Job{}, err
	}

	q := `INSERT INTO jobs (id, status, payload, retries)
		VALUES ($1, $2, $3::jsonb 0)
		RETURNING id, status, payload::text, retries, created_at, updated_at`

	var j Job
	err = db.QueryRow(q, id, status, payloadJSON).Scan(
		&j.ID, &j.Status, &j.Payload, &j.Retries, &j.CreatedAt, &j.UpdatedAt,
	)
	return j, err
}

func read_job(db *sql.DB, id uuid.UUID) (Job, error){
q := `
		SELECT id, status, payload::text, retries, created_at, updated_at
		FROM jobs
		WHERE id = $1
	`
	var j Job
	err := db.QueryRow(q, id).Scan(
		&j.ID, &j.Status, &j.Payload, &j.Retries, &j.CreatedAt, &j.UpdatedAt,
	)
	return j, err
}

func update_job(db *sql.DB, id uuid.UUID, req UpdateJobRequest) (Job, error){
	payloadJSON := (*string)(nil)
	if req.Payload != nil {
		p, err := jsonMarshal(req.Payload)
		if err != nil {
			return Job{}, err
		}
		payloadJSON = &p
	}

	q := `
		UPDATE jobs
		SET
			status = COALESCE($2, status),
			payload = COALESCE($3::jsonb, payload),
			retries = COALESCE($4, retries),
			updated_at = NOW()
		WHERE id = $1
		RETURNING id, status, payload::text, retries, created_at, updated_at
	`
	var j Job
	err := db.QueryRow(
		q,
		id,
		req.Status,
		payloadJSON,
		req.Retries,
	).Scan(&j.ID, &j.Status, &j.Payload, &j.Retries, &j.CreatedAt, &j.UpdatedAt)

	return j, err


}

func delete_job(db *sql.DB, id uuid.UUID) error {
	res, err := db.Exec(`DELETE FROM jobs WHERE id = $1`, id)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return sql.ErrNoRows
	}
	return nil
}


func jsonMarshal(v any) (string, error) {
	if v == nil {
		return "null", nil
	}
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
func main() {
	db, err := database.NewConnection()
	if err != nil{
		log.Fatalf(("failed to connect to datavase"))
	}
	defer db.Close()

	r := gin.Default()
	
	// Create a job
	r.POST("/jobs", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message":"pong",
			"id": id,
		})
	})

	// Read a job
	r.GET("/jobs/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message":"pong",
			"id": id,
		})
	})

	// Update a job
  	r.PATCH("/jobs/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":"pong",
		})
	})

	// Delete a job
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
