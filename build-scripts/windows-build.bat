@echo off
setlocal enabledelayedexpansion

REM Step 1: Build Macrame Go application for Windows
echo Building Macrame Go Application for Windows...

go build -ldflags "-H=windowsgui" -o Macrame.exe main.go

IF %ERRORLEVEL% NEQ 0 (
  echo Go build failed!
  exit /b %ERRORLEVEL%
) ELSE (
  echo Go build was successful!
)

REM Step 2: Build Macrame Vue UI
echo Moving to ui directory and building Vue UI
cd ui

echo Running npm install...
call npm install
IF %ERRORLEVEL% NEQ 0 (
    echo npm install failed!
    exit /b %ERRORLEVEL%
)

echo Running npm run build...
call npm run build
IF %ERRORLEVEL% NEQ 0 (
    echo npm run build failed!
    exit /b %ERRORLEVEL%
)

cd ..

REM Step 3: Cleanup root directory of build files.
echo Cleaning up root directory...

echo Removing directories: app, ui, and build-scripts
rmdir /s /q app
rmdir /s /q ui
rmdir /s /q build-scripts

for %%F in (*) do (
    set "file=%%~nxF"
    if /I not "!file!"=="Macrame.exe" if /I not "!file!"=="favicon.ico" if /I not "!file!"=="README.md" (
        echo Deleting !file!
        del /f /q "%%F"
    )
)

echo Build and cleanup complete.
