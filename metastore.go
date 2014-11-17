package metastore

import (
	"fmt"
)

type Metastore struct {
	data map[string]interface{}
}

func (m *Metastore) Get(name string) (interface{}, bool) {
	data, ok := m.data[name]
	if !ok {
		return nil, false
	}

	return data, true
}

func (m *Metastore) Fetch(name string) interface{} {
	data, ok := m.data[name]
	if !ok {
		panic("Unknown key: " + name)
	}

	return data
}

func (m *Metastore) Set(name string, data interface{}) {
	if nil == m.data {
		m.data = make(map[string]interface{})
	}

	m.data[name] = data
}

func (m *Metastore) Dump() {
	for k, v := range m.data {
		fmt.Printf("%s -> %#v\n", k, v)
	}
}
