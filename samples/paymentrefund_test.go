package samples

import (
	"github.com/stretchr/testify/assert"
	"github.com/erhmutlu/iyzipay-go/iyzipay/client"
	"github.com/erhmutlu/iyzipay-go/iyzipay/request"
	. "github.com/erhmutlu/iyzipay-go/iyzipay/request/model"
	"testing"
)

func TestPaymentRefund(t *testing.T) {
	options := RetrieveOptions()
	request := request.PaymentRefundRequest{}
	request.PaymentTransactionId = "12038420"
	request.Price = 0.2
	request.Ip = "127.0.0.1"
	request.ConversationId = "conversationId"
	request.Locale = TR

	response := client.PaymentRefund(request, options)
	assert.Equal(t, "success", response.Meta.Status)
	assert.Equal(t, "conversationId", response.Meta.ConversationId.Data)
	assert.Nil(t, response.Meta.ErrorCode)
	assert.Nil(t, response.Meta.ErrorMessage)
	assert.Nil(t, response.Meta.ErrorGroup)
}
