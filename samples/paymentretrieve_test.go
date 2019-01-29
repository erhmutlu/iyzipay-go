package samples_test

import (
	"github.com/stretchr/testify/assert"
	"iyzipay-go/iyzipay/client"
	"iyzipay-go/iyzipay/request"
	"iyzipay-go/samples"
	"testing"
)

func TestPaymentRetrieve(t *testing.T) {
	options := samples.RetrieveOptions()
	request := request.PaymentRetrieveRequest{PaymentId: "11132055", ConversationId: "conversationId", Locale: "tr"}

	result := client.PaymentRetrieve(request, options)
	assert.Equal(t, "TEST2", result)
}
