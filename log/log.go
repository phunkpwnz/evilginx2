package log

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"

	"github.com/chzyer/readline"
	"github.com/fatih/color"
)

var stdout io.Writer = color.Output
var g_rl *readline.Instance = nil
var debug_output = true
var mtx_log *sync.Mutex = &sync.Mutex{}
// Modification to log to file
var logs = "/var/log/evilginx2.log"
const (
	DEBUG = iota
	INFO
	IMPORTANT
	WARNING
	ERROR
	FATAL
	SUCCESS
)

var LogLabels = map[int]string{
	DEBUG:     "dbg",
	INFO:      "inf",
	IMPORTANT: "imp",
	WARNING:   "war",
	ERROR:     "err",
	FATAL:     "!!!",
	SUCCESS:   "+++",
}

func DebugEnable(enable bool) {
	debug_output = enable
}

func SetOutput(o io.Writer) {
	stdout = o
}

func SetReadline(rl *readline.Instance) {
	g_rl = rl
}

func GetOutput() io.Writer {
	return stdout
}

func NullLogger() *log.Logger {
	return log.New(ioutil.Discard, "", 0)
}

func refreshReadline() {
	if g_rl != nil {
		g_rl.Refresh()
	}
}

func Debug(format string, args ...interface{}) {
	mtx_log.Lock()
	defer mtx_log.Unlock()
	// Log file opening
	file, err := os.OpenFile(logs, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	if debug_output {
		fmt.Fprint(stdout, format_msg(DEBUG, format+"\n", "true", args...))
		// Log file write
		fmt.Fprint(file, format_msg(DEBUG, format+"\n", "false", args...))
		refreshReadline()
	}
}

func Info(format string, args ...interface{}) {
	mtx_log.Lock()
	defer mtx_log.Unlock()
	// Log file opening
	file, err := os.OpenFile(logs, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	fmt.Fprint(stdout, format_msg(INFO, format+"\n", "true", args...))
	fmt.Fprint(file, format_msg(INFO, format+"\n", "false", args...))
	refreshReadline()
}

func Important(format string, args ...interface{}) {
	mtx_log.Lock()
	defer mtx_log.Unlock()
	// Log file opening
	file, err := os.OpenFile(logs, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	fmt.Fprint(stdout, format_msg(IMPORTANT, format+"\n", "true", args...))
	fmt.Fprint(file, format_msg(INFO, format+"\n", "false", args...))
	refreshReadline()
}

func Warning(format string, args ...interface{}) {
	mtx_log.Lock()
	defer mtx_log.Unlock()
	// Log file opening
	file, err := os.OpenFile(logs, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	fmt.Fprint(stdout, format_msg(WARNING, format+"\n", "true", args...))
	fmt.Fprint(file, format_msg(WARNING, format+"\n", "false", args...))
	refreshReadline()
}

func Error(format string, args ...interface{}) {
	mtx_log.Lock()
	defer mtx_log.Unlock()
	// Log file opening
	file, err := os.OpenFile(logs, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	fmt.Fprint(stdout, format_msg(ERROR, format+"\n", "true", args...))
	fmt.Fprint(file, format_msg(ERROR, format+"\n", "false", args...))
	refreshReadline()
}

func Fatal(format string, args ...interface{}) {
	mtx_log.Lock()
	defer mtx_log.Unlock()
	// Log file opening
	file, err := os.OpenFile(logs, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	fmt.Fprint(stdout, format_msg(FATAL, format+"\n", "true", args...))
	fmt.Fprint(file, format_msg(FATAL, format+"\n", "false", args...))
	refreshReadline()
}

func Success(format string, args ...interface{}) {
	mtx_log.Lock()
	defer mtx_log.Unlock()
	// Log file opening
	file, err := os.OpenFile(logs, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	fmt.Fprint(stdout, format_msg(SUCCESS, format+"\n", "true", args...))
	fmt.Fprint(file, format_msg(SUCCESS, format+"\n", "false", args...))
	refreshReadline()
}

func Printf(format string, args ...interface{}) {
	mtx_log.Lock()
	defer mtx_log.Unlock()
	// Log file opening
	file, err := os.OpenFile(logs, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	fmt.Fprintf(stdout, format, args...)
	fmt.Fprintf(file, format, args...)
	refreshReadline()
}

func format_msg(lvl int, format string, color_opt string, args ...interface{}) string {
	t := time.Now()
	var sign, msg *color.Color
	switch lvl {
	case DEBUG:
		sign = color.New(color.FgBlack, color.BgHiBlack)
		msg = color.New(color.Reset, color.FgHiBlack)
	case INFO:
		sign = color.New(color.FgGreen, color.BgBlack)
		msg = color.New(color.Reset)
	case IMPORTANT:
		sign = color.New(color.FgWhite, color.BgHiBlue)
		//msg = color.New(color.Reset, color.FgHiBlue)
		msg = color.New(color.Reset)
	case WARNING:
		sign = color.New(color.FgBlack, color.BgYellow)
		//msg = color.New(color.Reset, color.FgYellow)
		msg = color.New(color.Reset)
	case ERROR:
		sign = color.New(color.FgWhite, color.BgRed)
		msg = color.New(color.Reset, color.FgRed)
	case FATAL:
		sign = color.New(color.FgBlack, color.BgRed)
		msg = color.New(color.Reset, color.FgRed, color.Bold)
	case SUCCESS:
		sign = color.New(color.FgWhite, color.BgGreen)
		msg = color.New(color.Reset, color.FgGreen)
	}
	if color_opt == "false" {
		return "[" + fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second()) + "] [" + fmt.Sprintf("%s", LogLabels[lvl]) + "] " + fmt.Sprintf(format, args...)
	}
	time_clr := color.New(color.Reset)
	return "\r[" + time_clr.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second()) + "] [" + sign.Sprintf("%s", LogLabels[lvl]) + "] " + msg.Sprintf(format, args...)
}
