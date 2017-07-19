# CHANGELOG

## v0.3.1

* **improvement**: log rotate for sidekick.log
* **feature**: remove syslog configuration file creations

## v0.3.0

* **feature**: restart consumer via api (_/consumers/restart_)
* **improvement**: add a bound at 100 for committed/deleted requested for protecting itself from nsq messages burst
* **restart**: restart killed haprox process on master too
* **feature**: at start waiting for nsqd process for producer
