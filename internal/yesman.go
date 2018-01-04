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
	log.Info("Stop yesman instance")
	return nil
}

// Delete this fake haproxy
func (hap YesmanHaproxy) Delete() error {
	log.Info("Delete yesman instance")
	return nil
}

// Fake return true
func (hap YesmanHaproxy) Fake() bool {
	return true
}

// RestartKilledHaproxy restart a fake yesman haproxy
func (hap YesmanHaproxy) RestartKilledHaproxy() error {
	log.Info("Yes sir!")
	return nil
}
