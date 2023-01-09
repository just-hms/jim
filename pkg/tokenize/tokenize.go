package tokenize

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

const (
	ARGUMENT_PATTERN = "(?:<)([^<>]*)(?:>)"
	OR_CHAR          = "|"
)

func Tokenize(text string, args []string) (string, error) {

	regex, err := regexp.Compile(ARGUMENT_PATTERN)

	if err != nil {
		return "", errors.New("Error generating the regex")
	}

	resultString := text

	for _, match := range regex.FindAllStringSubmatch(text, -1) {

		fullMatch := match[0]
		captureGroup := match[1]

		splitted := strings.Split(captureGroup, OR_CHAR)
		index, err := strconv.Atoi(splitted[0])

		if err != nil {
			return "", errors.New("Error parsing argument index")
		}

		// if it is passed replace it with the value
		if index < len(args) {
			resultString = strings.ReplaceAll(resultString, fullMatch, args[index])
			continue
		}

		// if not check if the default is provided
		if len(splitted) != 2 {
			return "", errors.New("Index out of bounds")
		}

		// if yes substitute the placeholder with the default value
		resultString = strings.ReplaceAll(resultString, fullMatch, splitted[1])
		continue

	}

	return resultString, nil
}
