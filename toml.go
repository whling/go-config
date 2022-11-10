package config

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"strings"
)

type Toml struct {
	Node
	value interface{} // *toml.Tree
}

func LoadTomlFile(filename string) (*Toml, error) {
	cfg, err := toml.LoadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Toml{value: cfg}, nil
}

func LoadToml(content string) (*Toml, error) {
	cfg, err := toml.Load(content)
	if err != nil {
		return nil, err
	}
	return &Toml{value: cfg}, nil
}

func (t *Toml) Type() Type {
	return TomlType
}

func (t *Toml) Value() interface{} {
	return t.value
}

func (t *Toml) Access(key string) Node {
	tree, ok := t.value.(*toml.Tree)
	if !ok {
		return nil
	}
	value := tree.GetPath(strings.Split(key, Delimiter))
	if value == nil {
		return nil
	}
	switch v := value.(type) {
	case *toml.Tree:
		return &Toml{value: v}
	case []*toml.Tree:
		if len(v) == 0 {
			return nil
		}
		return &Toml{value: v[len(v)-1]}
	default:
		return &Toml{value: v}
	}
}

func (t *Toml) AccessArray(key string) []Node {
	tree, ok := t.value.(*toml.Tree)
	if !ok {
		return nil
	}
	values := tree.GetPath(strings.Split(key, Delimiter))

	switch v := values.(type) {
	case []*toml.Tree:
		nodes := make([]Node, 0, len(v))
		for _, item := range v {
			nodes = append(nodes, &Toml{value: item})
		}
		return nodes
	case []interface{}:
		nodes := make([]Node, 0, len(v))
		for _, item := range v {
			nodes = append(nodes, &Toml{value: item})
		}
		return nodes
	default:
		return nil
	}
}

func (t *Toml) AccessMap(key string) map[string]Node {
	tree, ok := t.value.(*toml.Tree)
	if !ok {
		return nil
	}
	value := tree.GetPath(strings.Split(key, Delimiter))
	switch v := value.(type) {
	case *toml.Tree:
		r := make(map[string]Node)
		for _, k := range v.Keys() {
			r[k] = &Toml{value: v.Get(k)}
		}
		return r
	default:
		return nil
	}
}

func (t *Toml) String() string {
	switch v := t.value.(type) {
	case *toml.Tree:
		return v.String()
	default:
		return fmt.Sprintf("%v", v)
	}
}
