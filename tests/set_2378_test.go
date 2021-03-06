// Code is generated by ucloud-model, DO NOT EDIT IT.

package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet2378(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	ctx.SetVar("BucketName_2", ctx.Must(utest.Concat("apitest-", ctx.Must(utest.GetTimestamp(10)))))
	ctx.SetVar("original_type", "private")
	ctx.SetVar("tokenName", "test-auto-token")
	ctx.SetVar("allowed_ops_0", "TOKEN_ALLOW_WRITE")
	ctx.SetVar("allowed_ops_1", "TOKEN_ALLOW_READ")
	ctx.SetVar("allowed_ops_2", "TOKEN_ALLOW_DELETE")
	ctx.SetVar("allowed_ops_3", "TOKEN_ALLOW_LIST")
	ctx.SetVar("allowed_ops_4", "TOKEN_ALLOW_IOP")
	ctx.SetVar("allowedPrefixes", "test-auto")
	ctx.SetVar("mirror-old-prefix", "mirror-OldPrefix")
	ctx.SetVar("mirror-new-prefix", "mirror-NewPrefix")
	ctx.SetVar("DailyReport_StartTime", ctx.Must(utest.Calculate("-", ctx.Must(utest.GetTimestamp(10)), 2592000)))
	ctx.SetVar("DailyReport_EndTime", ctx.Must(utest.GetTimestamp(10)))

	testSet2378GetAvailableRegion00(&ctx)
	testSet2378GetFuctionAvailableRegion01(&ctx)
	testSet2378GetUFileDailyReport02(&ctx)
	testSet2378CreateBucket03(&ctx)
	testSet2378DescribeBucket04(&ctx)
	testSet2378BindBucketDomain05(&ctx)
	testSet2378DescribeBucketDomain06(&ctx)
	testSet2378UnbindBucketDomain07(&ctx)
	testSet2378UpdateBucket08(&ctx)
	testSet2378DescribeBucket09(&ctx)
	testSet2378DescribeMirrorRules10(&ctx)
	testSet2378CreateUFileToken11(&ctx)
	testSet2378DescribeUFileToken12(&ctx)
	testSet2378UpdateUFileToken13(&ctx)
	testSet2378DescribeUFileToken14(&ctx)
	testSet2378DeleteUFileToken15(&ctx)
	testSet2378DeleteBucket16(&ctx)
}

