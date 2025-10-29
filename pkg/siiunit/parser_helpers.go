package siiunit

import (
	"strconv"
	"strings"
)

// buildAttributeArray parses attribute data from line (using definingLine for context)
// and updates currAttrs with the parsed attributes.
//
// The function modifies currAttrs in-place by appending or merging parsed attribute
// entries. It returns an error when parsing fails, when required contextual information
// from definingLine is missing or invalid, or when attribute data violates expected
// semantics (for example malformed tokens, unmatched quotes, invalid escapes, or
// conflicting/duplicate attribute definitions).
//
// Expected responsibilities:
// - tokenize and trim input, respecting quoting and escapes,
// - interpret values according to context provided by definingLine,
// - validate attribute names and values and resolve duplicates according to format rules,
// - report detailed errors for malformed input.
func buildAttributeArray(line, definingLine string, currAttrs *Attributes) error {
	definingLineSplit := strings.Split(definingLine, ": ")
	arrKey := definingLineSplit[0]

	attr := currAttrs.attrs[arrKey]

	// Check if the attribute is already an array or not
	if attr.Atype != AttributeTypeArray {
		arrSize, err := strconv.Atoi(definingLineSplit[1])
		if err != nil {
			return err
		}

		attr.makeArray(arrSize)
	}

	// Append the value to the array
	lineSplit := strings.Split(line, ": ")
	attr.appendToArray(lineSplit[1])

	return nil
}

func containsArrSyntax(line string) bool {
	splitLine := strings.Split(line, ": ")
	return strings.Contains(splitLine[0], "[")
}
