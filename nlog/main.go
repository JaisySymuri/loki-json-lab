package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type LogEntry struct {
	UserExperience  string `json:"userExperience"`
	TransactionName string `json:"transactionName"`
	EventTimestamp  string `json:"eventTimestamp"`
}

var userExperiences = []string{"NORMAL", "SLOW", "ERROR"}
var transactionNames = []string{"/jw/web/help/guide", "/jw/web/help/faq"}

// Generate timestamps from now -15 minutes, every 10 seconds, total 60
func genTime() []time.Time {
	var timestamps []time.Time
	start := time.Now().Add(-15 * time.Minute)

	for i := 0; i < 60; i++ {
		timestamps = append(timestamps, start.Add(time.Duration(i*10)*time.Second))
	}

	return timestamps
}

func main() {
	rand.Seed(time.Now().UnixNano())
	filename := fmt.Sprintf("logs-%s.jsonl", time.Now().Format("2006-01-02T15-04-05"))

	outputFile, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// Pre-generate ordered timestamps
	timestamps := genTime()

	for _, ts := range timestamps {
		log := LogEntry{
			UserExperience:  userExperiences[rand.Intn(len(userExperiences))],
			TransactionName: transactionNames[rand.Intn(len(transactionNames))],
			EventTimestamp:  ts.Format(time.RFC3339Nano),
		}

		line, _ := json.Marshal(log)
		outputFile.WriteString(string(line) + "\n")
	}

	fmt.Printf("Generated logs.jsonl with %d logs.\n", len(timestamps))
}
