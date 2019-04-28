package samples

import (
	"github.com/stretchr/testify/assert"
	"iyzipay-go/iyzipay/client"
	"iyzipay-go/iyzipay/request"
	. "iyzipay-go/iyzipay/request/model"
	"iyzipay-go/iyzipay/util"
	"testing"
)

func TestCheckoutFormInitialize(t *testing.T) {
	options := RetrieveOptions()
	request := request.CheckoutFormInitializeRequest{}
	request.Locale = TR
	request.ConversationId = "123456789abc"
	request.Price = 0.3
	request.PaidPrice = 0.4
	request.Currency = TRY
	request.BasketId = "B67832"
	request.PaymentGroup = LISTING
	request.CallbackUrl = "https://www.merchant.com/callback"
	request.DebitCardAllowed = true
	request.PaymentWithNewCardEnabled = false
	request.ForceThreeDS = util.NewInt(0)

	request.EnabledInstallments = []int{1, 2, 3, 6, 9}

	buyer := Buyer{Id: "BY789", Name: "John", Surname: "Doe", GsmNumber: "+905350000000", Email: "email@email.com", IdentityNumber: "74300864791", LastLoginDate: "2015-10-05 12:43:35", RegistrationDate: "2013-04-21 15:12:09", RegistrationAddress: "iyziPark", Ip: "85.34.78.112", City: "Istanbul", Country: "Turkey", ZipCode: "11111"}
	request.Buyer = buyer

	shippingAddress := Address{ContactName: "Jane Doe", City: "Istanbul", Country: "Turkey", Address: "iyziPark", ZipCode: "11111"}
	request.ShippingAddress = shippingAddress

	billingAddress := Address{ContactName: "Jane Doe", City: "Istanbul", Country: "Turkey", Address: "iyziPark", ZipCode: "11111"}
	request.BillingAddress = billingAddress

	basketItem1 := BasketItem{Id: "BI101", Name: "Binocular", Category1: "Collectibles", Category2: "Accessories", ItemType: PHYSICAL, Price: 0.3}
	basketItems := []BasketItem{basketItem1}
	request.BasketItems = basketItems

	response := client.CheckoutFormInitialize(request, options)
	assert.Equal(t, "success", response.Meta.Status)
	assert.Equal(t, "123456789abc", response.Meta.ConversationId.Data)
}
