// Code is generated by ucloud-model, DO NOT EDIT IT.

package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/driver"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/functions"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/utils"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/validation"

	"github.com/ucloud/ucloud-sdk-go/ucloud"

	"github.com/ucloud/ucloud-sdk-go/services/uhost"

	"github.com/ucloud/ucloud-sdk-go/services/unet"

	"github.com/ucloud/ucloud-sdk-go/services/vpc"
)

func TestScenario614(t *testing.T) {
	spec.ParallelTest(t, &driver.Scenario{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Id: "614",
		Vars: func(scenario *driver.Scenario) map[string]interface{} {
			return map[string]interface{}{
				"Image_Id_cloud": "#{u_get_image_resource($Region,$Zone)}",
				"Region":         "cn-bj2",
				"Zone":           "cn-bj2-02",
			}
		},
		Owners: []string{"li.wei@ucloud.cn"},
		Title:  "新版NAT网关-natgw自动化回归-白名单-03-BGP线路",
		Steps: []*driver.Step{
			testStep614DescribeImage01,
			testStep614CreateVPC02,
			testStep614CreateSubnet03,
			testStep614CreateUHostInstance04,
			testStep614AllocateEIP05,
			testStep614DescribeFirewall06,
			testStep614CreateNATGW07,
			testStep614DescribeEIP08,
			testStep614GetAvailableResourceForWhiteList09,
			testStep614GetAvailableHostForWhiteList10,
			testStep614AddWhiteListResource11,
			testStep614DescribeWhiteListResource12,
			testStep614DeleteWhiteListResource13,
			testStep614DescribeWhiteList14,
			testStep614EnableWhiteList15,
			testStep614DeleteNATGW16,
			testStep614ReleaseEIP17,
			testStep614PoweroffUHostInstance18,
			testStep614TerminateUHostInstance19,
			testStep614DeleteSubnet20,
			testStep614DeleteVPC21,
		},
	})
}

