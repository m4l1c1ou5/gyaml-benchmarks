# Makefile for GYAML Benchmarks

.PHONY: all bench quick compare clean install-tools help

# Default target
all: help

## help: Display this help message
help:
	@echo "GYAML Benchmark Makefile"
	@echo ""
	@echo "Available targets:"
	@echo "  make bench          - Run full benchmark suite (5 iterations, 3s each)"
	@echo "  make quick          - Run quick benchmark (1 iteration, 1s each)"
	@echo "  make clean          - Clean benchmark results and profiles"
	@echo ""

## bench: Run full benchmark suite
bench:
	@echo "Running full benchmark suite..."
	@./run_benchmarks.sh

## quick: Run quick benchmark
quick:
	@echo "Running quick benchmark..."
	@./run_benchmarks.sh --quick

## clean: Clean benchmark results and profiles
clean:
	@echo "Cleaning benchmark results and profiles..."
	@rm -rf results/*.txt
	@rm -rf results/*.md
	@rm -f *.prof
	@rm -f *.test
	@echo "Done!"
