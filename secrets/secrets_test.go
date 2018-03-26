package secrets

import (
	"io/ioutil"
	"os"
	"testing"
)

type stubCrypt struct{}

func (s stubCrypt) Encrypt(val string) (string, error) {
	return val, nil
}
func (s stubCrypt) Decrypt(val string) (string, error) {
	return val, nil
}

func TestSecretAccess(t *testing.T) {
	file, _ := ioutil.TempFile(os.TempDir(), "vault")
	defer os.Remove(file.Name())

	s, err := New(file.Name(), stubCrypt{})

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

	s, _ := New(file.Name(), &stubCrypt{})

	s.Set("foo", "bar")

	if err := s.Save(); err != nil {
		t.Error("problem saving", err)
	}

	s2, _ := New(file.Name(), &stubCrypt{})

	val, _ := s2.Get("foo")

	if val != "bar" {
		t.Error("should equal bar", val)
	}
}

func TestSecretRemove(t *testing.T) {
	file, _ := ioutil.TempFile(os.TempDir(), "vault")
	defer os.Remove(file.Name())

	s, _ := New(file.Name(), &stubCrypt{})

	s.Set("foo", "bar")
	if err := s.Save(); err != nil {
		t.Error("problem saving", err)
	}

	s.Remove("foo")
	if err := s.Save(); err != nil {
		t.Error("problem saving", err)
	}

	s2, _ := New(file.Name(), &stubCrypt{})

	val, err := s2.Get("foo")

	if err == nil {
		t.Error("should have been removed", val)
	}
}

func TestKeys(t *testing.T) {
	file, _ := ioutil.TempFile(os.TempDir(), "vault")
	defer os.Remove(file.Name())

	s, _ := New(file.Name(), &stubCrypt{})

	s.Set("foo", "bar")
	s.Set("baz", "boo")

	val := s.Keys()

	if val[0] != "foo" && val[1] != "foo" {
		t.Error("should have foo key", val)
	}
	if val[0] != "baz" && val[1] != "baz" {
		t.Error("should have baz key", val)
	}
}
func TestNoFile(t *testing.T) {
	s, err := New("./dne", &stubCrypt{})
	if err != nil {
		t.Error("config doesn't exist, should have created it", err)
	}
	defer os.Remove("./dne")

	s.Set("foo", "bar")
	s.Save()

	s2, _ := New("./dne", &stubCrypt{})

	val, _ := s2.Get("foo")

	if val != "bar" {
		t.Error("should equal bar", val)
	}
}

func TestNoKey(t *testing.T) {
	file, _ := ioutil.TempFile(os.TempDir(), "vault")
	defer os.Remove(file.Name())

	s, _ := New(file.Name(), &stubCrypt{})

	_, err := s.Get("foo")
	if err == nil {
		t.Error("key doesn't exist, should have errored")
	}
}
