package siiunit

import (
	"fmt"
	"iter"
)

type Attributes struct {
	attrs map[string]*Attribute
}

func newAttributes() *Attributes {
	return &Attributes{
		attrs: make(map[string]*Attribute, 0),
	}
}

func (as *Attributes) addAttribute(key, val string) error {
	attr, err := newAttribute(val)
	if err != nil {
		return fmt.Errorf("failed to add attribute %s: %w", key, err)
	}

	as.attrs[key] = attr
	return nil
}

func (as *Attributes) Get(attrKey string) (Attribute, bool) {
	attr, ok := as.attrs[attrKey]
	return *attr, ok
}

// All returns an iterator over all attribute key-value pairs.
// Usage: for key, attr := range attrs.All() { ... }
func (as *Attributes) All() iter.Seq2[string, Attribute] {
	return func(yield func(string, Attribute) bool) {
		for k, v := range as.attrs {
			if !yield(k, *v) {
				return
			}
		}
	}
}
