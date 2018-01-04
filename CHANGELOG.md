# CHANGELOG

<<<<<<< HEAD
=======
## v0.3.4

* **ci**: improve docker image for including mono mode. Now only one instance is needed and to validate.

>>>>>>> ee7324d6056ced5646654d380f93c3b647fcd543
## v0.3.3

* **fix**: sudo command
* **change**: sudo by default


## v0.3.2

* **feature**: add a sudo option for running haproxy process


## v0.3.1

* **improvement**: log rotate for sidekick.log
* **feature**: remove syslog configuration file creations

## v0.3.0

* **feature**: restart consumer via api (_/consumers/restart_)
* **improvement**: add a bound at 100 for committed/deleted requested for protecting itself from nsq messages burst
* **restart**: restart killed haprox process on master too
* **feature**: at start waiting for nsqd process for producer
