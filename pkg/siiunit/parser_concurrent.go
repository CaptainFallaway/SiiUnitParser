package siiunit

import (
	"io"
	"strings"

	"golang.org/x/sync/errgroup"
)

// ParseAllUnitsConcurrent parses all units from the provided content using concurrent workers.
func ParseAllUnitsConcurrent(content io.Reader, opts ...ParserOption) ([]Unit, error) {
	options := getOptions(opts...)

	group := new(errgroup.Group)
	group.SetLimit(options.workerCount)

	unitDtos, err := parseDtos(content)
	if err != nil {
		return nil, err
	}

	units := make([]Unit, len(unitDtos))

	for i, dto := range unitDtos {
		group.Go(func() error {
			unit, err := parseUnitFromDto(dto)
			if err != nil {
				return err
			}

			units[i] = unit
			return nil
		})
	}

	return units, group.Wait()
}

func parseUnitFromDto(dto *unitDto) (Unit, error) {
	unit := Unit{
		Utype: dto.Utype,
		ID:    dto.ID,
		Attrs: *newAttributes(),
	}

	var prevLine string
	var definingFirstArrLine string // This will track the first line that defines an array attribute for multi-line arrays

	for _, line := range dto.Body {
		if containsArrSyntax(line) {
			if !containsArrSyntax(prevLine) {
				definingFirstArrLine = prevLine
			}

			err := buildAttributeArray(line, definingFirstArrLine, &unit.Attrs)
			if err != nil {
				return Unit{}, err
			}
		} else if strings.Contains(line, ": ") {
			splitLine := strings.Split(line, ": ")

			unit.Attrs.addAttribute(splitLine[0], splitLine[1])

			prevLine = line
		}
	}

	return unit, nil
}
