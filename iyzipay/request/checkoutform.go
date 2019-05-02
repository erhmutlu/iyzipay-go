package request

import (
	. "github.com/erhmutlu/iyzipay-go/iyzipay/request/model"
	. "github.com/erhmutlu/iyzipay-go/iyzipay/security"
)

type CheckoutFormRequest struct {
	BaseRequest
	Token string `json:"token,omitempty"`
}

func (checkoutFormRequest CheckoutFormRequest) ToPKIRequest() string {
	pki := PKIRequest{}.
		Append("locale", string(checkoutFormRequest.Locale)). //TODO: AppendBaseRequest gibi bir şey eklenebilir.
		Append("conversationId", checkoutFormRequest.ConversationId).
		Append("token", checkoutFormRequest.Token).
		ToString()

	return pki
}
