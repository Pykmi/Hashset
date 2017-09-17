package Hashset

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"errors"
)

const error_duplicate = "Cannot store duplicate values"

type HashsetMap map[interface{}]struct{}

type Hashset struct {
	set HashsetMap
}

func (h *Hashset) add(ob interface{}) error {
	hash, err := hash(ob)
	if err != nil {
		return err
	}

	if h.contains(hash) {
		return errors.New(error_duplicate)
	}

	h.set[hash] = struct{}{}
	return nil
}

func (h *Hashset) Add(ob interface{}) {
	err := h.add(ob)
	if err != nil {
		panic(err.Error())
	}
}

func (h *Hashset) AddAll(obs ...interface{}) {
	for _, ob := range obs {
		err := h.add(ob)
		if err != nil {
			panic(err.Error())
		}
	}
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
			panic(errors.New(error_duplicate))
		}

		internal.set[hash] = struct{}{}
	}

	return &internal
}

func (h *Hashset) Len() int {
	return len(h.set)
}
