package logger

/*
https://github.com/donnie4w
*/
import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	_VER string = "1.0.2"
)

const DATE_FORMAT = "2006-01-02"

type UNIT int64

const (
	_       = iota
	KB UNIT = 1 << (iota * 10)
	MB
	GB
	TB
)

type LEVEL int32

const (
	ALL LEVEL = iota
	DEBUG
	INFO
	WARN
	ERROR

	FATAL
	OFF
)

var (
	debugStr = map[LEVEL]string{
		DEBUG: "[D] ",
		INFO:  "[I] ",
		WARN:  "[\033[036;1mW\033[036;0m] ",
		ERROR: "[\033[031;1mE\033[031;0m] ",
		FATAL: "[\033[031;1mF\033[031;0m] ",
	}

	logLevel        LEVEL = 1
	maxFileSize     int64
	maxFileCount    int32
	dailyRolling    = true
	consoleAppender = true
	rollingFile     = false
	filePrefix      = ""
	logObj          *_FILE
	callDepth       = 3
	logFlag         = log.Ldate | log.Ltime | log.Llongfile
)

const (
	_ = iota
	ROLLINGDAILY
	ROLLINGFILE
)

func init() {
	if runtime.GOOS == "windows" {
		debugStr[WARN] = "[W] "
		debugStr[ERROR] = "[E] "
		debugStr[FATAL] = "[F] "
	}
}

type _FILE struct {
	dir      string
	filename string
	_suffix  int
	isCover  bool
	_date    *time.Time
	mu       *sync.RWMutex
	logfile  *os.File
	lg       *log.Logger
}

func SetPrefix(title string) {
	log.SetPrefix(title)
}

func SetLogFlag(flag int) {
	logFlag = flag
	if logObj != nil {
		logObj.lg.SetFlags(logFlag)
	}
}

func SetConsole(isConsole bool) {
	consoleAppender = isConsole
}

func SetFilePrefix(path string) {
	filePrefix = path
}

func SetLevel(_level LEVEL) {
	logLevel = _level
}

