package siiunit

import (
	"testing"
)

// TestNewAttributeString tests string attribute creation and retrieval
func TestNewAttributeString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "simple string",
			input:   `"hello world"`,
			want:    "hello world",
			wantErr: false,
		},
		{
			name:    "empty string",
			input:   `""`,
			want:    "",
			wantErr: false,
		},
		{
			name:    "string with special chars",
			input:   `"path/to/resource.pma"`,
			want:    "path/to/resource.pma",
			wantErr: false,
		},
		{
			name:    "token (unquoted)",
			input:   "mytoken",
			want:    "mytoken",
			wantErr: false,
		},
		{
			name:    "owner pointer",
			input:   ".some.nameless.unit",
			want:    ".some.nameless.unit",
			wantErr: false,
		},
		{
			name:    "link pointer",
			input:   "some.named.unit",
			want:    "some.named.unit",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attr, err := newAttribute(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAttribute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if attr.Atype != AttributeTypeString {
				t.Errorf("Expected AttributeTypeString, got %v", attr.Atype)
			}

			val, err := attr.String()
			if err != nil {
				t.Errorf("String() error = %v", err)
				return
			}

			if val != tt.want {
				t.Errorf("String() = %q, want %q", val, tt.want)
			}
		})
	}
}

// TestNewAttributeFloat tests float attribute creation and retrieval
func TestNewAttributeFloat(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    float64
		wantErr bool
	}{
		{
			name:    "positive float",
			input:   "1.5",
			want:    1.5,
			wantErr: false,
		},
		{
			name:    "negative float",
			input:   "-3.14",
			want:    -3.14,
			wantErr: false,
		},
		{
			name:    "zero float",
			input:   "0.0",
			want:    0.0,
			wantErr: false,
		},
		{
			name:    "large float",
			input:   "1000.5",
			want:    1000.5,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attr, err := newAttribute(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAttribute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if attr.Atype != AttributeTypeFloat {
				t.Errorf("Expected AttributeTypeFloat, got %v", attr.Atype)
			}

			val, err := attr.Float()
			if err != nil {
				t.Errorf("Float() error = %v", err)
				return
			}

			if val != tt.want {
				t.Errorf("Float() = %v, want %v", val, tt.want)
			}
		})
	}
}

// TestNewAttributeFloat2 tests float2 attribute creation and retrieval
func TestNewAttributeFloat2(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantArray [2]float64
		wantErr   bool
	}{
		{
			name:      "basic float2",
			input:     "(1.0, 2.0)",
			wantArray: [2]float64{1.0, 2.0},
			wantErr:   false,
		},
		{
			name:      "negative float2",
			input:     "(-5.5, 3.2)",
			wantArray: [2]float64{-5.5, 3.2},
			wantErr:   false,
		},
		{
			name:      "zero float2",
			input:     "(0.0, 0.0)",
			wantArray: [2]float64{0.0, 0.0},
			wantErr:   false,
		},
		{
			name:      "float2 with spaces",
			input:     "( 1.5 , 2.5 )",
			wantArray: [2]float64{1.5, 2.5},
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attr, err := newAttribute(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAttribute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if attr.Atype != AttributeTypeFloat2 {
				t.Errorf("Expected AttributeTypeFloat2, got %v", attr.Atype)
			}

			vals, _ := attr.Float2()
			if vals != tt.wantArray {
				t.Errorf("Float2() = %v, want %v", vals, tt.wantArray)
			}
		})
	}
}

// TestNewAttributeFloat3 tests float3 attribute creation and retrieval
func TestNewAttributeFloat3(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantArray [3]float64
		wantErr   bool
	}{
		{
			name:      "basic float3",
			input:     "(1.0, 2.0, 3.0)",
			wantArray: [3]float64{1.0, 2.0, 3.0},
			wantErr:   false,
		},
		{
			name:      "negative float3",
			input:     "(-1.5, -2.5, -3.5)",
			wantArray: [3]float64{-1.5, -2.5, -3.5},
			wantErr:   false,
		},
		{
			name:      "mixed float3",
			input:     "(5.5, -10.2, 3.1)",
			wantArray: [3]float64{5.5, -10.2, 3.1},
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attr, err := newAttribute(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAttribute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if attr.Atype != AttributeTypeFloat3 {
				t.Errorf("Expected AttributeTypeFloat3, got %v", attr.Atype)
			}

			vals, _ := attr.Float3()
			if vals != tt.wantArray {
				t.Errorf("Float3() = %v, want %v", vals, tt.wantArray)
			}
		})
	}
}

