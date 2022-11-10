package config

import (
	"time"
)

type TomlValue struct {
	*Toml
}

func (t *TomlValue) Has(key string) bool {
	if key == "" {
		return false
	}
	return t.Access(key) != nil
}

func (t *TomlValue) HasArray(key string) bool {
	if key == "" {
		return false
	}
	return t.AccessArray(key) != nil
}

func (t *TomlValue) HasMap(key string) bool {
	if key == "" {
		return false
	}
	return t.AccessMap(key) != nil
}

func (t *TomlValue) Bool(key string, def bool) bool {
	if key != "" {
		if n := t.Access(key); n != nil {
			if r, ok := n.Value().(bool); ok {
				return r
			}
		}
	}
	return def
}

func (t *TomlValue) Str(key string, def string) string {
	if key != "" {
		if n := t.Access(key); n != nil {
			if r, ok := n.Value().(string); ok {
				return r
			}
		}
	}
	return def
}

func (t *TomlValue) Int(key string, def int) int {
	return int(t.Int64(key, int64(def)))
}

func (t *TomlValue) Int64(key string, def int64) int64 {
	if key != "" {
		if n := t.Access(key); n != nil {
			if r, ok := n.Value().(int64); ok {
				return r
			}
		}
	}
	return def
}

func (t *TomlValue) Float64(key string, def float64) float64 {
	if key != "" {
		if n := t.Access(key); n != nil {
			if r, ok := n.Value().(float64); ok {
				return r
			}
		}
	}
	return def
}

func (t *TomlValue) Duration(key string, def time.Duration) time.Duration {
	if key != "" {
		if n := t.Access(key); n != nil {
			if str, ok := n.Value().(string); ok {
				if duration, err := time.ParseDuration(str); err == nil {
					return duration
				}
			}
		}
	}
	return def
}

func (t *TomlValue) BoolArray(key string) []bool {
	if key != "" {
		if n := t.AccessArray(key); n != nil {
			r := make([]bool, 0, len(n))
			for _, e := range n {
				if b, ok := e.Value().(bool); ok {
					r = append(r, b)
				}
			}
			return r
		}
	}
	return nil
}

func (t *TomlValue) StrArray(key string) []string {
	if key != "" {
		if n := t.AccessArray(key); n != nil {
			r := make([]string, 0, len(n))
			for _, e := range n {
				if b, ok := e.Value().(string); ok {
					r = append(r, b)
				}
			}
			return r
		}
	}
	return nil
}

func (t *TomlValue) IntArray(key string) []int {
	if key != "" {
		if array := t.AccessArray(key); array != nil {
			intSlice := make([]int, 0, len(array))
			for _, elem := range array {
				if v, ok := elem.Value().(int64); ok {
					intSlice = append(intSlice, int(v))
				}
			}
			return intSlice
		}
	}
	return nil
}

func (t *TomlValue) Int64Array(key string) []int64 {
	if key != "" {
		if n := t.AccessArray(key); n != nil {
			r := make([]int64, 0, len(n))
			for _, e := range n {
				if b, ok := e.Value().(int64); ok {
					r = append(r, b)
				}
			}
			return r
		}
	}
	return nil
}

func (t *TomlValue) Float64Array(key string) []float64 {
	if key != "" {
		if n := t.AccessArray(key); n != nil {
			r := make([]float64, 0, len(n))
			for _, e := range n {
				if b, ok := e.Value().(float64); ok {
					r = append(r, b)
				}
			}
			return r
		}
	}
	return nil
}

func (t *TomlValue) DurationArray(key string) []time.Duration {
	if key != "" {
		if n := t.AccessArray(key); n != nil {
			r := make([]time.Duration, 0, len(n))
			for _, e := range n {
				if b, ok := e.Value().(string); ok {
					duration, err := time.ParseDuration(b)
					if err == nil {
						r = append(r, duration)
					}
				}
			}
			return r
		}
	}
	return nil
}

func (t *TomlValue) ValueArray(key string) []Value {
	if key != "" {
		if n := t.AccessArray(key); n != nil {
			r := make([]Value, 0, len(n))
			for _, e := range n {
				toml, ok := e.(*Toml)
				if ok {
					r = append(r, &TomlValue{toml})
				}
			}
			return r
		}
	}
	return nil
}

func (t *TomlValue) ValueMap(key string) map[string]Value {
	if key != "" {
		if m := t.AccessMap(key); m != nil {
			r := make(map[string]Value, len(m))
			for k, v := range m {
				toml, ok := v.(*Toml)
				if ok {
					r[k] = &TomlValue{toml}
				}
			}
			return r
		}
	}
	return nil
}
