package redpack

import (
	"encoding/xml"
	"fmt"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	payment "github.com/ArtisanCloud/powerwechat/src/payment/kernel"
	"github.com/ArtisanCloud/powerwechat/src/payment/redpack/request"
)

type Client struct {
	*payment.BaseClient
}

func NewClient(app *payment.ApplicationPaymentInterface) *Client {
	return &Client{
		payment.NewBaseClient(app),
	}
}

// Query Red Pack.
// https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=13_6&index=5
func (comp *Client) Info(mchBillNO string) (interface{}, error) {
	config := (*comp.App).GetConfig()

	params := &object.StringMap{
		"mch_billno": mchBillNO,
		"appid":      config.GetString("app_id", ""),
		"bill_type":  "MCHT",
	}

	endpoint := comp.Wrap("mmpaymkttransfers/gethbinfo")
	result, err := comp.SafeRequest(endpoint, params, "POST", nil, false, nil)

	return result, err
}

// Send Miniprogram Normal redpack.
// https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=18_2&index=2
func (comp *Client) SendMiniProgramNormal(params *power.HashMap) (interface{}, error) {
	config := (*comp.App).GetConfig()

	externalRequest := (*comp.App).GetExternalRequest()
	clientIP := externalRequest.Host
	if (*params)["client_ip"] != nil && (*params)["client_ip"].(string) != "" {
		clientIP = (*params)["client_ip"].(string)
	}
	base := &object.HashMap{
		"total_num": 1,
		"client_ip": clientIP,
		"wxappid":   config.GetString("app_id", ""),
	}

	options := object.MergeHashMap(base, params.ToHashMap())

	endpoint := comp.Wrap("/mmpaymkttransfers/sendminiprogramhb")
	result, err := comp.SafeRequest(endpoint, nil, "POST", options, nil, nil)

	return result, err
}

// Send Normal redpack.
// https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=13_4&index=3
func (comp *Client) SendNormal(params *request.SendRedPack) (interface{}, error) {
	config := (*comp.App).GetConfig()

	externalRequest := (*comp.App).GetExternalRequest()
	clientIP := externalRequest.Host
	if params.ClientIp != "" {
		params.ClientIp = clientIP
	}
	if params.TotalNum <= 0 {
		params.TotalNum = 1
	}
	if params.Wxappid != "" {
		params.Wxappid = config.GetString("app_id", "")
	}

	fmt.Println(params)

	//
	xmlOptionsBytes, err := xml.Marshal(params)
	if err != nil {
		return nil, err
	}

	//options := string(xmlOptionsBytes)
	fmt.Println(string(xmlOptionsBytes))

	//endpoint := comp.Wrap("/mmpaymkttransfers/sendredpack")
	//result, err := comp.SafeRequest(endpoint, nil, "POST", params, nil, nil)

	return nil, nil
}

// Send Group redpack.
// https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=13_5&index=4
func (comp *Client) SendGroup(params *power.HashMap) (interface{}, error) {
	config := (*comp.App).GetConfig()

	base := &object.HashMap{
		"amt_type": "ALL_RAND",
		"wxappid":  config.GetString("app_id", ""),
	}

	options := object.MergeHashMap(base, params.ToHashMap())

	endpoint := comp.Wrap("/mmpaymkttransfers/sendgroupredpack")
	result, err := comp.SafeRequest(endpoint, nil, "POST", options, false, nil)

	return result, err
}


// Send Work WX redpack.
// https://work.weixin.qq.com/api/doc/90000/90135/90275
func (comp *Client) SendWorkWX(params *power.HashMap) (interface{}, error) {
	config := (*comp.App).GetConfig()

	base := &object.HashMap{
		"wxappid":  config.GetString("app_id", ""),
	}

	options := object.MergeHashMap(base, params.ToHashMap())

	endpoint := comp.Wrap("/mmpaymkttransfers/sendworkwxredpack")
	result, err := comp.SafeRequest(endpoint, nil, "POST", options, false, nil)

	return result, err
}


// Query Work WX redpack.
// https://work.weixin.qq.com/api/doc/90000/90135/90276
func (comp *Client) QueryWorkWX(params *power.HashMap) (interface{}, error) {
	config := (*comp.App).GetConfig()

	base := &object.HashMap{
		"wxappid":  config.GetString("app_id", ""),
	}

	options := object.MergeHashMap(base, params.ToHashMap())

	endpoint := comp.Wrap("/mmpaymkttransfers/queryworkwxredpack")
	result, err := comp.SafeRequest(endpoint, nil, "POST", options, false, nil)

	return result, err
}