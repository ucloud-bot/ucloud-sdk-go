// Code is generated by ucloud-model, DO NOT EDIT IT.

package ugn

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// UGN API Schema

// AttachUGNInstanceRequest is request schema for AttachUGNInstance action
type AttachUGNInstanceRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"true"`

	// 实例Id
	InstanceId *string `required:"true"`

	// 实例归属ProjectId
	InstanceProjectId *string `required:"true"`

	// 实例归属地域
	InstanceRegion *string `required:"true"`

	// 实例类型
	InstanceType *string `required:"true"`

	// 云联网Id
	UGNId *string `required:"true"`
}

// AttachUGNInstanceResponse is response schema for AttachUGNInstance action
type AttachUGNInstanceResponse struct {
	response.CommonBase

	// 返回码描述信息
	Message string
}

// NewAttachUGNInstanceRequest will create request of AttachUGNInstance action.
func (c *UGNClient) NewAttachUGNInstanceRequest() *AttachUGNInstanceRequest {
	req := &AttachUGNInstanceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: AttachUGNInstance

实例加入云联网
*/
func (c *UGNClient) AttachUGNInstance(req *AttachUGNInstanceRequest) (*AttachUGNInstanceResponse, error) {
	var err error
	var res AttachUGNInstanceResponse

	reqCopier := *req

	err = c.Client.InvokeAction("AttachUGNInstance", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// BatchDetachUGNInstanceRequest is request schema for BatchDetachUGNInstance action
type BatchDetachUGNInstanceRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"true"`

	// 实例id
	InstanceIds []string `required:"true"`

	// 云联网id
	UGNId *string `required:"true"`
}

// BatchDetachUGNInstanceResponse is response schema for BatchDetachUGNInstance action
type BatchDetachUGNInstanceResponse struct {
	response.CommonBase

	// 删除成功的实例id
	InstanceIds []string
}

