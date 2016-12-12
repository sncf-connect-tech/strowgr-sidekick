[![Build Status](https://travis-ci.org/voyages-sncf-technologies/strowgr-sidekick.svg?branch=0.2.x)](https://travis-ci.org/voyages-sncf-technologies/strowgr-sidekick) [![GitHub release](https://img.shields.io/github/release/voyages-sncf-technologies/strowgr-sidekick.svg)](https://github.com/voyages-sncf-technologies/strowgr-sidekick/releases/latest) [![Coverage Status](https://coveralls.io/repos/github/voyages-sncf-technologies/strowgr-sidekick/badge.svg)](https://coveralls.io/github/voyages-sncf-technologies/strowgr-sidekick) [![codecov](https://codecov.io/gh/voyages-sncf-technologies/strowgr-sidekick/branch/0.2.x/graph/badge.svg)](https://codecov.io/gh/voyages-sncf-technologies/strowgr-sidekick) [![Go Report Card](https://goreportcard.com/badge/github.com/voyages-sncf-technologies/strowgr-sidekick)](https://goreportcard.com/report/github.com/voyages-sncf-technologies/strowgr-sidekick)


# Build

```
$ GOPATH=$PWD go build sidekick
```

# Build with govvv

govvv is a tool for versioning binaries
 
```
$ export GOPATH=$PWD
$ go get github.com/ahmetalpbalkan/govvv
$ govvv build sidekick
$ ./sidekick -version
Version: 0.2.3-SNAPSHOT
Build date: 2016-12-01T22:18:43Z
GitCommit: 290f97d
GitBranch: nomaven
GitState: dirty
```


# Run


```
$ ./sidekick -h
Usage of ./sidekick:
  -config string
    	Configuration file (default "sidekick.conf")
  -fake string
    	Force response without reload for testing purpose. 'yesman': always say ok, 'drunk': random status/errors for entrypoint updates. Just for test purpose.
  -log-compact
    	compacting log
  -mono
    	only one haproxy instance which play slave/master roles.
  -verbose
    	Log in verbose mode
  -version
    	Print current version
```
