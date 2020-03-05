// Code is generated by ucloud-model, DO NOT EDIT IT.

package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/driver"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/utils"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/validation"

	"github.com/ucloud/ucloud-sdk-go/ucloud"

	"github.com/ucloud/ucloud-sdk-go/services/udisk"

	"github.com/ucloud/ucloud-sdk-go/services/uhost"
)

func TestScenario1935(t *testing.T) {
	spec.ParallelTest(t, &driver.Scenario{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Id: "1935",
		Vars: func(scenario *driver.Scenario) map[string]interface{} {
			return map[string]interface{}{
				"Password":                      "Z3VhbmxpeXVhbm1pbWExMjMhQCM=",
				"SnapshotSysName":               "snapshot-ARK-SYS-02",
				"SnapshotSysDesc":               "snapshot-ARK-SYS-02-desc",
				"SnapDiskType":                  "LocalBoot",
				"SnapshotDataNameModify":        "snapshot-ARK-DATA-02-modify",
				"SnapshotDataDescModify":        "snapshot-ARK-DATA-02-desc-Modify",
				"UhostName":                     "ARK-local-4",
				"SnapshotDataName":              "snapshot-ARK-DATA-02",
				"SnapshotDataDesc":              "snapshot-ARK-DATA-02-desc",
				"CreateFromTimeMachinePassword": "Z3VhbmxpeXVhbm1pbWExMjMhQCM=",
				"ImageID":                       "#{u_get_image_resource($Region,$Zone)}",
				"BootSize":                      20,
				"BootType":                      "LOCAL_NORMAL",
				"BootBackup":                    "DATAARK",
				"UHostType":                     "N2",
				"DiskSize":                      20,
				"DiskType":                      "LOCAL_NORMAL",
				"DiskBackup":                    "DATAARK",
				"Region":                        "cn-bj2",
				"Zone":                          "cn-bj2-02",
			}
		},
		Owners: []string{"maggie.an@ucloud.cn"},
		Title:  "主机功能之快照-方舟-N2-LOCAL_NORMAL-LOCAL_NORMAL-4",
		Steps: []*driver.Step{
			testStep1935DescribeImage01,
			testStep1935CreateUHostInstance02,
			testStep1935DescribeUHostInstance03,
			testStep1935DescribeUhostTmMeta04,
			testStep1935CreateUDisk05,
			testStep1935CreateUDisk06,
			testStep1935DescribeUDisk07,
			testStep1935DescribeUDisk08,
			testStep1935AttachUDisk09,
			testStep1935AttachUDisk10,
			testStep1935CreateUDiskSnapshot11,
			testStep1935DescribeSnapshot12,
			testStep1935StopUHostInstance13,
			testStep1935DescribeUHostInstance14,
			testStep1935RestoreUHostDisk15,
			testStep1935DescribeUHostInstance16,
			testStep1935PoweroffUHostInstance17,
			testStep1935DeleteSnapshot18,
			testStep1935DescribeUHostInstance19,
			testStep1935TerminateUHostInstance20,
			testStep1935DescribeUDisk21,
			testStep1935DescribeUDisk22,
			testStep1935DeleteUDisk23,
			testStep1935DeleteUDisk24,
		},
	})
}

var testStep1935DescribeImage01 = &driver.Step{
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
	FastFail:      true,
}

