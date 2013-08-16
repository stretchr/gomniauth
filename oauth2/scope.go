package oauth2

import (
	"strings"
)

func ParseScope(scope string) string {
	scope = strings.Replace(scope, ",", " ", -1)
	parts := strings.Split(scope, " ")
	cleaned := make([]string, 0)
	for _, part := range parts {
		if len(part) != 0 {
			cleaned = append(cleaned, part)
		}
	}

	return strings.Join(cleaned, " ")
}

func MergeScopes(scopes ...string) string {

	for i, scope := range scopes {
		scopes[i] = ParseScope(scope)
	}

	return strings.TrimSpace(strings.Join(scopes, " "))

}
