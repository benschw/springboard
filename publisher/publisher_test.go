package publisher

import (
	"testing"
	vaultapi "github.com/hashicorp/vault/api"
)

type secretsStoreStub struct {
	data map[string]string
}
func (s *secretsStoreStub) Keys() []string {
	keys := make([]string, 0, len(s.data))
	for k := range s.data {
		keys = append(keys, k)
	}

	return keys
}
func (s *secretsStoreStub) Get(key string) (string, error) {
	return s.data[key], nil
}

func TestEncryption(t *testing.T) {

	
	cfg := vaultapi.DefaultConfig()
	client, err := vaultapi.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	v := client.Logical()

	s := &secretsStoreStub{
		data: map[string]string{
			"foo": "bar",
			"baz": "boo",
		},
	}

	p := New(v, "secret/test")

	err = p.Push(s)
	if err != nil {
		t.Error(err)
	}
	result, _ := v.Read(p.path)

	if result.Data["foo"] != "bar" {
		t.Error("foo should be set to bar", result.Data)
	}
	if result.Data["baz"] != "boo" {
		t.Error("baz should be set to boo", result.Data)
	}
}

