package store

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/smvfal/metrics-processor/pkg/types"
)

func WriteMessage(data []byte) {

	var message types.Message
	err := json.Unmarshal(data, &message)
	if err != nil {
		log.Fatal(err)
	}

	functions := message.Functions
	nodes := message.Nodes

	for _, function := range functions {

		header := [][]string{{
			"timestamp",
			"replicas",
			"invocation_rate",
			"throughput",
			"processing_time",
			"response_time",
			"cold_start",
			"cpu",
			"mem",
		}}
		file := openFile("./data/"+function.Name+".csv", header)

		functionCpu := 0.0
		if function.Cpu != nil {
			for _, podCpu := range function.Cpu {
				functionCpu += podCpu
			}
			functionCpu /= float64(len(function.Cpu))
		}
		functionMem := 0.0
		if function.Mem != nil {
			for _, podMem := range function.Mem {
				functionMem += podMem
			}
			functionMem /= float64(len(function.Mem))
		}

		row := [][]string{{
			strconv.FormatInt(message.Timestamp, 10),
			strconv.FormatInt(int64(function.Replicas), 10),
			strconv.FormatFloat(function.InvocationRate, 'f', 4, 64),
			strconv.FormatFloat(function.Throughput, 'f', 4, 64),
			strconv.FormatFloat(function.ProcessingTime, 'f', 4, 64),
			strconv.FormatFloat(function.ResponseTime, 'f', 4, 64),
			strconv.FormatFloat(function.ColdStart, 'f', 4, 64),
			strconv.FormatFloat(functionCpu, 'f', 4, 64),
			strconv.FormatFloat(functionMem, 'f', 4, 64),
		}}

		// write row
		writeRecords(file, row)

		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}

	for _, node := range nodes {

		header := [][]string{{"timestamp", "cpu", "mem"}}
		file := openFile("./data/"+node.Name+".csv", header)

		row := [][]string{{
			strconv.FormatInt(message.Timestamp, 10),
			strconv.FormatFloat(node.Cpu, 'f', 4, 64),
			strconv.FormatFloat(node.Mem, 'f', 4, 64),
		}}

		// write row
		writeRecords(file, row)

		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}
}

func openFile(fileName string, header [][]string) *os.File {

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		// If the file doesn't exist, create it and write the header
		if os.IsNotExist(err) {
			log.Printf("File %v does not exist. Creating it...", fileName)
			file, err = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal(err)
			}
			// write the header
			writeRecords(file, header)

		} else {
			log.Fatal(err)
		}
	}
	return file
}

func writeRecords(file *os.File, records [][]string) {
	writer := csv.NewWriter(file)
	err := writer.WriteAll(records) // calls Flush internally
	if err != nil {
		log.Fatal(err)
	}
}
