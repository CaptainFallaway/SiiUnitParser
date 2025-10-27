package main

import (
	"bufio"
	"io"
	"strings"
)

// ReadAllDescriptors does panic (I'm to lazy to handle errors atm)
func ReadAllUnits(content io.Reader) ([]Unit, error) {
	scanner := bufio.NewScanner(content)

	var lines []Unit
	var currUnit *Unit
	var currAttrs []string
	var beginBlock = false

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if strings.Contains(line, "{") && strings.Contains(line, " : ") {
			beginBlock = true

			line = strings.TrimRight(line, " {")
			splitLine := strings.Split(line, " : ")

			currUnit = &Unit{
				utype: splitLine[0],
				uid:   splitLine[1],
			}

			continue
		}

		if currUnit == nil && currAttrs == nil {
			continue
		}

		if strings.Contains(line, "}") {
			lines = append(lines, *currUnit)
			currUnit = nil
		}

		if beginBlock {
			currUnit.attrs = append(currUnit.attrs, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
