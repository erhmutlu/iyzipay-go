package samples_test

import (
	"github.com/stretchr/testify/assert"
	"iyzipay-go/iyzipay/client"
	"iyzipay-go/iyzipay/request"
	"testing"
)

func TestPaymentCancel(t *testing.T) {
	options := RetrieveOptions()
	request := request.PaymentCancelRequest{}
	request.PaymentId = "11359041"
	request.Ip = "127.0.0.1"
	request.ConversationId = "conversationId"
	request.Locale = "tr"

	response := client.PaymentCancel(request, options)
	assert.Equal(t, "success", *response.Meta.Status)
	assert.Equal(t, "conversationId", *response.Meta.ConversationId)
	assert.Nil(t, response.Meta.ErrorCode)
	assert.Nil(t, response.Meta.ErrorMessage)
	assert.Nil(t, response.Meta.ErrorGroup)
}
