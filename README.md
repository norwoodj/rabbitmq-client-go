rabbitmq-client-go
==================
This codebase houses a [golang](https://golang.org) client library for interacting with [rabbitmq](https://www.rabbitmq.com).

Built on top of [amqp091-go](https://github.com/rabbitmq/amqp091-go), it provides higher-level patterns for
implementing consumers and producers of messages.

## Features
* Opinionated auto-declaration of routing topology
* Auto-publish to dead-letter queue on error
* Configurable message serialization strategy

## Disclaimer
This is a very new project, written by someone who does not purport to be an expert about rabbitmq. As such
it is entirely possible that:

* There are bugs in my code
* The routing topology that consumers/producers auto-declare is not in accordance with best practices
* Critical areas are lacking necessary configurability

This project was mainly written in order to provide a reusable higher-level rabbitmq interface for my
other project [hashbash](https://github.com/norwoodj/hashbash-backend-go). Please refer to that project
for example usage of this library and if something is not to your liking, please feel free to contribute.
