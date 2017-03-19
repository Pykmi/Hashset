package hashset

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"errors"
)

type HashsetMap map[interface{}]struct{}

type Hashset struct {
	set HashsetMap
}

func (h *Hashset) Add(ob interface{}) {
	hash, err := hash(ob)
	if err != nil {
		panic(err.Error())
	}

	if h.contains(hash) {
		panic(err.Error())
	}

	h.set[hash] = struct{}{}
}

func (h *Hashset) Contains(ob interface{}) bool {
	hash, err := hash(ob)
	if err != nil {
		panic(err.Error())
	}

	return h.contains(hash)
}

func (h *Hashset) contains(hash string) bool {
	if _, ok := h.set[hash]; ok {
		return true
	}

	return false
}

func hash(ob interface{}) (string, error) {
	bytes := bytes.Buffer{}
	data := gob.NewEncoder(&bytes)

	err := data.Encode(ob)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(bytes.Bytes()), nil
}

func New(values ...interface{}) *Hashset {
	internal := Hashset{make(HashsetMap)}
	for _, value := range values {
		hash, err := hash(value)
		if err != nil {
			panic(err.Error())
		}

		if internal.contains(hash) {
			panic(errors.New("Cannot store duplicate values"))
		}

		internal.set[hash] = struct{}{}
	}

	return &internal
}

func (h *Hashset) Len() int {
	return len(h.set)
}
