package config

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"testing"
)

var (
	TomlFilePath = "./test.toml"
	testToml     *Toml
)

func TestToml_LoadTomlFile(t *testing.T) {
	toml, err := LoadTomlFile(TomlFilePath)
	if err != nil {
		t.Fatal(err)
	}
	if toml.Type() != TomlType {
		t.Fatalf("file type err")
	}
	testToml = toml
}

func TestToml_Access(t *testing.T) {
	n1 := testToml.Access("integer.key1")
	if n1 == nil {
		t.Fatal()
	}
	if v, ok := n1.Value().(int64); !ok || v != 99 {
		t.Fatal()
	}

	n2 := testToml.Access("boolean.True")
	if n2 == nil {
		t.Fatal()
	}
	if b, ok := n2.Value().(bool); !ok || !b {
		t.Fatal()
	}

	n3 := testToml.Access("string_array")
	if n3 == nil {
		t.Fatal()
	}
	n4 := n3.Access("key2")
	if n4 == nil {
		t.Fatal()
	}

	n5 := testToml.Access("fruit")
	if n5 == nil {
		t.Fatal()
	}

	n6 := testToml.Access("products")
	if n6 == nil {
		t.Fatal()
	}
}

func TestToml_AccessArray(t *testing.T) {
	var (
		node  Node
		nodes []Node
	)

	fruitNode := testToml.AccessArray("fruit.name")
	if fruitNode != nil {
		t.Fatal() // fruit.name 对应的并不是array，而是取最后的那个值

	}
	nodes = testToml.AccessArray("products")
	if len(nodes) != 2 {
		t.Fatal(nodes)
	}

	_, ok := nodes[0].Value().(*toml.Tree)
	if !ok {
		t.Fatal()
	}

	node = nodes[0].Access("name")
	if v, ok := node.Value().(string); !ok || v != "Hammer" {
		t.Fatal(v)
	}
	node = nodes[1].Access("name")
	if v, ok := node.Value().(string); !ok || v != "Nail" {
		t.Fatal(v)
	}

	nodes = testToml.AccessArray("boolean_array.key")
	if len(nodes) != 4 {
		t.Fatal(nodes)
	}
}

func TestToml_AccessMap(t *testing.T) {
	var (
		nodeMap map[string]Node
	)
	nodeMap = testToml.AccessMap("boolean")
	if len(nodeMap) != 2 {
		t.Fatal(nodeMap)
	}
	if v, ok := nodeMap["True"]; !ok {
		t.Fatal(v)
	} else {
		if tv, ok := v.Value().(bool); !ok || !tv {
			t.Fatal(tv)
		}
	}
	if v, ok := nodeMap["False"]; !ok {
		t.Fatal(v)
	} else {
		if tv, ok := v.Value().(bool); !ok || tv {
			t.Fatal(tv)
		}
	}

	nodeMap = testToml.AccessMap("products.name")
	fmt.Println(nodeMap)
	if nodeMap != nil {
		t.Fatal(nodeMap)
	}

	nodeMap = testToml.AccessMap("test.server")
	if nodeMap == nil {
		t.Fatal(nodeMap)
	}
	if len(nodeMap) != 2 {
		t.Fatal(nodeMap)
	}
	if v, ok := nodeMap["timeout"]; !ok {
		t.Fatal(v)
	} else {
		node := v.Access("key2")
		if node == nil {
			t.Fatal(v)
		}
		if tv, ok := node.Value().(int64); !ok || tv != 123 {
			t.Fatal(tv)
		}
	}
}

func TestToml_AccessArrayAndAccessMap(t *testing.T) {
	fa := testToml.AccessArray("fruit")
	if len(fa) != 2 {
		t.Fatal()
	}
	fruitPhysicalTtMap := fa[0].AccessMap("physical.tt")
	if fruitPhysicalTtMap == nil || len(fruitPhysicalTtMap) != 2 {
		t.Fatal()
	}
	if s, ok := fruitPhysicalTtMap["color"].Value().(string); ok && s != "red" {
		t.Fatal()
	}

	fruitVarietyArray := fa[0].AccessArray("variety")
	if fruitVarietyArray == nil || len(fruitVarietyArray) != 2 {
		t.Fatal()
	}
	if s, ok := fruitVarietyArray[0].Access("name").Value().(string); ok && s != "red delicious" {
		t.Fatal()
	}
}
