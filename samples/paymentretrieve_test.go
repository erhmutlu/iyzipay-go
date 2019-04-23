package samples_test

import (
	"github.com/stretchr/testify/assert"
	"iyzipay-go/iyzipay/client"
	"iyzipay-go/iyzipay/request"
	"testing"
)

func TestPaymentRetrieve(t *testing.T) {
	options := RetrieveOptions()
	request := request.PaymentRetrieveRequest{PaymentId: "11354047", PaymentConversationId: "123456789abc", ConversationId: "conversationId", Locale: "tr"}

	payment := client.PaymentRetrieve(request, options)
	assert.Equal(t, "success", *payment.Meta.Status)
	assert.Equal(t, "conversationId", *payment.Meta.ConversationId)
	assert.Nil(t, payment.Meta.ErrorCode)
	assert.Nil(t, payment.Meta.ErrorMessage)
	assert.Nil(t, payment.Meta.ErrorGroup)
}
