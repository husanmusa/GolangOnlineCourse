package main

import (
	"log"
	"os"
	"time"
)

var loggerFileName = "logger.text"
var (
	InfoLogger  *log.Logger
	WarnLogger  *log.Logger
	DebugLogger *log.Logger
	ErrorLogger *log.Logger
	FatalLogger *log.Logger
)

// Logging the information in logger file for readability purpose.
func init() {
	logFile, err := os.OpenFile(loggerFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Unable to create Logger file:", err.Error())
		return
	}

	log.SetOutput(logFile)
	InfoLogger = log.New(logFile, "Info:", log.Ldate|log.Ltime|log.Lshortfile)
	WarnLogger = log.New(logFile, "Warn:", log.Ldate|log.Ltime|log.Llongfile)
	DebugLogger = log.New(logFile, "Debug:", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	ErrorLogger = log.New(logFile, "Error:", log.Ldate|log.Ltime|log.Lshortfile)
	FatalLogger = log.New(logFile, "Fatal:", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	InfoLogger.Println("Starting Service!")
	t := time.Now()
	InfoLogger.Printf("Time taken: %s \n", time.Since(t))
	WarnLogger.Print("Something WARN")
	ErrorLogger.Println("Error: Something went wrong!")
}
