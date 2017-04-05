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
	
	s.setValue("foo", "bar")
	
	val, err := s.getValue("foo")
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

	s.setValue("foo", "bar")

	if err := s.save(); err != nil {
		t.Error("problem saving", err)
	}

	s2, _ := New(file.Name())


	val, _ := s2.getValue("foo")

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

	_, err := s.getValue("foo")
	if err == nil {
		t.Error("key doesn't exist, should have errored")
	}
}
