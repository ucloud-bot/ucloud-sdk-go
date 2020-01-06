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
)

func TestScenario4420(t *testing.T) {
	spec.ParallelTest(t, &driver.Scenario{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Id: "4420",
		Vars: func(scenario *driver.Scenario) map[string]interface{} {
			return map[string]interface{}{
				"Region":             "cn-sh2",
				"Zone":               "cn-sh2-02",
				"ChargeType":         "Month",
				"CreateCPU":          1,
				"CreateMem":          1024,
				"Name":               "授权角色",
				"BootSize":           20,
				"BootType":           "LOCAL_NORMAL",
				"DiskSize":           20,
				"DiskType":           "LOCAL_NORMAL",
				"BootBackup":         "NONE",
				"DiskBackup":         "NONE",
				"MinimalCpuPlatform": "Intel/Auto",
				"MachineType":        "N",
				"SetId":              "~",
				"HostIp":             "~",
				"CharacterName":      "Uk8sServiceCharacter",
			}
		},
		Owners: []string{"maggie.an@ucloud.cn"},
		Title:  "Auto-Cloudinit-N-授权角色-LOCAL_NORMAL-LOCAL_NORMAL",
		Steps: []*driver.Step{
			testStep4420DescribeImage01,
			testStep4420DescribeImage02,
			testStep4420CreateUHostInstance03,
			testStep4420DescribeUHostInstance04,
			testStep4420DescribeUHostAttributes05,
			testStep4420DisassociateCharacterUHosts06,
			testStep4420DescribeUHostAttributes07,
			testStep4420AssociateCharacterWithUHosts08,
			testStep4420DescribeUHostAttributes09,
			testStep4420StopUHostInstance10,
			testStep4420DescribeUHostInstance11,
			testStep4420TerminateUHostInstance12,
		},
	})
}

