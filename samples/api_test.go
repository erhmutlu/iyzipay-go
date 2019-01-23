package samples_test

import (
	"github.com/stretchr/testify/assert"
	"iyzipay-go/iyzipay/client"
	"iyzipay-go/samples"
	"testing"
)

func TestApi(t *testing.T) {
	options := samples.RetrieveOptions()
	apiTestResponse := client.ApiTest(options)
	assert.Equal(t, "success", apiTestResponse.Status )
}