// NewBatchDetachUGNInstanceRequest will create request of BatchDetachUGNInstance action.
func (c *UGNClient) NewBatchDetachUGNInstanceRequest() *BatchDetachUGNInstanceRequest {
	req := &BatchDetachUGNInstanceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: BatchDetachUGNInstance

批量移除云联网中实例
*/
func (c *UGNClient) BatchDetachUGNInstance(req *BatchDetachUGNInstanceRequest) (*BatchDetachUGNInstanceResponse, error) {
	var err error
	var res BatchDetachUGNInstanceResponse

	reqCopier := *req

	err = c.Client.InvokeAction("BatchDetachUGNInstance", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// CreateInterRegionBandwidthRequest is request schema for CreateInterRegionBandwidth action
type CreateInterRegionBandwidthRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"true"`

	// 带宽（单位为Mb/s）
	Bandwidth *int `required:"true"`

	// 计费类型
	ChargeType *string `required:"true"`

	// 代金券Id
	CouponId *string `required:"false"`

	// 付费类型
	PayMode *string `required:"true"`

	// 购买时长
	Quantity *int `required:"true"`

	// 跨域带宽归属地域
	Region0 *string `required:"true"`

	// 跨域带宽归属地域
	Region1 *string `required:"true"`

	// 云联网Id
	UGNId *string `required:"true"`
}

// CreateInterRegionBandwidthResponse is response schema for CreateInterRegionBandwidth action
type CreateInterRegionBandwidthResponse struct {
	response.CommonBase

	// 跨域带宽Id
	InterRegionBandwidthId string

	// 返回码描述信息
	Message string
}

// NewCreateInterRegionBandwidthRequest will create request of CreateInterRegionBandwidth action.
func (c *UGNClient) NewCreateInterRegionBandwidthRequest() *CreateInterRegionBandwidthRequest {
	req := &CreateInterRegionBandwidthRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: CreateInterRegionBandwidth

购买跨域带宽
*/
func (c *UGNClient) CreateInterRegionBandwidth(req *CreateInterRegionBandwidthRequest) (*CreateInterRegionBandwidthResponse, error) {
	var err error
	var res CreateInterRegionBandwidthResponse

	reqCopier := *req

	err = c.Client.InvokeAction("CreateInterRegionBandwidth", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// CreateUGNRequest is request schema for CreateUGN action
type CreateUGNRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"true"`

	// 云联网名称，如果没有传入，默认值为“UGN”
	Name *string `required:"false"`

	// 云联网备注，如果没有传入的话，默认为“”
	Remark *string `required:"false"`

	// 业务组Id，如果没有传入，默认值为“Default”
	Tag *string `required:"false"`
}

// CreateUGNResponse is response schema for CreateUGN action
type CreateUGNResponse struct {
	response.CommonBase

	// 返回码描述信息
	Message string

	// 云联网Id
	UGNId string
}

// NewCreateUGNRequest will create request of CreateUGN action.
func (c *UGNClient) NewCreateUGNRequest() *CreateUGNRequest {
	req := &CreateUGNRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: CreateUGN

创建云联网
*/
func (c *UGNClient) CreateUGN(req *CreateUGNRequest) (*CreateUGNResponse, error) {
	var err error
	var res CreateUGNResponse

	reqCopier := *req

	err = c.Client.InvokeAction("CreateUGN", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DeleteInterRegionBandwidthRequest is request schema for DeleteInterRegionBandwidth action
type DeleteInterRegionBandwidthRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"true"`

	// 跨域带宽Id
	InterRegionBandwidthId *string `required:"true"`

	// 云联网Id
	UGNId *string `required:"true"`
}

// DeleteInterRegionBandwidthResponse is response schema for DeleteInterRegionBandwidth action
type DeleteInterRegionBandwidthResponse struct {
	response.CommonBase

	// 返回码描述信息
	Message string
}

// NewDeleteInterRegionBandwidthRequest will create request of DeleteInterRegionBandwidth action.
func (c *UGNClient) NewDeleteInterRegionBandwidthRequest() *DeleteInterRegionBandwidthRequest {
	req := &DeleteInterRegionBandwidthRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DeleteInterRegionBandwidth

删除跨域带宽
*/
func (c *UGNClient) DeleteInterRegionBandwidth(req *DeleteInterRegionBandwidthRequest) (*DeleteInterRegionBandwidthResponse, error) {
	var err error
	var res DeleteInterRegionBandwidthResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DeleteInterRegionBandwidth", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DeleteUGNRequest is request schema for DeleteUGN action
type DeleteUGNRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"true"`

	// 云联网Id
	UGNId *string `required:"true"`
}

// DeleteUGNResponse is response schema for DeleteUGN action
type DeleteUGNResponse struct {
	response.CommonBase

	// 返回码描述信息
	Message string
}

// NewDeleteUGNRequest will create request of DeleteUGN action.
func (c *UGNClient) NewDeleteUGNRequest() *DeleteUGNRequest {
	req := &DeleteUGNRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DeleteUGN

删除云联网
*/
func (c *UGNClient) DeleteUGN(req *DeleteUGNRequest) (*DeleteUGNResponse, error) {
	var err error
	var res DeleteUGNResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DeleteUGN", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DescribeInterRegionBandwidthRequest is request schema for DescribeInterRegionBandwidth action
type DescribeInterRegionBandwidthRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"true"`

	// 跨域带宽Id
	InterRegionBandwidthIds []string `required:"false"`

	// 数据分页值。默认为20
	Limit *int `required:"false"`

	// 数据偏移量。默认为0
	Offset *int `required:"false"`

	// 云联网Id
	UGNId *string `required:"true"`
}

// DescribeInterRegionBandwidthResponse is response schema for DescribeInterRegionBandwidth action
type DescribeInterRegionBandwidthResponse struct {
	response.CommonBase

	// 跨域带宽信息
	InterRegionBandwidths []InterRegionBandwidth

	// 返回码描述信息
	Message string

	// InterRegionBandwidths字段的数量
	TotalCount int
}

// NewDescribeInterRegionBandwidthRequest will create request of DescribeInterRegionBandwidth action.
func (c *UGNClient) NewDescribeInterRegionBandwidthRequest() *DescribeInterRegionBandwidthRequest {
	req := &DescribeInterRegionBandwidthRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DescribeInterRegionBandwidth

查询跨域带宽
*/
func (c *UGNClient) DescribeInterRegionBandwidth(req *DescribeInterRegionBandwidthRequest) (*DescribeInterRegionBandwidthResponse, error) {
	var err error
	var res DescribeInterRegionBandwidthResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribeInterRegionBandwidth", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DescribeUGNRequest is request schema for DescribeUGN action
type DescribeUGNRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"true"`

	// 数据分页值。默认为20
	Limit *int `required:"false"`

	// 数据偏移量。默认为0
	Offset *int `required:"false"`

	// 云联网Id
	UGNIds []string `required:"false"`
}

// DescribeUGNResponse is response schema for DescribeUGN action
type DescribeUGNResponse struct {
	response.CommonBase

	// 返回码描述信息
	Message string

	// UGNs字段的数量
	TotalCount int

	// 云联网信息
	UGNs []UGN
}

// NewDescribeUGNRequest will create request of DescribeUGN action.
func (c *UGNClient) NewDescribeUGNRequest() *DescribeUGNRequest {
	req := &DescribeUGNRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DescribeUGN

查询云联网
*/
func (c *UGNClient) DescribeUGN(req *DescribeUGNRequest) (*DescribeUGNResponse, error) {
	var err error
	var res DescribeUGNResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribeUGN", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DescribeUGNInstanceRequest is request schema for DescribeUGNInstance action
type DescribeUGNInstanceRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"true"`

	// 实例Id
	InstanceIds []string `required:"false"`

	// 数据分页值。默认为20
	Limit *int `required:"false"`

	// 数据偏移量。默认为0
	Offset *int `required:"false"`

	// 云联网Id
	UGNId *string `required:"true"`
}

// DescribeUGNInstanceResponse is response schema for DescribeUGNInstance action
type DescribeUGNInstanceResponse struct {
	response.CommonBase

	// 返回码描述信息
	Message string

	// UGNInstances字段的数量
	TotalCount int

	// 实例信息
	UGNInstances []Instance
}

// NewDescribeUGNInstanceRequest will create request of DescribeUGNInstance action.
func (c *UGNClient) NewDescribeUGNInstanceRequest() *DescribeUGNInstanceRequest {
	req := &DescribeUGNInstanceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DescribeUGNInstance

查询云联网实例
*/
func (c *UGNClient) DescribeUGNInstance(req *DescribeUGNInstanceRequest) (*DescribeUGNInstanceResponse, error) {
	var err error
	var res DescribeUGNInstanceResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribeUGNInstance", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DescribeUGNRegionListRequest is request schema for DescribeUGNRegionList action
type DescribeUGNRegionListRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"false"`

}

// DescribeUGNRegionListResponse is response schema for DescribeUGNRegionList action
type DescribeUGNRegionListResponse struct {
	response.CommonBase

	// 返回码描述信息
	Message string

	// 地域列表
	RegionList []string
}

// NewDescribeUGNRegionListRequest will create request of DescribeUGNRegionList action.
func (c *UGNClient) NewDescribeUGNRegionListRequest() *DescribeUGNRegionListRequest {
	req := &DescribeUGNRegionListRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DescribeUGNRegionList

获取ugn支持的地域
*/
func (c *UGNClient) DescribeUGNRegionList(req *DescribeUGNRegionListRequest) (*DescribeUGNRegionListResponse, error) {
	var err error
	var res DescribeUGNRegionListResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribeUGNRegionList", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DescribeUGNRouteRuleRequest is request schema for DescribeUGNRouteRule action
type DescribeUGNRouteRuleRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"true"`

	// 数据分页值。默认为20
	Limit *int `required:"false"`

	// 数据偏移量。默认为0
	Offset *int `required:"false"`

	// 路由规则Id
	RouteRuleIds []string `required:"false"`

	// 云联网Id
	UGNId *string `required:"true"`
}

// DescribeUGNRouteRuleResponse is response schema for DescribeUGNRouteRule action
type DescribeUGNRouteRuleResponse struct {
	response.CommonBase

	// 返回码描述信息
	Message string

	// UGNRouteRules字段的数量
	TotalCount int

	// 路由规则信息
	UGNRouteRules []RouteRule
}

// NewDescribeUGNRouteRuleRequest will create request of DescribeUGNRouteRule action.
func (c *UGNClient) NewDescribeUGNRouteRuleRequest() *DescribeUGNRouteRuleRequest {
	req := &DescribeUGNRouteRuleRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DescribeUGNRouteRule

查询路由规则
*/
func (c *UGNClient) DescribeUGNRouteRule(req *DescribeUGNRouteRuleRequest) (*DescribeUGNRouteRuleResponse, error) {
	var err error
	var res DescribeUGNRouteRuleResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribeUGNRouteRule", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DetachUGNInstanceRequest is request schema for DetachUGNInstance action
type DetachUGNInstanceRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"true"`

	// 实例Id
	InstanceId *string `required:"true"`

	// 实例类型
	InstanceType *string `required:"true"`

	// 云联网Id
	UGNId *string `required:"true"`
}

// DetachUGNInstanceResponse is response schema for DetachUGNInstance action
type DetachUGNInstanceResponse struct {
	response.CommonBase

	// 返回码描述信息
	Message string
}

// NewDetachUGNInstanceRequest will create request of DetachUGNInstance action.
func (c *UGNClient) NewDetachUGNInstanceRequest() *DetachUGNInstanceRequest {
	req := &DetachUGNInstanceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DetachUGNInstance

实例退出云联网
*/
func (c *UGNClient) DetachUGNInstance(req *DetachUGNInstanceRequest) (*DetachUGNInstanceResponse, error) {
	var err error
	var res DetachUGNInstanceResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DetachUGNInstance", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// ModifyInterRegionBandwidthRequest is request schema for ModifyInterRegionBandwidth action
type ModifyInterRegionBandwidthRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"true"`

	// 带宽（单位为Mb/s）
	Bandwidth *string `required:"true"`

	// 跨域带宽Id
	InterRegionBandwidthId *string `required:"true"`

	// 付费类型
	PayMode *string `required:"false"`

	// 云联网Id
	UGNId *string `required:"true"`
}

// ModifyInterRegionBandwidthResponse is response schema for ModifyInterRegionBandwidth action
type ModifyInterRegionBandwidthResponse struct {
	response.CommonBase

	// 返回码描述信息
	Message string
}

// NewModifyInterRegionBandwidthRequest will create request of ModifyInterRegionBandwidth action.
func (c *UGNClient) NewModifyInterRegionBandwidthRequest() *ModifyInterRegionBandwidthRequest {
	req := &ModifyInterRegionBandwidthRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: ModifyInterRegionBandwidth

修改跨域带宽
*/
func (c *UGNClient) ModifyInterRegionBandwidth(req *ModifyInterRegionBandwidthRequest) (*ModifyInterRegionBandwidthResponse, error) {
	var err error
	var res ModifyInterRegionBandwidthResponse

	reqCopier := *req

	err = c.Client.InvokeAction("ModifyInterRegionBandwidth", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// ModifyUGNAttributeRequest is request schema for ModifyUGNAttribute action
type ModifyUGNAttributeRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"true"`

	// 云联网名称
	Name *string `required:"false"`

	// 云联网备注
	Remark *string `required:"false"`

	// 云联网Id
	UGNId *string `required:"true"`
}

// ModifyUGNAttributeResponse is response schema for ModifyUGNAttribute action
type ModifyUGNAttributeResponse struct {
	response.CommonBase

	// 错误码描述信息
	Message string
}

// NewModifyUGNAttributeRequest will create request of ModifyUGNAttribute action.
func (c *UGNClient) NewModifyUGNAttributeRequest() *ModifyUGNAttributeRequest {
	req := &ModifyUGNAttributeRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: ModifyUGNAttribute

修改云联网属性
*/
func (c *UGNClient) ModifyUGNAttribute(req *ModifyUGNAttributeRequest) (*ModifyUGNAttributeResponse, error) {
	var err error
	var res ModifyUGNAttributeResponse

	reqCopier := *req

	err = c.Client.InvokeAction("ModifyUGNAttribute", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// PublishUGNRouteRuleRequest is request schema for PublishUGNRouteRule action
type PublishUGNRouteRuleRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"true"`

	// 归属实例路由规则Id
	DeriveRouteRuleId *string `required:"true"`

	// 归属实例路由表Id
	DeriveRouteTableId *string `required:"true"`

	// 实例Id
	InstanceId *string `required:"true"`

	// 云联网Id
	UGNId *string `required:"true"`
}

// PublishUGNRouteRuleResponse is response schema for PublishUGNRouteRule action
type PublishUGNRouteRuleResponse struct {
	response.CommonBase

	// 返回码描述信息
	Message string

	// 路由规则Id
	RouteRuleId string
}

// NewPublishUGNRouteRuleRequest will create request of PublishUGNRouteRule action.
func (c *UGNClient) NewPublishUGNRouteRuleRequest() *PublishUGNRouteRuleRequest {
	req := &PublishUGNRouteRuleRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: PublishUGNRouteRule

发布云联网路由规则
*/
func (c *UGNClient) PublishUGNRouteRule(req *PublishUGNRouteRuleRequest) (*PublishUGNRouteRuleResponse, error) {
	var err error
	var res PublishUGNRouteRuleResponse

	reqCopier := *req

	err = c.Client.InvokeAction("PublishUGNRouteRule", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// UnpublishUGNRouteRuleRequest is request schema for UnpublishUGNRouteRule action
type UnpublishUGNRouteRuleRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"true"`

	// 归属实例路由规则Id
	DeriveRouteRuleId *string `required:"true"`

	// 归属实例路由表Id
	DeriveRouteTableId *string `required:"true"`

	// 实例Id
	InstanceId *string `required:"true"`

	// 云联网Id
	UGNId *string `required:"true"`
}

// UnpublishUGNRouteRuleResponse is response schema for UnpublishUGNRouteRule action
type UnpublishUGNRouteRuleResponse struct {
	response.CommonBase

	// 返回码描述信息
	Message string
}

// NewUnpublishUGNRouteRuleRequest will create request of UnpublishUGNRouteRule action.
func (c *UGNClient) NewUnpublishUGNRouteRuleRequest() *UnpublishUGNRouteRuleRequest {
	req := &UnpublishUGNRouteRuleRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: UnpublishUGNRouteRule

取消发布云联网路由规则
*/
func (c *UGNClient) UnpublishUGNRouteRule(req *UnpublishUGNRouteRuleRequest) (*UnpublishUGNRouteRuleResponse, error) {
	var err error
	var res UnpublishUGNRouteRuleResponse

	reqCopier := *req

	err = c.Client.InvokeAction("UnpublishUGNRouteRule", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