var testStep4420DescribeImage01 = &driver.Step{
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
			"Limit":     200,
			"ImageType": "Base",
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeImage(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("ImageSet", step.Must(utils.GetValue(resp, "ImageSet")))
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

var testStep4420DescribeImage02 = &driver.Step{
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
			"ImageId":   step.Must(functions.SearchValue(step.Scenario.GetVar("ImageSet"), "OsName", "Debian 9.8 64位", "ImageId")),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeImage(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("ImageID", step.Must(utils.GetValue(resp, "ImageSet.0.ImageId")))
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

var testStep4420CreateUHostInstance03 = &driver.Step{
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
			"SetId":       step.Scenario.GetVar("SetId"),
			"Region":      step.Scenario.GetVar("Region"),
			"Quantity":    1,
			"Password":    "VXFhNzg5VGVzdCFAIyQ7LA==",
			"Name":        step.Scenario.GetVar("Name"),
			"Memory":      step.Scenario.GetVar("CreateMem"),
			"MachineType": step.Scenario.GetVar("MachineType"),
			"LoginMode":   "Password",
			"ImageId":     step.Scenario.GetVar("ImageID"),
			"HostIp":      step.Scenario.GetVar("HostIp"),
			"Disks": []map[string]interface{}{
				{
					"BackupType": step.Scenario.GetVar("BootBackup"),
					"IsBoot":     "True",
					"Size":       step.Scenario.GetVar("BootSize"),
					"Type":       step.Scenario.GetVar("BootType"),
				},
				{
					"BackupType": step.Scenario.GetVar("DiskBackup"),
					"IsBoot":     "False",
					"Size":       step.Scenario.GetVar("DiskSize"),
					"Type":       step.Scenario.GetVar("DiskType"),
				},
			},
			"ChargeType":    step.Scenario.GetVar("ChargeType"),
			"CharacterName": step.Scenario.GetVar("CharacterName"),
			"CPU":           step.Scenario.GetVar("CreateCPU"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateUHostInstance(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("hostId", step.Must(utils.GetValue(resp, "UHostIds.0")))
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
	Title:         "创建云主机",
	FastFail:      true,
}

var testStep4420DescribeUHostInstance04 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewDescribeUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone": step.Scenario.GetVar("Zone"),
			"UHostIds": []interface{}{
				step.Scenario.GetVar("hostId"),
			},
			"Region": step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeUHostInstance(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.CPU", step.Scenario.GetVar("CreateCPU"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.Memory", step.Scenario.GetVar("CreateMem"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.UHostId", step.Scenario.GetVar("hostId"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.Name", step.Scenario.GetVar("Name"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.BasicImageId", step.Scenario.GetVar("ImageID"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.State", "Running", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.0.IsBoot", "True", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.1.IsBoot", "False", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.0.DiskType", step.Scenario.GetVar("BootType"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.1.DiskType", step.Scenario.GetVar("DiskType"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.0.Size", step.Scenario.GetVar("BootSize"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.1.Size", step.Scenario.GetVar("DiskSize"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.MachineType", step.Scenario.GetVar("MachineType"), "str_eq"),
		}
	},
	StartupDelay:  time.Duration(60) * time.Second,
	MaxRetries:    60,
	RetryInterval: 30 * time.Second,
	Title:         "获取主机信息",
	FastFail:      true,
}

var testStep4420DescribeUHostAttributes05 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("DescribeUHostAttributes")
		err = req.SetPayload(map[string]interface{}{
			"UHostId": step.Scenario.GetVar("hostId"),
			"Region":  step.Scenario.GetVar("Region"),
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
			validation.Builtins.NewValidator("Action", "DescribeUHostAttributesResponse", "str_eq"),
			validation.Builtins.NewValidator("AttributeSet.CharacterNames", step.Scenario.GetVar("CharacterName"), "contains"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "查询主机cloudinit配置信息",
	FastFail:      false,
}

var testStep4420DisassociateCharacterUHosts06 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("DisassociateCharacterUHosts")
		err = req.SetPayload(map[string]interface{}{
			"UHostIds": []interface{}{
				step.Scenario.GetVar("hostId"),
			},
			"Region":        step.Scenario.GetVar("Region"),
			"CharacterName": step.Scenario.GetVar("CharacterName"),
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
			validation.Builtins.NewValidator("Action", "DisassociateCharacterUHostsResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "解除一台或多台虚拟机的实例角色",
	FastFail:      false,
}

var testStep4420DescribeUHostAttributes07 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("DescribeUHostAttributes")
		err = req.SetPayload(map[string]interface{}{
			"UHostId": step.Scenario.GetVar("hostId"),
			"Region":  step.Scenario.GetVar("Region"),
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
			validation.Builtins.NewValidator("Action", "DescribeUHostAttributesResponse", "str_eq"),
			validation.Builtins.NewValidator("AttributeSet.CharacterNames", step.Scenario.GetVar("CharacterName"), "object_not_contains"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "查询主机cloudinit配置信息",
	FastFail:      false,
}

var testStep4420AssociateCharacterWithUHosts08 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("AssociateCharacterWithUHosts")
		err = req.SetPayload(map[string]interface{}{
			"UHostIds": []interface{}{
				step.Scenario.GetVar("hostId"),
			},
			"Region":        step.Scenario.GetVar("Region"),
			"CharacterName": step.Scenario.GetVar("CharacterName"),
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
			validation.Builtins.NewValidator("Action", "AssociateCharacterWithUHostsResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "授权一台或多台虚拟机的实例角色",
	FastFail:      false,
}

var testStep4420DescribeUHostAttributes09 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("DescribeUHostAttributes")
		err = req.SetPayload(map[string]interface{}{
			"UHostId": step.Scenario.GetVar("hostId"),
			"Region":  step.Scenario.GetVar("Region"),
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
			validation.Builtins.NewValidator("Action", "DescribeUHostAttributesResponse", "str_eq"),
			validation.Builtins.NewValidator("AttributeSet.CharacterNames", step.Scenario.GetVar("CharacterName"), "contains"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "查询主机cloudinit配置信息",
	FastFail:      false,
}

var testStep4420StopUHostInstance10 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewStopUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("hostId"),
			"Region":  step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.StopUHostInstance(req)
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
	Title:         "关闭主机",
	FastFail:      false,
}

var testStep4420DescribeUHostInstance11 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewDescribeUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone": step.Scenario.GetVar("Zone"),
			"UHostIds": []interface{}{
				step.Scenario.GetVar("hostId"),
			},
			"Region": step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeUHostInstance(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.State", "Stopped", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(30) * time.Second,
	MaxRetries:    30,
	RetryInterval: 30 * time.Second,
	Title:         "获取主机信息",
	FastFail:      true,
}

var testStep4420TerminateUHostInstance12 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewTerminateUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":       step.Scenario.GetVar("Zone"),
			"UHostId":    step.Scenario.GetVar("hostId"),
			"ReleaseEIP": "true",
			"Region":     step.Scenario.GetVar("Region"),
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
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除云主机",
	FastFail:      false,
}
