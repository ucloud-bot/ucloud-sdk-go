// Code is generated by ucloud-model, DO NOT EDIT IT.

package udisk

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// RestoreUDiskRequest is request schema for RestoreUDisk action
type RestoreUDiskRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// 从指定的快照恢复
	SnapshotId *string `required:"false"`

	// 指定从方舟恢复的备份时间点
	SnapshotTime *int `required:"false"`

	// 需要恢复的盘ID
	UDiskId *string `required:"true"`
}

// RestoreUDiskResponse is response schema for RestoreUDisk action
type RestoreUDiskResponse struct {
	response.CommonBase
}

// NewRestoreUDiskRequest will create request of RestoreUDisk action.
func (c *UDiskClient) NewRestoreUDiskRequest() *RestoreUDiskRequest {
	req := &RestoreUDiskRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

// RestoreUDisk - 从备份恢复数据至UDisk
func (c *UDiskClient) RestoreUDisk(req *RestoreUDiskRequest) (*RestoreUDiskResponse, error) {
	var err error
	var res RestoreUDiskResponse

	reqCopier := *req

	err = c.Client.InvokeAction("RestoreUDisk", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
