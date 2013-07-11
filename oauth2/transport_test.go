package oauth2

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestTransport_Transport(t *testing.T) {

	transport := &Transport{
		Transport: nil,
	}

	assert.Equal(t, http.DefaultTransport, transport.transport())

	myTransport := new(http.Transport)
	transport = &Transport{
		Transport: myTransport,
	}

	assert.Equal(t, myTransport, transport.transport())

}

func TestTransport_Client(t *testing.T) {

	myTransport := new(http.Transport)

	transport := &Transport{
		Transport: myTransport,
	}

	var client *http.Client = transport.Client()

	if assert.NotNil(t, client) {
		assert.Equal(t, client.Transport, transport)
	}

}
