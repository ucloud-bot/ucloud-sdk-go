// Code is generated by ucloud-model, DO NOT EDIT IT.

package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/services/uhost"
	"github.com/ucloud/ucloud-sdk-go/services/unet"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/driver"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/functions"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/utils"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/validation"
)

func TestScenario302(t *testing.T) {
	spec.ParallelTest(t, &driver.Scenario{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Id: "302",
		Vars: func(scenario *driver.Scenario) map[string]interface{} {
			return map[string]interface{}{
				"tag":                "tag_api_test",
				"remark":             "remark_api_test",
				"fw_name_1":          scenario.Must(functions.Concat("fw_A_", scenario.Must(functions.GetTimestamp(13)), "_")),
				"fw_rule_1_protocol": "TCP",
				"fw_rule_1_port":     1111,
				"fw_rule_1_srcip":    "0.0.0.0/0",
				"fw_rule_1_action":   "ACCEPT",
				"fw_rule_1_priority": "HIGH",
				"uhost_name_1":       "firewall_api_test",
				"fw_rule_2_protocol": "UDP",
				"fw_rule_2_port":     2222,
				"fw_rule_2_srcip":    "10.0.0.0/8",
				"fw_rule_2_action":   "DROP",
				"fw_rule_2_priority": "LOW",
				"fw_name_2":          scenario.Must(functions.Concat("fw_B_", scenario.Must(functions.GetTimestamp(13)), "_")),
				"tag_2":              "tag_api_test_3",
				"remark_2":           "remark_api_test_3",
				"Image_Id":           "#{u_get_image_resource($Region,$Zone)}",
				"recommend_web":      "recommend web",
				"fw_rule_1":          "TCP|1111|0.0.0.0/0|ACCEPT|HIGH",
				"fw_rule_2":          "UDP|2222|10.0.0.0/8|DROP|LOW",
				"Region":             "cn-bj2",
				"Zone":               "cn-bj2-02",
			}
		},
		Owners: []string{"gemin.jiang@ucloud.cn"},
		Title:  "firewall自动化回归-防火墙-Firewall-01",
		Steps: []*driver.Step{
			testStep302DescribeImage01,
			testStep302DescribeFirewall02,
			testStep302CreateUHostInstance03,
			testStep302CreateFirewall04,
			testStep302DescribeFirewall05,
			testStep302DescribeFirewall06,
			testStep302GrantFirewall07,
			testStep302DescribeFirewallResource08,
			testStep302UpdateFirewall09,
			testStep302UpdateFirewallAttribute10,
			testStep302DescribeFirewall11,
			testStep302DescribeFirewall12,
			testStep302GrantFirewall13,
			testStep302PoweroffUHostInstance14,
			testStep302TerminateUHostInstance15,
			testStep302DescribeFirewall16,
			testStep302DeleteFirewall17,
			testStep302DescribeFirewall18,
		},
	})
}

