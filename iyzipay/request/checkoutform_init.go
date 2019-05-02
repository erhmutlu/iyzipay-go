package request

import (
	. "github.com/erhmutlu/iyzipay-go/iyzipay/request/model"
	. "github.com/erhmutlu/iyzipay-go/iyzipay/security"
	"github.com/erhmutlu/iyzipay-go/iyzipay/util"
)

type CheckoutFormInitializeRequest struct {
	BaseRequest
	Price                     float64      `json:"price,omitempty"`
	PaidPrice                 float64      `json:"paidPrice,omitempty"`
	BasketId                  string       `json:"basketId,omitempty"`
	PaymentGroup              PaymentGroup `json:"paymentGroup,omitempty"`
	PaymentSource             string       `json:"paymentSource,omitempty"`
	Currency                  Currency     `json:"currency,omitempty"`
	Buyer                     Buyer        `json:"buyer,omitempty"`
	ShippingAddress           Address      `json:"shippingAddress,omitempty"`
	BillingAddress            Address      `json:"billingAddress,omitempty"`
	BasketItems               []BasketItem `json:"basketItems,omitempty"`
	CallbackUrl               string       `json:"callbackUrl,omitempty"`
	ForceThreeDS              *int          `json:"forceThreeDS,omitempty"`
	CardUserKey               string       `json:"cardUserKey,omitempty"`
	PosOrderId                string       `json:"posOrderId,omitempty"`
	EnabledInstallments       []int        `json:"enabledInstallments,omitempty"`
	PaymentWithNewCardEnabled bool         `json:"paymentWithNewCardEnabled,omitempty"` //TODO: nil olma durumu
	DebitCardAllowed          bool         `json:"debitCardAllowed,omitempty"`          //TODO: nil olma durumu
}

func (checkoutFormInitializeRequest CheckoutFormInitializeRequest) ToPKIRequest() string {
	pki := PKIRequest{}.
		Append("locale", string(checkoutFormInitializeRequest.Locale)). //TODO: AppendBaseRequest gibi bir şey eklenebilir.
		Append("conversationId", checkoutFormInitializeRequest.ConversationId).
		Append("price", util.FormatPrimitiveFloat(checkoutFormInitializeRequest.Price)).
		Append("basketId", checkoutFormInitializeRequest.BasketId).
		Append("paymentGroup", string(checkoutFormInitializeRequest.PaymentGroup)).
		Append("buyer", checkoutFormInitializeRequest.Buyer.ToPKIRequest()).
		Append("shippingAddress", checkoutFormInitializeRequest.ShippingAddress.ToPKIRequest()).
		Append("billingAddress", checkoutFormInitializeRequest.BillingAddress.ToPKIRequest()).
		AppendSlice("basketItems", checkoutFormInitializeRequest.BasketItems).
		Append("callbackUrl", checkoutFormInitializeRequest.CallbackUrl).
		Append("paymentSource", checkoutFormInitializeRequest.PaymentSource).
		Append("currency", string(checkoutFormInitializeRequest.Currency)).
		Append("posOrderId", checkoutFormInitializeRequest.PosOrderId).
		Append("paidPrice", util.FormatPrimitiveFloat(checkoutFormInitializeRequest.PaidPrice)).
		Append("forceThreeDS", util.FormatInt(checkoutFormInitializeRequest.ForceThreeDS)).
		Append("cardUserKey", checkoutFormInitializeRequest.CardUserKey).
		AppendIntSlice("enabledInstallments", checkoutFormInitializeRequest.EnabledInstallments).
		Append("paymentWithNewCardEnabled", util.FormatPrimitiveBool(checkoutFormInitializeRequest.PaymentWithNewCardEnabled)).
		Append("debitCardAllowed", util.FormatPrimitiveBool(checkoutFormInitializeRequest.DebitCardAllowed)).
		ToString()

	return pki
}