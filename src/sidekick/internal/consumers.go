package internal

import (
	"github.com/bitly/go-nsq"
)

type Consumer struct {
	topic   string
	handler func(message *nsq.Message) error
	nsq     *nsq.Consumer
}

type Consumers struct {
	consumers map[string]Consumer
	nsqConfig nsq.Config
	sdkLogger SdkLogger
	config    Config
}

func NewConsumers(config Config, nsqConfig nsq.Config, sdkLogger SdkLogger) Consumers {
	return Consumers{
		consumers:    make(map[string]Consumer),
		config:       config,
		nsqConfig:    nsqConfig,
		sdkLogger:    sdkLogger,
	}
}

// create nsq consumer for given topic and with given handler to consumers and store in this consumer store
func (consumers Consumers) AddAndStartConsumer(topic string, handler func(message *nsq.Message) error) {
	consumers.consumers[topic] = Consumer{
		topic:  topic,
		handler:handler,
		nsq:    consumers.createConsumer(topic, handler),
	}
}

// stop all consumers and restart theim
func (consumers Consumers) RestartConsumers() {
	for _, consumer := range consumers.consumers {
		consumer.nsq.Stop()
		<-consumer.nsq.StopChan
		consumer.nsq = consumers.createConsumer(consumer.topic, consumer.handler)
	}
}

// create a consumer with a given topic and handler
func (consumers Consumers) createConsumer(topic string, handler func(message *nsq.Message) error) *nsq.Consumer {
	if consumer, err := nsq.NewConsumer(topic, consumers.config.Id, &consumers.nsqConfig); err == nil {
		consumer.SetLogger(consumers.sdkLogger, nsq.LogLevelWarning)

		consumer.AddConcurrentHandlers(nsq.HandlerFunc(handler), 6)
		if err = consumer.ConnectToNSQLookupds(consumers.config.LookupdAddresses); err != nil {
			panic(err)
		}
		return consumer
	} else {
		// error is probably from configuration
		panic(err)
	}
}
