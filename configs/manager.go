package configs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/aikesliu/utils/filecheck"
	"github.com/aikesliu/utils/log"
	"gopkg.in/yaml.v2"
)

type Manager struct {
	FileMgr         *filecheck.FileManager
	loadConfigIndex int
}

func New() *Manager {
	mgr := &Manager{}
	mgr.init()
	return mgr
}

func (m *Manager) init() {
	m.FileMgr = filecheck.NewFileManager()
}

func (m *Manager) Reset() {
	m.loadConfigIndex = 0
}

// 注册启动配置文件
func (m *Manager) RegisterStartFile(file string, cb func(f string)) {
	m.FileMgr.RegisterCheckFile(file, cb)
}

func (m *Manager) LoadFromData(data []byte, name string, v interface{}) error {
	if name == "" {
		return fmt.Errorf("file is nil")
	}
	idx := strings.LastIndex(name, ".")
	ex := name[idx+1:]
	// log.Debug("LoadConfig file: %v ex: %v", file, ex)
	log.I("load config [%d]: %v", m.loadConfigIndex, name)
	var err error
	if ex == "json" {
		err = json.Unmarshal(data, v)
	} else if ex == "yml" || ex == "yaml" {
		err = yaml.Unmarshal(data, v)
	} else {
		return fmt.Errorf("file: %v unknown file type: %v", name, ex)
	}
	if err != nil {
		return err
	}
	m.loadConfigIndex++

	if iConfig, ok := v.(IConfigLoad); ok {
		iConfig.OnLoad()
	}
	return nil
}

func (m *Manager) LoadConfig(file string, v interface{}) error {
	if file == "" {
		return fmt.Errorf("file is nil")
	}
	idx := strings.LastIndex(file, ".")
	ex := file[idx+1:]
	// log.Debug("LoadConfig file: %v ex: %v", file, ex)
	log.I("load config [%d]: %v", m.loadConfigIndex, file)
	var err error
	if ex == "json" {
		err = FileJSON(file, v)
	} else if ex == "yml" || ex == "yaml" {
		err = FileYML(file, v)
	} else {
		return fmt.Errorf("file: %v unknown file type: %v", file, ex)
	}
	if err != nil {
		return err
	}
	m.loadConfigIndex++

	if iConfig, ok := v.(IConfigLoad); ok {
		iConfig.OnLoad()
	}
	return nil
}

// FileYML parse yml data from file
// v will be a pointer of object
func FileYML(path string, v interface{}) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, v)
}

// FileJSON parse json data from file
func FileJSON(path string, v interface{}) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}
