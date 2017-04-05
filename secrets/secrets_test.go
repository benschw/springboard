package secrets


import (
	"testing"
	"io/ioutil"
	"os"
)


func TestSecretAccess(t *testing.T) {
	file, _ := ioutil.TempFile(os.TempDir(), "vault")
	defer os.Remove(file.Name())

	s, err := New(file.Name())

	if err != nil {
		t.Error("problem loading config: ", err)
	}
	
	s.Set("foo", "bar")
	
	val, err := s.Get("foo")
	if err != nil {
		t.Error("problem retreiving key: ", err)
	}

	if val != "bar" {
		t.Error("should equal bar", val)
	}

}

func TestSecretStorage(t *testing.T) {
	file, _ := ioutil.TempFile(os.TempDir(), "vault")
	defer os.Remove(file.Name())

	s, _ := New(file.Name())

	s.Set("foo", "bar")

	if err := s.Save(); err != nil {
		t.Error("problem saving", err)
	}

	s2, _ := New(file.Name())


	val, _ := s2.Get("foo")

	if val != "bar" {
		t.Error("should equal bar", val)
	}
}
func TestNoFile(t *testing.T) {
	_, err := New("./dne")
	
	if err == nil {
		t.Error("config doesn't exist, should have errored")
	}
}
func TestNoKey(t *testing.T) {
	file, _ := ioutil.TempFile(os.TempDir(), "vault")
	defer os.Remove(file.Name())

	s, _ := New(file.Name())

	_, err := s.Get("foo")
	if err == nil {
		t.Error("key doesn't exist, should have errored")
	}
}
