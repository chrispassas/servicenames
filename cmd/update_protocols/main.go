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
	var url = "https://www.iana.org/assignments/protocol-numbers/protocol-numbers-1.csv"

	var resp *http.Response
	var err error
	if resp, err = http.Get(url); err != nil {
		log.Fatalf("http.Get() error:%v", err)
	}

	if resp.StatusCode != 200 {
		var data []byte
		if data, err = io.ReadAll(resp.Body); err != nil {
			log.Fatalf("io.ReadAll() error:%v", err)
		}
		defer resp.Body.Close()
		log.Fatalf("resp.StatusCode:%d body:%s", resp.StatusCode, string(data))
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

	var protocols []servicenames.Protocol

	var badProtocol = 0
	for lineNumber, record := range records {
		// if lineNumber == 10 {
		// 	goto End
		// }
		if lineNumber == 0 {
			log.Printf("columns:%d", len(record))
			log.Printf("CSV Headers:%v", record)
		} else {

			var protocol int
			if protocol, err = strconv.Atoi(record[0]); err != nil {
				badProtocol++
				continue
			}

			protocols = append(protocols, servicenames.Protocol{
				Proto:   uint8(protocol),
				Keyword: strings.ToLower(strings.TrimSpace(record[1])),
			})
		}

	}
	// End:

	log.Printf("badProtocol count:%d", badProtocol)

	// log.Printf("services:%#v", services)

	var dataFileOutput = "package servicenames\n\n"

	dataFileOutput += "var protocolData = []Protocol{\n"

	for _, p := range protocols {
		dataFileOutput += fmt.Sprintf("\t{Proto:%d, Keyword:\"%s\"},\n",
			p.Proto, p.Keyword,
		)
	}

	dataFileOutput += "}"

	var formattedCode []byte
	if formattedCode, err = format.Source([]byte(dataFileOutput)); err != nil {
		log.Fatalf("format.Source() error:%v dataFileOutput:%s", err, dataFileOutput)
	}

	var f *os.File
	if f, err = os.OpenFile("../../protocoldata.go", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644); err != nil {
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
