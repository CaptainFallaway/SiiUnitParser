package siiunit

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func extractTupleValues(s string) []string {
	re := regexp.MustCompile(`\(([^)]+)\)`)
	matches := re.FindStringSubmatch(s)
	if len(matches) < 2 {
		return []string{}
	}

	parts := strings.Split(matches[1], ",")
	var result []string
	for _, p := range parts {
		trimmed := strings.TrimSpace(p)
		if trimmed != "" && trimmed != ";" {
			result = append(result, trimmed)
		}
	}

	return result
}

// parseValue parses the string value according to the attribute type
func (a *Attribute) parseValue(value string) error {
	value = strings.TrimSpace(value)

	switch a.Atype {
	case AttributeTypeString:
		// Remove quotes if present
		if strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`) {
			a.stringVal = value[1 : len(value)-1]
		} else {
			a.stringVal = value
		}

	case AttributeTypeFloat:
		if strings.HasPrefix(value, "&") {
			// IEEE754 hex format
			hexVal := value[1:]
			i64, err := strconv.ParseInt(hexVal, 16, 32)
			if err != nil {
				return fmt.Errorf("%w: %v", ErrParsingFailed, err)
			}
			a.floatVal = float64(math.Float32frombits(uint32(i64)))
		} else {
			f, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return fmt.Errorf("%w: %v", ErrParsingFailed, err)
			}
			a.floatVal = f
		}

	case AttributeTypeFloat2:
		vals := extractTupleValues(value)
		if len(vals) != 2 {
			return ErrParsingFailed
		}
		for i, v := range vals {
			f, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return fmt.Errorf("%w: %v", ErrParsingFailed, err)
			}
			a.float2Vals[i] = f
		}

	case AttributeTypeFloat3:
		vals := extractTupleValues(value)
		if len(vals) != 3 {
			return ErrParsingFailed
		}
		for i, v := range vals {
			f, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return fmt.Errorf("%w: %v", ErrParsingFailed, err)
			}
			a.float3Vals[i] = f
		}

	case AttributeTypeFloat4:
		vals := extractTupleValues(value)
		if len(vals) != 4 {
			return ErrParsingFailed
		}
		for i, v := range vals {
			f, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return fmt.Errorf("%w: %v", ErrParsingFailed, err)
			}
			a.float4Vals[i] = f
		}

	case AttributeTypePlacement:
		return a.parsePlacement(value)

	case AttributeTypeInt:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrParsingFailed, err)
		}
		a.intVal = i

	case AttributeTypeInt2:
		vals := extractTupleValues(value)
		if len(vals) != 2 {
			return ErrParsingFailed
		}
		for i, v := range vals {
			n, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return fmt.Errorf("%w: %v", ErrParsingFailed, err)
			}
			a.int2Vals[i] = n
		}

	case AttributeTypeInt3:
		vals := extractTupleValues(value)
		if len(vals) != 3 {
			return ErrParsingFailed
		}
		for i, v := range vals {
			n, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return fmt.Errorf("%w: %v", ErrParsingFailed, err)
			}
			a.int3Vals[i] = n
		}

	case AttributeTypeInt4:
		vals := extractTupleValues(value)
		if len(vals) != 4 {
			return ErrParsingFailed
		}
		for i, v := range vals {
			n, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return fmt.Errorf("%w: %v", ErrParsingFailed, err)
			}
			a.int4Vals[i] = n
		}

	case AttributeTypeBool:
		a.boolVal = value == "true"
	}

	return nil
}

func (a *Attribute) parsePlacement(value string) error {
	// Format: (x, y, z) (w; x, y, z)
	re := regexp.MustCompile(`^\s*\(\s*(-?\d+(?:\.\d+)?)\s*,\s*(-?\d+(?:\.\d+)?)\s*,\s*(-?\d+(?:\.\d+)?)\s*\)\s*\(\s*(-?\d+(?:\.\d+)?)\s*;\s*(-?\d+(?:\.\d+)?)\s*,\s*(-?\d+(?:\.\d+)?)\s*,\s*(-?\d+(?:\.\d+)?)\s*\)\s*$`)
	matches := re.FindStringSubmatch(value)

	if len(matches) != 8 {
		return ErrParsingFailed
	}

	// Parse position
	for i := 1; i <= 3; i++ {
		f, err := strconv.ParseFloat(matches[i], 64)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrParsingFailed, err)
		}
		a.placementPos[i-1] = f
	}

	// Parse rotation (quaternion)
	for i := 4; i <= 7; i++ {
		f, err := strconv.ParseFloat(matches[i], 64)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrParsingFailed, err)
		}
		a.placementRot[i-4] = f
	}

	return nil
}
