#!/bin/bash

# Load environment variables from .env
export $(grep -v '^#' .env | xargs)

# Run migration
migrate -path db/migrations -database "$DB_URL" down
if [ $? -ne 0 ]; then
    echo "Migration drop failed"
    exit 1
fi
# migrate -path db/migrations -database "$DB_URL" up
# if [ $? -ne 0 ]; then
#     echo "Migration up failed"
#     exit 1
# fi
echo "Migration completed successfully"