// TestNewAttributeFloat4 tests float4 attribute creation and retrieval
func TestNewAttributeFloat4(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantArray [4]float64
		wantErr   bool
	}{
		{
			name:      "basic float4",
			input:     "(1.0, 2.0, 3.0, 4.0)",
			wantArray: [4]float64{1.0, 2.0, 3.0, 4.0},
			wantErr:   false,
		},
		{
			name:      "quaternion-like float4",
			input:     "(1.0, 0.0, 0.0, 0.0)",
			wantArray: [4]float64{1.0, 0.0, 0.0, 0.0},
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attr, err := newAttribute(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAttribute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if attr.Atype != AttributeTypeFloat4 {
				t.Errorf("Expected AttributeTypeFloat4, got %v", attr.Atype)
			}

			vals, _ := attr.Float4()
			if vals != tt.wantArray {
				t.Errorf("Float4() = %v, want %v", vals, tt.wantArray)
			}
		})
	}
}

// TestNewAttributeInt tests int attribute creation and retrieval
func TestNewAttributeInt(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int64
		wantT   AttributeType
		wantErr bool
	}{
		{
			name:    "positive int",
			input:   "42",
			want:    42,
			wantT:   AttributeTypeInt,
			wantErr: false,
		},
		{
			name:    "negative int",
			input:   "-100",
			want:    -100,
			wantT:   AttributeTypeInt,
			wantErr: false,
		},
		{
			name:    "zero int",
			input:   "0",
			want:    0,
			wantT:   AttributeTypeInt,
			wantErr: false,
		},
		{
			name:    "large int",
			input:   "9223372036854775800",
			want:    9223372036854775800,
			wantT:   AttributeTypeInt,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attr, err := newAttribute(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAttribute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if attr.Atype != tt.wantT {
				t.Errorf("Expected type %v, got type %v", tt.wantT, attr.Atype)
			}

			val, err := attr.Int()
			if err != nil {
				t.Errorf("Int() error = %v", err)
				return
			}

			if val != tt.want {
				t.Errorf("Int() = %v, want %v", val, tt.want)
			}
		})
	}
}

// TestNewAttributeInt2 tests int2 attribute creation and retrieval
func TestNewAttributeInt2(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantArray [2]int64
		wantErr   bool
	}{
		{
			name:      "basic int2",
			input:     "(10, 22)",
			wantArray: [2]int64{10, 22},
			wantErr:   false,
		},
		{
			name:      "negative int2",
			input:     "(-5, 15)",
			wantArray: [2]int64{-5, 15},
			wantErr:   false,
		},
		{
			name:      "zero int2",
			input:     "(0, 0)",
			wantArray: [2]int64{0, 0},
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attr, err := newAttribute(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAttribute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if attr.Atype != AttributeTypeInt2 {
				t.Errorf("Expected AttributeTypeInt2, got %v", attr.Atype)
			}

			vals, _ := attr.Int2()
			if vals != tt.wantArray {
				t.Errorf("Int2() = %v, want %v", vals, tt.wantArray)
			}
		})
	}
}

