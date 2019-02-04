package response

type Meta struct {
	ConversationId *string `json:"conversationId"`
	Status         *string `json:"status"`
	ErrorCode      *string `json:"errorCode"`
	ErrorMessage   *string `json:"errorMessage"`
	ErrorGroup     *string `json:"errorGroup"`
	Locale         *string `json:"locale"`
	SystemTime     *int64  `json:"systemTime"`
}