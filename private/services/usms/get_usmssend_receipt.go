//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api USMS GetUSMSSendReceipt

package usms

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// GetUSMSSendReceiptRequest is request schema for GetUSMSSendReceipt action
type GetUSMSSendReceiptRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"false"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 发送短信时返回的SessionNo集合，SessionNoSet.0,SessionNoSet.1....格式
	SessionNoSet []string `required:"true"`
}

// GetUSMSSendReceiptResponse is response schema for GetUSMSSendReceipt action
type GetUSMSSendReceiptResponse struct {
	response.CommonBase

	// 错误描述
	Message string

	// 回执信息集合
	Data []ReceiptPerSession
}

// NewGetUSMSSendReceiptRequest will create request of GetUSMSSendReceipt action.
func (c *USMSClient) NewGetUSMSSendReceiptRequest() *GetUSMSSendReceiptRequest {
	req := &GetUSMSSendReceiptRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// GetUSMSSendReceipt - 获取短信发送回执信息。下游服务提供商回执信息返回会有一定延时，建议发送完短信以后，5-10分钟后再调用该接口拉取回执信息。若超过12小时未返回，则请联系技术支持确认原因
func (c *USMSClient) GetUSMSSendReceipt(req *GetUSMSSendReceiptRequest) (*GetUSMSSendReceiptResponse, error) {
	var err error
	var res GetUSMSSendReceiptResponse

	err = c.Client.InvokeAction("GetUSMSSendReceipt", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
