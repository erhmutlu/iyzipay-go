package samples

import (
	"github.com/stretchr/testify/assert"
	"github.com/erhmutlu/iyzipay-go/iyzipay/client"
	. "github.com/erhmutlu/iyzipay-go/iyzipay/request/model"
	"github.com/erhmutlu/iyzipay-go/iyzipay/request"
	"github.com/erhmutlu/iyzipay-go/iyzipay/util"
	"testing"
)

func TestMarketplacePaymentAuth(t *testing.T) {
	options := RetrieveOptions()
	request := request.PaymentAuthRequest{}
	request.Locale = TR
	request.ConversationId = "123456789abc"
	request.Price = 1
	request.PaidPrice = 1.2
	request.Currency = TRY
	request.Installment = 1
	request.BasketId = "B67832"
	request.PaymentChannel = WEB
	request.PaymentGroup = PRODUCT

	basketItem1 := BasketItem{Id: "BI101", Name: "Binocular", Category1: "Collectibles", Category2: "Accessories", ItemType: PHYSICAL, Price: 0.3, SubMerchantKey: "ha3us4v5mk2652kkjk5728cc4407an", SubMerchantPrice: 0.27}
	basketItem2 := BasketItem{Id: "BI102", Name: "Game code", Category1: "Game", Category2: "Online Game Items", ItemType: VIRTUAL, Price: 0.5, SubMerchantKey: "ha3us4v5mk2652kkjk5728cc4407an", SubMerchantPrice: 0.42}
	basketItem3 := BasketItem{Id: "BI103", Name: "Usb", Category1: "Electronics", Category2: "Usb / Cable", ItemType: PHYSICAL, Price: 0.2, SubMerchantKey: "ha3us4v5mk2652kkjk5728cc4407an", SubMerchantPrice: 0.18}
	basketItems := []BasketItem{basketItem1, basketItem2, basketItem3}
	request.BasketItems = basketItems

	paymentCard := PaymentCard{CardHolderName: "John Doe", CardNumber: "5528790000000008", ExpireYear: "2030", ExpireMonth: "12", Cvc: "123", RegisterCard: util.NewInt(0)}
	request.PaymentCard = paymentCard

	shippingAddress := Address{ContactName: "Jane Doe", City: "Istanbul", Country: "Turkey", Address: "iyziPark", ZipCode: "11111"}
	request.ShippingAddress = shippingAddress

	billingAddress := Address{ContactName: "Jane Doe", City: "Istanbul", Country: "Turkey", Address: "iyziPark", ZipCode: "11111"}
	request.BillingAddress = billingAddress

	buyer := Buyer{Id: "BY789", Name: "John", Surname: "Doe", GsmNumber: "+905350000000", Email: "email@email.com", IdentityNumber: "74300864791", LastLoginDate: "2015-10-05 12:43:35", RegistrationDate: "2013-04-21 15:12:09", RegistrationAddress: "iyziPark", Ip: "85.34.78.112", City: "Istanbul", Country: "Turkey", ZipCode: "11111"}
	request.Buyer = buyer

	payment := client.PaymentAuth(request, options)
	assert.Equal(t, "success", payment.Meta.Status)
	assert.Equal(t, "123456789abc", payment.Meta.ConversationId.Data)
}

func TestStandardMerchantPaymentAuth(t *testing.T) {
	options := RetrieveOptions()
	request := request.PaymentAuthRequest{}
	request.Locale = TR
	request.ConversationId = "123456789abc"
	request.Price = 1
	request.PaidPrice = 1.2
	request.Currency = TRY
	request.Installment = 1
	request.BasketId = "B67832"
	request.PaymentChannel = WEB
	request.PaymentGroup = PRODUCT

	basketItem1 := BasketItem{Id: "BI101", Name: "Binocular", Category1: "Collectibles", Category2: "Accessories", ItemType: PHYSICAL, Price: 0.3}
	basketItem2 := BasketItem{Id: "BI102", Name: "Game code", Category1: "Game", Category2: "Online Game Items", ItemType: VIRTUAL, Price: 0.5}
	basketItem3 := BasketItem{Id: "BI103", Name: "Usb", Category1: "Electronics", Category2: "Usb / Cable", ItemType: PHYSICAL, Price: 0.2}
	basketItems := []BasketItem{basketItem1, basketItem2, basketItem3}
	request.BasketItems = basketItems

	paymentCard := PaymentCard{CardHolderName: "John Doe", CardNumber: "5528790000000008", ExpireYear: "2030", ExpireMonth: "12", Cvc: "123", RegisterCard: util.NewInt(0)}
	request.PaymentCard = paymentCard

	shippingAddress := Address{ContactName: "Jane Doe", City: "Istanbul", Country: "Turkey", Address: "iyziPark", ZipCode: "11111"}
	request.ShippingAddress = shippingAddress

	billingAddress := Address{ContactName: "Jane Doe", City: "Istanbul", Country: "Turkey", Address: "iyziPark", ZipCode: "11111"}
	request.BillingAddress = billingAddress

	buyer := Buyer{Id: "BY789", Name: "John", Surname: "Doe", GsmNumber: "+905350000000", Email: "email@email.com", IdentityNumber: "74300864791", LastLoginDate: "2015-10-05 12:43:35", RegistrationDate: "2013-04-21 15:12:09", RegistrationAddress: "iyziPark", Ip: "85.34.78.112", City: "Istanbul", Country: "Turkey", ZipCode: "11111"}
	request.Buyer = buyer

	payment := client.PaymentAuth(request, options)
	assert.Equal(t, "success", payment.Meta.Status)
	assert.Equal(t, "123456789abc", payment.Meta.ConversationId.Data)
}

