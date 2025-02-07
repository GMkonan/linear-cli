#!/bin/bash

APP_NAME="linear-cli"
REPO_URL="https://github.com/GMkonan/linear-cli.git"
INSTALL_DIR="/usr/local/bin"
ENV_FILE=".env"
CONFIG_DIR="$HOME/.config/$APP_NAME"

if ! command -v go &> /dev/null; then
    echo "Go is not installed. Please install Go and try again."
    exit 1
fi

echo "Cloning repository..."
TEMP_DIR=$(mktemp -d)
git clone "$REPO_URL" "$TEMP_DIR"

echo "Building application..."
cd "$TEMP_DIR"
go build -o "$APP_NAME"

echo "Installing application to $INSTALL_DIR..."
sudo mv "$APP_NAME" "$INSTALL_DIR"

echo "Setting up environment file..."
mkdir -p "$CONFIG_DIR"
if [ -f "$TEMP_DIR/$ENV_FILE" ]; then
    echo "Copying $ENV_FILE to $CONFIG_DIR..."
    cp "$TEMP_DIR/$ENV_FILE" "$CONFIG_DIR/"
else
    echo "No $ENV_FILE found in the repository. Creating a new one..."
    touch "$CONFIG_DIR/$ENV_FILE"
    echo "# Add your environment variables here" > "$CONFIG_DIR/$ENV_FILE"
fi

echo "Cleaning up..."
rm -rf "$TEMP_DIR"

if command -v "$APP_NAME" &> /dev/null; then
    echo "Installation successful! You can now run '$APP_NAME' from anywhere."
    echo "Make sure to configure your environment variables in $CONFIG_DIR/$ENV_FILE."
else
    echo "Installation failed. Please check the script and try again."
    exit 1
fi