func SetRollingFile(fileDir, fileName string, maxNumber int32, maxSize int64, _unit UNIT) {
	maxFileCount = maxNumber
	maxFileSize = maxSize * int64(_unit)
	rollingFile = true
	dailyRolling = false
	mkdirlog(fileDir)
	logObj = &_FILE{dir: fileDir, filename: fileName, isCover: false, mu: new(sync.RWMutex)}
	logObj.mu.Lock()
	defer logObj.mu.Unlock()
	for i := 1; i <= int(maxNumber); i++ {
		if isExist(fileDir + "/" + fileName + "." + strconv.Itoa(i)) {
			logObj._suffix = i
		} else {
			break
		}
	}
	if !logObj.isMustRename() {
		logObj.logfile, _ = os.OpenFile(fileDir+"/"+fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		// logObj.lg = log.New(logObj.logfile, "", log.Ldate|log.Ltime|log.Lshortfile)
		logObj.lg = log.New(logObj.logfile, "", logFlag)
	} else {
		logObj.rename()
	}
	go fileMonitor()
}

func SetRollingDaily(fileDir, fileName string) {
	rollingFile = false
	dailyRolling = true
	t, _ := time.Parse(DATE_FORMAT, time.Now().Format(DATE_FORMAT))
	mkdirlog(fileDir)
	logObj = &_FILE{dir: fileDir, filename: fileName, _date: &t, isCover: false, mu: new(sync.RWMutex)}
	logObj.mu.Lock()
	defer logObj.mu.Unlock()

	if !logObj.isMustRename() {
		logObj.logfile, _ = os.OpenFile(fileDir+"/"+fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		// logObj.lg = log.New(logObj.logfile, "", log.Ldate|log.Ltime|log.Lshortfile)
		logObj.lg = log.New(logObj.logfile, "", logFlag)
	} else {
		logObj.rename()
	}
}

func mkdirlog(dir string) (e error) {
	_, er := os.Stat(dir)
	b := er == nil || os.IsExist(er)
	if !b {
		if err := os.MkdirAll(dir, 0666); err != nil {
			if os.IsPermission(err) {
				fmt.Println("create dir error:", err.Error())
				e = err
			}
		}
	}
	return
}

func console(depth int, s string) {
	if consoleAppender {
		_, file, line, _ := runtime.Caller(depth)
		if filePrefix != "" {
			file = strings.Replace(file, filePrefix, ".", 1)
		}
		t := time.Now()
		fmt.Printf("[%s|%d]%s %s:%d\n", t.Format("2006-01-02 15:04:05"), t.Unix(), s, file, line)
	}
}

func catchError() {
	if err := recover(); err != nil {
		log.Println("err", err)
	}
}

func Debug(lvl LEVEL, fmtStr string, a ...interface{}) {
	if dailyRolling {
		fileCheck()
	}
	defer catchError()
	if logObj != nil {
		logObj.mu.RLock()
		defer logObj.mu.RUnlock()
	}

	if logLevel <= lvl {
		fmtStr = debugStr[lvl] + fmtStr
		logStr := fmt.Sprintf(fmtStr, a...)
		if logObj != nil {
			logObj.lg.Output(callDepth, logStr)
		}
		console(callDepth, logStr)
		if lvl == FATAL {
			os.Exit(1)
		}
	}
}

func D(fmtStr string, a ...interface{}) {
	Debug(DEBUG, fmtStr, a...)
}
func I(fmtStr string, a ...interface{}) {
	Debug(INFO, fmtStr, a...)
}
func W(fmtStr string, a ...interface{}) {
	Debug(WARN, fmtStr, a...)
}
func E(fmtStr string, a ...interface{}) {
	Debug(ERROR, fmtStr, a...)
}
func F(fmtStr string, a ...interface{}) {
	Debug(FATAL, fmtStr, a...)
}

func (f *_FILE) isMustRename() bool {
	if dailyRolling {
		t, _ := time.Parse(DATE_FORMAT, time.Now().Format(DATE_FORMAT))
		if t.After(*f._date) {
			return true
		}
	} else {
		if maxFileCount > 1 {
			if fileSize(f.dir+"/"+f.filename) >= maxFileSize {
				return true
			}
		}
	}
	return false
}

func (f *_FILE) rename() {
	if dailyRolling {
		fn := f.dir + "/" + f._date.Format(DATE_FORMAT) + "." + f.filename
		if !isExist(fn) && f.isMustRename() {
			if f.logfile != nil {
				f.logfile.Close()
			}
			err := os.Rename(f.dir+"/"+f.filename, fn)
			if err != nil {
				f.lg.Println("rename err", err.Error())
			}
			t, _ := time.Parse(DATE_FORMAT, time.Now().Format(DATE_FORMAT))
			f._date = &t
			f.logfile, _ = os.Create(f.dir + "/" + f.filename)
			// f.lg = log.New(logObj.logfile, "", log.Ldate|log.Ltime|log.Lshortfile)
			f.lg = log.New(logObj.logfile, "", logFlag)
		}
	} else {
		f.coverNextOne()
	}
}

func (f *_FILE) nextSuffix() int {
	return int(f._suffix%int(maxFileCount) + 1)
}

func (f *_FILE) coverNextOne() {
	f._suffix = f.nextSuffix()
	if f.logfile != nil {
		f.logfile.Close()
	}
	if isExist(f.dir + "/" + f.filename + "." + strconv.Itoa(int(f._suffix))) {
		os.Remove(f.dir + "/" + f.filename + "." + strconv.Itoa(int(f._suffix)))
	}
	os.Rename(f.dir+"/"+f.filename, f.dir+"/"+f.filename+"."+strconv.Itoa(int(f._suffix)))
	f.logfile, _ = os.Create(f.dir + "/" + f.filename)
	// f.lg = log.New(logObj.logfile, "", log.Ldate|log.Ltime|log.Lshortfile)
	f.lg = log.New(logObj.logfile, "", logFlag)
}

func fileSize(file string) int64 {
	fmt.Println("fileSize", file)
	f, e := os.Stat(file)
	if e != nil {
		fmt.Println(e.Error())
		return 0
	}
	return f.Size()
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func fileMonitor() {
	timer := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-timer.C:
			fileCheck()
		}
	}
}

func fileCheck() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	if logObj != nil && logObj.isMustRename() {
		logObj.mu.Lock()
		defer logObj.mu.Unlock()
		logObj.rename()
	}
}
