# Block CSV to JSON Utility

This project is a simple Go command-line tool that converts blockchain
CSV data into structured JSON output and provides a viewer to inspect
the data.

## Project Structure

block-csv2json/
cmd/csv2json/main.go – CLI entry point
internal/validate – validation logic
internal/encode – encoding utilities
testdata/blocks.csv – sample blockchain data
viewer/index.html – web viewer for JSON output

## Running the Parser

Run the parser with:

go run ./cmd/csv2json --file=testdata/blocks.csv > out.json

This reads the CSV file and converts it into JSON format.

## CLI Summary Command

To print a summary of blocks:

go run ./cmd/csv2json --file=testdata/blocks.csv --summary

This prints:
- total number of blocks
- total number of transactions

## Running Tests

Run validation tests with:

go test ./...

## Running the Viewer

Start a local server:

python -m http.server 8000

Then open:

http://localhost:8000/viewer/

The viewer loads `out.json` and displays blockchain blocks in a table.

## Example JSON Output

[
  {
    "block_number": 1001,
    "timestamp": 1700000000,
    "tx_count": 120,
    "tx_hex_sample": "0xabc123"
  }
]
