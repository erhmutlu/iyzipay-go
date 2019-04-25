package request

import (
	. "iyzipay-go/iyzipay/request/model"
	. "iyzipay-go/iyzipay/security"
	"iyzipay-go/iyzipay/util"
	"strconv"
)

type PaymentPreAuthRequest struct {
	BaseRequest
	PaymentCreateRequest
}

func (paymentPreAuthRequest PaymentPreAuthRequest) ToPKIRequest() string {
	pki := PKIRequest{}.
		Append("locale", paymentPreAuthRequest.Locale). //TODO: AppendBaseRequest gibi bir şey eklenebilir.
		Append("conversationId", paymentPreAuthRequest.ConversationId).
		Append("price", util.FormatFloat(paymentPreAuthRequest.Price)).
		Append("paidPrice", util.FormatFloat(paymentPreAuthRequest.PaidPrice)).
		Append("installment", strconv.Itoa(paymentPreAuthRequest.Installment)).
		Append("paymentChannel", paymentPreAuthRequest.PaymentChannel).
		Append("basketId", paymentPreAuthRequest.BasketId).
		Append("paymentGroup", paymentPreAuthRequest.PaymentGroup).
		Append("paymentCard", paymentPreAuthRequest.PaymentCard.ToPKIRequest()).
		Append("buyer", paymentPreAuthRequest.Buyer.ToPKIRequest()).
		Append("shippingAddress", paymentPreAuthRequest.ShippingAddress.ToPKIRequest()).
		Append("billingAddress", paymentPreAuthRequest.BillingAddress.ToPKIRequest()).
		AppendSlice("basketItems", paymentPreAuthRequest.BasketItems).
		Append("paymentSource", paymentPreAuthRequest.PaymentSource).
		Append("currency", paymentPreAuthRequest.Currency).
		Append("posOrderId", paymentPreAuthRequest.PosOrderId).
		Append("connectorName", paymentPreAuthRequest.ConnectorName).
		Append("callbackUrl", paymentPreAuthRequest.CallbackUrl).
		ToString()

	return pki
}
