package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

const (
	_VER string = "1.0.2"

	DateFormat     = "2006-01-02"
	LongDateFormat = "2006/01/02 15:04:05"
)

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
)

func init() {
	if runtime.GOOS == "windows" {
		debugStr[WARN] = "[W] "
		debugStr[ERROR] = "[E] "
		debugStr[FATAL] = "[F] "
	}
}

type LoggerObj struct {
	// 日志目录
	dir string
	// 日志文件名称
	filename string
	// 日志前缀
	prefix string
	// 日志后缀索引
	suffix  int
	date    *time.Time
	mu      *sync.RWMutex
	logfile *os.File
	lg      *log.Logger

	level          LEVEL
	maxFileSize    int64
	maxFileCount   int
	isDailyRolling bool
	isConsole      bool
	isRollingFile  bool
	callDepth      int
	flags          int
	once           sync.Once
}

func (m *LoggerObj) SetPrefix(prefix string) {
	m.prefix = prefix
	if m.lg != nil {
		m.lg.SetPrefix(prefix)
	}
}
func (m *LoggerObj) SetFlags(flag int) {
	m.flags = flag
	if m.lg != nil {
		m.lg.SetFlags(flag)
	}
}
func (m *LoggerObj) SetConsole(isConsole bool) {
	m.isConsole = isConsole
}
func (m *LoggerObj) SetLevel(level LEVEL) {
	m.level = level
}
func (m *LoggerObj) newLog() {
	m.lg = log.New(m.logfile, "", m.flags)
	if m.prefix != "" {
		m.lg.SetPrefix(m.prefix)
	}
}

func (m *LoggerObj) SetRollingFile(fileDir, fileName string, maxNumber int, maxSize int64, unit UNIT) {
	mkDir(fileDir)
	m.dir = fileDir
	m.filename = fileName
	m.maxFileCount = maxNumber
	m.maxFileSize = maxSize * int64(unit)
	m.isRollingFile = true
	m.isDailyRolling = false
	m.mu = new(sync.RWMutex)

	m.mu.Lock()
	defer m.mu.Unlock()

	for i := 1; i <= maxNumber; i++ {
		if isExist(fileDir + "/" + fileName + "." + strconv.Itoa(i)) {
			m.suffix = i
		} else {
			break
		}
	}
	if !m.isMustRename() {
		m.logfile, _ = os.OpenFile(fileDir+"/"+fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		m.newLog()
	} else {
		m.rename()
	}
	go m.fileMonitor()
}

func (m *LoggerObj) SetRollingDaily(fileDir, fileName string) {
	mkDir(fileDir)
	m.isRollingFile = false
	m.isDailyRolling = true
	t, _ := time.Parse(DateFormat, time.Now().Format(DateFormat))
	m.dir = fileDir
	m.filename = fileName
	m.mu = new(sync.RWMutex)
	m.date = &t

	m.mu.Lock()
	defer m.mu.Unlock()

	if !m.isMustRename() {
		m.logfile, _ = os.OpenFile(fileDir+"/"+fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		m.newLog()
	} else {
		m.rename()
	}
}

func (m *LoggerObj) fileMonitor() {
	m.once.Do(func() {
		timer := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-timer.C:
				m.fileCheck()
			}
		}
	})
}
func (m *LoggerObj) fileCheck() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	if m.isMustRename() {
		m.mu.Lock()
		defer m.mu.Unlock()
		m.rename()
	}
}

func (m *LoggerObj) console(s string) {
	if !m.isConsole {
		return
	}
	t := time.Now()
	if m.flags&(log.Lshortfile|log.Llongfile) != 0 {
		_, file, line, _ := runtime.Caller(m.callDepth)
		if m.flags&log.Lshortfile != 0 {
			short := file
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					short = file[i+1:]
					break
				}
			}
			file = short
		}
		fmt.Printf("%s[%d] %s -%s:%d\n", t.Format("2006/01/02 15:04:05"), t.Unix(), s, file, line)
	} else {
		fmt.Printf("%s[%d] %s\n", t.Format("2006/01/02 15:04:05"), t.Unix(), s)
	}
}

func (m *LoggerObj) debug(lvl LEVEL, fmtStr string, a ...interface{}) {
	if m.isDailyRolling {
		m.fileCheck()
	}
	defer catchError()
	m.mu.RLock()
	defer m.mu.RUnlock()

	if m.level <= lvl {
		fmtStr = debugStr[lvl] + fmtStr
		logStr := fmt.Sprintf(fmtStr, a...)
		m.lg.Output(m.callDepth, logStr)
		m.console(logStr)
		if lvl == FATAL {
			os.Exit(1)
		}
	}
}

func (m *LoggerObj) isMustRename() bool {
	if m.isDailyRolling {
		t, _ := time.Parse(DateFormat, time.Now().Format(DateFormat))
		if t.After(*m.date) {
			return true
		}
	} else {
		if m.maxFileCount > 1 {
			if fileSize(m.dir+"/"+m.filename) >= m.maxFileSize {
				return true
			}
		}
	}
	return false
}

func (m *LoggerObj) rename() {
	if m.isDailyRolling {
		fn := m.dir + "/" + m.date.Format(DateFormat) + "." + m.filename
		if !isExist(fn) && m.isMustRename() {
			if m.logfile != nil {
				m.logfile.Close()
			}
			err := os.Rename(m.dir+"/"+m.filename, fn)
			if err != nil {
				m.lg.Println("rename err", err.Error())
			}
			t, _ := time.Parse(DateFormat, time.Now().Format(DateFormat))
			m.date = &t
			m.logfile, _ = os.Create(m.dir + "/" + m.filename)
			m.newLog()
		}
	} else {
		m.coverNextOne()
	}
}

func (m *LoggerObj) nextSuffix() int {
	return m.suffix%m.maxFileCount + 1
}

func (m *LoggerObj) coverNextOne() {
	m.suffix = m.nextSuffix()
	if m.logfile != nil {
		m.logfile.Close()
	}
	file := fmt.Sprintf("%s/%s.%d", m.dir, m.filename, m.suffix)
	if isExist(file) {
		os.Remove(file)
	}
	os.Rename(m.dir+"/"+m.filename, file)
	m.logfile, _ = os.Create(m.dir + "/" + m.filename)
	m.newLog()
}

func fileSize(file string) int64 {
	f, e := os.Stat(file)
	if e != nil {
		fmt.Println(e.Error())
		return 0
	}
	fmt.Println("fileSize", file, f.Size())
	return f.Size()
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func mkDir(dir string) (e error) {
	_, er := os.Stat(dir)
	b := er == nil || os.IsExist(er)
	if !b {
		if err := os.MkdirAll(dir, 0666); err != nil {
			if os.IsPermission(err) {
				fmt.Println("create dir failed:", err.Error())
				e = err
			}
		}
	}
	return
}

func catchError() {
	if err := recover(); err != nil {
		log.Println("err", err)
	}
}
