package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

// HapInstallation contains path to haproxy binary
type HapInstallation struct {
	Path string
}

// Config contains all sidekick configuration
type Config struct {
	LookupdAddresses []string                   `toml:"LookupdAddresses"`
	ProducerAddr     string                     `toml:"ProducerAddr"`
	ProducerRestAddr string                     `toml:"ProducerRestAddr"`
	ClusterID        string                     `toml:"ClusterId"`
	Port             int32                      `toml:"Port"`
	HapHome          string                     `toml:"HapHome"`
	ID               string                     `toml:"Id"`
	Hap              map[string]HapInstallation `toml:"Hap"`
	Sudo             bool                       `toml:"Sudo"`
}

// IsMaster answers true if vip points to current sidekick ip
func (config Config) IsMaster(vip string) (bool, error) {
	resp, err := http.Get(fmt.Sprintf("http://%s:%d/id", vip, config.Port))
	if err != nil {
		log.WithField("vip", vip).WithField("port", config.Port).WithError(err).Error("can't request in http the vip")
		return false, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	isMaster := string(body) == config.ID
	if log.GetLevel() == log.DebugLevel {
		log.WithField("id", string(body)).WithField("isMaster", isMaster).Debug("check if this instance is master")
	}
	return string(body) == config.ID, err
}

// DefaultConfig builds a default config
func DefaultConfig() *Config {
	return &Config{
		Port:      5000,
		HapHome:   "/HOME/hapadm",
		ClusterID: "default-name",
		Sudo:      true,
	}
}

// Header of nsq message
type Header struct {
	CorrelationID string `json:"correlationId"`
	Application   string `json:"application"`
	Platform      string `json:"platform"`
	Timestamp     int64  `json:"timestamp"`
	Source        string `json:"source"`
}

// Conf providing in nsq message from admin
type Conf struct {
	Haproxy []byte `json:"haproxy"`
	Syslog  []byte `json:"syslog"`
	Bind    string `json:"bind"`
	Version string `json:"haproxyVersion,omitempty"`
}

// EventMessage is main type for messages
type EventMessage struct {
	Header Header `json:"header"`
}

// EventMessageWithConf type of messages with Conf type additionally
type EventMessageWithConf struct {
	EventMessage
	Conf Conf `json:"conf,omitempty"`
}

// CloneWithConf clones an EventMessage with he header, a new source and a new timestamp and the conf
func (eventMessage EventMessageWithConf) CloneWithConf(source string) EventMessageWithConf {
	newMessage := EventMessageWithConf{
		Conf: eventMessage.Conf,
	}
	newMessage.Header = eventMessage.Header
	newMessage.Header.Source = source
	newMessage.Header.Timestamp = time.Now().UnixNano() / int64(time.Millisecond)
	return newMessage
}

// Context is a method for retrieving Context from an EventMessage
func (eventMessage EventMessage) Context() Context {
	return Context{
		CorrelationID: eventMessage.Header.CorrelationID,
		Timestamp:     eventMessage.Header.Timestamp,
		Application:   eventMessage.Header.Application,
		Platform:      eventMessage.Header.Platform,
	}
}

// Clone is method for cloning an EventMessage with just the header, a new source and a new timestamp
func (eventMessage EventMessage) Clone(source string) EventMessage {
	newMessage := EventMessage{
		Header: eventMessage.Header,
	}
	newMessage.Header.Source = source
	newMessage.Header.Timestamp = time.Now().UnixNano() / int64(time.Millisecond)
	return newMessage
}

// Context is for tracing current process and local processing
type Context struct {
	CorrelationID string `json:"correlationId"`
	Timestamp     int64  `json:"timestamp"`
	Application   string `json:"application"`
	Platform      string `json:"platform"`
}

// UpdateTimestamp update the timestamp of the current context
func (ctx Context) UpdateTimestamp() Context {
	ctx.Timestamp = time.Now().UnixNano() / int64(time.Millisecond)
	return ctx
}

// Fields for logging with context headers
func (ctx Context) Fields(fields log.Fields) *log.Entry {
	fields["correlationId"] = ctx.CorrelationID
	fields["timestamp"] = ctx.UpdateTimestamp().Timestamp
	fields["application"] = ctx.Application
	fields["platform"] = ctx.Platform

	return log.WithFields(fields)
}

// ReloadEvent is focus on reload event
type ReloadEvent struct {
	Message *EventMessageWithConf
	F       func(data *EventMessageWithConf) error
}

// Execute execute ReloadEvent
func (re *ReloadEvent) Execute() error {
	return re.F(re.Message)
}

// EventHandler handler on EventMessageWithConf
type EventHandler interface {
	HandleMessage(data *EventMessageWithConf) error
}

// HandlerFunc type
type HandlerFunc func(data *EventMessageWithConf) error

// HandleMessage handles EventMessageWithConf
func (h HandlerFunc) HandleMessage(m *EventMessageWithConf) error {
	return h(m)
}

// SdkLogger logger type
type SdkLogger struct {
	Logrus *log.Logger
}

// Output just calls a log
func (sdkLogger SdkLogger) Output(calldepth int, s string) error {
	log.WithField("type", "nsq driver").Info(s)
	return nil
}
