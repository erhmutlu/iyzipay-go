package samples

import (
	"github.com/stretchr/testify/assert"
	"github.com/erhmutlu/iyzipay-go/iyzipay/client"
	"github.com/erhmutlu/iyzipay-go/iyzipay/request"
	. "github.com/erhmutlu/iyzipay-go/iyzipay/request/model"
	"testing"
)

func TestPaymentRetrieve(t *testing.T) {
	options := RetrieveOptions()
	request := request.PaymentRetrieveRequest{}
	request.PaymentId = "11361181"
	request.PaymentConversationId = "123456789abc"
	request.ConversationId = "conversationId"
	request.Locale = TR

	payment := client.PaymentRetrieve(request, options)
	assert.Equal(t, "success", payment.Meta.Status)
	assert.Equal(t, "conversationId", payment.Meta.ConversationId.Data)
	assert.Nil(t, payment.Meta.ErrorCode)
	assert.Nil(t, payment.Meta.ErrorMessage)
	assert.Nil(t, payment.Meta.ErrorGroup)
}
