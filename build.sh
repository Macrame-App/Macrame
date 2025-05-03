#!/bin/bash

# Set the name of the build directory
BUILD_DIR="Macrame_$(date +'%m%d%H%M%S')"

# Build the Go application
go build -ldflags "-H=windowsgui" -o Macrame.exe main.go

# Build the frontend
cd fe
npm run build

cd ../builds

# Create the new build directory
mkdir $BUILD_DIR
mkdir $BUILD_DIR/macros
mkdir $BUILD_DIR/panels
mkdir $BUILD_DIR/panels/test_panel
mkdir $BUILD_DIR/public
mkdir $BUILD_DIR/public/assets

# Move the generated files to the new build directory
cp ../Macrame.exe $BUILD_DIR/Macrame.exe
cp ../favicon.ico $BUILD_DIR/favicon.ico
find ../macros -type f ! -name 'ED-*' -exec cp --parents {} "$BUILD_DIR/macros/" \;
cp -r ../panels/test_panel/* $BUILD_DIR/panels/test_panel/
mv ../public/* $BUILD_DIR/public/

# cp ../install.bat $BUILD_DIR/install.bat

powershell -Command "Compress-Archive -Path $BUILD_DIR/* -DestinationPath $BUILD_DIR.zip -Force"

# Print the path to the new build directory
echo "Build directory: ../$BUILD_DIR"