#!/bin/bash

# Load environment variables from .env
export $(grep -v '^#' .env | xargs)

# Run migration
migrate -path db/migrations -database "$DB_URL" up
