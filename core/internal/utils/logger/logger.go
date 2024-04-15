package logger

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"

	sdkpaths "github.com/flarehotspot/sdk/utils/paths"

	"github.com/flarehotspot/sdk/utils/wsv"
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

const (
	FLARELOG_METADATA_COUNT = 10
)

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

func openLogFile() (*os.File, error) {
	logdir := "/" + sdkpaths.TmpDir + "/logs"

	// ensure log file directory exists
	err := os.MkdirAll(logdir, 0700)
	if err != nil {
		log.Println("Error creating log directory", "error", err)
		return nil, err
	}

	// opening/creating log file
	logFile, err := os.OpenFile(logdir+"/"+LogFilename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Println("Error creating log file", "error", err)
		return nil, err
	}

	return logFile, nil
}

func colorize(colorCode int, v string) string {
	return fmt.Sprintf("\033[%sm%s%s", strconv.Itoa(colorCode), v, reset)
}

func GetCallerFileLine(calldepth int) (file string, line int) {
	calldepth++

	_, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		log.Println("Cannot retrieve caller")
	}

	return
}

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func GetLogLines() int {
	// get app logs file path
	logdir := "/" + sdkpaths.TmpDir + "/logs"

	// open logs
	file, err := os.Open(logdir + "/" + LogFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// get log's lines count
	logLines, err := lineCounter(file)
	if err != nil {
		log.Fatal("error counting lines", err)
	}

	return logLines
}

// reverse file scanner
type ReverseScanner struct {
	r   io.ReaderAt
	pos int
	err error
	buf []byte
}

func NewReverseScanner(r io.ReaderAt, pos int) *ReverseScanner {
	return &ReverseScanner{r: r, pos: pos}
}

func (s *ReverseScanner) readMore() {
	if s.pos == 0 {
		s.err = io.EOF
		return
	}
	size := 1024
	if size > s.pos {
		size = s.pos
	}
	s.pos -= size
	buf2 := make([]byte, size, size+len(s.buf))

	// ReadAt attempts to read full buff!
	_, s.err = s.r.ReadAt(buf2, int64(s.pos))
	if s.err == nil {
		s.buf = append(buf2, s.buf...)
	}
}

func (s *ReverseScanner) Line() (line string, start int, err error) {
	if s.err != nil {
		return "", 0, s.err
	}
	for {
		lineStart := bytes.LastIndexByte(s.buf, '\n')
		if lineStart >= 0 {
			// We have a complete line:
			var line string
			line, s.buf = string(dropCR(s.buf[lineStart+1:])), s.buf[:lineStart]
			return line, s.pos + lineStart + 1, nil
		}
		// Need more data:
		s.readMore()
		if s.err != nil {
			if s.err == io.EOF {
				if len(s.buf) > 0 {
					return string(dropCR(s.buf)), 0, nil
				}
			}
			return "", 0, s.err
		}
	}
}

// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

func ReadLogsReverse() ([]map[string]any, error) {
	var logs []map[string]any

	logdir := "/" + sdkpaths.TmpDir + "/logs"
	file, err := os.Open(logdir + "/" + LogFilename)
	if err != nil {
		log.Fatal(err)
	}
	fi, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reverseScanner := NewReverseScanner(file, int(fi.Size()))

	// read the empty log line
	_, _, err = reverseScanner.Line()
	if err != io.EOF {
		log.Fatal(err)
	}

	for {
		line, _, err := reverseScanner.Line()

		// done reading the file
		if err == io.EOF {
			break
		}

		// if something goes wrong in reading
		if err != nil {
			log.Fatal("Error:", err)
			break
		}

		// read of line successful
		dataInLine, err := wsv.ParseLineAsArray(line)
		if err != nil {
			log.Println("error parsing raw log file to wsv: ", err)
			return nil, err
		}

		fmt.Println("current data line", dataInLine)

		parsedlog, err := parseLog(dataInLine)
		if err != nil {
			log.Println("error parsing log file to flarelog format: ", err)
			return nil, err
		}

		logs = append(logs, parsedlog)
	}

	return logs, nil
}

// ----

func ReadLogs(start int, end int) ([]map[string]any, error) {
	var logs []map[string]any

	// get app logs file path
	logdir := "/" + sdkpaths.TmpDir + "/logs"

	// open logs
	file, err := os.Open(logdir + "/" + LogFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rd := bufio.NewReader(file)

	currLine := 0

	// TODO: make it concurrent
	for {
		l, err := rd.ReadString('\n')

		if currLine < start {
			currLine++
			continue
		}

		// file has no content left
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		// read of line successful
		dataInLine, err := wsv.ParseLineAsArray(l)
		if err != nil {
			log.Println("error parsing raw log file to wsv: ", err)
			return nil, err
		}

		parsedlog, err := parseLog(dataInLine)
		if err != nil {
			log.Println("error parsing log file to flarelog format: ", err)
			return nil, err
		}

		logs = append(logs, parsedlog)

		if currLine >= end {
			break
		}

		currLine++
	}

	return logs, nil
}

func parseLog(logLine []string) (map[string]any, error) {
	logLength := len(logLine)

	// check if valid flare log file
	if logLength < FLARELOG_METADATA_COUNT {
		return nil, errors.New("invalid flarelog format")
	}

	// get file/packages
	var pkgs []string

	pathRaw := logLine[9] // raw file path
	j := 0
	for i := 0; i < len(pathRaw); i++ {
		if pathRaw[i] == '/' {
			pkgs = append(pkgs, pathRaw[j:i])
			j = i + 1
			continue
		}
	}
	pkgs = append(pkgs, pathRaw[j:])

	plugin := pkgs[2]
	filename := pkgs[len(pkgs)-1]
	filepluginpath := strings.Join(pkgs[3:], "/")

	var body any
	// check if log has body
	if logLength > FLARELOG_METADATA_COUNT {
		body = logLine[FLARELOG_METADATA_COUNT+1:]
	}

	log := map[string]any{
		"level":          logLine[0],
		"title":          logLine[1],
		"year":           logLine[2],
		"month":          logLine[3],
		"day":            logLine[4],
		"hour":           logLine[5],
		"min":            logLine[6],
		"sec":            logLine[7],
		"nano":           logLine[8],
		"fullpath":       logLine[9],
		"plugin":         plugin,
		"filepluginpath": filepluginpath,
		"filename":       filename,
		"line":           logLine[10],
		"body":           body,
	}

	return log, nil
}

func LogToConsole(file string, line int, level int, title string, body ...any) {
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

func LogToFile(file string, line int, level int, title string, body ...any) {
	// log file format:
	// level title YYYY M d H m s n file line
	// "key" "value" // if any
	// --

	f, err := openLogFile()
	if err != nil {
		log.Println("Failed to create log file", "error", err)
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
