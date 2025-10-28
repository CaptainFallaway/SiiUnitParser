package siiunit

import "strings"

type Unit struct {
	Utype string
	ID    string
	Attrs Attributes
}

func (u Unit) String() string {
	var sb strings.Builder
	sb.WriteString(u.Utype)
	sb.WriteString(" : ")
	sb.WriteString(u.ID)
	sb.WriteString(" {\n")
	for key, attr := range u.Attrs.All() {
		sb.WriteString("\t")
		sb.WriteString(key)
		sb.WriteString(": ")
		sb.WriteString(attr.Printable())
		sb.WriteString("\n")
	}
	sb.WriteString("}")
	return sb.String()
}
