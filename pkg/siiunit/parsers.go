package siiunit

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func makeArray(line, prevLine string, currAttrs *Attributes) error {
	prevLineSplit := strings.Split(prevLine, ": ")
	arrKey := prevLineSplit[0]

	attr := currAttrs.attrs[arrKey]

	// Check if the attribute is already an array or not
	if attr.Atype != AttributeTypeArray {
		arrSize, err := strconv.Atoi(prevLineSplit[1])
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

// ReadAllDescriptors does panic (I'm to lazy to handle errors atm)
func ParseAllUnits(content io.Reader) ([]Unit, error) {
	scanner := bufio.NewScanner(content)

	var units []Unit

	var prevLine string
	var firstArrLine string
	var currUnit *Unit
	var currAttrs *Attributes
	var beginBlock = false

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if strings.Contains(line, "{") && strings.Contains(line, " : ") {
			beginBlock = true

			line = strings.TrimRight(line, " {")
			splitLine := strings.Split(line, " : ")

			currUnit = &Unit{
				Utype: splitLine[0],
				ID:    splitLine[1],
			}

			currAttrs = newAttributes()

			continue
		}

		if strings.Contains(line, "}") && beginBlock {
			currUnit.Attrs = *currAttrs
			units = append(units, *currUnit)
			currUnit = nil
			beginBlock = false
		}

		if beginBlock && containsArrSyntax(line) {
			if !containsArrSyntax(prevLine) {
				firstArrLine = prevLine
			}

			err := makeArray(line, firstArrLine, currAttrs)
			if err != nil {
				return nil, err
			}

			continue
		}

		if beginBlock {
			if !strings.Contains(line, ": ") {
				continue
			}

			splitLine := strings.Split(line, ": ")

			currAttrs.addAttribute(splitLine[0], splitLine[1])
		}

		prevLine = line
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return units, nil
}
