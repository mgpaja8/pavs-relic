package valueobjects

import "errors"

type CompanyName string

func NewCompanyName(companyName string) (CompanyName, error) {
	if len(companyName) > 60 {
		return "", errors.New("company name too long")
	}

	if len(companyName) < 2 {
		return "", errors.New("company name too short")
	}

	return CompanyName(companyName), nil
}

func (cn CompanyName) String() string {
	return string(cn)
}
