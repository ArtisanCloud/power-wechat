package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseWebDriveSpaceShare struct {
	*response.ResponseWork

	SpaceShareURL string `json:"space_share_url"`
}
