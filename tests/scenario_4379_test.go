// Code is generated by ucloud-model, DO NOT EDIT IT.

package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/services/uhost"
	"github.com/ucloud/ucloud-sdk-go/services/ulb"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/driver"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/utils"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/validation"
)

func TestScenario4379(t *testing.T) {
	spec.ParallelTest(t, &driver.Scenario{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Id: "4379",
		Vars: func(scenario *driver.Scenario) map[string]interface{} {
			return map[string]interface{}{
				"Image_Id_cloud": "#{u_get_image_resource($Region,$Zone)}",
				"saopaulo_image": "uimage-1bkjka",
				"Region":         "cn-bj2",
				"Zone":           "cn-bj2-02",
			}
		},
		Owners: []string{"li.wei@ucloud.cn"},
		Title:  "外网-ULB7自动化回归-内容转发-03",
		Steps: []*driver.Step{
			testStep4379DescribeImage01,
			testStep4379CreateUHostInstance02,
			testStep4379CreateULB03,
			testStep4379CreateVServer04,
			testStep4379DescribeVServer05,
			testStep4379AllocateBackend06,
			testStep4379DescribeVServer07,
			testStep4379CreatePolicy08,
			testStep4379UpdatePolicy09,
			testStep4379DeletePolicy10,
			testStep4379DeleteULB11,
			testStep4379PoweroffUHostInstance12,
			testStep4379TerminateUHostInstance13,
		},
	})
}

var testStep4379DescribeImage01 = &driver.Step{
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

var testStep4379CreateUHostInstance02 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewCreateUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":        step.Scenario.GetVar("Zone"),
			"Tag":         "Default",
			"Region":      step.Scenario.GetVar("Region"),
			"Password":    "VXFhNzg5VGVzdCFAIyQ7LA==",
			"Name":        "ulb-host",
			"Memory":      1024,
			"MachineType": "N",
			"LoginMode":   "Password",
			"ImageId":     step.Scenario.GetVar("Image_Id_cloud"),
			"Disks": []map[string]interface{}{
				{
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

		step.Scenario.SetVar("UHostId_01", step.Must(utils.GetValue(resp, "UHostIds.0")))
		step.Scenario.SetVar("IP_01", step.Must(utils.GetValue(resp, "IPs.0")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(5) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "创建云主机",
	FastFail:      false,
}

var testStep4379CreateULB03 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewCreateULBRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"ULBName":   "测试",
			"Tag":       "Default",
			"Region":    step.Scenario.GetVar("Region"),
			"InnerMode": "No",
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateULB(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("ULBId", step.Must(utils.GetValue(resp, "ULBId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(60) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "创建负载均衡",
	FastFail:      false,
}

var testStep4379CreateVServer04 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewCreateVServerRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"VServerName":     "vserver-test",
			"ULBId":           step.Scenario.GetVar("ULBId"),
			"Region":          step.Scenario.GetVar("Region"),
			"Protocol":        "HTTP",
			"PersistenceType": "UserDefined",
			"PersistenceInfo": "huangchao",
			"Method":          "Roundrobin",
			"ListenType":      "RequestProxy",
			"FrontendPort":    80,
			"ClientTimeout":   60,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateVServer(req)
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
	StartupDelay:  time.Duration(5) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "创建VServer",
	FastFail:      false,
}

var testStep4379DescribeVServer05 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewDescribeVServerRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"ULBId":  step.Scenario.GetVar("ULBId"),
			"Region": step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeVServer(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("VServerId", step.Must(utils.GetValue(resp, "DataSet.0.VServerId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("DataSet", 1, "len_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "获取VServer信息",
	FastFail:      false,
}

var testStep4379AllocateBackend06 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewAllocateBackendRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"VServerId":    step.Scenario.GetVar("VServerId"),
			"ULBId":        step.Scenario.GetVar("ULBId"),
			"ResourceType": "UHost",
			"ResourceId":   step.Scenario.GetVar("UHostId_01"),
			"Region":       step.Scenario.GetVar("Region"),
			"Port":         80,
			"Enabled":      1,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.AllocateBackend(req)
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
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "添加后端实例",
	FastFail:      false,
}

var testStep4379DescribeVServer07 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewDescribeVServerRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"VServerId": step.Scenario.GetVar("VServerId"),
			"ULBId":     step.Scenario.GetVar("ULBId"),
			"Region":    step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeVServer(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("BackendId", step.Must(utils.GetValue(resp, "DataSet.0.BackendSet.0.BackendId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.BackendSet", 1, "len_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "获取VServer信息",
	FastFail:      false,
}

var testStep4379CreatePolicy08 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewCreatePolicyRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"VServerId": step.Scenario.GetVar("VServerId"),
			"ULBId":     step.Scenario.GetVar("ULBId"),
			"Type":      "Domain",
			"Region":    step.Scenario.GetVar("Region"),
			"Match":     "www.test.com",
			"BackendId": []interface{}{
				step.Scenario.GetVar("BackendId"),
			},
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreatePolicy(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("PolicyId", step.Must(utils.GetValue(resp, "PolicyId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "创建转发策略",
	FastFail:      false,
}

var testStep4379UpdatePolicy09 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewUpdatePolicyRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"VServerId": step.Scenario.GetVar("VServerId"),
			"ULBId":     step.Scenario.GetVar("ULBId"),
			"Type":      "Domain",
			"Region":    step.Scenario.GetVar("Region"),
			"PolicyId":  step.Scenario.GetVar("PolicyId"),
			"Match":     "www.testgai.com",
			"BackendId": []interface{}{
				step.Scenario.GetVar("BackendId"),
			},
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.UpdatePolicy(req)
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
	StartupDelay:  time.Duration(5) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "更新内容转发规则",
	FastFail:      false,
}

var testStep4379DeletePolicy10 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewDeletePolicyRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"VServerId": step.Scenario.GetVar("VServerId"),
			"Region":    step.Scenario.GetVar("Region"),
			"PolicyId":  step.Scenario.GetVar("PolicyId"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DeletePolicy(req)
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
	StartupDelay:  time.Duration(5) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "删除转发策略",
	FastFail:      false,
}

var testStep4379DeleteULB11 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewDeleteULBRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"ULBId":  step.Scenario.GetVar("ULBId"),
			"Region": step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DeleteULB(req)
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
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "删除负载均衡",
	FastFail:      false,
}

var testStep4379PoweroffUHostInstance12 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewPoweroffUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("UHostId_01"),
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
	StartupDelay:  time.Duration(60) * time.Second,
	MaxRetries:    3,
	RetryInterval: 30 * time.Second,
	Title:         "模拟主机掉电",
	FastFail:      false,
}

var testStep4379TerminateUHostInstance13 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewTerminateUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("UHostId_01"),
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
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除云主机",
	FastFail:      false,
}
