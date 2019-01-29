package samples_test

import (
	"iyzipay-go/iyzipay/model"
	"iyzipay-go/iyzipay/request"
	"testing"
)

func TestPaymentAuth(t *testing.T) {
	//options := samples.RetrieveOptions()
	request := request.PaymentAuthRequest{}
	request.Locale = "tr"
	request.ConversationId = "123456789"
	request.Price = 1.1
	request.PaidPrice = 1.2
	request.Currency = "TRY" //TODO: structa çıkılabilir
	request.Installment = 1
	request.BasketId = "B67832"
	request.PaymentChannel = "WEB" //TODO: structa çıkılabilir
	request.PaymentGroup = "PRODUCT" //TODO: structa çıkılabilir

	basketItem1 := model.BasketItem{Id: "BI101", Name: "Binocular", Category1: "Collectibles", Category2: "Accessories", ItemType: "PHYSICAL", Price: 0.3 }
	basketItem2 := model.BasketItem{Id: "BI102", Name: "Game code", Category1: "Game", Category2: "Online Game Items", ItemType: "VIRTUAL", Price: 0.5 }
	basketItem3 := model.BasketItem{Id: "BI103", Name: "Usb", Category1: "Electronics", Category2: "Usb / Cable", ItemType: "PHYSICAL", Price: 0.2 }
	basketItems := []model.BasketItem{basketItem1, basketItem2, basketItem3}
	request.BasketItems = basketItems

	paymentCard := model.PaymentCard{CardHolderName: "John Doe", CardNumber: "5528790000000008", ExpireYear: "2030", ExpireMonth: "12", Cvc: "123", RegisterCard: 0}
	request.PaymentCard = paymentCard

	shippingAddress := model.Address{ContactName: "Jane Doe", City: "Istanbul", Country: "Turkey", Address: "iyziPark", ZipCode: "11111"}
	request.ShippingAddress = shippingAddress

	billingAddress := model.Address{ContactName: "Jane Doe", City: "Istanbul", Country: "Turkey", Address: "iyziPark", ZipCode: "11111"}
	request.BillingAddress = billingAddress

	buyer := model.Buyer{Id: "BY789", Name: "John", Surname: "Doe", GsmNumber: "+905350000000", Email: "email@email.com", IdentityNumber: "74300864791", LastLoginDate: "2015-10-05 12:43:35", RegistrationDate: "2013-04-21 15:12:09", RegistrationAddress: "iyziPark", Ip: "85.34.78.112", City: "Istanbul", Country: "Turkey", ZipCode: "11111"}
	request.Buyer = buyer

	println(request.ToPKIRequest())
}
