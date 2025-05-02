package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/lib/pq"
	"github.com/segmentio/kafka-go"
	"log"
	"net/http"
	"os"
)

var (
	rdb *redis.Client
	kafkaWriter *kafka.Writer
)

func main() {
	// Initialize Redis
	rdb = redis.NewClient(&redis.Options{
		Addr:     getEnv("REDIS_URL", "localhost:6379"),
		Password: "",
		DB:       0,
	})

	// Initialize Kafka
	kafkaWriter = &kafka.Writer{
		Addr:     kafka.TCP(getEnv("KAFKA_URL", "localhost:9092")),
		Topic:    "orders",
		Balancer: &kafka.LeastBytes{},
	}

	// Create Gin router
	r := gin.Default()

	// Routes
	r.GET("/health", healthCheck)
	r.POST("/orders", createOrder)

	// Start server
	port := getEnv("PORT", "8080")
	log.Printf("Server running on port %s", port)
	log.Fatal(r.Run(":" + port))
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}

func createOrder(c *gin.Context) {
	// Order creation logic will be implemented here
	c.JSON(http.StatusOK, gin.H{"message": "Order created"})
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
