@echo off

set ruleName="Macrame LAN Access"
set exePath=%~dp0be\Macrame.exe

:: Check if rule exists
netsh advfirewall firewall show rule name=%ruleName% >nul 2>&1
if %errorlevel%==1 (
    netsh advfirewall firewall add rule name=%ruleName% dir=in action=allow program=%exePath% protocol=tcp profile=private enabled=true
    echo Firewall rule '%ruleName%' added for %exePath%
) else (
    echo Firewall rule '%ruleName%' already exists
)

:: Navigate to the "be" directory
cd /d "%~dp0be"

echo Moved to Backend directory

:: Run setup.exe to generate configuration and necessary files
start /wait Setup.exe

:: Run Caddy to generate certificates and serve content
:: start /wait caddy.exe start --config CaddyFile

:: taskkill /f /im caddy.exe

:: Now start macrame.exe
start Macrame.exe

:: End of script
exit