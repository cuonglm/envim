package envim

import (
	"os"
	"reflect"
	"testing"
)

var testKey = "foo"
var testVal = "bar"

func setUp() {
	os.Clearenv()
	_ = os.Setenv(testKey, testVal)
}

func TestSet(t *testing.T) {
	setUp()

	err := Set(testKey, "")
	if err != nil {
		t.Fatalf("Set(): %+v", err)
	}

	setInvalidTests := []struct {
		k, v string
	}{
		{"", ""},
		{"=", ""},
		{"\x00", ""},
		{"key", "\x00"},
	}

	for _, tt := range setInvalidTests {
		err := Set(tt.k, tt.v)
		if err == nil {
			t.Fatalf("Set(%q, %q) should be error", tt.k, tt.v)
		}
	}
}

func TestUnset(t *testing.T) {
	setUp()

	err := Unset(testKey)
	if err != nil {
		t.Fatalf("Unset(): %+v", err)
	}

	if IsSet(testKey) {
		t.Fatalf("%q should be unset", testKey)
	}
}

func TestGet(t *testing.T) {
	setUp()

	getTests := []struct {
		k, v string
	}{
		{testKey, testVal},
		{"fooo", ""},
	}

	for _, tt := range getTests {
		got := Get(tt.k)
		if got != tt.v {
			t.Fatalf("Get(%q) should return %q, got %q", tt.k, tt.v, got)
		}
	}
}

func TestIsSet(t *testing.T) {
	setUp()

	isSetTests := []struct {
		key      string
		expected bool
	}{
		{testKey, true},
		{"NonExistedKey", false},
	}

	for _, tt := range isSetTests {
		if IsSet(tt.key) != tt.expected {
			t.Fatalf("IsSet(%q) should return %v", tt.key, tt.expected)
		}
	}
}

func TestClear(t *testing.T) {
	setUp()

	Clear()
	m := Map()

	if len(m) != 0 {
		t.Fatalf("Clear() should clean environment variables, got %+v", m)
	}
}

func TestMap(t *testing.T) {
	setUp()

	m := Map()

	typ := reflect.TypeOf(m).Kind()

	if typ != reflect.Map {
		t.Fatalf("Map(): wrong type %T", typ)
	}
}

func TestMapWithPrefix(t *testing.T) {
	mapWithPrefixTests := []struct {
		k, v, prefix string
		length       int
	}{
		{"fooo", "", "foo", 2},
		{"bar", "", "bar", 1},
	}

	for _, tt := range mapWithPrefixTests {
		setUp()
		_ = Set(tt.k, tt.v)
		m := MapWithPrefix(tt.prefix)
		if len(m) != tt.length {
			t.Fatalf("MapWithPrefix(): wrong prefix, %+v", m)
		}
	}
}

func TestFromMap(t *testing.T) {
	fromMapTests := []struct {
		fromMap map[string]string
		length  int
	}{
		{map[string]string{"fooo": "", "baz": ""}, 3},
		{map[string]string{"fooo": "", "": ""}, 2},
	}
	for _, tt := range fromMapTests {
		setUp()
		FromMap(tt.fromMap)
		m := Map()
		if len(m) != tt.length {
			t.Fatalf("FromMap(): %+v - %d - %d", m, len(m), tt.length)
		}
	}
}
