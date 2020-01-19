package rabbitmq

type Config struct {
	// Connection fields
	Hostname    string
	Port        int
	Username    string
	Password    string
	VirtualHost string

	// Routing Configuration
	DeadLetterExchangeName string
	DeadLetterQueueSuffix  string

	// Naming Configuration
	QueueNamingStrategy QueueNamingStrategy
}

func NewConfig(
	hostname string,
	port int,
	username string,
	password string,
	virtualHost string,
) *Config {
	return &Config{
		Hostname:               hostname,
		Port:                   port,
		Username:               username,
		Password:               password,
		VirtualHost:            virtualHost,
		DeadLetterExchangeName: DefaultDeadLetterExchangeName,
		DeadLetterQueueSuffix:  DefaultDeadLetterQueueSuffix,
		QueueNamingStrategy:    DefaultQueueNamingStrategy{},
	}
}
