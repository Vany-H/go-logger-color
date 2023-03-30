// Packages wrapper for log messages and errors
//
// func: Log, Error, LogTime, Warning
package logger

// import fmt, os, time modules
import (
	"fmt"
	"log"
	"os"
	"time"
)

// Log message in green color with date and pid.
func Log(message string, param ...any) {
	mess := fmt.Sprintf(message, param...)
	logMessage := fmt.Sprintf(
		"\x1b[32m[Go] %v –	LOG\x1b[33m:\x1b[32m %v\x1b[0m",
		os.Getegid(), // Get id of current process
		mess,
	)
	log.Print(logMessage)
}

// Log any what you want in green color with date and pid.
func LogAny(objs ...any) {
	mess := ""
	for i := 0; i < len(objs); i++ {
		mess = fmt.Sprintln(objs...)
	}
	logMessage := fmt.Sprintf(
		"\x1b[32m[Go] %v –	LOG\x1b[33m:\x1b[32m %v\x1b[0m",
		os.Getegid(), // Get id of current process
		mess,
	)
	log.Print(logMessage)
}

// Log current time
func LogTime() {
	currentTime := time.Now()
	Log(currentTime.Format("2006.01.02, 15:04:05"))
}

// Log error messages in red color with date and pid.
func Error(err interface{}) {
	logMessage := fmt.Sprintf(
		"\x1b[31m[Go] %v –	\x1b[31mERROR %v\x1b[0m",
		os.Getegid(),
		err,
	)
	log.Print(logMessage)
}

// Log error messages in red color with date and pid.
func FatalError(err interface{}) {
	logMessage := fmt.Sprintf(
		"\x1b[31m[Go] %v –	\x1b[31mERROR %v\x1b[0m",
		os.Getegid(),
		err,
	)
	log.Fatal(logMessage)
}

// Log error messages in yellow color with date and pid.
func Warning(message string, param ...any) {
	mess := fmt.Sprintf(message, param...)
	logMessage := fmt.Sprintf(
		"\x1b[33m[Go] %v –	WARNING %v\x1b[0m",
		os.Getegid(),
		mess,
	)
	log.Print(logMessage)
}

// Shields message in green color
func GreenString(mess any) string {
	message := fmt.Sprintln(mess)
	return fmt.Sprintf("\x1b[32m%v\x1b[0m", message)
}

// Shields message in red color
func RedString(mess any) string {
	message := fmt.Sprintln(mess)
	return fmt.Sprintf("\x1b[31m%v\x1b[0m", message)
}

// Shields message in yellow color
func YellowString(mess any) string {
	message := fmt.Sprintln(mess)
	return fmt.Sprintf("\x1b[33m%v\x1b[0m", message)
}

// Shields message in blue color
func BlueString(mess any) string {
	message := fmt.Sprintln(mess)
	return fmt.Sprintf("\x1b[36m%v\x1b[0m", message)
}
