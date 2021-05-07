package log

import (
	"log"
	"runtime"
	"testing"
	"time"
)

func print(i int) {
	D("a: %v, %v, %v", 1, "xl", []int32{1, 2, 3})
	E("a: %v, %v, %v", 1, "xl", []int32{1, 2, 3})
	I("a: %v, %v, %v", 1, "xl", []int32{1, 2, 3})
	W("a: %v, %v, %v", 1, "xl", []int32{1, 2, 3})
	// F("a: %v, %v, %v", 1, "xl", []int32{1, 2, 3})
	D("a: %v, %v, %v", 1, "xl", []int32{1, 2, 3})
}

func Test(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 指定是否控制台打印，默认为true
	SetConsole(true)
	// 指定日志文件备份方式为文件大小的方式
	// 第一个参数为日志文件存放目录
	// 第二个参数为日志文件命名
	// 第三个参数为备份文件最大数量
	// 第四个参数为备份文件大小
	// 第五个参数为文件大小的单位
	// SetRollingFile("d:/logtest", "test.log", 10, 5, KB)

	// 指定日志文件备份方式为日期的方式
	// 第一个参数为日志文件存放目录
	// 第二个参数为日志文件命名
	SetRollingDaily("./logtest", "test.log")

	// 指定日志级别  ALL，DEBUG，INFO，WARN，ERROR，FATAL，OFF 级别由低到高
	// 一般习惯是测试阶段为debug，生成环境为info以上
	SetLevel(DEBUG)

	SetFlags(log.Lshortfile | log.Ltime | log.Ldate)

	print(1)
	// for i := 10000; i > 0; i-- {
	// 	go print(i)
	// 	time.Sleep(1000 * time.Millisecond)
	// }
	// time.Sleep(15 * time.Second)

	l := New()
	// l.SetConsole(true)
	// l.SetRollingDaily("./logtest", "test-1.log")
	l.SetRollingFile("./logtest", "test-2.log", 10, 1, KB)
	// l.SetRollingFile("./logtest", "test-2.log", 3, 5, KB)
	l.SetLevel(DEBUG)
	l.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
	// l.SetPrefix("[fish] ")

	for i := 0; i < 1000; i++ {
		l.D("a: %v, %v, %v", i, "xl100", []int32{2, 3, 4, 5})
		l.E("a: %v, %v, %v", i, "xl100", []int32{2, 3, 4, 5})
		l.I("a: %v, %v, %v", i, "xl100", []int32{2, 3, 4, 5})
		l.W("a: %v, %v, %v", i, "xl100", []int32{2, 3, 4, 5})
		// l.F("a: %v, %v, %v", 100, "xl100", []int32{2, 3, 4, 5})
		time.Sleep(time.Millisecond * 100)
	}
}
