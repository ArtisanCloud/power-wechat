package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseAccountList struct {
	*response.ResponseWork

	AccountList []*power.HashMap `json:"account_list"`
}