// TestNewAttributeInt3 tests int3 attribute creation and retrieval
func TestNewAttributeInt3(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantArray [3]int64
		wantErr   bool
	}{
		{
			name:      "basic int3",
			input:     "(10, 22, 33)",
			wantArray: [3]int64{10, 22, 33},
			wantErr:   false,
		},
		{
			name:      "negative int3",
			input:     "(-1, -2, -3)",
			wantArray: [3]int64{-1, -2, -3},
			wantErr:   false,
		},
		{
			name:      "mixed int3",
			input:     "(5, -10, 3)",
			wantArray: [3]int64{5, -10, 3},
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attr, err := newAttribute(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAttribute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if attr.Atype != AttributeTypeInt3 {
				t.Errorf("Expected AttributeTypeInt3, got %v", attr.Atype)
			}

			vals, _ := attr.Int3()
			if vals != tt.wantArray {
				t.Errorf("Int3() = %v, want %v", vals, tt.wantArray)
			}
		})
	}
}

// TestNewAttributeInt4 tests int4 attribute creation and retrieval
func TestNewAttributeInt4(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantArray [4]int64
		wantErr   bool
	}{
		{
			name:      "basic int4",
			input:     "(10, 22, 33, 44)",
			wantArray: [4]int64{10, 22, 33, 44},
			wantErr:   false,
		},
		{
			name:      "negative int4",
			input:     "(-1, -2, -3, -4)",
			wantArray: [4]int64{-1, -2, -3, -4},
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attr, err := newAttribute(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAttribute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if attr.Atype != AttributeTypeInt4 {
				t.Errorf("Expected AttributeTypeInt4, got %v", attr.Atype)
			}

			vals, _ := attr.Int4()
			if vals != tt.wantArray {
				t.Errorf("Int4() = %v, want %v", vals, tt.wantArray)
			}
		})
	}
}

// TestNewAttributePlacement tests placement attribute creation and retrieval
func TestNewAttributePlacement(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantPos [3]float64
		wantRot [4]float64
		wantErr bool
	}{
		{
			name:    "basic placement",
			input:   "(0, 0, 0) (1; 0, 0, 0)",
			wantPos: [3]float64{0.0, 0.0, 0.0},
			wantRot: [4]float64{1.0, 0.0, 0.0, 0.0},
			wantErr: false,
		},
		{
			name:    "placement with position",
			input:   "(5.5, 10.2, -3.1) (1; 0, 0, 0)",
			wantPos: [3]float64{5.5, 10.2, -3.1},
			wantRot: [4]float64{1.0, 0.0, 0.0, 0.0},
			wantErr: false,
		},
		{
			name:    "placement with rotation",
			input:   "(0, 0, 0) (0.7071; 0.7071, 0, 0)",
			wantPos: [3]float64{0.0, 0.0, 0.0},
			wantRot: [4]float64{0.7071, 0.7071, 0.0, 0.0},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attr, err := newAttribute(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAttribute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if attr.Atype != AttributeTypePlacement {
				t.Errorf("Expected AttributeTypePlacement, got %v", attr.Atype)
			}

			pos, rot, _ := attr.Placement()

			const epsilon = 0.0001
			if !floatArraysEqual(pos[:], tt.wantPos[:], epsilon) {
				t.Errorf("Placement pos = %v, want %v", pos, tt.wantPos)
			}

			if !floatArraysEqual(rot[:], tt.wantRot[:], epsilon) {
				t.Errorf("Placement rot = %v, want %v", rot, tt.wantRot)
			}
		})
	}
}

// TestNewAttributeBool tests bool attribute creation and retrieval
func TestNewAttributeBool(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    bool
		wantErr bool
	}{
		{
			name:    "true",
			input:   "true",
			want:    true,
			wantErr: false,
		},
		{
			name:    "false",
			input:   "false",
			want:    false,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attr, err := newAttribute(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAttribute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if attr.Atype != AttributeTypeBool {
				t.Errorf("Expected AttributeTypeBool, got %v", attr.Atype)
			}

			val, err := attr.Bool()
			if err != nil {
				t.Errorf("Bool() error = %v", err)
				return
			}

			if val != tt.want {
				t.Errorf("Bool() = %v, want %v", val, tt.want)
			}
		})
	}
}

