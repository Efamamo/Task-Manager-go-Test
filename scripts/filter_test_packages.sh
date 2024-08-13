#!/bin/bash

# List of packages to exclude (adjust as needed)
exclude_packages=(
  "github.com/Task-Management-go/Tests/task_tests/task_repo"
  "github.com/Task-Management-go/Tests/user_tests/user_repo"
)

# Find all packages with .go files
all_packages=$(go list ./Tests)

# Filter out the excluded packages
for exclude in "${exclude_packages[@]}"; do
  all_packages=$(echo "$all_packages" | grep -v "^${exclude}$")
done

# Convert to a space-separated list
test_packages_list=$(echo "$all_packages" | tr '\n' ' ')

# Print the list of packages
echo $test_packages_list