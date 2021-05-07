package filecheck

import (
	"os"
	"strings"
	"time"

	"github.com/aikesliu/utils/log"
)

type fileCheckInfo struct {
	cb          func(filePath string)
	lastModTime time.Time
}

type FileManager struct {
	checkFileInfoMap map[string]*fileCheckInfo
	ticker           *time.Ticker
}

func NewFileManager() *FileManager {
	m := new(FileManager)
	m.checkFileInfoMap = make(map[string]*fileCheckInfo)
	return m
}

func (m *FileManager) ClearCheckFile() {
	m.checkFileInfoMap = make(map[string]*fileCheckInfo)
}

func (m *FileManager) RegisterCheckFile(file string, cb func(f string)) {
	if strings.Compare(file, "") == 0 || cb == nil {
		log.E("register failed: file path: %v nil or cb: %p nil", file, cb)
		return
	}
	if m.checkFileInfoMap == nil {
		m.checkFileInfoMap = make(map[string]*fileCheckInfo)
	}
	fileInfo, _ := os.Stat(file)
	fci := &fileCheckInfo{
		cb: cb,
	}
	if fileInfo == nil {
		log.W("file: %v not exist", file)
	} else {
		fci.lastModTime = fileInfo.ModTime()
	}
	m.checkFileInfoMap[file] = fci

	log.I("register check file: %v", file)
}

func (m *FileManager) Update() {
	for k, v := range m.checkFileInfoMap {
		fileInfo, _ := os.Stat(k)
		if fileInfo == nil {
			// log.Debug("warning, file: %v not exist", k)
			continue
		}
		// 修改时间
		modTime := fileInfo.ModTime()
		// 有修改
		if v.lastModTime.Unix() != modTime.Unix() {
			v.lastModTime = modTime
			v.cb(k)
		}
	}
}

func (m *FileManager) Stop() {
	if m == nil || m.ticker == nil {
		return
	}
	m.ticker.Stop()
	m.ticker = nil
}

func (m *FileManager) StartUpdate(d time.Duration) {
	if d <= 0 {
		return
	}
	m.Stop()
	m.ticker = time.NewTicker(d)
	go func() {
		for range m.ticker.C {
			m.Update()
		}
	}()
}