var testStep302DescribeImage01 = &driver.Step{
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

		step.Scenario.SetVar("Image_Id", step.Must(utils.GetValue(resp, "ImageSet.0.ImageId")))
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

var testStep302DescribeFirewall02 = &driver.Step{
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

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取防火墙信息",
	FastFail:      false,
}

var testStep302CreateUHostInstance03 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewCreateUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":        step.Scenario.GetVar("Zone"),
			"Tag":         step.Scenario.GetVar("tag"),
			"Region":      step.Scenario.GetVar("Region"),
			"Password":    "VXFhNzg5VGVzdCFAIyQ7LA==",
			"Name":        step.Scenario.GetVar("uhost_name_1"),
			"Memory":      1024,
			"MachineType": "N",
			"LoginMode":   "Password",
			"ImageId":     step.Scenario.GetVar("Image_Id"),
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

		step.Scenario.SetVar("uhost_id1", step.Must(utils.GetValue(resp, "UHostIds.0")))
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

var testStep302CreateFirewall04 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewCreateFirewallRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Tag": step.Scenario.GetVar("tag"),
			"Rule": []interface{}{
				step.Scenario.GetVar("fw_rule_1"),
			},
			"Remark": step.Scenario.GetVar("remark"),
			"Region": step.Scenario.GetVar("Region"),
			"Name":   step.Must(functions.Concat(step.Scenario.GetVar("fw_name_1"), step.Scenario.GetVar("Zone"))),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateFirewall(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("fw_id1", step.Must(utils.GetValue(resp, "FWId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(180) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "创建防火墙",
	FastFail:      true,
}

var testStep302DescribeFirewall05 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewDescribeFirewallRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region": step.Scenario.GetVar("Region"),
			"FWId":   step.Scenario.GetVar("fw_id1"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeFirewall(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.FWId", step.Scenario.GetVar("fw_id1"), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Name", step.Must(functions.Concat(step.Scenario.GetVar("fw_name_1"), step.Scenario.GetVar("Zone"))), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Tag", step.Scenario.GetVar("tag"), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Remark", step.Scenario.GetVar("remark"), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.ResourceCount", 0, "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Type", "user defined", "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Rule.0.ProtocolType", step.Scenario.GetVar("fw_rule_1_protocol"), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Rule.0.DstPort", step.Scenario.GetVar("fw_rule_1_port"), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Rule.0.SrcIP", step.Scenario.GetVar("fw_rule_1_srcip"), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Rule.0.RuleAction", step.Scenario.GetVar("fw_rule_1_action"), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Rule.0.Priority", step.Scenario.GetVar("fw_rule_1_priority"), "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    0,
	RetryInterval: 0 * time.Second,
	Title:         "获取防火墙信息",
	FastFail:      false,
}

var testStep302DescribeFirewall06 = &driver.Step{
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

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取防火墙信息",
	FastFail:      false,
}

var testStep302GrantFirewall07 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewGrantFirewallRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"ResourceType": "UHost",
			"ResourceId":   step.Scenario.GetVar("uhost_id1"),
			"Region":       step.Scenario.GetVar("Region"),
			"FWId":         step.Scenario.GetVar("fw_id1"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.GrantFirewall(req)
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
	MaxRetries:    0,
	RetryInterval: 0 * time.Second,
	Title:         "应用防火墙",
	FastFail:      false,
}

var testStep302DescribeFirewallResource08 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewDescribeFirewallResourceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region": step.Scenario.GetVar("Region"),
			"FWId":   step.Scenario.GetVar("fw_id1"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeFirewallResource(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("ResourceSet.0.Name", step.Scenario.GetVar("uhost_name_1"), "str_eq"),
			validation.Builtins.NewValidator("ResourceSet.0.ResourceType", "uhost", "str_eq"),
			validation.Builtins.NewValidator("ResourceSet.0.ResourceID", step.Scenario.GetVar("uhost_id1"), "str_eq"),
			validation.Builtins.NewValidator("TotalCount", 1, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    0,
	RetryInterval: 0 * time.Second,
	Title:         "获取防火墙绑定资源",
	FastFail:      false,
}

var testStep302UpdateFirewall09 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewUpdateFirewallRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Rule": []interface{}{
				step.Scenario.GetVar("fw_rule_2"),
			},
			"Region": step.Scenario.GetVar("Region"),
			"FWId":   step.Scenario.GetVar("fw_id1"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.UpdateFirewall(req)
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
	MaxRetries:    0,
	RetryInterval: 0 * time.Second,
	Title:         "更新防火墙规则",
	FastFail:      false,
}

var testStep302UpdateFirewallAttribute10 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewUpdateFirewallAttributeRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Tag":    step.Scenario.GetVar("tag_2"),
			"Remark": step.Scenario.GetVar("remark_2"),
			"Region": step.Scenario.GetVar("Region"),
			"Name":   step.Must(functions.Concat(step.Scenario.GetVar("fw_name_2"), step.Scenario.GetVar("Zone"))),
			"FWId":   step.Scenario.GetVar("fw_id1"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.UpdateFirewallAttribute(req)
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
	MaxRetries:    0,
	RetryInterval: 0 * time.Second,
	Title:         "更新防火墙属性",
	FastFail:      false,
}

var testStep302DescribeFirewall11 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewDescribeFirewallRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region": step.Scenario.GetVar("Region"),
			"FWId":   step.Scenario.GetVar("fw_id1"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeFirewall(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.FWId", step.Scenario.GetVar("fw_id1"), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Name", step.Must(functions.Concat(step.Scenario.GetVar("fw_name_2"), step.Scenario.GetVar("Zone"))), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Tag", step.Scenario.GetVar("tag_2"), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Remark", step.Scenario.GetVar("remark_2"), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.ResourceCount", 1, "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Type", "user defined", "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Rule.0.ProtocolType", step.Scenario.GetVar("fw_rule_2_protocol"), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Rule.0.DstPort", step.Scenario.GetVar("fw_rule_2_port"), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Rule.0.SrcIP", step.Scenario.GetVar("fw_rule_2_srcip"), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Rule.0.RuleAction", step.Scenario.GetVar("fw_rule_2_action"), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Rule.0.Priority", step.Scenario.GetVar("fw_rule_2_priority"), "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    0,
	RetryInterval: 0 * time.Second,
	Title:         "获取防火墙信息",
	FastFail:      false,
}

var testStep302DescribeFirewall12 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewDescribeFirewallRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region": step.Scenario.GetVar("Region"),
			"Offset": 0,
			"Limit":  20,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeFirewall(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("fw_dataset", step.Must(utils.GetValue(resp, "DataSet")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取防火墙信息",
	FastFail:      false,
}

var testStep302GrantFirewall13 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewGrantFirewallRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"ResourceType": "UHost",
			"ResourceId":   step.Scenario.GetVar("uhost_id1"),
			"Region":       step.Scenario.GetVar("Region"),
			"FWId":         step.Must(functions.SearchValue(step.Scenario.GetVar("fw_dataset"), "Type", step.Scenario.GetVar("recommend_web"), "FWId")),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.GrantFirewall(req)
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
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "应用防火墙",
	FastFail:      false,
}

var testStep302PoweroffUHostInstance14 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewPoweroffUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("uhost_id1"),
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
	StartupDelay:  time.Duration(30) * time.Second,
	MaxRetries:    3,
	RetryInterval: 60 * time.Second,
	Title:         "模拟主机掉电",
	FastFail:      false,
}

var testStep302TerminateUHostInstance15 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewTerminateUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("uhost_id1"),
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
	StartupDelay:  time.Duration(90) * time.Second,
	MaxRetries:    3,
	RetryInterval: 10 * time.Second,
	Title:         "删除云主机",
	FastFail:      false,
}

var testStep302DescribeFirewall16 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewDescribeFirewallRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region": step.Scenario.GetVar("Region"),
			"FWId":   step.Scenario.GetVar("fw_id1"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeFirewall(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.ResourceCount", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(60) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取防火墙信息",
	FastFail:      false,
}

var testStep302DeleteFirewall17 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewDeleteFirewallRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region": step.Scenario.GetVar("Region"),
			"FWId":   step.Scenario.GetVar("fw_id1"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DeleteFirewall(req)
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
	StartupDelay:  time.Duration(30) * time.Second,
	MaxRetries:    10,
	RetryInterval: 1 * time.Second,
	Title:         "删除防火墙",
	FastFail:      true,
}

var testStep302DescribeFirewall18 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewDescribeFirewallRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region": step.Scenario.GetVar("Region"),
			"Offset": 0,
			"Limit":  10,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeFirewall(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("DataSet", step.Scenario.GetVar("fw_id1"), "object_not_contains"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取防火墙信息",
	FastFail:      false,
}
