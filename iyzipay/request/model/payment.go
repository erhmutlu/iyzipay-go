package model

type PaymentCreateRequest struct {
	Price           float64      `json:"price,omitempty"`
	PaidPrice       float64      `json:"paidPrice,omitempty"`
	Installment     int          `json:"installment,omitempty"`
	PaymentChannel  string       `json:"paymentChannel,omitempty"`
	BasketId        string       `json:"basketId,omitempty"`
	PaymentGroup    string       `json:"paymentGroup,omitempty"`
	PaymentCard     PaymentCard  `json:"paymentCard,omitempty"`
	Buyer           Buyer        `json:"buyer,omitempty"`
	ShippingAddress Address      `json:"shippingAddress,omitempty"`
	BillingAddress  Address      `json:"billingAddress,omitempty"`
	BasketItems     []BasketItem `json:"basketItems,omitempty"`
	PaymentSource   string       `json:"paymentSource,omitempty"`
	Currency        Currency     `json:"currency,omitempty"`
	PosOrderId      string       `json:"posOrderId,omitempty"`
	ConnectorName   string       `json:"connectorName,omitempty"`
	CallbackUrl     string       `json:"callbackUrl,omitempty"`
}