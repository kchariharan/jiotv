@echo off
setlocal EnableDelayedExpansion

:: ============================================================
::  JioTV Go — Build Script for Windows Command Prompt
::  Usage:
::    build.bat                  -> local dev build (no optimizations)
::    build.bat prod             -> production build (-s -w -trimpath)
::    build.bat deps             -> download Go dependencies only
::    build.bat run              -> build (local) then run the server
::    build.bat run prod         -> build (prod) then run the server
::    build.bat clean            -> delete the build\ output folder
::    build.bat docker           -> package the application in a Docker image
:: ============================================================

:: Required for Go 1.25 experimental features
set GOEXPERIMENT=jsonv2,greenteagc

:: Output binary location
set OUT_DIR=build
set BINARY=%OUT_DIR%\jiotv_go.exe

:: ── Parse first argument ──────────────────────────────────────
set MODE=%1
if "%MODE%"=="" set MODE=local

:: ── deps: just download modules ──────────────────────────────
if /I "%MODE%"=="deps" (
    echo [INFO] Downloading Go dependencies...
    go mod tidy
    if !ERRORLEVEL! NEQ 0 ( echo [ERROR] go mod tidy failed & exit /b 1 )
    echo [OK] Dependencies downloaded.
    goto :EOF
)

:: ── clean: remove build folder ────────────────────────────────
if /I "%MODE%"=="clean" (
    echo [INFO] Cleaning build output...
    if exist "%OUT_DIR%" rmdir /s /q "%OUT_DIR%"
    echo [OK] Build folder removed.
    goto :EOF
)

:: ── Validate Go is installed ─────────────────────────────────
where go >nul 2>&1
if !ERRORLEVEL! NEQ 0 (
    echo [ERROR] Go is not installed or not in PATH.
    echo         Download Go 1.25+ from: https://go.dev/dl/
    exit /b 1
)

:: ── Ensure output directory exists ───────────────────────────
if not exist "%OUT_DIR%" mkdir "%OUT_DIR%"

:: ── Download dependencies ─────────────────────────────────────
echo [INFO] Downloading Go dependencies...
go mod tidy
if !ERRORLEVEL! NEQ 0 ( echo [ERROR] go mod tidy failed & exit /b 1 )

:: ── Build ─────────────────────────────────────────────────────
if /I "%MODE%"=="docker" (
    echo [INFO] Packaging application in Docker image 'jiotv_go_local'...
    docker build -t jiotv_go_local .
    if !ERRORLEVEL! NEQ 0 ( echo [ERROR] Docker build failed & exit /b 1 )
    echo [INFO] Exporting Docker image to 'jiotv_go.tar'...
    docker save -o jiotv_go.tar jiotv_go_local
    if !ERRORLEVEL! NEQ 0 ( echo [ERROR] Docker export failed & exit /b 1 )
    echo [OK] Docker image packaged and exported successfully. You can now copy 'jiotv_go.tar' to your destination server.
    goto :EOF
) else if /I "%MODE%"=="prod" (
    echo [INFO] Building PRODUCTION binary...
    go build -ldflags="-s -w" -trimpath -o "%BINARY%" .
) else if /I "%MODE%"=="local" (
    echo [INFO] Building LOCAL/DEV binary...
    go build -o "%BINARY%" .
) else if /I "%MODE%"=="run" (
    :: Second arg decides prod or local for "run" mode
    set BUILD_TYPE=%2
    if "!BUILD_TYPE!"=="" set BUILD_TYPE=local
    if /I "!BUILD_TYPE!"=="prod" (
        echo [INFO] Building PRODUCTION binary for run...
        go build -ldflags="-s -w" -trimpath -o "%BINARY%" .
    ) else (
        echo [INFO] Building LOCAL binary for run...
        go build -o "%BINARY%" .
    )
) else (
    echo [ERROR] Unknown argument: %MODE%
    echo.
    echo Usage:
    echo   build.bat              - local dev build
    echo   build.bat prod         - production build
    echo   build.bat deps         - download dependencies only
    echo   build.bat run          - local build then run
    echo   build.bat run prod     - production build then run
    echo   build.bat clean        - remove build\ folder
    echo   build.bat docker       - package the application in a Docker image
    exit /b 1
)

if !ERRORLEVEL! NEQ 0 (
    echo [ERROR] Build failed.
    exit /b 1
)

echo [OK] Binary ready: %BINARY%

:: ── Run (only when MODE=run) ──────────────────────────────────
if /I "%MODE%"=="run" (
    echo [INFO] Starting JioTV Go server at http://localhost:5001
    echo        Press Ctrl+C to stop.
    echo.
    "%BINARY%" serve
)

endlocal
