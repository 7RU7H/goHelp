


// InfoLogger.Printf("Something noteworthy happened\n")
// WarningLogger.Printf("There is something you should know about\n")
// ErrorLogger.Printf("Something went wrong\n")
func initaliseLogging() error {
	now := time.Now().UTC()
	dateFormatted := now.Format("2006-01-01")
	nameBuilder := strings.Builder{}
	nameBuilder.WriteString(dateFormatted)
	nameBuilder.WriteString(".log")
	file, err := os.OpenFile(nameBuilder.String(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0667)
	checkError(err, 0, 0) // No logging = sadness

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	return nil
}

InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
WarningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
ErrorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)