func TestMarketplacePaymentAuthWithRegisteredCard(t *testing.T) {
	options := RetrieveOptions()
	request := request.PaymentAuthRequest{}
	request.Locale = TR
	request.ConversationId = "123456789abc"
	request.Price = 1
	request.PaidPrice = 1.2
	request.Currency = TRY
	request.Installment = 1
	request.BasketId = "B67832"
	request.PaymentChannel = WEB
	request.PaymentGroup = PRODUCT

	basketItem1 := BasketItem{Id: "BI101", Name: "Binocular", Category1: "Collectibles", Category2: "Accessories", ItemType: PHYSICAL, Price: 0.3, SubMerchantKey: "ha3us4v5mk2652kkjk5728cc4407an", SubMerchantPrice: 0.27}
	basketItem2 := BasketItem{Id: "BI102", Name: "Game code", Category1: "Game", Category2: "Online Game Items", ItemType: VIRTUAL, Price: 0.5, SubMerchantKey: "ha3us4v5mk2652kkjk5728cc4407an", SubMerchantPrice: 0.42}
	basketItem3 := BasketItem{Id: "BI103", Name: "Usb", Category1: "Electronics", Category2: "Usb / Cable", ItemType: PHYSICAL, Price: 0.2, SubMerchantKey: "ha3us4v5mk2652kkjk5728cc4407an", SubMerchantPrice: 0.18}
	basketItems := []BasketItem{basketItem1, basketItem2, basketItem3}
	request.BasketItems = basketItems

	paymentCard := PaymentCard{CardToken: "CFo1rOg+FHucYgW/bnZM32OJ+/I=", CardUserKey: "+xCZCA5Wi5wVWxpkA4Qaypaa3OY="}
	request.PaymentCard = paymentCard

	shippingAddress := Address{ContactName: "Jane Doe", City: "Istanbul", Country: "Turkey", Address: "iyziPark", ZipCode: "11111"}
	request.ShippingAddress = shippingAddress

	billingAddress := Address{ContactName: "Jane Doe", City: "Istanbul", Country: "Turkey", Address: "iyziPark", ZipCode: "11111"}
	request.BillingAddress = billingAddress

	buyer := Buyer{Id: "BY789", Name: "John", Surname: "Doe", GsmNumber: "+905350000000", Email: "email@email.com", IdentityNumber: "74300864791", LastLoginDate: "2015-10-05 12:43:35", RegistrationDate: "2013-04-21 15:12:09", RegistrationAddress: "iyziPark", Ip: "85.34.78.112", City: "Istanbul", Country: "Turkey", ZipCode: "11111"}
	request.Buyer = buyer

	payment := client.PaymentAuth(request, options)
	assert.Equal(t, "success", payment.Meta.Status)
	assert.Equal(t, "123456789abc", payment.Meta.ConversationId.Data)
}

func TestStandardMerchantPaymentAuthWithRegisteredCard(t *testing.T) {
	options := RetrieveOptions()
	request := request.PaymentAuthRequest{}
	request.Locale = TR
	request.ConversationId = "123456789abc"
	request.Price = 1
	request.PaidPrice = 1.2
	request.Currency = TRY
	request.Installment = 1
	request.BasketId = "B67832"
	request.PaymentChannel = WEB
	request.PaymentGroup = PRODUCT

	basketItem1 := BasketItem{Id: "BI101", Name: "Binocular", Category1: "Collectibles", Category2: "Accessories", ItemType: PHYSICAL, Price: 0.3}
	basketItem2 := BasketItem{Id: "BI102", Name: "Game code", Category1: "Game", Category2: "Online Game Items", ItemType: VIRTUAL, Price: 0.5}
	basketItem3 := BasketItem{Id: "BI103", Name: "Usb", Category1: "Electronics", Category2: "Usb / Cable", ItemType: PHYSICAL, Price: 0.2}
	basketItems := []BasketItem{basketItem1, basketItem2, basketItem3}
	request.BasketItems = basketItems

	paymentCard := PaymentCard{CardToken: "CFo1rOg+FHucYgW/bnZM32OJ+/I=", CardUserKey: "+xCZCA5Wi5wVWxpkA4Qaypaa3OY="}
	request.PaymentCard = paymentCard

	shippingAddress := Address{ContactName: "Jane Doe", City: "Istanbul", Country: "Turkey", Address: "iyziPark", ZipCode: "11111"}
	request.ShippingAddress = shippingAddress

	billingAddress := Address{ContactName: "Jane Doe", City: "Istanbul", Country: "Turkey", Address: "iyziPark", ZipCode: "11111"}
	request.BillingAddress = billingAddress

	buyer := Buyer{Id: "BY789", Name: "John", Surname: "Doe", GsmNumber: "+905350000000", Email: "email@email.com", IdentityNumber: "74300864791", LastLoginDate: "2015-10-05 12:43:35", RegistrationDate: "2013-04-21 15:12:09", RegistrationAddress: "iyziPark", Ip: "85.34.78.112", City: "Istanbul", Country: "Turkey", ZipCode: "11111"}
	request.Buyer = buyer

	payment := client.PaymentAuth(request, options)
	assert.Equal(t, "success", payment.Meta.Status)
	assert.Equal(t, "123456789abc", payment.Meta.ConversationId.Data)
}