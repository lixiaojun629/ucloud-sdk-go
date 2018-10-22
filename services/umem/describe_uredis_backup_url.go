//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UMem DescribeURedisBackupURL

package umem

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeURedisBackupURLRequest is request schema for DescribeURedisBackupURL action
type DescribeURedisBackupURLRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"true"`

	// 备份ID
	BackupId *int `required:"true"`

	// 是否是跨机房URedis(默认false)
	RegionFlag *bool `required:"false"`
}

// DescribeURedisBackupURLResponse is response schema for DescribeURedisBackupURL action
type DescribeURedisBackupURLResponse struct {
	response.CommonBase

	// 备份文件公网的地址
	BackupURL string
}

// NewDescribeURedisBackupURLRequest will create request of DescribeURedisBackupURL action.
func (c *UMemClient) NewDescribeURedisBackupURLRequest() *DescribeURedisBackupURLRequest {
	req := &DescribeURedisBackupURLRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeURedisBackupURL - 获取主备Redis备份下载链接
func (c *UMemClient) DescribeURedisBackupURL(req *DescribeURedisBackupURLRequest) (*DescribeURedisBackupURLResponse, error) {
	var err error
	var res DescribeURedisBackupURLResponse

	err = c.client.InvokeAction("DescribeURedisBackupURL", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