var testStep1935CreateUHostInstance02 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewCreateUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":      step.Scenario.GetVar("Zone"),
			"UHostType": step.Scenario.GetVar("UHostType"),
			"Region":    step.Scenario.GetVar("Region"),
			"Password":  "VXFhNzg5VGVzdCFAIyQ7LA==",
			"Name":      step.Scenario.GetVar("UhostName"),
			"Memory":    1024,
			"LoginMode": "Password",
			"ImageId":   step.Scenario.GetVar("ImageID"),
			"Disks": []map[string]interface{}{
				map[string]interface{}{
					"BackupType": step.Scenario.GetVar("BootBackup"),
					"IsBoot":     "True",
					"Size":       step.Scenario.GetVar("BootSize"),
					"Type":       step.Scenario.GetVar("BootType"),
				},
				map[string]interface{}{
					"BackupType": step.Scenario.GetVar("DiskBackup"),
					"IsBoot":     "False",
					"Size":       step.Scenario.GetVar("DiskSize"),
					"Type":       step.Scenario.GetVar("DiskType"),
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

var testStep1935DescribeUHostInstance03 = &driver.Step{
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
			validation.Builtins.NewValidator("UHostSet.0.State", "Running", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.TimemachineFeature", "yes", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.0.BackupType", "DATAARK", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(100) * time.Second,
	MaxRetries:    60,
	RetryInterval: 60 * time.Second,
	Title:         "获取主机信息",
	FastFail:      true,
}

var testStep1935DescribeUhostTmMeta04 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("DescribeUhostTmMeta")
		err = req.SetPayload(map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UhostId": step.Scenario.GetVar("hostId"),
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
			validation.Builtins.NewValidator("Action", "DescribeUhostTmMetaResponse", "str_eq"),
			validation.Builtins.NewValidator("UtmStatus", "normal", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    100,
	RetryInterval: 60 * time.Second,
	Title:         "查询主机方舟状态",
	FastFail:      true,
}

var testStep1935CreateUDisk05 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewCreateUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":         step.Scenario.GetVar("Zone"),
			"UDataArkMode": "Yes",
			"Size":         1,
			"Region":       step.Scenario.GetVar("Region"),
			"Name":         "test-for-ark",
			"DiskType":     "DataDisk",
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateUDisk(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("UDiskId", step.Must(utils.GetValue(resp, "UDiskId.0")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "CreateUDiskResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "创建云硬盘",
	FastFail:      true,
}

var testStep1935CreateUDisk06 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewCreateUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":     step.Scenario.GetVar("Zone"),
			"Size":     1,
			"Region":   step.Scenario.GetVar("Region"),
			"Name":     "ssd",
			"DiskType": "SSDDataDisk",
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateUDisk(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("UDiskId2", step.Must(utils.GetValue(resp, "UDiskId.0")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "CreateUDiskResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "创建云硬盘",
	FastFail:      true,
}

var testStep1935DescribeUDisk07 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewDescribeUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("UDiskId"),
			"Region":  step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeUDisk(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "DescribeUDiskResponse", "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Status", "Available", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    30,
	RetryInterval: 10 * time.Second,
	Title:         "获取云硬盘列表",
	FastFail:      true,
}

var testStep1935DescribeUDisk08 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewDescribeUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("UDiskId2"),
			"Region":  step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeUDisk(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "DescribeUDiskResponse", "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Status", "Available", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取云硬盘列表",
	FastFail:      true,
}

var testStep1935AttachUDisk09 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewAttachUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("hostId"),
			"UDiskId": step.Scenario.GetVar("UDiskId"),
			"Region":  step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.AttachUDisk(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "AttachUDiskResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "挂载云硬盘",
	FastFail:      true,
}

var testStep1935AttachUDisk10 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewAttachUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("hostId"),
			"UDiskId": step.Scenario.GetVar("UDiskId2"),
			"Region":  step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.AttachUDisk(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "AttachUDiskResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "挂载云硬盘",
	FastFail:      true,
}

var testStep1935CreateUDiskSnapshot11 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewCreateUDiskSnapshotRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("UDiskId"),
			"Region":  step.Scenario.GetVar("Region"),
			"Name":    "test-for-ark",
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateUDiskSnapshot(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("SnapshotId", step.Must(utils.GetValue(resp, "SnapshotId.0")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "CreateUDiskSnapshotResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "创建快照",
	FastFail:      true,
}

var testStep1935DescribeSnapshot12 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("DescribeSnapshot")
		err = req.SetPayload(map[string]interface{}{
			"Zone": step.Scenario.GetVar("Zone"),
			"SnapshotIds": []interface{}{
				step.Scenario.GetVar("SnapshotId"),
			},
			"Region": step.Scenario.GetVar("Region"),
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
			validation.Builtins.NewValidator("Action", "DescribeSnapshotResponse", "str_eq"),
			validation.Builtins.NewValidator("UHostSnapshotSet.0.State", "Normal", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    30,
	RetryInterval: 10 * time.Second,
	Title:         "描述快照",
	FastFail:      true,
}

var testStep1935StopUHostInstance13 = &driver.Step{
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
			validation.Builtins.NewValidator("Action", "StopUHostInstanceResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "关闭主机",
	FastFail:      true,
}

var testStep1935DescribeUHostInstance14 = &driver.Step{
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
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    30,
	RetryInterval: 10 * time.Second,
	Title:         "获取主机信息",
	FastFail:      true,
}

var testStep1935RestoreUHostDisk15 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("RestoreUHostDisk")
		err = req.SetPayload(map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("hostId"),
			"SnapshotIds": []interface{}{
				step.Scenario.GetVar("SnapshotId"),
			},
			"Region": step.Scenario.GetVar("Region"),
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
			validation.Builtins.NewValidator("Action", "RestoreUHostDiskResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(200) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "恢复主机磁盘",
	FastFail:      true,
}

var testStep1935DescribeUHostInstance16 = &driver.Step{
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
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    60,
	RetryInterval: 30 * time.Second,
	Title:         "获取主机信息",
	FastFail:      true,
}

var testStep1935PoweroffUHostInstance17 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewPoweroffUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("hostId"),
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
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "模拟主机掉电",
	FastFail:      true,
}

var testStep1935DeleteSnapshot18 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("DeleteSnapshot")
		err = req.SetPayload(map[string]interface{}{
			"Zone":       step.Scenario.GetVar("Zone"),
			"SnapshotId": step.Scenario.GetVar("SnapshotId"),
			"Region":     step.Scenario.GetVar("Region"),
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
			validation.Builtins.NewValidator("Action", "DeleteSnapshotResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除快照",
	FastFail:      true,
}

var testStep1935DescribeUHostInstance19 = &driver.Step{
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
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    30,
	RetryInterval: 10 * time.Second,
	Title:         "获取主机信息",
	FastFail:      true,
}

var testStep1935TerminateUHostInstance20 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewTerminateUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("hostId"),
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
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(5) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除云主机",
	FastFail:      true,
}

var testStep1935DescribeUDisk21 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewDescribeUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("UDiskId"),
			"Region":  step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeUDisk(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "DescribeUDiskResponse", "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Status", "Available", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    30,
	RetryInterval: 10 * time.Second,
	Title:         "获取云硬盘列表",
	FastFail:      true,
}

var testStep1935DescribeUDisk22 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewDescribeUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("UDiskId2"),
			"Region":  step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeUDisk(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "DescribeUDiskResponse", "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Status", "Available", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    30,
	RetryInterval: 10 * time.Second,
	Title:         "获取云硬盘列表",
	FastFail:      true,
}

var testStep1935DeleteUDisk23 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewDeleteUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("UDiskId"),
			"Region":  step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DeleteUDisk(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "DeleteUDiskResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(5) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除云硬盘",
	FastFail:      true,
}

var testStep1935DeleteUDisk24 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewDeleteUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("UDiskId2"),
			"Region":  step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DeleteUDisk(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "DeleteUDiskResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除云硬盘",
	FastFail:      true,
}
