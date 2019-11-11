// Code is generated by ucloud-model, DO NOT EDIT IT.

package udisk

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeUDiskUpgradePriceRequest is request schema for DescribeUDiskUpgradePrice action
type DescribeUDiskUpgradePriceRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// UDisk 类型: DataDisk（普通数据盘），SSDDataDisk（SSD数据盘），SystemDisk（普通系统盘），SSDSystemDisk（SSD系统盘），RSSDDataDisk（RSSD数据盘），默认值（DataDisk）
	DiskType *string `required:"false"`

	// 云主机机型（V2.0），枚举值["N", "C", "G", "O", "OM"]。参考[[api:uhost-api:uhost_type|云主机机型说明]]。
	MachineType *string `required:"false"`

	// 购买UDisk大小,单位:GB,普通数据盘：范围[1~8000]；SSD数据盘：范围[1~8000]；普通系统盘：范围[1~8000]；SSD系统盘：范围[1~4000]；RSSD数据盘：范围[1~32000]。
	Size *int `required:"true"`

	// 升级目标UDisk ID
	SourceId *string `required:"true"`

	// 是否打开数据方舟, 打开"Yes",关闭"No", 默认关闭
	UDataArkMode *string `required:"false"`
}

// DescribeUDiskUpgradePriceResponse is response schema for DescribeUDiskUpgradePrice action
type DescribeUDiskUpgradePriceResponse struct {
	response.CommonBase

	// 用户折后价 (对应计费CustomPrice)
	OriginalPrice int

	// 价格
	Price int
}

// NewDescribeUDiskUpgradePriceRequest will create request of DescribeUDiskUpgradePrice action.
func (c *UDiskClient) NewDescribeUDiskUpgradePriceRequest() *DescribeUDiskUpgradePriceRequest {
	req := &DescribeUDiskUpgradePriceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeUDiskUpgradePrice - 获取UDisk升级价格信息
func (c *UDiskClient) DescribeUDiskUpgradePrice(req *DescribeUDiskUpgradePriceRequest) (*DescribeUDiskUpgradePriceResponse, error) {
	var err error
	var res DescribeUDiskUpgradePriceResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribeUDiskUpgradePrice", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
