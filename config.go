package config

const (
	TomlType = Type("Toml")
	TomlJson = Type("Json")

	Delimiter = "."
)

type Type string

type Node interface {
	Type() Type
	Value() interface{}
	Access(string) Node
	AccessArray(string) []Node
	AccessMap(string) map[string]Node
	String() string
}
