package common

import (
	"fmt"
	"github.com/stretchr/objx"
	"strconv"
)

const (
	CredentialsKeyID string = "id"
)

// Credentials represent data that describes information needed
// to make authenticated requests.
type Credentials struct {
	objx.Map
}

var EmptyCredentials *Credentials = nil

// PublicData gets the storable data from this credentials object.
func (c *Credentials) PublicData(options map[string]interface{}) (publicData interface{}, err error) {

	// ensure the ID is a string
	idValue := c.Map.Get(CredentialsKeyID).Data()
	var idStringValue string
	switch idValue.(type) {
	case float64:
		idStringValue = strconv.FormatFloat(idValue.(float64), 'g', -1, 64)
	case string:
		idStringValue = idValue.(string)
	default:
		idStringValue = fmt.Sprintf("%v", idValue)
	}
	c.Map.Set(CredentialsKeyID, idStringValue)

	return c.Map, nil
}
