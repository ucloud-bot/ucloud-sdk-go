// Code is generated by ucloud-model, DO NOT EDIT IT.

package udisk

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// RecoverUDiskRequest is request schema for RecoverUDisk action
type RecoverUDiskRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// Year , Month, Dynamic 默认: Dynamic
	ChargeType *string `required:"false"`

	// 购买时长 默认: 1
	Quantity *int `required:"false"`

	// 云硬盘资源ID
	UDiskId *string `required:"true"`
}

// RecoverUDiskResponse is response schema for RecoverUDisk action
type RecoverUDiskResponse struct {
	response.CommonBase
}

// NewRecoverUDiskRequest will create request of RecoverUDisk action.
func (c *UDiskClient) NewRecoverUDiskRequest() *RecoverUDiskRequest {
	req := &RecoverUDiskRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// RecoverUDisk - 从回收站中恢复云硬盘
func (c *UDiskClient) RecoverUDisk(req *RecoverUDiskRequest) (*RecoverUDiskResponse, error) {
	var err error
	var res RecoverUDiskResponse

	reqCopier := *req

	err = c.Client.InvokeAction("RecoverUDisk", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
