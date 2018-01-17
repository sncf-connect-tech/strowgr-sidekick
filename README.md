[![Build Status](https://travis-ci.org/voyages-sncf-technologies/strowgr-sidekick.svg?branch=0.2.x)](https://travis-ci.org/voyages-sncf-technologies/strowgr-sidekick) [![GitHub release](https://img.shields.io/github/release/voyages-sncf-technologies/strowgr-sidekick.svg)](https://github.com/voyages-sncf-technologies/strowgr-sidekick/releases/latest) [![Coverage Status](https://coveralls.io/repos/github/voyages-sncf-technologies/strowgr-sidekick/badge.svg)](https://coveralls.io/github/voyages-sncf-technologies/strowgr-sidekick) [![codecov](https://codecov.io/gh/voyages-sncf-technologies/strowgr-sidekick/branch/0.2.x/graph/badge.svg)](https://codecov.io/gh/voyages-sncf-technologies/strowgr-sidekick) [![Go Report Card](https://goreportcard.com/badge/github.com/voyages-sncf-technologies/strowgr-sidekick)](https://goreportcard.com/report/github.com/voyages-sncf-technologies/strowgr-sidekick)


# Build

```
$ GOPATH=$PWD go build sidekick
```

# Build with go 
  1. SET (GOARCH,GOOS) to (amd64,linux) for linux and (GOARCH,GOOS) to (amd64,windows) for windows.
  
  You can:
  
  2. either run 
    - go install -a github.com/voyages-sncf-technologies/strowgr-sidekick/cmd/sidekick.
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

Local Desk Installation.
  0. Prequisite: Install go if not present.
    0.1. Set GOROOT the installation folder. 
    0.2. add GOROOT/bin to PATH.
  1. If no existing workspace, create one (== Empty Directory).
  2. Set GOPATH to this workspace.
    2.1.  ADD GOPATH/bin to PATH
  3.  If not present, get go dep: go get -u -v github.com/golang/dep/cmd/dep. 
    3.1 dep executable is installed at GOPATH/bin
  4. Get the projet: run: go get -v github.com/voyages-sncf-technologies/strowgr-sidekick
  5. Go to the project at GOPATH/src/github.com/voyages-sncf-technologies/strowgr-sidekick
  7. Install dependencies, run: dep ensure to install dependencies (vendor sub folder created)
  8. To build a native executable
    8.1 Run go install -a github.com/voyages-sncf-technologies/strowgr-sidekick/cmd/sidekick.
      sidekick executable will be installed in GOPATH/bin, so in the PATH
    8.2 At root of the project only, you can also run: 
      8.2.1 For windows user: go build -o sidekick.exe sidekick executable will be installed in GOPATH/bin
      8.2.2 For linux user: go build -o sidekick executable will be installed in GOPATH/bin
      8.2.3 You can arbitraly choose the value for -o option, but need .exe extension for windows run

  9. To build a docker image in local desk
    9.1 the executable you build has to be embedded in docker linux image, so you will have to change GOOS to linux if necessary (and reset it after)
       
  
  - WARN: in windows command line, the window could close after executing, so you can run: start go get -u -v github.com/golang/dep/cmd/dep      