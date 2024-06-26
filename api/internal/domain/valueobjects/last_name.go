package valueobjects

import "errors"

type LastName string

func NewLastName(lastName string) (LastName, error) {
	if len(lastName) > 20 {
		return "", errors.New("last name too long")
	}

	if len(lastName) < 2 {
		return "", errors.New("last name too short")
	}

	return LastName(lastName), nil
}

func (ln LastName) String() string {
	return string(ln)
}
