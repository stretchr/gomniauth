package oauth2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApprovalPrompts(t *testing.T) {

	assert.Equal(t, "force", ApprovalPromptForce)
	assert.Equal(t, "auto", ApprovalPromptAuto)

}

func TestAccessTypes(t *testing.T) {

	assert.Equal(t, "online", AccessTypeOnline)
	assert.Equal(t, "offline", AccessTypeOffline)

}
