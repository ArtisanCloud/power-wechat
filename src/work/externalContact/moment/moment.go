package moment

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/request"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 获取企业全部的发表列表
// https://work.weixin.qq.com/api/doc/90000/90135/93333
func (comp *Client) GetMomentList(params *request.RequestGetMomentList) (*response.ResponseGetMomentList, error) {

	result := &response.ResponseGetMomentList{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/get_moment_list", params, nil, nil, result)

	return result, err
}

// 客户朋友圈规则组管理
// https://work.weixin.qq.com/api/doc/90000/90135/94890
func (comp *Client) MomentStrategyList(cursor string, limit int) (*response.ResponseMomentStrategyList, error) {

	result := &response.ResponseMomentStrategyList{}

	options := &object.HashMap{
		"cursor": cursor,
		"limit":  limit,
	}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/moment_strategy/list", options, nil, nil, result)

	return result, err
}
