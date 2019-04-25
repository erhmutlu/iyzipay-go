package request

import (
	. "iyzipay-go/iyzipay/request/model"
	. "iyzipay-go/iyzipay/security"
	"iyzipay-go/iyzipay/util"
	"strconv"
)

type PaymentAuthRequest struct {
	BaseRequest
	PaymentCreateRequest
}

func (paymentAuthRequest PaymentAuthRequest) ToPKIRequest() string {
	pki := PKIRequest{}.
		Append("locale", paymentAuthRequest.Locale).  //TODO: AppendBaseRequest gibi bir şey eklenebilir.
		Append("conversationId", paymentAuthRequest.ConversationId).
		Append("price", util.FormatFloat(paymentAuthRequest.Price)).
		Append("paidPrice", util.FormatFloat(paymentAuthRequest.PaidPrice)).
		Append("installment", strconv.Itoa(paymentAuthRequest.Installment)).
		Append("paymentChannel", paymentAuthRequest.PaymentChannel).
		Append("basketId", paymentAuthRequest.BasketId).
		Append("paymentGroup", paymentAuthRequest.PaymentGroup).
		Append("paymentCard", paymentAuthRequest.PaymentCard.ToPKIRequest()).
		Append("buyer", paymentAuthRequest.Buyer.ToPKIRequest()).
		Append("shippingAddress", paymentAuthRequest.ShippingAddress.ToPKIRequest()).
		Append("billingAddress", paymentAuthRequest.BillingAddress.ToPKIRequest()).
		AppendSlice("basketItems", paymentAuthRequest.BasketItems).
		Append("paymentSource", paymentAuthRequest.PaymentSource).
		Append("currency", paymentAuthRequest.Currency).
		Append("posOrderId", paymentAuthRequest.PosOrderId).
		Append("connectorName", paymentAuthRequest.ConnectorName).
		Append("callbackUrl", paymentAuthRequest.CallbackUrl).
		ToString()

	return pki
}
