package plugins

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"time"

	sdkpaths "github.com/flarehotspot/sdk/utils/paths"
	wsv "github.com/flarehotspot/sdk/utils/wsv"
)

const (
	reset = "\033[0m"

	darkGray    = 90
	lightRed    = 91
	lightYellow = 93
	lightBlue   = 94
)

const (
	LogFilename = "app.log"
	Infoprefix  = "[INFO] "
	Debugprefix = "[DEBUG] "
	Errorprefix = "[ERROR] "
)

const (
	Info  = 0
	Debug = 1
	Error = 2
)

type LoggerApi struct{}

func NewLoggerApi() *LoggerApi {
	// TODO: create log rotation
	// run a go routine for log rotation

	// TODO: create log retention
	// run a go routine for log retention

	return &LoggerApi{}
}

func itoa(i int, wid int) int {
	num := i
	d := 1

	for i >= 10 {
		q := i / 10
		i = q
		d++
	}

	return num / int(math.Pow10(d-wid))
}

func getLevelAsStr(level int) string {
	switch level {
	case 0:
		return "INFO"
	case 1:
		return "DEBUG"
	case 2:
		return "ERROR"
	}

	return "INFO"
}

func colorizeLevel(level int) string {
	var color int
	switch level {
	case 0:
		color = lightBlue
	case 1:
		color = lightYellow
	case 2:
		color = lightRed
	}
	return colorize(color, getLevelAsStr(level))
}

var std = NewLoggerApi()

func openLogFile() (*os.File, error) {
	logdir := "/" + sdkpaths.TmpDir + "/logs"

	// ensure log file directory exists
	err := os.MkdirAll(logdir, 0700)
	if err != nil {
		std.Info("Error creating log directory", "error", err)
		return nil, err
	}

	// opening/creating log file
	logFile, err := os.OpenFile(logdir+"/"+LogFilename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		std.Error("Error creating log file", "error", err)
		return nil, err
	}

	return logFile, nil
}

func colorize(colorCode int, v string) string {
	return fmt.Sprintf("\033[%sm%s%s", strconv.Itoa(colorCode), v, reset)
}

