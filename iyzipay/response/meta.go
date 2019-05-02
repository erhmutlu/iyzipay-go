package response

import (
	. "github.com/erhmutlu/iyzipay-go/iyzipay/request/model"
	"github.com/erhmutlu/iyzipay-go/iyzipay/response/nullable"
)

type Meta struct {
	ConversationId *nullable.String `json:"conversationId"`
	Status         string           `json:"status"`
	ErrorCode      *nullable.String `json:"errorCode"`
	ErrorMessage   *nullable.String `json:"errorMessage"`
	ErrorGroup     *nullable.String `json:"errorGroup"`
	Locale         Locale           `json:"locale"`
	SystemTime     *nullable.Int64  `json:"systemTime"`
}