func testSet2378GetAvailableRegion00(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iufileClient.NewGetAvailableRegionRequest()

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iufileClient.GetAvailableRegion(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "GetAvailableRegionResponse", "str_eq"),
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

func testSet2378GetFuctionAvailableRegion01(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iufileClient.NewGetFuctionAvailableRegionRequest()

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iufileClient.GetFuctionAvailableRegion(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "GetFuctionAvailableRegionResponse", "str_eq"),
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

func testSet2378GetUFileDailyReport02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iufileClient.NewGetUFileDailyReportRequest()

	ctx.NoError(utest.SetReqValue(req, "StartTime", ctx.GetVar("DailyReport_StartTime")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "EndTime", ctx.GetVar("DailyReport_EndTime")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iufileClient.GetUFileDailyReport(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "GetUFileDailyReportResponse", "str_eq"),
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

func testSet2378CreateBucket03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ufileClient.NewCreateBucketRequest()

	ctx.NoError(utest.SetReqValue(req, "Type", ctx.GetVar("original_type")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "BucketName", ctx.GetVar("BucketName_2")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ufileClient.CreateBucket(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "CreateBucketResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Fatal(err)
		ctx.T.FailNow()

	}

	ctx.Vars["bucketId"] = ctx.Must(utest.GetValue(resp, "BucketId"))
}

func testSet2378DescribeBucket04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ufileClient.NewDescribeBucketRequest()

	ctx.NoError(utest.SetReqValue(req, "BucketName", ctx.GetVar("BucketName_2")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ufileClient.DescribeBucket(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "DescribeBucketResponse", "str_eq"),
			ctx.NewValidator("DataSet.0.BucketName", ctx.GetVar("BucketName_2"), "str_eq"),
			ctx.NewValidator("DataSet.0.Type", "private", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Fatal(err)
		ctx.T.FailNow()

	}

	ctx.Vars["bindDomain"] = ctx.Must(utest.GetValue(resp, "DataSet.0.Domain.Src.0"))
}

func testSet2378BindBucketDomain05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iufileClient.NewBindBucketDomainRequest()

	ctx.NoError(utest.SetReqValue(req, "Domain", ctx.GetVar("bindDomain")))

	ctx.NoError(utest.SetReqValue(req, "BucketName", ctx.GetVar("BucketName_2")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iufileClient.BindBucketDomain(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "BindBucketDomainResponse", "str_eq"),
			ctx.NewValidator("BucketId", ctx.GetVar("bucketId"), "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Fatal(err)
		ctx.T.FailNow()

	}

}

func testSet2378DescribeBucketDomain06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iufileClient.NewDescribeBucketDomainRequest()

	ctx.NoError(utest.SetReqValue(req, "BucketName", ctx.GetVar("BucketName_2")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iufileClient.DescribeBucketDomain(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "DescribeBucketDomainResponse", "str_eq"),
			ctx.NewValidator("DataSet.0.Domain", ctx.GetVar("bindDomain"), "str_eq"),
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

func testSet2378UnbindBucketDomain07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iufileClient.NewUnbindBucketDomainRequest()

	ctx.NoError(utest.SetReqValue(req, "Domain", ctx.GetVar("bindDomain")))

	ctx.NoError(utest.SetReqValue(req, "BucketName", ctx.GetVar("BucketName_2")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iufileClient.UnbindBucketDomain(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "UnbindBucketDomainResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Fatal(err)
		ctx.T.FailNow()

	}

}

func testSet2378UpdateBucket08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ufileClient.NewUpdateBucketRequest()

	ctx.NoError(utest.SetReqValue(req, "Type", "public"))

	ctx.NoError(utest.SetReqValue(req, "BucketName", ctx.GetVar("BucketName_2")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ufileClient.UpdateBucket(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "UpdateBucketResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Fatal(err)
		ctx.T.FailNow()

	}

}

func testSet2378DescribeBucket09(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ufileClient.NewDescribeBucketRequest()

	ctx.NoError(utest.SetReqValue(req, "BucketName", ctx.GetVar("BucketName_2")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ufileClient.DescribeBucket(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "DescribeBucketResponse", "str_eq"),
			ctx.NewValidator("DataSet.0.Type", "public", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Fatal(err)
		ctx.T.FailNow()

	}

}

func testSet2378DescribeMirrorRules10(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iufileClient.NewDescribeMirrorRulesRequest()

	ctx.NoError(utest.SetReqValue(req, "BucketName", ctx.GetVar("BucketName_2")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iufileClient.DescribeMirrorRules(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "DescribeMirrorRulesResponse", "str_eq"),
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

func testSet2378CreateUFileToken11(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ufileClient.NewCreateUFileTokenRequest()

	ctx.NoError(utest.SetReqValue(req, "TokenName", ctx.GetVar("tokenName")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "AllowedPrefixes", "*"))

	ctx.NoError(utest.SetReqValue(req, "AllowedOps", "TOKEN_ALLOW_READ"))

	ctx.NoError(utest.SetReqValue(req, "AllowedBuckets", ctx.GetVar("BucketName_2")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ufileClient.CreateUFileToken(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "CreateUFileTokenResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

	ctx.Vars["tokenId"] = ctx.Must(utest.GetValue(resp, "TokenId"))
}

func testSet2378DescribeUFileToken12(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ufileClient.NewDescribeUFileTokenRequest()

	ctx.NoError(utest.SetReqValue(req, "TokenId", ctx.GetVar("tokenId")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ufileClient.DescribeUFileToken(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "DescribeUFileTokenResponse", "str_eq"),
			ctx.NewValidator("DataSet.0.TokenName", ctx.GetVar("tokenName"), "str_eq"),
			ctx.NewValidator("DataSet.0.AllowedOps.0", "TOKEN_ALLOW_READ", "str_eq"),
			ctx.NewValidator("DataSet.0.AllowedPrefixes.0", "*", "str_eq"),
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

func testSet2378UpdateUFileToken13(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ufileClient.NewUpdateUFileTokenRequest()

	ctx.NoError(utest.SetReqValue(req, "TokenId", ctx.GetVar("tokenId")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "AllowedPrefixes", ctx.GetVar("allowedPrefixes")))

	ctx.NoError(utest.SetReqValue(req, "AllowedOps", ctx.GetVar("allowed_ops_0"), ctx.GetVar("allowed_ops_1"), ctx.GetVar("allowed_ops_2"), ctx.GetVar("allowed_ops_3"), ctx.GetVar("allowed_ops_4")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ufileClient.UpdateUFileToken(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "UpdateUFileTokenResponse", "str_eq"),
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

func testSet2378DescribeUFileToken14(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ufileClient.NewDescribeUFileTokenRequest()

	ctx.NoError(utest.SetReqValue(req, "TokenId", ctx.GetVar("tokenId")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ufileClient.DescribeUFileToken(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "DescribeUFileTokenResponse", "str_eq"),
			ctx.NewValidator("DataSet.0.TokenName", ctx.GetVar("tokenName"), "str_eq"),
			ctx.NewValidator("DataSet.0.AllowedPrefixes.0", ctx.GetVar("allowedPrefixes"), "str_eq"),
			ctx.NewValidator("DataSet.0.AllowedOps.0", "TOKEN_ALLOW_READ", "str_eq"),
			ctx.NewValidator("DataSet.0.AllowedOps.1", "TOKEN_ALLOW_WRITE", "str_eq"),
			ctx.NewValidator("DataSet.0.AllowedOps.2", "TOKEN_ALLOW_DELETE", "str_eq"),
			ctx.NewValidator("DataSet.0.AllowedOps.3", "TOKEN_ALLOW_LIST", "str_eq"),
			ctx.NewValidator("DataSet.0.AllowedOps.4", "TOKEN_ALLOW_IOP", "str_eq"),
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

func testSet2378DeleteUFileToken15(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ufileClient.NewDeleteUFileTokenRequest()

	ctx.NoError(utest.SetReqValue(req, "TokenId", ctx.GetVar("tokenId")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ufileClient.DeleteUFileToken(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "DeleteUFileTokenResponse", "str_eq"),
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

func testSet2378DeleteBucket16(ctx *utest.TestContext) {
	time.Sleep(time.Duration(30) * time.Second)

	req := ufileClient.NewDeleteBucketRequest()

	ctx.NoError(utest.SetReqValue(req, "BucketName", ctx.GetVar("BucketName_2")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ufileClient.DeleteBucket(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "DeleteBucketResponse", "str_eq"),
		},
		MaxRetries:    20,
		RetryInterval: 3 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}
