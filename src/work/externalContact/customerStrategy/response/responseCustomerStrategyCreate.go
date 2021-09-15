package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseCustomerStrategyCreate struct {
	*response.ResponseWork

	StrategyID *power.HashMap `json:"strategy_id"`
}
