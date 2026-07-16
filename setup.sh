#!/bin/bash

# Variables
REPO_URL="https://github.com/DayalMukati/hr-pharma-hlf.git"
REPO_BRANCH="solution"   # TEST BRANCH: filled-in solution chaincode
TARGET_DIR="$(pwd)"

# Clone the repository
echo "Cloning repository from $REPO_URL (branch: $REPO_BRANCH)..."
git clone --branch "$REPO_BRANCH" $REPO_URL $TARGET_DIR

# Check if clone was successful
if [ $? -ne 0 ]; then
    echo "Failed to clone repository. Exiting..."
    exit 1
fi

echo "Setup complete. Files downloaded."
