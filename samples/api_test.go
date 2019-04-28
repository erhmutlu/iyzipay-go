package samples

import (
	"github.com/stretchr/testify/assert"
	"iyzipay-go/iyzipay/client"
	"testing"
)

func TestApi(t *testing.T) {
	options := RetrieveOptions()
	apiTestResponse := client.ApiTest(options)
	assert.Equal(t, "success", apiTestResponse.Status)
}
