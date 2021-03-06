// Code is generated by ucloud-model, DO NOT EDIT IT.

package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet4214(t *testing.T) {
	t.Skip()
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	ctx.SetVar("EndTime", ctx.Must(utest.GetTimestamp(10)))
	ctx.SetVar("BeginTime", ctx.Must(utest.Calculate("-", ctx.Must(utest.GetTimestamp(10)), 3600)))
	ctx.SetVar("Region", "cn")
	ctx.SetVar("Zone", "zone-01")

	testSet4214CreateVMInstance00(&ctx)
	testSet4214DescribeVMInstance01(&ctx)
	testSet4214DescribeMetric02(&ctx)
	testSet4214StopVMInstance03(&ctx)
	testSet4214DescribeVMInstance04(&ctx)
	testSet4214DeleteVMInstance05(&ctx)
}

func testSet4214CreateVMInstance00(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ucloudstackClient.NewCreateVMInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "WANSGID", "sg-1al_S_tbN"))

	ctx.NoError(utest.SetReqValue(req, "VPCID", "vpc-1al_S_tbN"))

	ctx.NoError(utest.SetReqValue(req, "VMType", "Normal"))

	ctx.NoError(utest.SetReqValue(req, "SubnetID", "subnet-1al_S_tbN"))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "Quantity", 1))

	ctx.NoError(utest.SetReqValue(req, "Password", "ucloud.cn"))

	ctx.NoError(utest.SetReqValue(req, "Name", "sdk-test"))

	ctx.NoError(utest.SetReqValue(req, "Memory", 2048))

	ctx.NoError(utest.SetReqValue(req, "ImageID", "cn-image-centos-74"))

	ctx.NoError(utest.SetReqValue(req, "DataDiskType", "Normal"))

	ctx.NoError(utest.SetReqValue(req, "DataDiskSpace", 0))

	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Month"))

	ctx.NoError(utest.SetReqValue(req, "CPU", 1))

	ctx.NoError(utest.SetReqValue(req, "BootDiskType", "Normal"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ucloudstackClient.CreateVMInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "CreateVMInstanceResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

	ctx.Vars["VMID"] = ctx.Must(utest.GetValue(resp, "VMID"))
}

func testSet4214DescribeVMInstance01(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ucloudstackClient.NewDescribeVMInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "VMIDs", ctx.GetVar("VMID")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ucloudstackClient.DescribeVMInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "DescribeVMInstanceResponse", "str_eq"),
			ctx.NewValidator("Infos.0.State", "Running", "str_eq"),
		},
		MaxRetries:    30,
		RetryInterval: 10 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet4214DescribeMetric02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ucloudstackClient.NewDescribeMetricRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "ResourceID", ctx.GetVar("VMID")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "MetricName", "DiskReadOps"))

	ctx.NoError(utest.SetReqValue(req, "EndTime", ctx.GetVar("EndTime")))

	ctx.NoError(utest.SetReqValue(req, "BeginTime", ctx.GetVar("BeginTime")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ucloudstackClient.DescribeMetric(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "DescribeMetricResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet4214StopVMInstance03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ucloudstackClient.NewStopVMInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "VMID", ctx.GetVar("VMID")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ucloudstackClient.StopVMInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "StopVMInstanceResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet4214DescribeVMInstance04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ucloudstackClient.NewDescribeVMInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "VMIDs", ctx.GetVar("VMID")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ucloudstackClient.DescribeVMInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "DescribeVMInstanceResponse", "str_eq"),
			ctx.NewValidator("Infos.0.State", "Stopped", "str_eq"),
		},
		MaxRetries:    30,
		RetryInterval: 10 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet4214DeleteVMInstance05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ucloudstackClient.NewDeleteVMInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "VMID", ctx.GetVar("VMID")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ucloudstackClient.DeleteVMInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "DeleteVMInstanceResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}
}
