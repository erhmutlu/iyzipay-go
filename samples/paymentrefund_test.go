package samples_test

import (
	"github.com/stretchr/testify/assert"
	"iyzipay-go/iyzipay/client"
	"iyzipay-go/iyzipay/request"
	"testing"
)

func TestPaymentRefund(t *testing.T) {
	options := RetrieveOptions()
	request := request.PaymentRefundRequest{PaymentTransactionId: "12030211", Price: 0.2, Ip: "127.0.0.1", ConversationId: "conversationId", Locale: "tr"}

	response := client.PaymentRefund(request, options)
	assert.Equal(t, "success", *response.Meta.Status)
	assert.Equal(t, "conversationId", *response.Meta.ConversationId)
	assert.Nil(t, response.Meta.ErrorCode)
	assert.Nil(t, response.Meta.ErrorMessage)
	assert.Nil(t, response.Meta.ErrorGroup)
}
