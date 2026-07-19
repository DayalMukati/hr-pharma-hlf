#!/bin/bash

# Variables
REPO_URL="https://github.com/DayalMukati/hr-pharma-hlf.git"
TARGET_DIR="$(pwd)"

# Clone the repository
echo "Cloning repository from $REPO_URL..."
git clone $REPO_URL $TARGET_DIR

# Check if clone was successful
if [ $? -ne 0 ]; then
    echo "Failed to clone repository. Exiting..."
    exit 1
fi

# Remove the clone's git metadata. The challenge IDE initializes its own git
# repository in this directory; leaving our .git behind makes its "git init" hit
# an existing repo and its "git remote add origin" fail with
# "remote origin already exists".
rm -rf "$TARGET_DIR/.git"

echo "Setup complete. Files downloaded."
