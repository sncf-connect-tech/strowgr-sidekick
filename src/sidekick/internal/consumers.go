package internal

import (
	"github.com/bitly/go-nsq"
	log "github.com/Sirupsen/logrus"
	"time"
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

// create nsq consumer for given topic and with given handler to consumers and store in this consumer store.
// if topic has already a consumer, it will be overriden by the new one.
func (consumers Consumers) PutAndStartConsumer(topic string, handler func(message *nsq.Message) error) {
	consumers.consumers[topic] = Consumer{
		topic:  topic,
		handler:handler,
		nsq:    consumers.createConsumer(topic, handler),
	}
}

// stop all consumers and restart theim
func (consumers Consumers) RestartConsumers() {
	log.Info("restart all nsq consumers")
	for _, consumer := range consumers.consumers {
		log.WithField("topic", consumer.topic).Info("stop nsq consumer")
		consumer.nsq.Stop()
		// waiting for NSQ stop with 5 seconds of timeout
		timeout := time.Tick(5 * time.Second)
		stop := false
		for !stop {
			select {
			case <-consumer.nsq.StopChan:
			case <-timeout:
				log.WithField("topic", consumer.topic).Warn("timeout on close consumer")
				stop = true
				break
			default:
				time.Sleep(100 * time.Millisecond)
			}
		}
		// start a new consumer with the same topic and the same handler
		consumers.PutAndStartConsumer(consumer.topic, consumer.handler)
		log.WithField("topic", consumer.topic).Info("consumer on this topic has been restarted")
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
		log.WithField("topic", topic).Info("create consumer")
		return consumer
	} else {
		// error is probably from configuration
		panic(err)
	}
}
