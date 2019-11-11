// Code is generated by ucloud-model, DO NOT EDIT IT.

package udisk

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// CreateUDiskSnapshotRequest is request schema for CreateUDiskSnapshot action
type CreateUDiskSnapshotRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// Year , Month, Dynamic 默认: Dynamic  (已废弃)
	ChargeType *string `required:"false"`

	// 快照描述
	Comment *string `required:"false"`

	// 快照名称
	Name *string `required:"true"`

	// 购买时长 默认: 1  (已废弃)
	Quantity *int `required:"false"`

	// 快照的UDisk的Id
	UDiskId *string `required:"true"`
}

// CreateUDiskSnapshotResponse is response schema for CreateUDiskSnapshot action
type CreateUDiskSnapshotResponse struct {
	response.CommonBase

	// 快照Id
	SnapshotId []string
}

// NewCreateUDiskSnapshotRequest will create request of CreateUDiskSnapshot action.
func (c *UDiskClient) NewCreateUDiskSnapshotRequest() *CreateUDiskSnapshotRequest {
	req := &CreateUDiskSnapshotRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

// CreateUDiskSnapshot - 创建snapshot快照
func (c *UDiskClient) CreateUDiskSnapshot(req *CreateUDiskSnapshotRequest) (*CreateUDiskSnapshotResponse, error) {
	var err error
	var res CreateUDiskSnapshotResponse

	reqCopier := *req

	err = c.Client.InvokeAction("CreateUDiskSnapshot", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
