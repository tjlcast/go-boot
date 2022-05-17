package utils

import (
	"log"

	"github.com/Unknwon/goconfig"
	"github.com/fsnotify/fsnotify"
)

const (
	DEFAULT_SECTION = goconfig.DEFAULT_SECTION
)

var cfgMap = make(map[string]map[string]string)
var filePath string

func InitConfig(inFilePath string) (err error) {
	err = parse(inFilePath)
	watchConfig(inFilePath)
	return
}

func GetAllCfg() (cfgMap map[string]map[string]string) {
	return
}

func GetSec(sectionKey string) (section map[string]string, ok bool) {
	section, ok = cfgMap[sectionKey]
	return
}

func parse(inFilePath string) (err error) {
	cfg, err := goconfig.LoadConfigFile(inFilePath)
	if err != nil {
		log.Fatal("Fail to load file: ", err)
	}
	// mark the path
	filePath = inFilePath

	sectionList := cfg.GetSectionList()
	for _, section := range sectionList {
		cfgMap[section] = make(map[string]string)
		keyList := cfg.GetKeyList(section)
		for _, key := range keyList {
			cfgMap[section][key], err = cfg.GetValue(section, key)
			if err != nil {
				log.Fatal("Fail to get section: ", section, " val: ", key)
			}
		}
	}
	return err
}

func watchConfig(filePaths ...string) {
	go func() {
		watcher, err := fsnotify.NewWatcher()
		defer func() {
			if watcher != nil {
				watcher.Close()
			}
		}()
		if err != nil {
			log.Fatal("Config watch error: ", err)
		}

		for _, fpath := range filePaths {
			err := watcher.Add(fpath)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("WatchConfig: " + fpath)
		}

		for {
			select {
			case event := <-watcher.Events:
				{
					if event.Op&fsnotify.Write == fsnotify.Write {
						_ = reloadAllCfg()
					}
				}
			case err := <-watcher.Errors:
				{
					log.Fatal("Error: ", err)
					return
				}
			}
		}
	}()
}

func reloadAllCfg() (err error) {
	err = parse(filePath)
	return
}
