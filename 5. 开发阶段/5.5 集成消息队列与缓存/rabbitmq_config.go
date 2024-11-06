package config

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// RabbitMQConfig holds the configuration for RabbitMQ connection.
type RabbitMQConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Queue    string
}

// NewRabbitMQConfig creates a new RabbitMQ configuration instance.
func NewRabbitMQConfig() *RabbitMQConfig {
	return &RabbitMQConfig{
		Host:     "localhost",
		Port:     5672,
		Username: "guest",
		Password: "guest",
		Queue:    "shop-app-queue",
	}
}

// Connect establishes a connection to RabbitMQ.
func (r *RabbitMQConfig) Connect() (*amqp.Connection, error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%d/%s", r.Host, r.Port, r.Queue))
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	return conn, err
}

// Channel creates a new channel on the connection.
func (r *RabbitMQConfig) Channel(conn *amqp.Connection) (*amqp.Channel, error) {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	return ch, err
}