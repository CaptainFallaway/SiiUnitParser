package siiunit

import (
	"bufio"
	"io"
	"strings"
)

type unitDto struct {
	Utype string
	ID    string
	Body  []string
}

func parseDtos(content io.Reader) ([]*unitDto, error) {
	scanner := bufio.NewScanner(content)

	var currDto *unitDto
	var dtos []*unitDto

	inBlock := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.Contains(line, "{") && strings.Contains(line, " : ") {
			inBlock = true

			line = strings.TrimRight(line, " {")
			splitLine := strings.Split(line, " : ")

			currDto = &unitDto{
				Utype: splitLine[0],
				ID:    splitLine[1],
				Body:  make([]string, 0),
			}

			continue
		}

		if inBlock {
			currDto.Body = append(currDto.Body, line)
		}

		if strings.Contains(line, "}") && inBlock {
			dtos = append(dtos, currDto)
			currDto = nil
			inBlock = false
		}
	}

	return dtos, scanner.Err()
}
