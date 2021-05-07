package log

type ILog interface {
	SetCallDepth(d int)
	SetFlags(f int)
	Debug(format string, a ...interface{})
	Release(format string, a ...interface{})
	Error(format string, a ...interface{})
	Fatal(format string, a ...interface{})
	Close()
}

var (
	l *Logger
)

func init() {
	l = New()
	// l.SetCallDepth(3)
}

func New() *Logger {
	logger := &Logger{
		LoggerObj: &LoggerObj{},
	}
	logger.SetCallDepth(3)
	return logger
}

type Logger struct {
	*LoggerObj
}

func (m *Logger) SetCallDepth(d int) {
	m.LoggerObj.callDepth = d
}

func (m *Logger) SetFlags(f int) {
	m.LoggerObj.SetFlags(f)
}

func (m *Logger) D(fmtStr string, a ...interface{}) {
	m.LoggerObj.debug(DEBUG, fmtStr, a...)
}
func (m *Logger) I(fmtStr string, a ...interface{}) {
	m.LoggerObj.debug(INFO, fmtStr, a...)
}
func (m *Logger) W(fmtStr string, a ...interface{}) {
	m.LoggerObj.debug(WARN, fmtStr, a...)
}
func (m *Logger) E(fmtStr string, a ...interface{}) {
	m.LoggerObj.debug(ERROR, fmtStr, a...)
}
func (m *Logger) F(fmtStr string, a ...interface{}) {
	m.LoggerObj.debug(FATAL, fmtStr, a...)
}

func (m *Logger) Debug(format string, a ...interface{}) {
	m.LoggerObj.debug(DEBUG, format, a...)
}

func (m *Logger) Release(format string, a ...interface{}) {
	m.LoggerObj.debug(WARN, format, a...)
}

func (m *Logger) Error(format string, a ...interface{}) {
	m.LoggerObj.debug(ERROR, format, a...)
}

func (m *Logger) Fatal(format string, a ...interface{}) {
	m.LoggerObj.debug(FATAL, format, a...)
}

func (m *Logger) Close() {
}

func SetFlags(flag int) {
	l.SetFlags(flag)
}

func SetConsole(isConsole bool) {
	l.SetConsole(isConsole)
}

func SetLevel(level LEVEL) {
	l.SetLevel(level)
}

func SetRollingFile(fileDir, fileName string, maxNumber int, maxSize int64, unit UNIT) {
	l.SetRollingFile(fileDir, fileName, maxNumber, maxSize, unit)
}

func SetRollingDaily(fileDir, fileName string) {
	l.SetRollingDaily(fileDir, fileName)
}

func D(fmtStr string, a ...interface{}) {
	l.debug(DEBUG, fmtStr, a...)
}
func I(fmtStr string, a ...interface{}) {
	l.debug(INFO, fmtStr, a...)
}
func W(fmtStr string, a ...interface{}) {
	l.debug(WARN, fmtStr, a...)
}
func E(fmtStr string, a ...interface{}) {
	l.debug(ERROR, fmtStr, a...)
}
func F(fmtStr string, a ...interface{}) {
	l.debug(FATAL, fmtStr, a...)
}
