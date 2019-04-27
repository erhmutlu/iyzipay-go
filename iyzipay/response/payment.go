package response

import (
	. "iyzipay-go/iyzipay/request/model"
	"iyzipay-go/iyzipay/response/nullable"
)

type PaymentRetrieveResponse struct {
	*Meta
	*PaymentInfo
}

type PaymentInfo struct {
	Price                        float64          `json:"price"`
	PaidPrice                    float64          `json:"paidPrice"`
	Currency                     Currency         `json:"currency"`
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
	CardToken                    *nullable.String `json:"cardToken"`
	CardUserKey                  *nullable.String `json:"cardUserKey"`
	BinNumber                    string           `json:"binNumber"`
	BasketId                     string           `json:"basketId"`
	ConnectorName                *nullable.String `json:"connectorName"`
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
	PaidPrice                     *nullable.Float  `json:"paidPrice"`
	MerchantCommissionRate        float64          `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount  float64          `json:"merchantCommissionRateAmount"`
	IyziCommissionRateAmount      float64          `json:"iyziCommissionRateAmount"`
	IyziCommissionFee             float64          `json:"iyziCommissionFee"`
	BlockageRate                  float64          `json:"blockageRate"`
	BlockageRateAmountMerchant    float64          `json:"blockageRateAmountMerchant"`
	BlockageRateAmountSubMerchant float64          `json:"blockageRateAmountSubMerchant"`
	BlockageResolvedDate          string           `json:"blockageResolvedDate"`
	SubMerchantKey                *nullable.String `json:"subMerchantKey"`
	SubMerchantPrice              *nullable.Float  `json:"subMerchantPrice"`
	SubMerchantPayoutRate         *nullable.Float  `json:"subMerchantPayoutRate"`
	SubMerchantPayoutAmount       *nullable.Float  `json:"subMerchantPayoutAmount"`
	MerchantPayoutAmount          float64          `json:"merchantPayoutAmount"`
	ConvertedPayout               *ConvertedPayout `json:"convertedPayout"`
}

type ConvertedPayout struct {
	PaidPrice                     float64         `json:"paidPrice"`
	IyziCommissionRateAmount      float64         `json:"iyziCommissionRateAmount"`
	IyziCommissionFee             float64         `json:"iyziCommissionFee"`
	BlockageRateAmountMerchant    float64         `json:"blockageRateAmountMerchant"`
	BlockageRateAmountSubMerchant float64         `json:"blockageRateAmountSubMerchant"`
	SubMerchantPayoutAmount       *nullable.Float `json:"subMerchantPayoutAmount"`
	MerchantPayoutAmount          float64         `json:"merchantPayoutAmount"`
	IyziConversionRate            *nullable.Float `json:"iyziConversionRate"`
	IyziConversionRateAmount      *nullable.Float `json:"iyziConversionRateAmount"`
	Currency                      Currency        `json:"currency"`
}
