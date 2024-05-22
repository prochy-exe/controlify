@echo off

if exist "build" (
    rmdir /S /Q "build"
)

call :build_for_platform windows 386
call :build_for_platform windows arm

exit /b

:build_for_platform
setlocal
set GOOS=%1
set GOARCH=%2

if "%GOARCH%"=="arm" (
    set HUMANARCH=arm
) else if "%GOARCH%"=="386" (
    set HUMANARCH=x86
) else (
    set HUMANARCH=%GOARCH%
)

echo Building for %GOOS% %HUMANARCH%...
go build -ldflags="-s -w -X main.isCLI=true" -o build\%GOOS%-%HUMANARCH%\controlify_cli.exe controlify.go
go build -ldflags="-s -w -H=windowsgui" -o build\%GOOS%-%HUMANARCH%\controlify_tray.exe controlify.go
copy config.json build\%GOOS%-%HUMANARCH%\config.json > nul 2>&1
endlocal