package samples_test

import (
	"github.com/stretchr/testify/assert"
	"iyzipay-go/iyzipay/client"
	"iyzipay-go/iyzipay/model"
	"iyzipay-go/iyzipay/request"
	"math/big"
	"testing"
)

func TestMarketplacePaymentAuth(t *testing.T) {
	options := RetrieveOptions()
	request := request.PaymentAuthRequest{}
	request.Locale = "tr"
	request.ConversationId = "123456789abc"
	request.Price = big.NewFloat(1)
	request.PaidPrice = big.NewFloat(1.2)
	request.Currency = "TRY" //TODO: constanta çıkılabilir
	request.Installment = 1
	request.BasketId = "B67832"
	request.PaymentChannel = "WEB"   //TODO: constanta çıkılabilir
	request.PaymentGroup = "PRODUCT" //TODO: constanta çıkılabilir

	basketItem1 := model.BasketItem{Id: "BI101", Name: "Binocular", Category1: "Collectibles", Category2: "Accessories", ItemType: "PHYSICAL", Price: big.NewFloat(0.3), SubMerchantKey: model.NewNullString("ha3us4v5mk2652kkjk5728cc4407an"), SubMerchantPrice: big.NewFloat(0.27)}
	basketItem2 := model.BasketItem{Id: "BI102", Name: "Game code", Category1: "Game", Category2: "Online Game Items", ItemType: "VIRTUAL", Price: big.NewFloat(0.5), SubMerchantKey: model.NewNullString("ha3us4v5mk2652kkjk5728cc4407an"), SubMerchantPrice: big.NewFloat(0.42)}
	basketItem3 := model.BasketItem{Id: "BI103", Name: "Usb", Category1: "Electronics", Category2: "Usb / Cable", ItemType: "PHYSICAL", Price: big.NewFloat(0.2), SubMerchantKey: model.NewNullString("ha3us4v5mk2652kkjk5728cc4407an"), SubMerchantPrice: big.NewFloat(0.18)}
	basketItems := []model.BasketItem{basketItem1, basketItem2, basketItem3}
	request.BasketItems = basketItems

	paymentCard := model.PaymentCard{CardHolderName: "John Doe", CardNumber: "5528790000000008", ExpireYear: "2030", ExpireMonth: "12", Cvc: "123"}
	request.PaymentCard = paymentCard

	shippingAddress := model.Address{ContactName: "Jane Doe", City: "Istanbul", Country: "Turkey", Address: "iyziPark", ZipCode: "11111"}
	request.ShippingAddress = shippingAddress

	billingAddress := model.Address{ContactName: "Jane Doe", City: "Istanbul", Country: "Turkey", Address: "iyziPark", ZipCode: "11111"}
	request.BillingAddress = billingAddress

	buyer := model.Buyer{Id: "BY789", Name: "John", Surname: "Doe", GsmNumber: "+905350000000", Email: "email@email.com", IdentityNumber: "74300864791", LastLoginDate: "2015-10-05 12:43:35", RegistrationDate: "2013-04-21 15:12:09", RegistrationAddress: "iyziPark", Ip: "85.34.78.112", City: "Istanbul", Country: "Turkey", ZipCode: "11111"}
	request.Buyer = buyer

	payment := client.PaymentAuth(request, options)
	assert.Equal(t, "success", *payment.Meta.Status)
	assert.Equal(t, "123456789", *payment.Meta.ConversationId)

	//var x model.BasketItem
	//strr := `{"id":"BI101","price":"0.3","name":"Binocular","category1":"Collectibles","category2":"Accessories","itemType":"PHYSICAL","subMerchantKey": "test", "subMerchantPrice":"0.27"}`
	//json.Unmarshal([]byte(strr), &x)
	//fmt.Printf("%+v", x.SubMerchantKey)
}
