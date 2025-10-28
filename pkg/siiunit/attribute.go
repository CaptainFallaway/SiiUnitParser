package siiunit

import (
	"errors"
	"fmt"
)

// AttributeType represents the type of a SII unit attribute
type AttributeType int

const (
	AttributeTypeString AttributeType = iota
	AttributeTypeFloat
	AttributeTypeFloat2
	AttributeTypeFloat3
	AttributeTypeFloat4
	AttributeTypePlacement
	AttributeTypeInt
	AttributeTypeInt2
	AttributeTypeInt3
	AttributeTypeInt4
	AttributeTypeBool
	AttributeTypeArray
)

var attributeTypeNames = map[AttributeType]string{
	AttributeTypeString:    "string",
	AttributeTypeFloat:     "float",
	AttributeTypeFloat2:    "float2",
	AttributeTypeFloat3:    "float3",
	AttributeTypeFloat4:    "float4",
	AttributeTypePlacement: "placement",
	AttributeTypeInt:       "int",
	AttributeTypeInt2:      "int2",
	AttributeTypeInt3:      "int3",
	AttributeTypeInt4:      "int4",
	AttributeTypeBool:      "bool",
	AttributeTypeArray:     "array",
}

// Attribute represents a SII unit attribute with its type and values
type Attribute struct {
	Atype AttributeType

	// For scalar and basic types
	stringVal string
	floatVal  float64
	intVal    int64
	boolVal   bool

	// For vector types
	float2Vals [2]float64
	float3Vals [3]float64
	float4Vals [4]float64
	int2Vals   [2]int64
	int3Vals   [3]int64
	int4Vals   [4]int64

	// For placement (position + rotation)
	placementPos [3]float64
	placementRot [4]float64

	// For arrays
	arrayVals     []Attribute
	arrayElemType AttributeType
}

var (
	ErrInvalidType   = errors.New("invalid attribute type for this operation")
	ErrNotAnArray    = errors.New("attribute is not an array")
	ErrParsingFailed = errors.New("failed to parse attribute value")
)

// newAttribute creates a new Attribute by parsing the string value
func newAttribute(value string) (*Attribute, error) {
	atype := detectAttributeType(value)
	attr := &Attribute{Atype: atype}

	err := attr.parseValue(value)
	if err != nil {
		return nil, err
	}

	return attr, nil
}

// makeArray marks this attribute as an array and initializes the array slice
func (a *Attribute) makeArray(size int) error {
	if size < 0 {
		return errors.New("array size cannot be negative")
	}
	a.Atype = AttributeTypeArray
	a.arrayVals = make([]Attribute, 0, size)
	return nil
}

// appendToArray parses the string value and appends it to the array
func (a *Attribute) appendToArray(value string) error {
	if a.Atype != AttributeTypeArray {
		return ErrNotAnArray
	}

	// Parse the string into a new attribute
	attr, err := newAttribute(value)
	if err != nil {
		return fmt.Errorf("failed to parse array element: %w", err)
	}

	// Set element type on first append
	if len(a.arrayVals) == 0 {
		a.arrayElemType = attr.Atype
	} else if attr.Atype != a.arrayElemType {
		return fmt.Errorf("cannot append %s to array of %s", attributeTypeNames[attr.Atype], attributeTypeNames[a.arrayElemType])
	}

	a.arrayVals = append(a.arrayVals, *attr)
	return nil
}

// Arr returns the array of attributes
func (a *Attribute) Arr() ([]Attribute, error) {
	if a.Atype != AttributeTypeArray {
		return nil, ErrNotAnArray
	}
	return a.arrayVals, nil
}

// String retrieves the string value
func (a *Attribute) String() (string, error) {
	if a.Atype != AttributeTypeString {
		return "", ErrInvalidType
	}
	return a.stringVal, nil
}

// Float retrieves the float value
func (a *Attribute) Float() (float64, error) {
	if a.Atype != AttributeTypeFloat {
		return 0, ErrInvalidType
	}
	return a.floatVal, nil
}

// Float2 retrieves both float2 values
func (a *Attribute) Float2() ([2]float64, error) {
	if a.Atype != AttributeTypeFloat2 {
		return [2]float64{}, ErrInvalidType
	}
	return a.float2Vals, nil
}

// Float3 retrieves all float3 values
func (a *Attribute) Float3() ([3]float64, error) {
	if a.Atype != AttributeTypeFloat3 {
		return [3]float64{}, ErrInvalidType
	}
	return a.float3Vals, nil
}

// Float4 retrieves all float4 values
func (a *Attribute) Float4() ([4]float64, error) {
	if a.Atype != AttributeTypeFloat4 {
		return [4]float64{}, ErrInvalidType
	}
	return a.float4Vals, nil
}

// Placement retrieves both position and rotation
func (a *Attribute) Placement() ([3]float64, [4]float64, error) {
	if a.Atype != AttributeTypePlacement {
		return [3]float64{}, [4]float64{}, ErrInvalidType
	}
	return a.placementPos, a.placementRot, nil
}

// Int retrieves the int value
func (a *Attribute) Int() (int64, error) {
	if a.Atype != AttributeTypeInt {
		return 0, ErrInvalidType
	}
	return a.intVal, nil
}

// Int2 retrieves both int2 values
func (a *Attribute) Int2() ([2]int64, error) {
	if a.Atype != AttributeTypeInt2 {
		return [2]int64{}, ErrInvalidType
	}
	return a.int2Vals, nil
}

// Int3 retrieves all int3 values
func (a *Attribute) Int3() ([3]int64, error) {
	if a.Atype != AttributeTypeInt3 {
		return [3]int64{}, ErrInvalidType
	}
	return a.int3Vals, nil
}

// Int4 retrieves all int4 values
func (a *Attribute) Int4() ([4]int64, error) {
	if a.Atype != AttributeTypeInt4 {
		return [4]int64{}, ErrInvalidType
	}
	return a.int4Vals, nil
}

// Bool retrieves the bool value
func (a *Attribute) Bool() (bool, error) {
	if a.Atype != AttributeTypeBool {
		return false, ErrInvalidType
	}
	return a.boolVal, nil
}

// TypeName returns the string name of the attribute type
func (a *Attribute) TypeName() string {
	name, exists := attributeTypeNames[a.Atype]
	if !exists {
		return "unknown"
	}
	return name
}

// Returns a printable representation of the attribute value
func (a *Attribute) Printable() string {
	switch a.Atype {
	case AttributeTypeString:
		return a.stringVal
	case AttributeTypeFloat:
		return fmt.Sprintf("%f", a.floatVal)
	case AttributeTypeInt:
		return fmt.Sprintf("%d", a.intVal)
	case AttributeTypeBool:
		return fmt.Sprintf("%t", a.boolVal)
	case AttributeTypeArray:
		return fmt.Sprintf("Array[%d] of %s", len(a.arrayVals), attributeTypeNames[a.arrayElemType])
	default:
		return fmt.Sprintf("<%s>", attributeTypeNames[a.Atype])
	}
}
