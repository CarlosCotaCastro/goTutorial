package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// CodeExecutionRequest represents a request to execute Go code
type CodeExecutionRequest struct {
	Code string `json:"code"`
}

// CodeExecutionResponse represents the response from code execution
type CodeExecutionResponse struct {
	Output string `json:"output"`
	Error  string `json:"error,omitempty"`
}

// UserProgress represents user's learning progress
type UserProgress struct {
	UserID      string  `json:"user_id"`
	LessonID    int     `json:"lesson_id"`
	Completed   bool    `json:"completed"`
	CompletedAt *string `json:"completed_at,omitempty"`
}

// WebSocket connection manager
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for development
	},
}

func main() {
	// Initialize execution service
	executionService = NewCodeExecutionService()

	// Initialize database
	var err error
	database, err = NewDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	// Initialize Gin router
	r := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"}
	r.Use(cors.New(config))

	// API routes
	api := r.Group("/api")
	{
		// Health check
		api.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "healthy"})
		})

		// Code execution endpoint
		api.POST("/execute", executeCode)

		// Lessons endpoints
		api.GET("/lessons", getLessons)
		api.GET("/lessons/:id", getLesson)

		// Progress endpoints
		api.GET("/progress/:user_id", getUserProgress)
		api.POST("/progress", updateProgress)

		// WebSocket endpoint for real-time features
		api.GET("/ws", handleWebSocket)
	}

	// Serve static files (for production)
	r.Static("/static", "./static")

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Go Tutorial Server starting on port %s", port)
	log.Fatal(r.Run(":" + port))
}

// Global execution service
var executionService *CodeExecutionService

// Global database instance
var database *Database

// executeCode handles Go code execution requests
func executeCode(c *gin.Context) {
	var req CodeExecutionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Execute code using the execution service
	response, err := executionService.ExecuteCodeFallback(req.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// getUserProgress returns user's learning progress
func getUserProgress(c *gin.Context) {
	userID := c.Param("user_id")

	// Get progress from database
	progress, err := database.GetUserProgress(userID)
	if err != nil {
		log.Printf("Error getting user progress: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user progress"})
		return
	}

	// Always return an empty array if no progress found
	if len(progress) == 0 {
		progress = []UserProgress{}
	}

	c.JSON(http.StatusOK, progress)
}

// updateProgress updates user's learning progress
func updateProgress(c *gin.Context) {
	var progress UserProgress
	if err := c.ShouldBindJSON(&progress); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update progress in database
	if err := database.UpdateUserProgress(progress); err != nil {
		log.Printf("Error updating user progress: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update progress"})
		return
	}

	log.Printf("Progress updated for user %s, lesson %d: completed=%v",
		progress.UserID, progress.LessonID, progress.Completed)

	c.JSON(http.StatusOK, gin.H{"message": "Progress updated successfully"})
}

// handleWebSocket handles WebSocket connections for real-time features
func handleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}
	defer conn.Close()

	// Handle WebSocket messages
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("WebSocket read error: %v", err)
			break
		}

		// Echo the message back (for now)
		if err := conn.WriteMessage(messageType, message); err != nil {
			log.Printf("WebSocket write error: %v", err)
			break
		}
	}
}
