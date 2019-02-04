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

	payment := client.PaymentRetrieve(request, options)
	assert.Equal(t, "success", *payment.Meta.Status)
	assert.Equal(t, "conversationId", *payment.Meta.ConversationId)
	assert.Nil(t, payment.Meta.ErrorCode)
	assert.Nil(t, payment.Meta.ErrorMessage)
	assert.Nil(t, payment.Meta.ErrorGroup)

	println(*payment.PaymentInfo.PaymentItems[0].PaidPrice == 0.36)
	println(*payment.PaymentInfo.PaymentItems[0].SubMerchantKey)
	println(*payment.PaymentInfo.PaymentItems[0].SubMerchantPayoutAmount)
}
