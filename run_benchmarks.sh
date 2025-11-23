#!/bin/bash

# GYAML Benchmark Runner
# This script runs benchmarks and generates a formatted report

set -e

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}╔═══════════════════════════════════════════════════════════════╗${NC}"
echo -e "${BLUE}║          GYAML Benchmark Suite                                ║${NC}"
echo -e "${BLUE}╚═══════════════════════════════════════════════════════════════╝${NC}"
echo ""

# Get system information
echo -e "${GREEN}System Information:${NC}"
echo "Go Version: $(go version)"

if [[ "$OSTYPE" == "darwin"* ]]; then
    echo "System: macOS"
    system_profiler SPHardwareDataType | grep -E "Model Name|Chip|Memory" | sed 's/^/  /'
elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
    echo "System: Linux"
    lscpu | grep -E "Model name|CPU\(s\):" | sed 's/^/  /'
    free -h | grep Mem | awk '{print "  Memory: " $2}'
fi

echo ""
echo -e "${GREEN}Running benchmarks...${NC}"
echo ""

# Create results directory
RESULTS_DIR="results"
mkdir -p "$RESULTS_DIR"

# Timestamp for results
TIMESTAMP=$(date +"%Y%m%d_%H%M%S")
RESULT_FILE="$RESULTS_DIR/benchmark_${TIMESTAMP}.txt"

# Run benchmarks
if [ "$1" == "--quick" ]; then
    echo "Running quick benchmark (count=1)..."
    go test -bench . -benchmem -benchtime=1s -count=1 | tee "$RESULT_FILE"
else
    echo "Running full benchmark (count=5, 3s each)..."
    go test -bench . -benchmem -benchtime=3s -count=5 | tee "$RESULT_FILE"
fi

echo ""
echo -e "${GREEN}Benchmark complete!${NC}"
echo -e "Results saved to: ${YELLOW}$RESULT_FILE${NC}"
echo ""

# Generate summary
echo -e "${BLUE}╔═══════════════════════════════════════════════════════════════╗${NC}"
echo -e "${BLUE}║          Quick Summary                                        ║${NC}"
echo -e "${BLUE}╚═══════════════════════════════════════════════════════════════╝${NC}"
echo ""

# Extract GYAML vs yaml.v3 comparisons
echo -e "${GREEN}Performance Comparison:${NC}"
echo ""

# Simple comparison
gyaml_get=$(grep "BenchmarkGYAMLGet-" "$RESULT_FILE" | head -1 | awk '{print $3, $4}')
yaml_map=$(grep "BenchmarkYAMLv3UnmarshalMap-" "$RESULT_FILE" | head -1 | awk '{print $3, $4}')
yaml_struct=$(grep "BenchmarkYAMLv3UnmarshalStruct-" "$RESULT_FILE" | head -1 | awk '{print $3, $4}')

echo "  GYAML Get:              $gyaml_get"
echo "  yaml.v3 Unmarshal Map:  $yaml_map"
echo "  yaml.v3 Unmarshal Struct: $yaml_struct"
echo ""
