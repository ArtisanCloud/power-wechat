package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseAppChatCreate struct {
	*response.ResponseWork

	ChatID string `json:"chatid"`
}
