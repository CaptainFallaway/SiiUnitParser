package siiunit

import (
	"regexp"
	"strconv"
	"strings"
)

// detectAttributeType determines the attribute type from a string value
func detectAttributeType(value string) AttributeType {
	value = strings.TrimSpace(value)

	// String type (quoted)
	if strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`) {
		return AttributeTypeString
	}

	// Boolean
	if value == "true" || value == "false" {
		return AttributeTypeBool
	}

	// Placement format: (x, y, z) (w; x, y, z)
	if isPlacementFormat(value) {
		return AttributeTypePlacement
	}

	// Int4: (x, y, z, w) - all integers
	if isInt4Format(value) {
		return AttributeTypeInt4
	}

	// Int3: (x, y, z) - all integers
	if isInt3Format(value) {
		return AttributeTypeInt3
	}

	// Int2: (x, y) - both integers
	if isInt2Format(value) {
		return AttributeTypeInt2
	}

	// Float4: (x, y, z, w)
	if isFloat4Format(value) {
		return AttributeTypeFloat4
	}

	// Float3: (x, y, z)
	if isFloat3Format(value) {
		return AttributeTypeFloat3
	}

	// Float2: (x, y)
	if isFloat2Format(value) {
		return AttributeTypeFloat2
	}

	// IEEE754 hex float (starts with &)
	if strings.HasPrefix(value, "&") {
		return AttributeTypeFloat
	}

	// Numeric integer - try to parse as int or uint
	if _, err := strconv.ParseInt(value, 10, 64); err == nil {
		return AttributeTypeInt
	}

	// Numeric float
	if _, err := strconv.ParseFloat(value, 64); err == nil {
		return AttributeTypeFloat
	}

	// Anything else is a string (includes tokens, owner_ptr, link_ptr, resource_tie)
	return AttributeTypeString
}

// Helper functions for format detection
func isPlacementFormat(s string) bool {
	re := regexp.MustCompile(`^\s*\(\s*-?\d+(?:\.\d+)?\s*,\s*-?\d+(?:\.\d+)?\s*,\s*-?\d+(?:\.\d+)?\s*\)\s*\(\s*-?\d+(?:\.\d+)?\s*;\s*-?\d+(?:\.\d+)?\s*,\s*-?\d+(?:\.\d+)?\s*,\s*-?\d+(?:\.\d+)?\s*\)\s*$`)
	return re.MatchString(s)
}

func isFloat4Format(s string) bool {
	parts := extractTupleValues(s)
	if len(parts) != 4 {
		return false
	}
	for _, p := range parts {
		if _, err := strconv.ParseFloat(p, 64); err != nil {
			return false
		}
	}
	return true
}

func isInt4Format(s string) bool {
	parts := extractTupleValues(s)
	if len(parts) != 4 {
		return false
	}
	for _, p := range parts {
		if _, err := strconv.ParseInt(p, 10, 64); err != nil {
			return false
		}
	}
	return true
}

func isFloat3Format(s string) bool {
	parts := extractTupleValues(s)
	if len(parts) != 3 {
		return false
	}
	for _, p := range parts {
		if _, err := strconv.ParseFloat(p, 64); err != nil {
			return false
		}
	}
	return true
}

func isInt3Format(s string) bool {
	parts := extractTupleValues(s)
	if len(parts) != 3 {
		return false
	}
	for _, p := range parts {
		if _, err := strconv.ParseInt(p, 10, 64); err != nil {
			return false
		}
	}
	return true
}

func isInt2Format(s string) bool {
	parts := extractTupleValues(s)
	if len(parts) != 2 {
		return false
	}
	for _, p := range parts {
		if _, err := strconv.ParseInt(p, 10, 64); err != nil {
			return false
		}
	}
	return true
}

func isFloat2Format(s string) bool {
	parts := extractTupleValues(s)
	if len(parts) != 2 {
		return false
	}
	for _, p := range parts {
		if _, err := strconv.ParseFloat(p, 64); err != nil {
			return false
		}
	}
	return true
}
