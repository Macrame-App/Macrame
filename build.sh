#!/bin/bash

# Set the name of the build directory
BUILD_DIR="Macrame_$(date +'%m%d%H%M%S')"

# Build the Go application
cd be
go build -o Macrame.exe main.go
go build -o Setup.exe setup/setup.go

# Build the frontend
cd ../fe
npm run build

cd ../builds

# Create the new build directory
mkdir $BUILD_DIR
mkdir $BUILD_DIR/be
mkdir $BUILD_DIR/macros
mkdir $BUILD_DIR/panels
mkdir $BUILD_DIR/public

# Move the generated files to the new build directory
cp ../be/Macrame.exe $BUILD_DIR/be/Macrame.exe
cp ../be/Setup.exe $BUILD_DIR/be/Setup.exe
cp -r ../macros/* $BUILD_DIR/macros/
cp -r ../panels/* $BUILD_DIR/panels/
mv ../public/* $BUILD_DIR/public/

cp ../install.bat $BUILD_DIR/install.bat
cp ../Macrame.lnk $BUILD_DIR/Macrame.lnk

powershell -Command "Compress-Archive -Path $BUILD_DIR/* -DestinationPath $BUILD_DIR.zip -Force"

# Print the path to the new build directory
echo "Build directory: ../$BUILD_DIR"