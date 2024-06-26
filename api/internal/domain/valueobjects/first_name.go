package valueobjects

import "errors"

type FirstName string

func NewFirstName(firstName string) (FirstName, error) {
	if len(firstName) > 20 {
		return "", errors.New("first name too long")
	}

	if len(firstName) < 2 {
		return "", errors.New("first name too short")
	}

	return FirstName(firstName), nil
}

func (fn FirstName) String() string {
	return string(fn)
}
