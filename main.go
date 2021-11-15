package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"

	"github.com/apsdehal/go-logger"
)

type Person struct {
	Id         string `json:"firstname"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Avatar     string `json:"avatar"`
	IP_Address string `json:"ip_address"`
}

func main() {

	log, err := logger.New("csv-go-parser")
	if err != nil {
		panic(err) // Check for error
	}

	csvFile, _ := os.Open("MOCK_DATA.csv")
	reader := csv.NewReader(csvFile)
	var people []Person
	for {

		line, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.FatalF("File reader error", err)
		}

		people = append(people, Person{
			Id:         line[0],
			FirstName:  line[1],
			LastName:   line[2],
			Email:      line[3],
			Avatar:     line[4],
			IP_Address: line[5],
		})

	}

	file, _ := json.MarshalIndent(people, "", "")
	if fileError := ioutil.WriteFile("csvjson_output.json", file, 0644); fileError != nil {
		log.Errorf("json out failed", fileError)
	}

	log.Debug(string(file))
}
