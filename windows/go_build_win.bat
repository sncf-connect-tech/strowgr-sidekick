echo @off
call script.conf.win.bat
echo @on
REM deprecated go build -o sidekick.exe github.com/voyages-sncf-technologies/strowgr-sidekick/cmd/sidekick
go install -a github.com/voyages-sncf-technologies/strowgr-sidekick/cmd/sidekick