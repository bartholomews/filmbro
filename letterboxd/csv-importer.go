package letterboxd

import (
	"encoding/csv"
	"github.com/bartholomews/filmbro/mubi"
	"log"
	"os"
	"strconv"
)

// CreateCsvImport - see https://letterboxd.com/about/importing-data
func CreateCsvImport(entries []mubi.DiaryEntry) {
	// create a file
	file, err := os.Create("letterboxd-diary-import.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			os.Exit(1)
		}
	}(file)

	// initialize csv writer
	writer := csv.NewWriter(file)

	// Write first row
	firstRow := []string{"Title", "Year", "WatchedDate", "Rating"}
	_ = writer.Write(firstRow)

	var rows [][]string
	for _, entry := range entries {
		var rating = ""
		if entry.Rating != nil {
			rating = strconv.Itoa(*entry.Rating)
		}
		row := []string{entry.Film.Title, strconv.Itoa(entry.Film.Year), entry.WatchedDate, rating}
		rows = append(rows, row)
	}

	_ = writer.WriteAll(rows)

	defer writer.Flush()
}
