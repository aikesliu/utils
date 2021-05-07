package configs

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/aikesliu/utils/filecheck"
	"github.com/aikesliu/utils/log"
	"gopkg.in/yaml.v2"
)

type Manager struct {
	configs map[string]interface{}
	// 记录注册顺序，保证配置可以按照顺序加载
	names           []string
	FileMgr         *filecheck.FileManager
	loadConfigIndex int
}

func New() *Manager {
	mgr := &Manager{}
	mgr.init()
	return mgr
}

func (m *Manager) init() {
	m.configs = make(map[string]interface{})
	m.FileMgr = filecheck.NewFileManager()
}

func (m *Manager) Reset() {
	m.loadConfigIndex = 0
}

// 注册启动配置文件
func (m *Manager) RegisterStartFile(file string, cb func(f string)) {
	m.FileMgr.RegisterCheckFile(file, cb)
}

// 注册启动配置文件中配置名称对应的结构体指针
func (m *Manager) Register(name string, v interface{}) {
	if _, ok := m.configs[name]; ok {
		log.E("Register failed: name: %v config has registered", name)
		return
	}
	m.configs[name] = v
	m.names = append(m.names, name)
}

func (m *Manager) LoadConfig(file string, v interface{}) bool {
	if file == "" {
		return false
	}
	idx := strings.LastIndex(file, ".")
	ex := file[idx+1:]
	// log.Debug("LoadConfig file: %v ex: %v", file, ex)
	log.I("load config %d: %v", m.loadConfigIndex, file)
	var err error
	if ex == "json" {
		err = FileJSON(file, v)
	} else if ex == "yml" || ex == "yaml" {
		err = FileYML(file, v)
	} else {
		log.E("LoadConfig failed: file: %v unknown file type: %v", file, ex)
		return false
	}
	if err != nil {
		log.E("LoadConfig failed: %v", err)
		return false
	}
	m.loadConfigIndex++

	if iConfig, ok := v.(IConfigLoad); ok {
		iConfig.OnLoad()
	}
	return true
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
