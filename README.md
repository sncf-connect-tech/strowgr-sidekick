[![Build Status](https://travis-ci.org/voyages-sncf-technologies/strowgr-sidekick.svg?branch=0.2.x)](https://travis-ci.org/voyages-sncf-technologies/strowgr-sidekick) [![GitHub release](https://img.shields.io/github/release/voyages-sncf-technologies/strowgr-sidekick.svg)](https://github.com/voyages-sncf-technologies/strowgr-sidekick/releases/latest) [![Coverage Status](https://coveralls.io/repos/github/voyages-sncf-technologies/strowgr-sidekick/badge.svg)](https://coveralls.io/github/voyages-sncf-technologies/strowgr-sidekick) [![codecov](https://codecov.io/gh/voyages-sncf-technologies/strowgr-sidekick/branch/0.2.x/graph/badge.svg)](https://codecov.io/gh/voyages-sncf-technologies/strowgr-sidekick) [![Go Report Card](https://goreportcard.com/badge/github.com/voyages-sncf-technologies/strowgr-sidekick)](https://goreportcard.com/report/github.com/voyages-sncf-technologies/strowgr-sidekick)


# Build

```
$ GOPATH=$PWD go build sidekick
```

# Build with go 
  1. SET (GOARCH,GOOS) to (amd64,linux) for linux and (GOARCH,GOOS) to (amd64,windows) for windows.
  
  You can:
  
  2. either run 
    - go install -a github.com/voyages-sncf-technologies/strowgr-sidekick.
    - The executable is at $GOPATH/bin or $GOPATH/bin/$GOOS_$GOARCH and is named strowgr_sidekick.
  
  3. At root of your project: 
    - go build -o sidekick
    - The executable is at the root of the project and is named sidekick.


$ ./sidekick -version
Version: 0.2.3-SNAPSHOT
Build date: 2016-12-01T22:18:43Z
GitCommit: 290f97d
GitBranch: nomaven
GitState: dirty
```

# Build with docker container
```bash
docker build -f Dockerfile.build -t strowgr/sidekick-builder:$(cat VERSION) .
docker run -v $PWD:/go -v $PWD:/bin strowgr/sidekick-builder:$(cat VERSION)
```
Binary is generated here: ```./sidekick```.

# Build docker image with haproxy instances
```bash
docker build -t strowgr/sidekick:$(cat VERSION) .
```

# Run


```
$ ./sidekick -h
Usage of ./sidekick:
  -config string
    	Configuration file (default "sidekick.conf")
  -fake string
    	Force response without reload for testing purpose. 'yesman': always say ok, 'drunk': random status/errors for entrypoint updates. Just for test purpose.
  -generate-config
    	generate a default config file to the standard output
  -log-compact
    	compacting log format
  -mono
    	only one haproxy instance which play slave/master roles.
  -verbose
    	Log in verbose mode
  -version
    	Print current version
```

# Release

For a release branch created like this:

```bash
git checkout -b 0.1.0 develop
```

Execute following command to create a release:

```bash
./release.sh 0.1.0
```

Local Desk Installation
  1.  If no existing workspace, create one (== Empty Directory).
  2.  Set GOPATH to this workspace.
  3.   Think about set GOOS and GOARCH, depending of your OS: (GOARCH,GOOS)  = (amd64,linux) or (amd64,windows)
  4.   If not present, get go dep: go get -u github.com/golang/dep/cmd/dep
    - in windows command line, the window could close after executing, so you can run: start go get -u github.com/golang/dep/cmd/dep 
  5.  Get the projet: run: go get github.com/voyages-sncf-technologies/strowgr-sidekick (start github.com/voyages-sncf-technologies/strowgr-sidekick) 