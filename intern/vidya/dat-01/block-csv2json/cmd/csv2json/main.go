package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"block-csv2json/internal/validate"
)

type Block struct {
	BlockNumber int    `json:"block_number"`
	Timestamp   int64  `json:"timestamp"`
	TxCount     int    `json:"tx_count"`
	TxHexSample string `json:"tx_hex_sample"`
}

func main() {
	filePath := flag.String("file", "", "path to CSV file")
    summary := flag.Bool("summary", false, "show block summary")
    flag.Parse()


	if *filePath == "" {
		fmt.Println("Please provide --file flag")
		os.Exit(1)
	}

	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		os.Exit(1)
	}

	var blocks []Block

	for i, row := range records {
		if i == 0 {
			continue // skip header
		}

		blockNumber, err := validate.BlockNumber(row[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		timestamp, err := validate.Timestamp(row[1])
		if err != nil {
			fmt.Println(err)
			continue
		}

		txCount, err := validate.TxCount(row[2])
		if err != nil {
			fmt.Println(err)
			continue
		}

		block := Block{
			BlockNumber: blockNumber,
			Timestamp:   timestamp,
			TxCount:     txCount,
			TxHexSample: row[3],
		}

		blocks = append(blocks, block)
	}

	output, err := json.MarshalIndent(blocks, "", "  ")
	if err != nil {
		fmt.Println("Error creating JSON:", err)
		os.Exit(1)
	}

	if *summary {
	totalTx := 0
	for _, b := range blocks {
		totalTx += b.TxCount
	}

	fmt.Println("Block count:", len(blocks))
	fmt.Println("Total transactions:", totalTx)
	return
}

fmt.Println(string(output))

}
