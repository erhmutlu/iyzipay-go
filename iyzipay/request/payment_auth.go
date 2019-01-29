package request

import (
	. "iyzipay-go/iyzipay/model"
	. "iyzipay-go/iyzipay/security"
	"strconv"
)

type PaymentAuthRequest struct {
	Locale          string      `json:"locale,omitempty"`
	ConversationId  string      `json:"conversationId,omitempty"`
	Price           float64     `json:"price,omitempty"`
	PaidPrice       float64     `json:"paidPrice,omitempty"`
	Installment     int         `json:"installment,omitempty"`
	PaymentChannel  string      `json:"paymentChannel,omitempty"`
	BasketId        string      `json:"basketId,omitempty"`
	PaymentGroup    string      `json:"paymentGroup,omitempty"`
	PaymentCard     PaymentCard `json:"paymentCard,omitempty"`
	Buyer           Buyer       `json:"buyer,omitempty"`
	ShippingAddress Address     `json:"shippingAddress,omitempty"`
	BillingAddress  Address     `json:"billingAddress,omitempty"`
	BasketItems     []BasketItem `json:"basketItems,omitempty"`
	PaymentSource   string      `json:"paymentSource,omitempty"`
	Currency        string      `json:"currency,omitempty"`
	PosOrderId      string      `json:"posOrderId,omitempty"`
	ConnectorName   string      `json:"connectorName,omitempty"`
	CallbackUrl     string      `json:"callbackUrl,omitempty"`
}

func (paymentAuthRequest PaymentAuthRequest) ToPKIRequest() string {
	pki := PKIRequest{}.
		Append("locale", paymentAuthRequest.Locale).
		Append("conversationId", paymentAuthRequest.ConversationId).
		Append("price", strconv.FormatFloat(paymentAuthRequest.Price, 'f', -1, 64)).
		Append("paidPrice", strconv.FormatFloat(paymentAuthRequest.PaidPrice, 'f', -1, 64)).
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
