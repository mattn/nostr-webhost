package relays

import (
	"os"
	"strings"
)

const PATH = ".nostr_relays"

func AddRelay(relayURL string) error {
	file, err := os.OpenFile(PATH, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(relayURL + "\n")
	return err
}

func RemoveRelay(targetURL string) error {
	content, err := os.ReadFile(PATH)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")
	var newLines []string

	for _, line := range lines {
		if line != targetURL {
			newLines = append(newLines, line)
		}
	}

	newContent := strings.Join(newLines, "\n")
	err = os.WriteFile(PATH, []byte(newContent), 0644)

	return err
}

func GetAllRelays() ([]string, error) {
	content, err := os.ReadFile(PATH)
	if err != nil {
		return nil, err
	}

	data := strings.Split(string(content), "\n")
	lines := []string{}
	for _, line := range data {
		if len(line) > 1 {
			lines = append(lines, line)
		}
	}
	return lines, nil
}
