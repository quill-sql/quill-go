package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/quill-sql/quill-go/pkg/core"
	"github.com/quill-sql/quill-go/pkg/models"
)

func main() {
	// Set Gin mode to release in prod
	// gin.SetMode(gin.ReleaseMode)

	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DB_URL")
	var dbURLPtr *string
	if dbURL != "" {
		dbURLPtr = &dbURL
	}

	quillParams := models.QuillClientParams{
		PrivateKey:               os.Getenv("PRIVATE_KEY"),
		DatabaseConnectionString: dbURLPtr,
		DatabaseConfig:           map[string]string{}, // TODO: Add database configuration
		DatabaseType:             os.Getenv("DB_TYPE"),
	}
	quill, err := core.NewQuill(quillParams)
	if err != nil {
		log.Fatalf("Error creating Quill client: %v", err)
	}

	router := gin.Default()

	// CORS configuration
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	// Health check endpoint
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	})

	// Quill POST endpoint
	router.POST("/quill", func(c *gin.Context) {
		var body models.QuillRequest
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		orgID := body.OrgID
		if body.Metadata.OrgID != nil && orgID == nil {
			id := body.Metadata.OrgID.(string)
			orgID = &id

		}
		if orgID == nil {
			orgID = new(string)
			*orgID = "*"
		}
		params := models.QuillQueryParams{
			OrgId:    *orgID,
			Metadata: body.Metadata,
		}

		result, err := quill.Query(params)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, result)
	})

	router.POST("/quill-filtered", func(c *gin.Context) {
		var body models.QuillRequest
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		orgID := body.OrgID
		if body.Metadata.OrgID != nil && orgID == nil {
			id := body.Metadata.OrgID.(string)
			orgID = &id

		}
		if orgID == nil {
			orgID = new(string)
			*orgID = "*"
		}
		params := models.QuillQueryParams{
			OrgId:    *orgID,
			Metadata: body.Metadata,
			Filters: &[]models.Filter{
				{
					FilterType: models.STRING_FILTER,
					Operator:   models.STRING_IS_EXACTLY,
					Value:      "unpaid",
					Field:      "status",
					Table:      "subscriptions",
				},
			},
		}

		result, err := quill.Query(params)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, result)
	})

	router.Run(":3003")
}
