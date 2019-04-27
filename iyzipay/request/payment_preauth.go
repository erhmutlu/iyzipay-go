package request

import (
	. "iyzipay-go/iyzipay/request/model"
	. "iyzipay-go/iyzipay/security"
	"iyzipay-go/iyzipay/util"
)

type PaymentPreAuthRequest struct {
	BaseRequest
	PaymentCreateRequest
}

func (paymentPreAuthRequest PaymentPreAuthRequest) ToPKIRequest() string {
	pki := PKIRequest{}.
		Append("locale", string(paymentPreAuthRequest.Locale)).
		Append("conversationId", paymentPreAuthRequest.ConversationId).
		Append("price", util.FormatPrimitiveFloat(paymentPreAuthRequest.Price)).
		Append("paidPrice", util.FormatPrimitiveFloat(paymentPreAuthRequest.PaidPrice)).
		Append("installment", util.FormatPrimitiveInt(paymentPreAuthRequest.Installment)).
		Append("paymentChannel", string(paymentPreAuthRequest.PaymentChannel)).
		Append("basketId", paymentPreAuthRequest.BasketId).
		Append("paymentGroup", paymentPreAuthRequest.PaymentGroup).
		Append("paymentCard", paymentPreAuthRequest.PaymentCard.ToPKIRequest()).
		Append("buyer", paymentPreAuthRequest.Buyer.ToPKIRequest()).
		Append("shippingAddress", paymentPreAuthRequest.ShippingAddress.ToPKIRequest()).
		Append("billingAddress", paymentPreAuthRequest.BillingAddress.ToPKIRequest()).
		AppendSlice("basketItems", paymentPreAuthRequest.BasketItems).
		Append("paymentSource", paymentPreAuthRequest.PaymentSource).
		Append("currency", string(paymentPreAuthRequest.Currency)).
		Append("posOrderId", paymentPreAuthRequest.PosOrderId).
		Append("connectorName", paymentPreAuthRequest.ConnectorName).
		Append("callbackUrl", paymentPreAuthRequest.CallbackUrl).
		ToString()

	return pki
}
