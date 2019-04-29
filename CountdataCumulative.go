package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	rows := readOrders("showdata.csv")
	rows = calculate(rows)
	writeOrders("showdata.csv", rows)
}

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
func calculate(rows [][]string) [][]string {

	sum := 0


	for i := range rows {





		price, err := strconv.Atoi(strings.Replace(rows[i][1], ".", "", -1))
		if err != nil {
			log.Fatalf("Cannot retrieve price of %s: %s\n", err)

		}




		rows[i] = append(rows[i])


		sum += price









	}

	rows = append(rows, []string{"sum", intToFloatString(sum)})


	return rows
}

func intToFloatString(n int) string {
	intgr := n/100
	frac := n- intgr *100
	return fmt.Sprintf("%d.%d", intgr, frac)
}



func writeOrders(name string, rows [][]string) {

	f, err := os.Create(name)
	if err != nil {
		log.Fatalf("Cannot open '%s': %s\n", name, err.Error())
	}

	defer func() {
		e := f.Close()
		if e != nil {
			log.Fatalf("Cannot close '%s': %s\n", name, e.Error())
		}
	}()

	w := csv.NewWriter(f)
	err = w.WriteAll(rows)
}