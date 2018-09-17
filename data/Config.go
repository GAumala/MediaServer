package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
	"sort"
)

type Config struct {
	Port      int
	Verbose   bool
	VideoDirs []string
}

func (c Config) videoDirsString() string {
	var buffer bytes.Buffer
	buffer.WriteString("[\n")
	for i := 0; i < len(c.VideoDirs); i++ {
		buffer.WriteString("      \"")
		buffer.WriteString(c.VideoDirs[i])
		buffer.WriteString("\"\n")
	}
	buffer.WriteString("    ]\n")
	return buffer.String()
}

func (c Config) String() string {
	return fmt.Sprintf("{\n  port: %d\n  verbose: %t\n  VideoDirs: %s}",
		c.Port, c.Verbose, c.videoDirsString())
}

func defaultVideoDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}

	return usr.HomeDir + "/Videos/"
}

func ParseJsonConfig(configPath string) (*Config, error) {
	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	config := DefaultConfig()
	json.Unmarshal(b, config)
	sort.Strings(config.VideoDirs)

	if config.Verbose {
		fmt.Printf("Using custom config at: %s\n%s\n", configPath, config)
	}
	return config, nil
}

func DefaultConfig() *Config {
	return &Config{
		Port:      8080,
		Verbose:   true,
		VideoDirs: []string{defaultVideoDir()},
	}
}
