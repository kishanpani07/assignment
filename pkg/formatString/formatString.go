package formatString

import (
	"strings"
)

func FormatString(rawString string) []Dependency {
	var dependencies []Dependency = []Dependency{}

	rawStringSplit := strings.Split(rawString, newLineString)
	var filteredRawStringSplit []string

	for _, rawStringPart := range rawStringSplit {
		if strings.Contains(rawStringPart, arrowString) {
			filteredRawStringSplit = append(filteredRawStringSplit, rawStringPart)
		}
	}

	for _, filteredRawStringPart := range filteredRawStringSplit {
		var filteredRawStringPartCleaned string
		if strings.Contains(filteredRawStringPart, squareBracketStartString) {
			filteredRawStringPartCleaned = filteredRawStringPart[:strings.Index(filteredRawStringPart, squareBracketStartString)]
		} else {
			filteredRawStringPartCleaned = filteredRawStringPart
		}
		filteredRawStringPartSplit := strings.Split(filteredRawStringPartCleaned, arrowSeparatorString)
		parent, child := filteredRawStringPartSplit[0], filteredRawStringPartSplit[1]
		dependency := Dependency{
			Parent: removeDoubleQuotes(parent),
			Child:  removeDoubleQuotes(child),
		}
		dependencies = append(dependencies, dependency)
	}

	return dependencies
}

func removeDoubleQuotes(rawString string) string {
	return strings.Trim(strings.ReplaceAll(rawString, doubleQuotesString, nullString), spaceString)
}
