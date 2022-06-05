package readers

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"
)

type CSV struct{}

type Data struct {
	Date  string
	Value float64
}

func NewCSV() *CSV {
	return &CSV{}
}

func (c *CSV) Read(filename string, header bool, dateIdx, fieldIdx int) (data []Data, err error) {

	f, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	r := csv.NewReader(strings.NewReader(string(f)))
	for {

		var record []string
		record, err = r.Read()
		if err != nil {
			if err == io.EOF {
				return data, nil
			}
			return
		}
		if header {
			header = false
			continue
		}

		var value float64
		value, err = strconv.ParseFloat(record[fieldIdx], 64)
		if err != nil {
			return
		}
		d := Data{
			Date:  record[dateIdx],
			Value: value,
		}

		data = append(data, d)
	}

	return
}
