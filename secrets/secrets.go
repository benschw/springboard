package secrets

import (
	"gopkg.in/yaml.v2"
	"os"
	"io/ioutil"
	"errors"
)

func New(path string) (*Secrets, error) {
	if _, err := os.Stat(path); err != nil {
		return nil, errors.New("config path not valid")
	}

	ymlData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	data := make([]SecretsEntry, 0)

	if err = yaml.Unmarshal([]byte(ymlData), &data); err != nil {
		return nil, err
	}

	return &Secrets{path: path, data: data}, nil
}

type Secrets struct {
	path string
	data []SecretsEntry
}



func (s *Secrets) Set(key string, value string) {
	for i := 0; i < len(s.data); i++ {
		if s.data[i].Key == key {
			s.data[i].Value = value
			return
		}
	}
	s.data = append(s.data, SecretsEntry{Key: key, Value: value})
}
func (s *Secrets) Get(key string) (string, error) {
	for i := 0; i < len(s.data); i++ {
		if s.data[i].Key == key {
			return s.data[i].Value, nil
		}
	}
	return "", errors.New("key not found")
}
func (s *Secrets) Save() error {
	b, err := yaml.Marshal(s.data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(s.path, b, 0644)
}

type SecretsEntry struct {
	Key string
	Value string
}

