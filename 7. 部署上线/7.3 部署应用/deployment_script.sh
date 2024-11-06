#!/bin/bash

# Define variables
APP_NAME="shop-app"
DEPLOYMENT_DIR="/path/to/deployment"
REPO_URL="https://github.com/your-repo/$APP_NAME.git"
BRANCH="main"

# Clone the repository
if [ ! -d "$DEPLOYMENT_DIR/$APP_NAME" ]; then
    git clone "$REPO_URL" "$DEPLOYMENT_DIR/$APP_NAME"
fi

# Navigate to the application directory
cd "$DEPLOYMENT_DIR/$APP_NAME"

# Checkout the specified branch
git checkout "$BRANCH"

# Pull the latest changes from the remote repository
git pull origin "$BRANCH"

# Build the application
go build -o /usr/local/bin/$APP_NAME

# Start the application
/usr/local/bin/$APP_NAME &

# Output the PID of the running application
echo "Application started with PID $!"