// TODO: read logs
func ReadLogs() {
	// TODO : this method should only be accessible by the core
	// It should handle file renaming in case of log rotation
	logdir := "/" + sdkpaths.TmpDir + "/logs"

	file, err := os.Open(logdir + "/" + LogFilename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rd := bufio.NewReader(file)

	for {
		l, err := rd.ReadString('\n')

		if err == io.EOF {
			fmt.Println(l)
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(l)
	}

	// scanner := bufio.NewScanner(file)
	// scanner.Split(bufio.ScanLines)
	// var txtlines []string

	// for scanner.Scan() {
	// 	// txtlines = append(txtlines, scanner.Text())

	// 	level, title, year, month, day, hour, min, sec, nano, file, line, body, err := ParseLogLine(scanner.Text())
	// 	if err != nil {
	// 		fmt.Println("level", level)
	// 		fmt.Println("title", title)
	// 		fmt.Println("year", year)
	// 		fmt.Println("month", month)
	// 		fmt.Println("day", day)
	// 		fmt.Println("hour", hour)
	// 		fmt.Println("min", min)
	// 		fmt.Println("sec", sec)
	// 		fmt.Println("nano", nano)
	// 		fmt.Println("file", file)
	// 		fmt.Println("line", line)
	// 		fmt.Println("body", body)
	// 	}
	// }
}

// TODO : parse read lines
func ParseLogLine(logLine string) (level int, title string, year int, month int, day int, hour int, min int, sec int, nano int, file string, fileline int, body []string, err error) {
	values, err := wsv.ParseLineAsArray(logLine)

	if err != nil {
		return
	}

	if len(values) >= 11 {
		body = values[11:]
	}

	return
}

func (l *LoggerApi) Info(title string, body ...any) error {
	calldepth := 1
	level := 0

	file, line := getCallerFileLine(calldepth)

	logToConsole(file, line, level, title, body...)
	logToFile(file, line, level, title, body...)

	return nil
}

func (l *LoggerApi) Debug(title string, body ...any) error {
	// TODO : remove this function
	// just for testing
	ReadLogs()

	calldepth := 1
	level := 1

	file, line := getCallerFileLine(calldepth)

	logToConsole(file, line, level, title, body...)
	logToFile(file, line, level, title, body...)

	return nil
}

func (l *LoggerApi) Error(title string, body ...any) error {
	calldepth := 1
	level := 2

	file, line := getCallerFileLine(calldepth)

	logToConsole(file, line, level, title, body...)
	logToFile(file, line, level, title, body...)

	return nil
}

func getCallerFileLine(calldepth int) (file string, line int) {
	calldepth++

	_, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		std.Error("Cannot retrieve caller")
	}

	// short := file
	// for i := len(file) - 1; i > 0; i-- {
	// 	if file[i] == '/' {
	// 		short = file[i+1:]
	// 		break
	// 	}
	// }
	// file = short

	return
}

func logToConsole(file string, line int, level int, title string, body ...any) {
	// date and time data
	now := time.Now()
	hour, min, sec := now.Clock()
	year, month, day := now.Date()
	nano := itoa(now.Nanosecond(), 3)

	metadata := fmt.Sprintf("[%s:%d %d/%d/%d %d:%d:%d.%d]", file, line, year, month, day, hour, min, sec, nano)
	content := colorize(darkGray, metadata)
	content = fmt.Sprintf("%s\n%s %s", content, colorizeLevel(level), title)

	// adding all body key-value pairs if any
	for i, arg := range body {
		value := reflect.ValueOf(arg)
		var str string

		switch value.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			str = fmt.Sprintf("%d", value.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			str = fmt.Sprintf("%d", value.Uint())
		case reflect.Float32, reflect.Float64:
			str = fmt.Sprintf("%f", value.Float())
		case reflect.String:
			str = value.String()
		// Add cases for other types as needed
		default:
			str = fmt.Sprintf("%v", arg)
		}

		// if i is last and is even,
		// means that the value is not given
		if i == len(body)-1 && i%2 == 0 {
			content = fmt.Sprintf("%s\n  \"%s\": -", content, str)
			break
		}

		// if i is key
		if i%2 == 0 {
			content = fmt.Sprintf("%s\n  \"%v\": ", content, str)
			continue
		}

		// if i is value
		content = fmt.Sprintf("%s\"%s\"", content, str)
	}

	fmt.Println(content)
}

func logToFile(file string, line int, level int, title string, body ...any) {
	// log file format:
	// level title YYYY M d H m s n file line
	// "key" "value" // if any
	// --

	f, err := openLogFile()
	if err != nil {
		std.Error("Failed to create log file", "error", err)
		panic(err)
	}
	defer f.Close()

	var content [][]string

	// date and time data
	now := time.Now()
	hour, min, sec := now.Clock()
	year, month, day := now.Date()
	nano := itoa(now.Nanosecond(), 3)

	var logInfo []string
	logInfo = append(logInfo, strconv.Itoa(level), title, strconv.Itoa(year), strconv.Itoa(int(month)), strconv.Itoa(day), strconv.Itoa(hour), strconv.Itoa(min), strconv.Itoa(sec), strconv.Itoa(nano), file, strconv.Itoa(line))

	for _, arg := range body {
		value := reflect.ValueOf(arg)
		var str string

		switch value.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			str = fmt.Sprintf("%d", value.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			str = fmt.Sprintf("%d", value.Uint())
		case reflect.Float32, reflect.Float64:
			str = fmt.Sprintf("%f", value.Float())
		case reflect.String:
			str = value.String()
		// Add cases for other types as needed
		default:
			str = fmt.Sprintf("%v", arg)
		}

		logInfo = append(logInfo, str)
	}

	content = append(content, logInfo)

	f.WriteString(wsv.Serialize(content) + "\n")
}
