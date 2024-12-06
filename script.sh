#!/usr/bin/env bash

# Function to create a new day directory and copy template files
function create_day() {
    if [ -z "$1" ]; then
        echo "Usage: $0 <year/day>"
        exit 1
    fi

    # Extract the year and day from the input
    local input="$1"
    local year=$(echo "$input" | cut -d'/' -f1)
    local day=$(echo "$input" | cut -d'/' -f2)

    # Create the new directory
    local new_dir="$year/$day"
    mkdir -p "$new_dir"

    # Copy template files to the new directory
    local template_dir="template"
    if [ -d "$template_dir" ]; then
        cp -r "$template_dir/"* "$new_dir/"
        echo "$new_dir/main_test.go created."
    else
        echo "Template directory '$template_dir' does not exist."
        exit 1
    fi
}

# Call the function with the provided argument
create_day "$1"