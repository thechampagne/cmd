@echo off
for /r %%i in (*.go) do (
    go build -o ./bin/%%~ni.exe %%i
)