package main

import (
	"encoding/csv"

	"fmt"
	"io"
	"log"
	"os"

)

func main() {
	datatype := "sum"
	if datatype == "sum" {


		csvfile1, err := os.Open("sum.csv")

		if err != nil {
			fmt.Println(err)
			return
		}
		defer csvfile1.Close()

		reader := csv.NewReader(csvfile1)

		csvfile2, err := os.Create("showdata.csv")

		if err != nil {
			fmt.Println(err)
			return
		}
		writer := csv.NewWriter(csvfile2)

		for i := 0; i  < 5; i++ {


			record, err := reader.Read()

			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println(err)
				return
			}
			err = writer.Write(record)

			if err != nil {
				fmt.Println(err)
				return
			}
		}


		writer.Flush()
		err = csvfile2.Close()
		if err != nil {
			fmt.Println(err)
			return
		}


	}
	if datatype == "cumulative" {csvfile1, err := os.Open("cumulative.csv")

		if err != nil {
			fmt.Println(err)
			return
		}
		defer csvfile1.Close()

		reader := csv.NewReader(csvfile1)

		csvfile2, err := os.Create("showdata.csv")

		if err != nil {
			fmt.Println(err)
			return
		}
		writer := csv.NewWriter(csvfile2)

		for i := 0; i  < 19000; i++ {


			record, err := reader.Read()

			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println(err)
				return
			}
			err = writer.Write(record)

			if err != nil {
				fmt.Println(err)
				return
			}
		}


		writer.Flush()
		err = csvfile2.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}}

func readOrders(name string) [][]string {

	f, err := os.Open(name)

	if err != nil {
		log.Fatalf("Cannot open '%s': %s\n", name, err.Error())
	}

	defer f.Close()

	r := csv.NewReader(f)


	r.Comma = ';'

	rows, err := r.ReadAll()

	if err != nil {
		log.Fatalln("Cannot read CSV data:", err.Error())
	}

	return rows
}



