package envy

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

var varRe *regexp.Regexp

func getKeyValue(line []byte) (string, string, error) {
	matches := varRe.FindSubmatch(line)
	if matches == nil {
		return "", "", fmt.Errorf("Improperly formated variable declaration: %s", line)
	}
	k := string(matches[1])
	v := string(matches[2])

	// remove wrapping quotes
	if v[0] == '"' && v[len(v)-1] == '"' {
		v = v[1 : len(v)-1]
	}

	return k, v, nil
}

func loadEnvVars(envFilePath string) error {
	varLines, err := ioutil.ReadFile(envFilePath)
	if err != nil {
		panic(err)
	}
	envVars := bytes.Split(varLines, []byte("\n"))
	for _, v := range envVars {
		if len(v) == 0 {
			continue
		}
		key, value, err := getKeyValue(v)
		if err != nil {
			return err
		}
		if err := os.Setenv(key, value); err != nil {
			return fmt.Errorf("Error setting pair %s:%s: %s", key, value, err)
		}
	}
	return nil
}

func init() {
	varRe = regexp.MustCompile(`([^=]*)=(.*)`)
}

// Load will look for a .env file and attempt to load its variables as
// environment variables
func Load() error {
	_, err := os.Stat(".env")
	if err != nil {
		return fmt.Errorf("No .env file detected")
	}
	if err = loadEnvVars(".env"); err != nil {
		return fmt.Errorf("Error loading .env file: %s", err)
	}
	return nil
}

// LoadFiles takes a slice of filepaths and attempts to load their variables
// into the environment
func LoadFiles(filePaths []string) error {
	for _, envFile := range filePaths {
		if err := loadEnvVars(envFile); err != nil {
			return fmt.Errorf("Error load file %s: %s", envFile, err)
		}
	}
	return nil
}
