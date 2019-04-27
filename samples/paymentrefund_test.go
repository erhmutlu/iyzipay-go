package samples_test

import (
	"github.com/stretchr/testify/assert"
	"iyzipay-go/iyzipay/client"
	"iyzipay-go/iyzipay/request"
	. "iyzipay-go/iyzipay/request/model"
	"testing"
)

func TestPaymentRefund(t *testing.T) {
	options := RetrieveOptions()
	request := request.PaymentRefundRequest{}
	request.PaymentTransactionId = "12038286"
	request.Price = 0.2
	request.Ip = "127.0.0.1"
	request.ConversationId = "conversationId"
	request.Locale = TR

	response := client.PaymentRefund(request, options)
	assert.Equal(t, "success", *response.Meta.Status)
	assert.Equal(t, "conversationId", *response.Meta.ConversationId)
	assert.Nil(t, response.Meta.ErrorCode)
	assert.Nil(t, response.Meta.ErrorMessage)
	assert.Nil(t, response.Meta.ErrorGroup)
}
