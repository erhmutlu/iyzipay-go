package samples_test

import (
	"github.com/stretchr/testify/assert"
	"iyzipay-go/iyzipay/client"
	"iyzipay-go/iyzipay/request"
	. "iyzipay-go/iyzipay/request/model"
	"testing"
)

func TestPaymentPostAuth(t *testing.T) {
	options := RetrieveOptions()
	request := request.PaymentPostAuthRequest{}
	request.Locale = TR
	request.ConversationId = "123456789abc"
	request.PaymentId = "11361093"
	request.PaidPrice = 1.2
	request.Currency = TRY
	request.Ip = "127.0.0.1"

	payment := client.PaymentPostAuth(request, options)
	assert.Equal(t, "success", *payment.Meta.Status)
	assert.Equal(t, "123456789abc", *payment.Meta.ConversationId)
}
