#!/bin/bash

# run_test.sh - Run bot-go with test repositories
#
# Usage:
#   ./run_test.sh <repo-name> [options]
#
# Options:
#   --build-index    Build the code graph index for the repository
#   --test-dump      Dump the code graph for testing/debugging
#   --clean          Clean up the index from different DBs
#   --head           Use git HEAD mode for faster indexing (with --build-index)
#   --all            Run all options: build-index, test-dump, clean
#   --help           Show this help message
#
# Examples:
#   ./run_test.sh python-calculator --build-index
#   ./run_test.sh go-calculator --build-index --head
#   ./run_test.sh typescript-calculator --test-dump
#   ./run_test.sh java-modern-calculator --clean
#   ./run_test.sh java8-calculator --all

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
BOT_GO_BIN="${SCRIPT_DIR}/bin/bot-go"
APP_CONFIG="${SCRIPT_DIR}/config/app.yaml"
SOURCE_CONFIG="${SCRIPT_DIR}/tests/source.yaml"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Available test repositories
AVAILABLE_REPOS=(
    "python-calculator"
    "go-calculator"
    "typescript-calculator"
    "java-modern-calculator"
    "java8-calculator"
)

show_help() {
    echo "Usage: $0 <repo-name> [options]"
    echo ""
    echo "Run bot-go with test repositories from tests/repos/"
    echo ""
    echo "Available repositories:"
    for repo in "${AVAILABLE_REPOS[@]}"; do
        echo "  - $repo"
    done
    echo ""
    echo "Options:"
    echo "  --build-index    Build the code graph index for the repository"
    echo "  --test-dump      Dump the code graph for testing/debugging"
    echo "  --clean          Clean up the index from different DBs"
    echo "  --head           Use git HEAD mode for faster indexing (with --build-index)"
    echo "  --all            Run all options: build-index, test-dump, clean"
    echo "  --help           Show this help message"
    echo ""
    echo "Examples:"
    echo "  $0 python-calculator --build-index"
    echo "  $0 go-calculator --build-index --head"
    echo "  $0 typescript-calculator --test-dump"
    echo "  $0 java-modern-calculator --clean"
    echo "  $0 java8-calculator --all"
}

# Check if bot-go binary exists
check_binary() {
    if [[ ! -f "$BOT_GO_BIN" ]]; then
        echo -e "${YELLOW}Bot-go binary not found. Building...${NC}"
        cd "$SCRIPT_DIR"
        make build
        if [[ ! -f "$BOT_GO_BIN" ]]; then
            echo -e "${RED}Failed to build bot-go binary${NC}"
            exit 1
        fi
        echo -e "${GREEN}Build complete${NC}"
    fi
}

# Validate repository name
validate_repo() {
    local repo="$1"
    for r in "${AVAILABLE_REPOS[@]}"; do
        if [[ "$r" == "$repo" ]]; then
            return 0
        fi
    done
    return 1
}

# Main execution
main() {
    # Check for help flag first
    for arg in "$@"; do
        if [[ "$arg" == "--help" || "$arg" == "-h" ]]; then
            show_help
            exit 0
        fi
    done

    # Check if repo name is provided
    if [[ $# -lt 1 ]]; then
        echo -e "${RED}Error: Repository name is required${NC}"
        echo ""
        show_help
        exit 1
    fi

    REPO_NAME="$1"
    shift

    # Validate repository name
    if ! validate_repo "$REPO_NAME"; then
        echo -e "${RED}Error: Unknown repository '$REPO_NAME'${NC}"
        echo ""
        echo "Available repositories:"
        for repo in "${AVAILABLE_REPOS[@]}"; do
            echo "  - $repo"
        done
        exit 1
    fi

    # Parse options
    DO_BUILD_INDEX=false
    DO_TEST_DUMP=false
    DO_CLEAN=false
    USE_HEAD=false

    while [[ $# -gt 0 ]]; do
        case "$1" in
            --build-index)
                DO_BUILD_INDEX=true
                shift
                ;;
            --test-dump)
                DO_TEST_DUMP=true
                shift
                ;;
            --clean)
                DO_CLEAN=true
                shift
                ;;
            --head)
                USE_HEAD=true
                shift
                ;;
            --all)
                DO_BUILD_INDEX=true
                DO_TEST_DUMP=true
                DO_CLEAN=true
                shift
                ;;
            *)
                echo -e "${RED}Error: Unknown option '$1'${NC}"
                show_help
                exit 1
                ;;
        esac
    done

    # Check if at least one action is specified
    if ! $DO_BUILD_INDEX && ! $DO_TEST_DUMP && ! $DO_CLEAN; then
        echo -e "${RED}Error: At least one action must be specified${NC}"
        echo "Use --build-index, --test-dump, --clean, or --all"
        exit 1
    fi

    # Check binary exists
    check_binary

    echo -e "${BLUE}========================================${NC}"
    echo -e "${BLUE}Running tests for: ${GREEN}$REPO_NAME${NC}"
    echo -e "${BLUE}========================================${NC}"

    # Build index
    if $DO_BUILD_INDEX; then
        echo ""
        echo -e "${YELLOW}>>> Building index for $REPO_NAME${NC}"

        BUILD_CMD="$BOT_GO_BIN -app=$APP_CONFIG -source=$SOURCE_CONFIG -build-index=$REPO_NAME"
        if $USE_HEAD; then
            BUILD_CMD="$BUILD_CMD --head"
        fi

        echo -e "${BLUE}Command: $BUILD_CMD${NC}"
        eval "$BUILD_CMD"
        echo -e "${GREEN}<<< Index build complete${NC}"
    fi

    # Test dump
    if $DO_TEST_DUMP; then
        echo ""
        echo -e "${YELLOW}>>> Dumping code graph for $REPO_NAME${NC}"

        DUMP_CMD="$BOT_GO_BIN -app=$APP_CONFIG -source=$SOURCE_CONFIG -test-dump=$REPO_NAME"
        echo -e "${BLUE}Command: $DUMP_CMD${NC}"
        eval "$DUMP_CMD"
        echo -e "${GREEN}<<< Test dump complete${NC}"
    fi

    # Clean
    if $DO_CLEAN; then
        echo ""
        echo -e "${YELLOW}>>> Cleaning index for $REPO_NAME${NC}"

        CLEAN_CMD="$BOT_GO_BIN -app=$APP_CONFIG -source=$SOURCE_CONFIG -test-dump=$REPO_NAME --clean"
        echo -e "${BLUE}Command: $CLEAN_CMD${NC}"
        eval "$CLEAN_CMD"
        echo -e "${GREEN}<<< Clean complete${NC}"
    fi

    echo ""
    echo -e "${GREEN}========================================${NC}"
    echo -e "${GREEN}All operations completed for $REPO_NAME${NC}"
    echo -e "${GREEN}========================================${NC}"
}

main "$@"
