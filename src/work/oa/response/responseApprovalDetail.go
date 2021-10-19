package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseApprovalDetail struct {
	*response.ResponseWork
	Info power.HashMap `json:"info"`
}
