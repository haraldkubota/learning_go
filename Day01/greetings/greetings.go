package greetings
import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {

	// If no name was give, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	// Failing test
	//message := fmt.Sprintf(randomFormat(), "xx")
	return message, nil;
}

// Hellos returns a map tat associates each of the named people with a greeting message

func Hellos(names []string) (map[string]string, error) {
	// Create the map
	messages := make(map[string]string)
	
	// Loop through names
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string {
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}