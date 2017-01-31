package internal

import log "github.com/Sirupsen/logrus"

// YesmanHaproxy is an haproxy fake process which always return yes
type YesmanHaproxy struct {
	role    string
	context Context
	fake    bool
}

// ApplyConfiguration always successes
func (hap YesmanHaproxy) ApplyConfiguration(data *EventMessageWithConf) (int, error) {
	return SUCCESS, nil
}

// Stop this fake haproxy
func (hap YesmanHaproxy) Stop() error {
	hap.context.Fields(log.Fields{}).Info("Stop yesman instance")
	return nil
}

// Delete this fake haproxy
func (hap YesmanHaproxy) Delete() error {
	hap.context.Fields(log.Fields{}).Info("Delete yesman instance")
	return nil
}

// Fake return true
func (hap YesmanHaproxy) Fake() bool {
	return true
}
