//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UMem DescribeURedisPrice

package umem

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeURedisPriceRequest is request schema for DescribeURedisPrice action
type DescribeURedisPriceRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"true"`

	// 量大小,单位:GB  取值范围[1-32]
	Size *int `required:"true"`

	// 计费模式，Year， Month， Dynamic；如果不指定，则一次性获取三种计费
	ChargeType *string `required:"false"`

	// 计费模式为Dynamic时，购买的时长, 默认为1
	Quantity *int `required:"false"`

	// 是否是跨机房URedis(默认false)
	RegionFlag *bool `required:"false"`

	// 产品类型：MS_Redis（标准主备版），S_Redis（从库），默认为MS_Redis
	ProductType *string `required:"false"`
}

// DescribeURedisPriceResponse is response schema for DescribeURedisPrice action
type DescribeURedisPriceResponse struct {
	response.CommonBase

	// 价格 参数见 UMemPriceSet
	DataSet []URedisPriceSet
}

// NewDescribeURedisPriceRequest will create request of DescribeURedisPrice action.
func (c *UMemClient) NewDescribeURedisPriceRequest() *DescribeURedisPriceRequest {
	req := &DescribeURedisPriceRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeURedisPrice - 取uredis价格信息
func (c *UMemClient) DescribeURedisPrice(req *DescribeURedisPriceRequest) (*DescribeURedisPriceResponse, error) {
	var err error
	var res DescribeURedisPriceResponse

	err = c.client.InvokeAction("DescribeURedisPrice", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