var testStep614DescribeImage01 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewDescribeImageRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":      step.Scenario.GetVar("Zone"),
			"Region":    step.Scenario.GetVar("Region"),
			"OsType":    "Linux",
			"ImageType": "Base",
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeImage(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("Image_Id_cloud", step.Must(utils.GetValue(resp, "ImageSet.0.ImageId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "DescribeImageResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取镜像列表",
	FastFail:      false,
}

var testStep614CreateVPC02 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewCreateVPCRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region": step.Scenario.GetVar("Region"),
			"Network": []interface{}{
				"172.16.0.0/12",
			},
			"Name": "vpc-natgw-bgp",
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateVPC(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("VPCId", step.Must(utils.GetValue(resp, "VPCId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    3,
	RetryInterval: 10 * time.Second,
	Title:         "创建VPC",
	FastFail:      false,
}

var testStep614CreateSubnet03 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewCreateSubnetRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"VPCId":      step.Scenario.GetVar("VPCId"),
			"SubnetName": "natgw-s1-bgp",
			"Subnet":     "172.16.0.0",
			"Region":     step.Scenario.GetVar("Region"),
			"Netmask":    21,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateSubnet(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("SubnetId", step.Must(utils.GetValue(resp, "SubnetId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    3,
	RetryInterval: 10 * time.Second,
	Title:         "创建子网",
	FastFail:      false,
}

var testStep614CreateUHostInstance04 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewCreateUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":        step.Scenario.GetVar("Zone"),
			"VPCId":       step.Scenario.GetVar("VPCId"),
			"Tag":         "Default",
			"SubnetId":    step.Scenario.GetVar("SubnetId"),
			"Region":      step.Scenario.GetVar("Region"),
			"Password":    "VXFhNzg5VGVzdCFAIyQ7LA==",
			"Name":        "natgw-s1-bgp",
			"Memory":      1024,
			"MachineType": "N",
			"LoginMode":   "Password",
			"ImageId":     step.Scenario.GetVar("Image_Id_cloud"),
			"Disks": []map[string]interface{}{
				map[string]interface{}{
					"IsBoot": "True",
					"Size":   20,
					"Type":   "LOCAL_NORMAL",
				},
			},
			"CPU": 1,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateUHostInstance(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("UHostIds_s1", step.Must(utils.GetValue(resp, "UHostIds.0")))
		step.Scenario.SetVar("IPs_s1", step.Must(utils.GetValue(resp, "IPs.0")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    3,
	RetryInterval: 10 * time.Second,
	Title:         "创建云主机",
	FastFail:      false,
}

var testStep614AllocateEIP05 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewAllocateEIPRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Tag":          "Default",
			"Region":       step.Scenario.GetVar("Region"),
			"Quantity":     1,
			"PayMode":      "Bandwidth",
			"OperatorName": "Bgp",
			"Name":         "natgw-bgp",
			"ChargeType":   "Month",
			"Bandwidth":    2,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.AllocateEIP(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("EIPId", step.Must(utils.GetValue(resp, "EIPSet.0.EIPId")))
		step.Scenario.SetVar("EIP", step.Must(utils.GetValue(resp, "EIPSet.0.EIPAddr.0.IP")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(180) * time.Second,
	MaxRetries:    3,
	RetryInterval: 10 * time.Second,
	Title:         "申请弹性IP",
	FastFail:      false,
}

var testStep614DescribeFirewall06 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewDescribeFirewallRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region": step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeFirewall(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("FWId", step.Must(utils.GetValue(resp, "DataSet")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    3,
	RetryInterval: 10 * time.Second,
	Title:         "获取防火墙信息",
	FastFail:      false,
}

var testStep614CreateNATGW07 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewCreateNATGWRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"VPCId": step.Scenario.GetVar("VPCId"),
			"Tag":   "Default",
			"SubnetworkIds": []interface{}{
				step.Scenario.GetVar("SubnetId"),
			},
			"Remark":     "bgp",
			"Region":     step.Scenario.GetVar("Region"),
			"NATGWName":  "natgw-bgp",
			"IfOpen":     0,
			"FirewallId": step.Must(functions.SearchValue(step.Scenario.GetVar("FWId"), "Type", "recommend web", "FWId")),
			"EIPIds": []interface{}{
				step.Scenario.GetVar("EIPId"),
			},
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateNATGW(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "创建NatGateway",
	FastFail:      false,
}

var testStep614DescribeEIP08 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewDescribeEIPRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region": step.Scenario.GetVar("Region"),
			"EIPIds": []interface{}{
				step.Scenario.GetVar("EIPId"),
			},
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeEIP(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("NATGWId", step.Must(utils.GetValue(resp, "EIPSet.0.Resource.ResourceID")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("EIPSet.0.Resource.ResourceType", "natgw", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "获取弹性IP信息",
	FastFail:      false,
}

var testStep614GetAvailableResourceForWhiteList09 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewGetAvailableResourceForWhiteListRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region":  step.Scenario.GetVar("Region"),
			"NATGWId": step.Scenario.GetVar("NATGWId"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.GetAvailableResourceForWhiteList(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    3,
	RetryInterval: 10 * time.Second,
	Title:         "获得可添加白名单的资源列表",
	FastFail:      false,
}

var testStep614GetAvailableHostForWhiteList10 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("GetAvailableHostForWhiteList")
		err = req.SetPayload(map[string]interface{}{
			"Region":  step.Scenario.GetVar("Region"),
			"NATGWId": step.Scenario.GetVar("NATGWId"),
		})
		if err != nil {
			return nil, err
		}
		resp, err := client.GenericInvoke(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "GetAvailableHostForWhiteListResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取可用于白名单的uhost",
	FastFail:      false,
}

var testStep614AddWhiteListResource11 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewAddWhiteListResourceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"ResourceIds": []interface{}{
				step.Scenario.GetVar("UHostIds_s1"),
			},
			"Region":  step.Scenario.GetVar("Region"),
			"NATGWId": step.Scenario.GetVar("NATGWId"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.AddWhiteListResource(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    3,
	RetryInterval: 10 * time.Second,
	Title:         "添加白名单资源",
	FastFail:      false,
}

var testStep614DescribeWhiteListResource12 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewDescribeWhiteListResourceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region": step.Scenario.GetVar("Region"),
			"NATGWIds": []interface{}{
				step.Scenario.GetVar("NATGWId"),
			},
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeWhiteListResource(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    3,
	RetryInterval: 10 * time.Second,
	Title:         "描述白名单资源的详细信息",
	FastFail:      false,
}

var testStep614DeleteWhiteListResource13 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewDeleteWhiteListResourceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"ResourceIds": []interface{}{
				step.Scenario.GetVar("UHostIds_s1"),
			},
			"Region":  step.Scenario.GetVar("Region"),
			"NATGWId": step.Scenario.GetVar("NATGWId"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DeleteWhiteListResource(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    3,
	RetryInterval: 10 * time.Second,
	Title:         "删除白名单列表资源",
	FastFail:      false,
}

var testStep614DescribeWhiteList14 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("DescribeWhiteList")
		err = req.SetPayload(map[string]interface{}{
			"Region": step.Scenario.GetVar("Region"),
			"NATGWIds": []interface{}{
				step.Scenario.GetVar("NATGWId"),
			},
		})
		if err != nil {
			return nil, err
		}
		resp, err := client.GenericInvoke(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    3,
	RetryInterval: 10 * time.Second,
	Title:         "获取NatGateway白名单列表",
	FastFail:      false,
}

var testStep614EnableWhiteList15 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewEnableWhiteListRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region":  step.Scenario.GetVar("Region"),
			"NATGWId": step.Scenario.GetVar("NATGWId"),
			"IfOpen":  1,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.EnableWhiteList(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    3,
	RetryInterval: 10 * time.Second,
	Title:         "修改 NatGateWay白名单开关",
	FastFail:      false,
}

var testStep614DeleteNATGW16 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewDeleteNATGWRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region":  step.Scenario.GetVar("Region"),
			"NATGWId": step.Scenario.GetVar("NATGWId"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DeleteNATGW(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    3,
	RetryInterval: 10 * time.Second,
	Title:         "删除NatGateway",
	FastFail:      false,
}

var testStep614ReleaseEIP17 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewReleaseEIPRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region": step.Scenario.GetVar("Region"),
			"EIPId":  step.Scenario.GetVar("EIPId"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.ReleaseEIP(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{}
	},
	StartupDelay:  time.Duration(5) * time.Second,
	MaxRetries:    3,
	RetryInterval: 10 * time.Second,
	Title:         "释放弹性IP",
	FastFail:      false,
}

var testStep614PoweroffUHostInstance18 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewPoweroffUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("UHostIds_s1"),
			"Region":  step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.PoweroffUHostInstance(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{}
	},
	StartupDelay:  time.Duration(5) * time.Second,
	MaxRetries:    3,
	RetryInterval: 10 * time.Second,
	Title:         "模拟主机掉电",
	FastFail:      false,
}

var testStep614TerminateUHostInstance19 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewTerminateUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("UHostIds_s1"),
			"Region":  step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.TerminateUHostInstance(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{}
	},
	StartupDelay:  time.Duration(60) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "删除云主机",
	FastFail:      false,
}

var testStep614DeleteSubnet20 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewDeleteSubnetRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"SubnetId": step.Scenario.GetVar("SubnetId"),
			"Region":   step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DeleteSubnet(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{}
	},
	StartupDelay:  time.Duration(5) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "删除子网",
	FastFail:      false,
}

var testStep614DeleteVPC21 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewDeleteVPCRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"VPCId":  step.Scenario.GetVar("VPCId"),
			"Region": step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DeleteVPC(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "删除VPC",
	FastFail:      false,
}
