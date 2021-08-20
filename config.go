package gpractice

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// Cfg contains configuration properties
type Cfg struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`

	Db struct {
		Type     string `yaml:"type"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
		UserName string `yaml:"userName"`
		UserPass string `yaml:"userPass"`
	} `yaml:"db"`
}

// LoadCfg reads yaml file, decodes and returns *Cfg
func LoadCfg(file string) (*Cfg, error) {
	f, err := os.Open(path(file))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Cfg
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func path(file string) string {
	if strings.HasPrefix(file, "~/") {
		usr, _ := user.Current()
		dir := usr.HomeDir
		result := filepath.Join(dir, strings.Replace(file, "~/", "", 1))
		log.Printf("path: %s", result)
		return result
	}
	return file
}
