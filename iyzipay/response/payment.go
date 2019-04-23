package response

import "iyzipay-go/iyzipay/response/nullable"

type PaymentRetrieveResponse struct {
	*Meta
	*PaymentInfo
}

type PaymentInfo struct {
	Price                        float64          `json:"price"`
	PaidPrice                    float64          `json:"paidPrice"`
	Currency                     string           `json:"currency"`
	Installment                  int              `json:"installment"`
	PaymentId                    string           `json:"paymentId"`
	PaymentStatus                string           `json:"paymentStatus"`
	FraudStatus                  int              `json:"fraudStatus"`
	MerchantCommissionRate       float64          `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount float64          `json:"merchantCommissionRateAmount"`
	IyziCommissionRateAmount     float64          `json:"iyziCommissionRateAmount"`
	IyziCommissionFee            float64          `json:"iyziCommissionFee"`
	CardType                     string           `json:"cardType"`
	CardAssociation              string           `json:"cardAssociation"`
	CardFamily                   string           `json:"cardFamily"`
	CardToken                    string           `json:"cardToken"`
	CardUserKey                  string           `json:"cardUserKey"`
	BinNumber                    string           `json:"binNumber"`
	BasketId                     string           `json:"basketId"`
	ConnectorName                string           `json:"connectorName"`
	AuthCode                     string           `json:"authCode"`
	Phase                        string           `json:"phase"`
	LastFourDigits               string           `json:"lastFourDigits"`
	PosOrderId                   *nullable.String `json:"posOrderId"`
	HostReference                string           `json:"hostReference"`
	PaymentItems                 []PaymentItem    `json:"itemTransactions"`
}

type PaymentItem struct {
	ItemId                        string           `json:"itemId"`
	PaymentTransactionId          string           `json:"paymentTransactionId"`
	TransactionStatus             int              `json:"transactionStatus"`
	Price                         float64          `json:"price"`
	PaidPrice                     float64          `json:"paidPrice"`
	MerchantCommissionRate        float64          `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount  float64          `json:"merchantCommissionRateAmount"`
	IyziCommissionRateAmount      float64          `json:"iyziCommissionRateAmount"`
	IyziCommissionFee             float64          `json:"iyziCommissionFee"`
	BlockageRate                  float64         `json:"blockageRate"`
	BlockageRateAmountMerchant    float64         `json:"blockageRateAmountMerchant"`
	BlockageRateAmountSubMerchant float64         `json:"blockageRateAmountSubMerchant"`
	BlockageResolvedDate          string          `json:"blockageResolvedDate"`
	SubMerchantKey                *nullable.String `json:"subMerchantKey"`
	SubMerchantPrice              *nullable.Float  `json:"subMerchantPrice"`
	SubMerchantPayoutRate         *nullable.Float  `json:"subMerchantPayoutRate"`
	SubMerchantPayoutAmount       *nullable.Float  `json:"subMerchantPayoutAmount"`
	MerchantPayoutAmount          float64         `json:"merchantPayoutAmount"`
	ConvertedPayout               *ConvertedPayout `json:"convertedPayout"`
}

//TODO: buradaki pointerlerin üstünden geç
type ConvertedPayout struct {
	PaidPrice                     *float64 `json:"paidPrice"`
	IyziCommissionRateAmount      *float64 `json:"iyziCommissionRateAmount"`
	IyziCommissionFee             *float64 `json:"iyziCommissionFee"`
	BlockageRateAmountMerchant    *float64 `json:"blockageRateAmountMerchant"`
	BlockageRateAmountSubMerchant *float64 `json:"blockageRateAmountSubMerchant"`
	SubMerchantPayoutAmount       *float64 `json:"subMerchantPayoutAmount"`
	MerchantPayoutAmount          *float64 `json:"merchantPayoutAmount"`
	IyziConversionRate            *float64 `json:"iyziConversionRate"`
	IyziConversionRateAmount      *float64 `json:"iyziConversionRateAmount"`
	Currency                      *string  `json:"currency"`
}