// TestTypeDetection tests correct type detection for edge cases
func TestTypeDetection(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantType AttributeType
	}{
		{
			name:     "string vs token",
			input:    "mytoken",
			wantType: AttributeTypeString,
		},
		{
			name:     "string vs owner ptr",
			input:    ".namespace.item",
			wantType: AttributeTypeString,
		},
		{
			name:     "string vs link ptr",
			input:    "namespace.item.part",
			wantType: AttributeTypeString,
		},
		{
			name:     "float2 vs int2",
			input:    "(1.5, 2.5)",
			wantType: AttributeTypeFloat2,
		},
		{
			name:     "int2 parse",
			input:    "(1, 2)",
			wantType: AttributeTypeInt2,
		},
		{
			name:     "negative triggers int",
			input:    "-42",
			wantType: AttributeTypeInt,
		},
		{
			name:     "positive triggers uint",
			input:    "42",
			wantType: AttributeTypeInt,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attr, err := newAttribute(tt.input)
			if err != nil {
				t.Errorf("NewAttribute() error = %v", err)
				return
			}

			if attr.Atype != tt.wantType {
				t.Errorf("Expected %v, got %v", tt.wantType, attr.Atype)
			}
		})
	}
}

// TestArrayOperations tests array creation and appending
func TestArrayOperations(t *testing.T) {
	t.Run("create array and append floats", func(t *testing.T) {
		attr := &Attribute{}
		err := attr.makeArray(3)
		if err != nil {
			t.Errorf("makeArray() error = %v", err)
			return
		}

		if attr.Atype != AttributeTypeArray {
			t.Error("Expected attribute to be an array")
		}

		// Append float values
		err = attr.appendToArray("1.5")
		if err != nil {
			t.Errorf("AppendToArray() error = %v", err)
		}

		err = attr.appendToArray("2.5")
		if err != nil {
			t.Errorf("AppendToArray() error = %v", err)
		}

		err = attr.appendToArray("3.5")
		if err != nil {
			t.Errorf("AppendToArray() error = %v", err)
		}

		arr, _ := attr.Arr()
		if len(arr) != 3 {
			t.Errorf("Arr() length = %d, want 3", len(arr))
		}

		if len(arr) != 3 {
			t.Errorf("ArrayLen() = %d, want 3", len(arr))
		}

		// Verify values
		for i, expected := range []float64{1.5, 2.5, 3.5} {
			val, err := arr[i].Float()
			if err != nil {
				t.Errorf("Float() error at index %d: %v", i, err)
			}
			if val != expected {
				t.Errorf("arr[%d].Float() = %v, want %v", i, val, expected)
			}
		}
	})

	t.Run("create array and append ints", func(t *testing.T) {
		attr := &Attribute{}
		attr.makeArray(2)

		err := attr.appendToArray("100")
		if err != nil {
			t.Errorf("AppendToArray() error = %v", err)
		}

		err = attr.appendToArray("-50")
		if err != nil {
			t.Errorf("AppendToArray() error = %v", err)
		}

		arr, _ := attr.Arr()
		for i, expected := range []int64{100, -50} {
			val, err := arr[i].Int()
			if err != nil {
				t.Errorf("Int() error at index %d: %v", i, err)
			}
			if val != expected {
				t.Errorf("arr[%d].Int() = %v, want %v", i, val, expected)
			}
		}
	})

	t.Run("append wrong type to array", func(t *testing.T) {
		attr := &Attribute{}
		attr.makeArray(2)

		// Append a float first to set the type
		err := attr.appendToArray("1.5")
		if err != nil {
			t.Errorf("AppendToArray() error = %v", err)
		}

		// Try to append an int - should fail
		err = attr.appendToArray("42")
		if err == nil {
			t.Error("Expected error when appending int to float array, got nil")
		}
	})

	t.Run("append to non-array", func(t *testing.T) {
		attr := &Attribute{}
		attr.Atype = AttributeTypeFloat
		attr.floatVal = 5.0

		err := attr.appendToArray("1.5")
		if err != ErrNotAnArray {
			t.Errorf("Expected ErrNotAnArray, got %v", err)
		}
	})

	t.Run("Arr() on non-array", func(t *testing.T) {
		attr := &Attribute{}
		attr.Atype = AttributeTypeFloat
		attr.floatVal = 5.0

		_, err := attr.Arr()
		if err != ErrNotAnArray {
			t.Errorf("Expected ErrNotAnArray, got %v", err)
		}
	})
}

