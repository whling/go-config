package config

import (
	"testing"
	"time"
)

var (
	testTomlValue *TomlValue
)

func TestTomlValue_Init(t *testing.T) {
	TestToml_LoadTomlFile(t)
	testTomlValue = &TomlValue{
		testToml,
	}
}

func TestTomlValue_Has(t *testing.T) {
	if !testTomlValue.Has("test.server") {
		t.Fatal()
	}
	if testTomlValue.Has("string.str3") {
		t.Fatal()
	}
}

func TestValue_HasArray(t *testing.T) {
	if !testTomlValue.HasArray("products") {
		t.Fatal()
	}
	if testTomlValue.HasArray("boolean_array") {
		t.Fatal()
	}
	if testTomlValue.HasArray("boolean_array.key.true") {
		t.Fatal()
	}
}

func TestValue_HasMap(t *testing.T) {
	if !testTomlValue.HasMap("test.server") {
		t.Fatal()
	}
	if testTomlValue.HasMap("test.client.tt") {
		t.Fatal()
	}
	if testTomlValue.HasMap("test.server.timeout.key1") {
		t.Fatal()
	}
}

func TestValue_Value(t *testing.T) {
	if !testTomlValue.Bool("boolean.True", false) {
		t.Fatal()
	}
	if testTomlValue.Bool("boolean.False", true) {
		t.Fatal()
	}
	if testTomlValue.Bool("boolean.nil", false) {
		t.Fatal()
	}
	if v := testTomlValue.Float64("test.client.name.float_val", 1.0); v != float64(3.1415926) {
		t.Fatal(v)
	}
	if v := testTomlValue.Float64("test.client.name.float2_val", 1.23); v != float64(1.23) {
		t.Fatal(v)
	}
	if v := testTomlValue.Duration("test.golang.timeout", 100*time.Millisecond); v != 300*time.Millisecond {
		t.Fatal(v)
	}
	if v := testTomlValue.Duration("test.golang.timeout2", 100*time.Millisecond); v != 100*time.Millisecond {
		t.Fatal(v)
	}
}

func TestValue_ValueArray(t *testing.T) {
	v := testTomlValue.BoolArray("boolean_array.key")
	if len(v) != 4 || v[2] {
		t.Fatal(v)
	}

	integerKey1 := testTomlValue.Int("integer.key1", 0)
	if integerKey1 != 99 {
		t.Fatal()
	}

	if testTomlValue.BoolArray("boolean_array.key2") != nil {
		t.Fatal()
	}

	if arr := testTomlValue.ValueArray("string_array.key2"); arr == nil {
		t.Fatal(arr)
	} else {
		if len(arr) != 3 {
			t.Fatal(arr)
		}
		if tv, ok := arr[1].Value().(string); !ok || tv != "pear" {
			t.Fatal(tv)
		}
		if tv, ok := arr[2].Value().(string); !ok || tv != "banana" {
			t.Fatal(tv)
		}
	}

	if intArray := testTomlValue.IntArray("int_array.key"); intArray == nil {
		t.Fatal()
	} else {
		if len(intArray) == 0 {
			t.Fatal()
		}
		if intArray[2] != 5 {
			t.Fatal()
		}
	}
}

func TestValue_ValueMap(t *testing.T) {
	vmap := testTomlValue.ValueMap("test.server")
	if vmap == nil || len(vmap) != 2 {
		t.Fatal(vmap)
	}

	item, ok := vmap["name"]
	if !ok {
		t.Fatal(vmap)
	}
	if item.Str("key1", "str") != "str1" {
		t.Fatal(item)
	}
}
