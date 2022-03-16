package screens

import "regexp"

type ExpectedVariable struct {
	Variable string
	Regex    *string
}

func (ex ExpectedVariable) IsValid(response string) bool {
	if ex.Regex == nil {
		return true
	}
	match, error := regexp.MatchString(*ex.Regex, response)
	if error != nil {
		return false
	}
	return match
}
func (ex ExpectedVariable) Compare(ex2 ExpectedVariable) bool {
	return ex.Variable == ex2.Variable
}
