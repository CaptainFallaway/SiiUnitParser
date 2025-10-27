package main

const (
	AttributeTypeTest = iota
	AttributeTypeList
	AttributeTypeNumber
)

type Attribute struct {
	atype   string
	content string
}

func (a *Attribute) String() string {
	return a.content
}

type Attributes struct {
	attrs map[string]Attribute
}

func (as *Attributes) AddAttribute()

type Unit struct {
	utype string
	uid   string
	attrs Attributes
}
