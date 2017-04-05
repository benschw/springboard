package secrets

import (
	"gopkg.in/yaml.v2"
	"os"
	"io/ioutil"
	"errors"
)

type crypt interface {
	Encrypt(string) (string, error)
	Decrypt(string) (string, error)
}

func New(path string, crypt crypt) (*Secrets, error) {
	if _, err := os.Stat(path); err != nil {
		ioutil.WriteFile(path, []byte{}, 0644)
	}

	ymlData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	data := make([]SecretsEntry, 0)

	if err = yaml.Unmarshal([]byte(ymlData), &data); err != nil {
		return nil, err
	}

	return &Secrets{path: path, crypt: crypt, data: data}, nil
}

type Secrets struct {
	path string
	crypt crypt
	data []SecretsEntry
}
func (s *Secrets) Keys() []string {
	keys := make([]string, len(s.data))
	for i := 0; i < len(s.data); i++ {
		keys[i] = s.data[i].Key
	}
	return keys
}
func (s *Secrets) Set(key string, value string) {
	for i := 0; i < len(s.data); i++ {
		if s.data[i].Key == key {
			val, _ := s.crypt.Encrypt(value)
			s.data[i].Value = val
			return
		}
	}
	s.data = append(s.data, SecretsEntry{Key: key, Value: value})
}
func (s *Secrets) Get(key string) (string, error) {
	for i := 0; i < len(s.data); i++ {
		if s.data[i].Key == key {
			return s.crypt.Decrypt(s.data[i].Value)
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

