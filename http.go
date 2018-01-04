/*
 *  Copyright (C) 2016 VSCT
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */
package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"net"
	"net/http"
	"strconv"
)

type RestApi struct {
	properties *Config
	consumers  Consumers
	listener   *StoppableListener
}

func NewRestApi(properties *Config, consumers Consumers) *RestApi {
	api := &RestApi{
		properties: properties,
		consumers:  consumers,
	}
	return api
}

func (api *RestApi) Start() error {
	sm := http.NewServeMux()
	sm.HandleFunc("/id", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "%s", api.properties.ID)
	})
	sm.HandleFunc("/ismaster", func(writer http.ResponseWriter, request *http.Request) {
		isMaster, err := api.properties.IsMaster(request.URL.Query().Get("vip"))
		if err == nil {
			fmt.Fprintf(writer, "%s", strconv.FormatBool(isMaster))
		} else {
			fmt.Fprint(writer, "can't know if there is master")
		}
	})

	sm.HandleFunc("/consumers/restart", func(writer http.ResponseWriter, request *http.Request) {
		api.consumers.RestartConsumers()
		fmt.Fprintln(writer, "consumer restarted")

	})

	listener, err := net.Listen("tcp4", fmt.Sprintf(":%d", api.properties.Port))
	if err != nil {
		log.WithError(err).Fatal(err)
	}
	api.listener, err = NewListener(listener)
	if err != nil {
		return err
	}

	log.WithField("port", api.properties.Port).Info("start http server")
	http.Serve(api.listener, sm)

	return nil
}

func (api *RestApi) Stop() {
	if api.listener != nil {
		api.listener.Stop()
	}
}
