package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"go/format"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/chrispassas/servicenames"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	var url = "https://www.iana.org/assignments/service-names-port-numbers/service-names-port-numbers.csv"

	var resp *http.Response
	var err error
	if resp, err = http.Get(url); err != nil {
		log.Fatalf("http.Get() error:%v", err)
	}

	if resp.StatusCode != 200 {
		log.Fatalf("resp.StatusCode:%d", resp.StatusCode)
	}

	var data []byte
	if data, err = io.ReadAll(resp.Body); err != nil {
		log.Fatalf("io.ReadAll() error:%v", err)
	}
	defer resp.Body.Close()

	r := bytes.NewReader(data)

	csvReader := csv.NewReader(r)

	var records [][]string
	if records, err = csvReader.ReadAll(); err != nil {
		log.Fatalf("csvReader.ReadAll() error:%v", err)
	}

	var services []servicenames.Service

	var badPort = 0
	var serviceMissingCount = 0
	var protocol int
	for lineNumber, record := range records {

		if lineNumber == 0 {
			log.Printf("columns:%d", len(record))
			log.Printf("CSV Headers:%v", record)
		} else {

			var port int
			if port, err = strconv.Atoi(record[1]); err != nil {
				badPort++
				continue
			}

			if len(record[0]) == 0 {
				serviceMissingCount++
				continue
			}

			switch strings.ToLower(strings.TrimSpace(record[2])) {
			case "tcp":
				protocol = 6
			case "udp":
				protocol = 17
			case "sctp":
				protocol = 132
			case "dccp":
				protocol = 33
			}

			// log.Printf("service:%s", record[0])
			// log.Printf("port:%d", port)
			// log.Printf("transport:%s", record[2])
			// log.Printf("description:%s", record[3])
			// log.Printf("line:%d------------------------------------", lineNumber)
			description := strings.ToLower(strings.TrimSpace(record[3]))
			description = strings.ReplaceAll(description, "\n", " ")
			services = append(services, servicenames.Service{
				Protocol:    uint8(protocol),
				Port:        uint16(port),
				Name:        strings.ToLower(strings.TrimSpace(record[0])),
				Description: description,
			})
		}

	}

	log.Printf("badPort count:%d", badPort)
	log.Printf("serviceMissingCount:%d", serviceMissingCount)

	// log.Printf("services:%#v", services)

	var dataFileOutput = "package servicenames\n\n"

	dataFileOutput += "var serviceData = []Service{\n"

	for _, s := range services {
		dataFileOutput += fmt.Sprintf("\t{Protocol:%d, Port:%d, Name:\"%s\", Description:`%s`},\n",
			s.Protocol, s.Port, s.Name, s.Description,
		)
	}

	dataFileOutput += "}"

	var formattedCode []byte
	if formattedCode, err = format.Source([]byte(dataFileOutput)); err != nil {
		log.Fatalf("format.Source() error:%v", err)
	}

	var f *os.File
	if f, err = os.OpenFile("../../servicedata.go", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644); err != nil {
		log.Fatalf("os.OpenFile() error:%v", err)
	}
	defer f.Close()

	if _, err = f.Write(formattedCode); err != nil {
		log.Fatalf("f.Write() error:%v", err)
	}

	// csvLines := bytes.Split(data, []byte("\n"))

	// for lineNumber, line := range csvLines {
	// 	if lineNumber == 0 {

	// 	} else {
	// 		lineStr := string(line)
	// 		strings.Split(lineStr, sep)
	// 		csv.
	// 	}
	// }

	/*
		Stream download xml
		write out Service{} slice into data.go (overwrite the whole file)
	*/
}
