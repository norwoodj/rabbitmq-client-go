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
	port int,
	virtualHost string,
	username string,
	password string,
) (*ServerConnection, error) {
	censoredRabbitDsn := formatRabbitMqDsn(
		host,
		port,
		virtualHost,
		username,
		"********",
	)

	log.Infof("Connecting to RabbitMQ server %s", censoredRabbitDsn)

	rabbitDsn := formatRabbitMqDsn(
		host,
		port,
		virtualHost,
		username,
		password,
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
	port int,
	virtualHost string,
	username string,
	password string,
) string {
	return fmt.Sprintf(
		"amqp://%s:%s@%s:%d/%s",
		username,
		password,
		hostname,
		port,
		virtualHost,
	)
}
