#!/bin/bash

# Set to true to actually delete indices, false for dry run
DELETE_MODE=true

# Function to delete an index
delete_index() {
    local index=$1
    if [ "$DELETE_MODE" = true ]; then
        echo "Deleting index: $index"
        curl -s -X DELETE "http://10.9.36.111:42792/$index" > /dev/null
        if [ $? -eq 0 ]; then
            echo "✓ Successfully deleted $index"
        else
            echo "✗ Failed to delete $index"
        fi
    else
        echo "[DRY RUN] Would delete index: $index"
    fi
}

# Get indices and process them
indices=$(curl -s -X GET "http://10.9.36.111:42792/_alias?pretty" 2>/dev/null | \
jq -r '
    to_entries[] 
    | select(.value.aliases == {})
    | .key
    | select(test("\\d+$") and (contains("car_exp") | not) and (contains("carmodel2") | not))
')

if [ -z "$indices" ]; then
    echo "No matching indices found"
    exit 0
fi

echo "Found indices to process:"
echo "$indices" | while read -r index; do
    if [ ! -z "$index" ]; then
        delete_index "$index"
    fi
done

if [ "$DELETE_MODE" = false ]; then
    echo -e "\nThis was a dry run. Set DELETE_MODE=true to actually delete indices."
fi