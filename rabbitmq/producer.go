package rabbitmq

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	BaseClient
	messageSerializer MessageSerializer
}

func NewProducer(
	connection *ServerConnection,
	serializer MessageSerializer,
	exchangeName string,
	exchangeType string,
	routingKey string,
	declareRoutingTopology bool,
) (*Producer, error) {
	baseClient, err := newBaseRabbitmqClient(
		connection,
		exchangeName,
		exchangeType,
		routingKey,
		declareRoutingTopology,
	)

	if err != nil {
		return nil, err
	}

	return &Producer{
		BaseClient:        baseClient,
		messageSerializer: serializer,
	}, nil
}

func (producer *Producer) PublishMessage(msg interface{}) error {
	serializedMsg, err := producer.messageSerializer.SerializeMessage(msg)
	if err != nil {
		return fmt.Errorf("failed to serialize message for publishing: %s", err)
	}

	if err := producer.channel.Publish(
		producer.exchangeName,
		producer.routingKey,
		false,
		false,
		amqp.Publishing{ContentType: producer.messageSerializer.GetContentType(), Body: serializedMsg},
	); err != nil {
		if err := producer.createNewChannel(); err != nil {
			return fmt.Errorf("failed to publish and failed to re-establish channel: %s", err)
		}

		return err
	}

	return nil
}
