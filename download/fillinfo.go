package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	// Read CSV from stdin.
	in, err := ioutil.ReadFile("courses-annotated.csv")
	if err != nil {
		log.Fatalf("Error opening courses-annotated.csv: %s", err)
	}
	r := csv.NewReader(bytes.NewReader(in))
	w := csv.NewWriter(os.Stdout)

	for {
		// For every row
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error parsing CSV: %s", err)
		}

		// If record already has a Field, then write it back verbatim
		if record[3] != "" {
			w.Write(record)
			continue
		}

		// Run extraction script w/ URL
		fmt.Fprintf(os.Stderr, "Extracting info for: %s\n", record[2])
		cmd := exec.Command("bash", "./extract_info.sh", record[2])
		out, err := cmd.Output()
		// Ensure there is no error
		if err != nil || len(out) == 0 {
			fmt.Fprintf(os.Stderr, "Error extracting %s: %s", record[2], err)
			w.Write(record)
			continue
		}
		//fmt.Fprintf(os.Stderr, "Output: %s", out)
		// Extract 3 strings from output and repopulate CSV data
		info := bytes.Split(out, []byte("\n"))
		//fmt.Fprintf(os.Stderr, "Field: %s\n", info[0])
		record[3] = strings.Trim(string(info[0]), "\"")
		//fmt.Fprintf(os.Stderr, "Name of University/Institute: %s\n", info[2])
		record[4] = strings.Trim(string(info[2]), "\"")
		//fmt.Fprintf(os.Stderr, "Introduction: %s", info[1])
		record[5] = strings.Trim(string(info[1]), "\"")
		record[5] = strings.Replace(record[5], "\\n", "\n", -1)
		record[5] = strings.Replace(record[5], "\\t", "\t", -1)
		time.Sleep(500 * time.Millisecond)

		// Write back CSV entry to stdout
		w.Write(record)
	}
	w.Flush()
}
