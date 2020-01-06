package rabbitmq

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type ServerConnection struct {
	connection *amqp.Connection
	config     *Config
}

func (connection *ServerConnection) Close() {
	connection.connection.Close()
}

func NewServerConnection(
	config *Config,
	host string,
	username string,
	password string,
	port int,
) (*ServerConnection, error) {
	censoredRabbitDsn := formatRabbitMqDsn(
		host,
		username,
		"********",
		port,
	)

	log.Infof("Connecting to RabbitMQ server %s", censoredRabbitDsn)

	rabbitDsn := formatRabbitMqDsn(
		host,
		username,
		password,
		port,
	)

	connection, err := amqp.Dial(rabbitDsn)
	if err != nil {
		return nil, err
	}

	return &ServerConnection{
		config:     config,
		connection: connection,
	}, nil
}

func formatRabbitMqDsn(
	hostname string,
	username string,
	password string,
	port int,
) string {
	return fmt.Sprintf(
		"amqp://%s:%s@%s:%d/",
		username,
		password,
		hostname,
		port,
	)
}
