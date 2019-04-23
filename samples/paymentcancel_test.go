package samples_test

import (
	"github.com/stretchr/testify/assert"
	"iyzipay-go/iyzipay/client"
	"iyzipay-go/iyzipay/request"
	"testing"
)

func TestPaymentCancel(t *testing.T) {
	options := RetrieveOptions()
	request := request.PaymentCancelRequest{PaymentId: "11354069", Ip: "127.0.0.1", ConversationId: "conversationId", Locale: "tr"}

	payment := client.PaymentCancel(request, options)
	assert.Equal(t, "success", *payment.Meta.Status)
	assert.Equal(t, "conversationId", *payment.Meta.ConversationId)
	assert.Nil(t, payment.Meta.ErrorCode)
	assert.Nil(t, payment.Meta.ErrorMessage)
	assert.Nil(t, payment.Meta.ErrorGroup)
}
