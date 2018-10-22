//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UMem ResizeURedisGroup

package umem

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// ResizeURedisGroupRequest is request schema for ResizeURedisGroup action
type ResizeURedisGroupRequest struct {
	request.CommonBase

	// 组ID
	GroupId *string `required:"true"`

	// 内存大小, 单位:GB (需要大于原size,且小于等于32) 目前仅支持1/2/4/8/16/32 G 六种容量规格
	Size *int `required:"true"`

	// 代金券ID 请参考DescribeCoupon接口
	CouponId *int `required:"false"`
}

// ResizeURedisGroupResponse is response schema for ResizeURedisGroup action
type ResizeURedisGroupResponse struct {
	response.CommonBase
}

// NewResizeURedisGroupRequest will create request of ResizeURedisGroup action.
func (c *UMemClient) NewResizeURedisGroupRequest() *ResizeURedisGroupRequest {
	req := &ResizeURedisGroupRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// ResizeURedisGroup - 调整主备redis容量
func (c *UMemClient) ResizeURedisGroup(req *ResizeURedisGroupRequest) (*ResizeURedisGroupResponse, error) {
	var err error
	var res ResizeURedisGroupResponse

	err = c.client.InvokeAction("ResizeURedisGroup", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
