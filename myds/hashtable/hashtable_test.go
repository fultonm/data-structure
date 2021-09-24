package hashtable

import (
	"testing"
)

var tables = []struct {
	key   string
	value string
}{
	{key: "foo", value: "1"},
	{key: "bar", value: "2"},
	{key: "baz", value: "3"},
	{key: "bonk", value: "4"},
}

func TestAdd(t *testing.T) {
	ht := New[string](23, HashFunc, QuadraticProbe)
	ht.Insert(tables[0].key, tables[0].value)
	if ht.Len() != 1 {
		t.Fatalf("Expected hashtable to have length 1, has length %d", ht.len)
	}
	// myString := *ht.GetString("alice")
	// if myString != input {
	// 	t.Fatalf("Expected %v, got %v", input, myString)
	// }
}

func TestGet(t *testing.T) {
	ht := New[string](7, HashFunc, QuadraticProbe)
	for _, e := range tables {
		ht.Insert(e.key, e.value)
	}
	for _, e := range tables {
		retrieved := *ht.Get(e.key)
		if retrieved != e.value {
			t.Fatalf("Expected %v, got %v", e.value, retrieved)
		}
	}
}

func TestOverload(t *testing.T) {
	ht := New[string](3, HashFunc, QuadraticProbe)
	for i := 0; i < 3; i++ {
		ht.Insert(tables[i].key, tables[i].value)
	}
	ht, err := ht.Insert(tables[3].key, tables[3].value)
	if err == nil {
		t.Fatalf("Expected error, got no error")
	}
}

func TestGetNotExist(t *testing.T) {
	ht := New[string](3, HashFunc, QuadraticProbe)
	ht.Insert("foo", "bar")
	element := ht.Get("baz")
	t.Fatalf(" Anything %v", element)
}
