@echo off
setlocal

REM Set paths relative to the current working directory

$ruleName = "Macrame LAN Access"
$exePath = Join-Path $PSScriptRoot "be\Macrame.exe"

# Check if rule exists
$existingRule = Get-NetFirewallRule -DisplayName $ruleName -ErrorAction SilentlyContinue

if (-not $existingRule) {
    New-NetFirewallRule -DisplayName $ruleName `
                        -Direction Inbound `
                        -Action Allow `
                        -Program $exePath `
                        -Protocol TCP `
                        -Profile Private `
                        -Enabled True

    Write-Host "Firewall rule '$ruleName' added for $exePath"
} else {
    Write-Host "Firewall rule '$ruleName' already exists"
}
:: Navigate to the "be" directory
cd /d "%~dp0be"

:: Run setup.exe to generate configuration and necessary files
start /wait Setup.exe

:: start /wait caddy.exe fmt --overwrite

:: Run Caddy to generate certificates and serve content
:: start /wait caddy.exe start --config CaddyFile

:: taskkill /f /im caddy.exe

:: Now start macrame.exe
start Macrame.exe

:: End of script
exit