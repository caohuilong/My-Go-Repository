package main

import (
	"encoding/csv"
	"fmt"
	"github.com/tealeg/xlsx"
	"io"
	"log"
	"os"
)

func main() {
	ipFile, err := xlsx.OpenFile("C:/Users/cao13/Desktop/公网IP列表.xlsx")
	if err != nil {
		log.Printf("can not open the file, err: %s", err.Error())
		return
	}
	ipMap := make(map[string]bool)
	for i, row := range ipFile.Sheets[0].Rows {
		if i == 0 {
			continue
		}
		ipMap[row.Cells[0].String()] = true
	}

	fs1, err := os.Open("C:/Users/cao13/Desktop/ip_sale_226.csv")
	if err != nil {
		log.Printf("can not open the file, err: %s", err.Error())
		return
	}
	defer fs1.Close()

	count := 0
	ipSaleFile := csv.NewReader(fs1)
	for {
		row, err := ipSaleFile.Read()
		if err != nil && err != io.EOF {
			log.Printf("can not read, err: %s", err.Error())
		}
		if err == io.EOF {
			break
		}
		if !ipMap[row[0]] {
			fmt.Println(row[0])
			count++
		}
	}
	fmt.Println(count)
}
