package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // Return a greeting that embeds the name in a message.
	if name == "" {
		return "", errors.New("empty name")
	}
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

// Hello returns a greeting for the named person.
func Hellos(names []string) (map[string]string, error) {
    // Keep track of message for every name in names
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {

	formats := []string {
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]

}