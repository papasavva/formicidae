package formicidae

import (
	"errors"
	"io/ioutil"
	"strings"
)

func UpdateVariable(filename string, name string, value string) (string, error) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", errors.New("cannot read the file")
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		const separator = "="

		if strings.Contains(line, separator) {
			parts := strings.Split(line, separator)

			if name == parts[0] {
				lines[i] = parts[0] + separator + value
			}
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(filename, []byte(output), 0644)
	
	if err != nil {
		return "", errors.New("cannot write to the file")
	}

	return output, nil
}
