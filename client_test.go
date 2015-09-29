package clarifai

import (
	"testing"

	"github.com/cheekybits/is"
)

var tests = []struct {
	input_ID        string
	expected_ID     string
	input_Secret    string
	expected_Secret string
}{
	{"testID", "testID", "testSecret", "testSecret"},
}

func TestClientInit(t *testing.T) {
	is := is.New(t)
	client := InitClient("testing", "testing")
	is.OK(client)
}

func TestClientGetters(t *testing.T) {
	is := is.New(t)

	for _, test := range tests {
		client := InitClient(test.input_ID, test.input_Secret)
		is.Equal(client.getClientID(), test.expected_ID)
		is.Equal(client.getClientSecret(), test.expected_Secret)
	}
}
