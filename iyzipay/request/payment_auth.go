package request

import (
	. "iyzipay-go/iyzipay/request/model"
	. "iyzipay-go/iyzipay/security"
	"iyzipay-go/iyzipay/util"
)

type PaymentAuthRequest struct {
	BaseRequest
	PaymentCreateRequest
}

func (paymentAuthRequest PaymentAuthRequest) ToPKIRequest() string {
	pki := PKIRequest{}.
		Append("locale", string(paymentAuthRequest.Locale)). //TODO: AppendBaseRequest gibi bir şey eklenebilir.
		Append("conversationId", paymentAuthRequest.ConversationId).
		Append("price", util.FormatPrimitiveFloat(paymentAuthRequest.Price)).
		Append("paidPrice", util.FormatPrimitiveFloat(paymentAuthRequest.PaidPrice)).
		Append("installment", util.FormatPrimitiveInt(paymentAuthRequest.Installment)).
		Append("paymentChannel", string(paymentAuthRequest.PaymentChannel)).
		Append("basketId", paymentAuthRequest.BasketId).
		Append("paymentGroup", string(paymentAuthRequest.PaymentGroup)).
		Append("paymentCard", paymentAuthRequest.PaymentCard.ToPKIRequest()).
		Append("buyer", paymentAuthRequest.Buyer.ToPKIRequest()).
		Append("shippingAddress", paymentAuthRequest.ShippingAddress.ToPKIRequest()).
		Append("billingAddress", paymentAuthRequest.BillingAddress.ToPKIRequest()).
		AppendSlice("basketItems", paymentAuthRequest.BasketItems).
		Append("paymentSource", paymentAuthRequest.PaymentSource).
		Append("currency", string(paymentAuthRequest.Currency)).
		Append("posOrderId", paymentAuthRequest.PosOrderId).
		Append("connectorName", paymentAuthRequest.ConnectorName).
		Append("callbackUrl", paymentAuthRequest.CallbackUrl).
		ToString()

	return pki
}
