//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDB DeleteUDBParamGroup

package udb

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DeleteUDBParamGroupRequest is request schema for DeleteUDBParamGroup action
type DeleteUDBParamGroupRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"true"`

	// 参数组id,可通过DescribeUDBParamGroup获取
	GroupId *int `required:"true"`

	// 是否属于地域级别
	RegionFlag *bool `required:"false"`
}

// DeleteUDBParamGroupResponse is response schema for DeleteUDBParamGroup action
type DeleteUDBParamGroupResponse struct {
	response.CommonBase
}

// NewDeleteUDBParamGroupRequest will create request of DeleteUDBParamGroup action.
func (c *UDBClient) NewDeleteUDBParamGroupRequest() *DeleteUDBParamGroupRequest {
	req := &DeleteUDBParamGroupRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DeleteUDBParamGroup - 删除配置参数组
func (c *UDBClient) DeleteUDBParamGroup(req *DeleteUDBParamGroupRequest) (*DeleteUDBParamGroupResponse, error) {
	var err error
	var res DeleteUDBParamGroupResponse

	err = c.client.InvokeAction("DeleteUDBParamGroup", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}