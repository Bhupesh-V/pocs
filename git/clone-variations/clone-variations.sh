#!/usr/bin/env bash

# Ensure a repository URL is provided
if [ -z "$1" ]; then
  echo "Usage: $0 <repo_url> [branch_name]"
  exit 1
fi

REPO_URL="$1"
BRANCH_NAME="${2:-main}"  # Default to 'main' if not provided
BASE_DIR="clones"

echo "Using branch: $BRANCH_NAME"

# remove previous clones
rm -rf "${BASE_DIR:?}"/*

plain_clone() {
  local dir="$BASE_DIR/plain"
  [ -d "$dir" ] || git clone "$REPO_URL" "$dir"
}

bare_clone() {
  local dir="$BASE_DIR/bare"
  [ -d "$dir" ] || git clone --bare "$REPO_URL" "$dir"
}

bare_blobless_clone() {
  local dir="$BASE_DIR/bare-blobless"
  [ -d "$dir" ] || git clone --bare --filter=blob:none "$REPO_URL" "$dir"
}

shallow_clone() {
  local dir="$BASE_DIR/shallow"
  [ -d "$dir" ] || git clone --depth 1 "$REPO_URL" "$dir"
}

treeless_bare_clone() {
  local dir="$BASE_DIR/bare-treeless"
  [ -d "$dir" ] || git clone --bare --filter=tree:0 "$REPO_URL" "$dir"
}

single_branch_clone() {
  local dir="$BASE_DIR/single-branch"
  [ -d "$dir" ] || git clone --single-branch --branch "$BRANCH_NAME" "$REPO_URL" "$dir"
}

blobless_shallow_clone() {
  local dir="$BASE_DIR/blobless-shallow"
  [ -d "$dir" ] || git clone --bare --filter=blob:none --depth 1 "$REPO_URL" "$dir"
}

blobless_single_branch_clone() {
  local dir="$BASE_DIR/blobless-single-branch"
  [ -d "$dir" ] || git clone --bare --filter=blob:none --single-branch --branch "$BRANCH_NAME" "$REPO_URL" "$dir"
}

treeless_single_branch_clone() {
  local dir="$BASE_DIR/treeless-single-branch"
  [ -d "$dir" ] || git clone --bare --filter=tree:0 --single-branch --branch "$BRANCH_NAME" "$REPO_URL" "$dir"
}

export REPO_URL BRANCH_NAME BASE_DIR
export -f plain_clone bare_clone bare_blobless_clone shallow_clone treeless_bare_clone single_branch_clone blobless_shallow_clone blobless_single_branch_clone treeless_single_branch_clone
parallel --progress ::: plain_clone bare_clone bare_blobless_clone shallow_clone treeless_bare_clone single_branch_clone blobless_shallow_clone blobless_single_branch_clone treeless_single_branch_clone

# results
du -h -d 0 "$BASE_DIR"/* | sort -h