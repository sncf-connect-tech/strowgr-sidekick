@echo off
call script.conf.linux.bat
@echo on

REM build the executable, with GOSO=linux 
go install -a github.com/voyages-sncf-technologies/strowgr-sidekick/cmd/sidekick

REM COPY built executable at project root
COPY %GOPATH%\bin\linux_amd64\sidekick %GOPATH%\src\github.com\voyages-sncf-technologies\strowgr-sidekick\sidekick

Rem build docker image
docker build -t strowgr/sidekick:%S_VERSION% --no-cache -f  %GOPATH%\src/github.com/voyages-sncf-technologies\strowgr-sidekick\Dockerfile %GOPATH%\src\github.com\voyages-sncf-technologies\strowgr-sidekick