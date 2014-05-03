package gojx

import (
	"strconv"
)

type Index struct {
	Name       string           `json:"name"`
	Type       string           `json:"type"`
	StringData map[string][]int `json:"string_dat"`
	StringKeys []string         `json:"strings"`
	IntData    map[string][]int `json:"int_dat"`
	IntKeys    []int            `json:"ints"`
}

func (i *Index) updateKeys() {
	if i.Type == INDEX_INT {
		keys := []int{}
		for k, _ := range i.IntData {
			kInt, _ := strconv.Atoi(k)
			keys = append(keys, kInt)
		}
		i.IntKeys = keys
		return
	}
	if i.Type == INDEX_STRING {
		keys := []string{}
		for k, _ := range i.StringData {
			keys = append(keys, k)
		}
		i.StringKeys = keys
		return
	}
}

func (i *Index) Insert(value string, id int) {
	if i.Type == INDEX_INT {
		if _, ok := i.IntData[value]; !ok {
			i.IntData[value] = []int{id}
		} else {
			i.IntData[value] = append(i.IntData[value], id)
		}
		i.updateKeys()
		return
	}
	if i.Type == INDEX_STRING {
		if _, ok := i.StringData[value]; !ok {
			i.StringData[value] = []int{id}
		} else {
			i.StringData[value] = append(i.StringData[value], id)
		}
		i.updateKeys()
		return
	}
}

// Create new index with name and type.
// Index support int and string type.
func NewIndex(name string, idxType string) *Index {
	idx := &Index{Name: name, Type: idxType}
	if idx.Type == INDEX_INT {
		idx.IntData = make(map[string][]int)
		idx.IntKeys = make([]int, 0)
	}
	if idx.Type == INDEX_STRING {
		idx.StringData = make(map[string][]int)
		idx.StringKeys = make([]string, 0)
	}
	return idx
}
