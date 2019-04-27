package model

type BaseRequest struct {
	Locale         Locale `json:"locale,omitempty"`
	ConversationId string `json:"conversationId,omitempty"`
}
