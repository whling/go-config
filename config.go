package config

import "time"

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

type Value interface {
	Node

	Has(string) bool
	HasArray(string) bool
	HasMap(string) bool

	Bool(key string, def bool) bool
	Str(key string, def string) string
	Int64(key string, def int64) int64
	Float64(key string, def float64) float64
	Duration(key string, def time.Duration) time.Duration

	BoolArray(string) []bool
	StrArray(string) []string
	Int64Array(string) []int64
	Float64Array(string) []float64
	DurationArray(string) []time.Duration

	ValueArray(string) []Value
	ValueMap(string) map[string]Value
}
