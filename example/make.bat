@echo off

setlocal

if exist make.bat goto ok
echo make.bat must be run from its folder
goto end

:ok

set OLDGOPATH=%GOPATH%
set GOPATH=%~dp0

go install monitor

set GOPATH=%OLDGOPATH%

:end
echo finished