// TestTypeErrors tests calling getters on wrong types
func TestTypeErrors(t *testing.T) {
	tests := []struct {
		name    string
		attr    *Attribute
		getter  func(*Attribute) error
		wantErr bool
	}{
		{
			name: "Float() on string attribute",
			attr: func() *Attribute {
				a, _ := newAttribute(`"hello"`)
				return a
			}(),
			getter: func(a *Attribute) error {
				_, err := a.Float()
				return err
			},
			wantErr: true,
		},
		{
			name: "String() on float attribute",
			attr: func() *Attribute {
				a, _ := newAttribute("1.5")
				return a
			}(),
			getter: func(a *Attribute) error {
				_, err := a.String()
				return err
			},
			wantErr: true,
		},
		{
			name: "Bool() on int attribute",
			attr: func() *Attribute {
				a, _ := newAttribute("-42")
				return a
			}(),
			getter: func(a *Attribute) error {
				_, err := a.Bool()
				return err
			},
			wantErr: true,
		},
		{
			name: "Float2() on float attribute",
			attr: func() *Attribute {
				a, _ := newAttribute("1.5")
				return a
			}(),
			getter: func(a *Attribute) error {
				_, err := a.Float2()
				return err
			},
			wantErr: true,
		},
		{
			name: "Int() on uint attribute",
			attr: func() *Attribute {
				a, _ := newAttribute("42")
				return a
			}(),
			getter: func(a *Attribute) error {
				_, err := a.Int()
				return err
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.getter(tt.attr)
			if (err != nil) != tt.wantErr {
				t.Errorf("getter error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestTypeName tests the TypeName method
func TestTypeName(t *testing.T) {
	tests := []struct {
		attr *Attribute
		want string
	}{
		{
			attr: func() *Attribute {
				a, _ := newAttribute(`"hello"`)
				return a
			}(),
			want: "string",
		},
		{
			attr: func() *Attribute {
				a, _ := newAttribute("1.5")
				return a
			}(),
			want: "float",
		},
		{
			attr: func() *Attribute {
				a, _ := newAttribute("true")
				return a
			}(),
			want: "bool",
		},
		{
			attr: func() *Attribute {
				a, _ := newAttribute("-42")
				return a
			}(),
			want: "int",
		},
		{
			attr: func() *Attribute {
				a, _ := newAttribute("(1, 2)")
				return a
			}(),
			want: "int2",
		},
	}

	for _, tt := range tests {
		name := tt.attr.TypeName()
		if name != tt.want {
			t.Errorf("TypeName() = %q, want %q", name, tt.want)
		}
	}
}

// TestPlacementComponents tests individual placement component access
func TestPlacementComponents(t *testing.T) {
	attr, _ := newAttribute("(5.5, 10.2, -3.1) (1; 0, 0, 0)")

	t.Run("PlacementPos", func(t *testing.T) {
		pos, _, err := attr.Placement()
		if err != nil {
			t.Errorf("PlacementPos() error = %v", err)
			return
		}

		wantPos := [3]float64{5.5, 10.2, -3.1}
		if pos != wantPos {
			t.Errorf("PlacementPos() = %v, want %v", pos, wantPos)
		}
	})

	t.Run("PlacementRot", func(t *testing.T) {
		_, rot, err := attr.Placement()
		if err != nil {
			t.Errorf("PlacementRot() error = %v", err)
			return
		}

		wantRot := [4]float64{1.0, 0.0, 0.0, 0.0}
		if rot != wantRot {
			t.Errorf("PlacementRot() = %v, want %v", rot, wantRot)
		}
	})
}

// Helper function for floating point array comparison
func floatArraysEqual(a, b []float64, epsilon float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if absFloat64(a[i]-b[i]) > epsilon {
			return false
		}
	}
	return true
}

// Helper function for absolute value of float64
func absFloat64(f float64) float64 {
	if f < 0 {
		return -f
	}
	return f
}
