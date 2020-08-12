// Code is generated by ucloud-model, DO NOT EDIT IT.

package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/services/udisk"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/driver"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/functions"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/utils"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/validation"
)

func TestScenario526(t *testing.T) {
	spec.ParallelTest(t, &driver.Scenario{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Id: "526",
		Vars: func(scenario *driver.Scenario) map[string]interface{} {
			return map[string]interface{}{
				"DiskType":     "DataDisk",
				"Size":         1,
				"UDataArkMode": "Yes",
				"Name":         "udisk_fz",
				"Region":       "cn-bj2",
				"Zone":         "cn-bj2-02",
			}
		},
		Owners: []string{"chenoa.chen@ucloud.cn"},
		Title:  "UDisk-普通方舟盘_03",
		Steps: []*driver.Step{
			testStep526DescribeUDiskPrice01,
			testStep526CheckUDiskAllowance02,
			testStep526CreateUDisk03,
			testStep526DescribeUDisk04,
			testStep526CreateUDiskSnapshot05,
			testStep526DescribeUDisk06,
			testStep526DescribeSnapshot07,
			testStep526CreateUDiskSnapshot08,
			testStep526DescribeSnapshot09,
			testStep526CreateUDiskSnapshot10,
			testStep526DescribeSnapshot11,
			testStep526CreateUDiskSnapshot12,
			testStep526CloneUDiskSnapshot13,
			testStep526CloneUDiskSnapshot14,
			testStep526RestoreUDisk15,
			testStep526DescribeUDisk16,
			testStep526DeleteSnapshot17,
			testStep526DescribeSnapshot18,
			testStep526DeleteSnapshot19,
			testStep526DescribeSnapshot20,
			testStep526DeleteSnapshot21,
			testStep526DescribeSnapshot22,
			testStep526DeleteUDisk23,
			testStep526DescribeUDisk24,
			testStep526DeleteUDisk25,
			testStep526DescribeUDisk26,
			testStep526DeleteUDisk27,
		},
	})
}

var testStep526DescribeUDiskPrice01 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewDescribeUDiskPriceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":         step.Scenario.GetVar("Zone"),
			"UDataArkMode": step.Scenario.GetVar("UDataArkMode"),
			"Size":         step.Scenario.GetVar("Size"),
			"Region":       step.Scenario.GetVar("Region"),
			"Quantity":     1,
			"DiskType":     step.Scenario.GetVar("DiskType"),
			"ChargeType":   "Month",
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeUDiskPrice(req)
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
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    0,
	RetryInterval: 0 * time.Second,
	Title:         "获取云硬盘价格",
	FastFail:      false,
}

var testStep526CheckUDiskAllowance02 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("CheckUDiskAllowance")
		err = req.SetPayload(map[string]interface{}{
			"Zone":     step.Scenario.GetVar("Zone"),
			"Size":     10,
			"Region":   step.Scenario.GetVar("Region"),
			"DiskType": step.Scenario.GetVar("DiskType"),
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
			validation.Builtins.NewValidator("Action", "CheckUDiskAllowanceResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "检查UDisk资源余量",
	FastFail:      false,
}

var testStep526CreateUDisk03 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewCreateUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":         step.Scenario.GetVar("Zone"),
			"UDataArkMode": step.Scenario.GetVar("UDataArkMode"),
			"Size":         step.Scenario.GetVar("Size"),
			"Region":       step.Scenario.GetVar("Region"),
			"Quantity":     0,
			"Name":         step.Must(functions.Concat("auto_", step.Scenario.GetVar("Name"))),
			"DiskType":     step.Scenario.GetVar("DiskType"),
			"ChargeType":   "Month",
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateUDisk(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("udisk_id", step.Must(utils.GetValue(resp, "UDiskId.0")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(1) * time.Second,
	MaxRetries:    0,
	RetryInterval: 0 * time.Second,
	Title:         "创建云硬盘",
	FastFail:      false,
}

var testStep526DescribeUDisk04 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewDescribeUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("udisk_id"),
			"Region":  step.Scenario.GetVar("Region"),
			"Offset":  0,
			"Limit":   100,
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
			validation.Builtins.NewValidator("DataSet.0.Status", "Available", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    0,
	RetryInterval: 0 * time.Second,
	Title:         "获取云硬盘列表",
	FastFail:      false,
}

var testStep526CreateUDiskSnapshot05 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewCreateUDiskSnapshotRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("udisk_id"),
			"Region":  step.Scenario.GetVar("Region"),
			"Name":    "snapshot_01",
			"Comment": "comment_01",
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateUDiskSnapshot(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("snapshot_id", step.Must(utils.GetValue(resp, "SnapshotId.0")))
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
	Title:         "创建快照",
	FastFail:      false,
}

var testStep526DescribeUDisk06 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewDescribeUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("udisk_id"),
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
			validation.Builtins.NewValidator("DataSet.0.SnapshotLimit", 3, "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Status", "Available", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(60) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取云硬盘列表",
	FastFail:      false,
}

var testStep526DescribeSnapshot07 = &driver.Step{
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
				step.Scenario.GetVar("snapshot_id"),
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
			validation.Builtins.NewValidator("UHostSnapshotSet.0.State", "Normal", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(60) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "描述快照",
	FastFail:      false,
}

var testStep526CreateUDiskSnapshot08 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewCreateUDiskSnapshotRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("udisk_id"),
			"Region":  step.Scenario.GetVar("Region"),
			"Name":    "snapshot_02",
			"Comment": "comment_01",
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateUDiskSnapshot(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("snapshot_id1", step.Must(utils.GetValue(resp, "SnapshotId.0")))
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
	Title:         "创建快照",
	FastFail:      false,
}

var testStep526DescribeSnapshot09 = &driver.Step{
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
				step.Scenario.GetVar("snapshot_id1"),
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
			validation.Builtins.NewValidator("UHostSnapshotSet.0.State", "Normal", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(60) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "描述快照",
	FastFail:      false,
}

var testStep526CreateUDiskSnapshot10 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewCreateUDiskSnapshotRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("udisk_id"),
			"Region":  step.Scenario.GetVar("Region"),
			"Name":    "snapshot_03",
			"Comment": "comment_01",
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateUDiskSnapshot(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("snapshot_id2", step.Must(utils.GetValue(resp, "SnapshotId.0")))
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
	Title:         "创建快照",
	FastFail:      false,
}

var testStep526DescribeSnapshot11 = &driver.Step{
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
				step.Scenario.GetVar("snapshot_id2"),
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
			validation.Builtins.NewValidator("UHostSnapshotSet.0.State", "Normal", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(60) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "描述快照",
	FastFail:      false,
}

var testStep526CreateUDiskSnapshot12 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewCreateUDiskSnapshotRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("udisk_id"),
			"Region":  step.Scenario.GetVar("Region"),
			"Name":    "snapshot_04",
			"Comment": "comment_01",
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateUDiskSnapshot(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 16999, "gt"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "创建快照",
	FastFail:      false,
}

var testStep526CloneUDiskSnapshot13 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewCloneUDiskSnapshotRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":         step.Scenario.GetVar("Zone"),
			"UDataArkMode": "Yes",
			"SourceId":     step.Scenario.GetVar("snapshot_id"),
			"Size":         step.Scenario.GetVar("Size"),
			"Region":       step.Scenario.GetVar("Region"),
			"Quantity":     0,
			"Name":         "clone_from_kz_fz",
			"ChargeType":   "Month",
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CloneUDiskSnapshot(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("udisk_id_fromkz_fz", step.Must(utils.GetValue(resp, "UDiskId.0")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(120) * time.Second,
	MaxRetries:    6,
	RetryInterval: 5 * time.Second,
	Title:         "克隆快照",
	FastFail:      false,
}

var testStep526CloneUDiskSnapshot14 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewCloneUDiskSnapshotRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":         step.Scenario.GetVar("Zone"),
			"UDataArkMode": "No",
			"SourceId":     step.Scenario.GetVar("snapshot_id"),
			"Size":         step.Scenario.GetVar("Size"),
			"Region":       step.Scenario.GetVar("Region"),
			"Quantity":     0,
			"Name":         "clone_from_kz_nofz",
			"ChargeType":   "Month",
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CloneUDiskSnapshot(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("udisk_id_fromkz_nofz", step.Must(utils.GetValue(resp, "UDiskId.0")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(130) * time.Second,
	MaxRetries:    6,
	RetryInterval: 5 * time.Second,
	Title:         "克隆快照",
	FastFail:      false,
}

var testStep526RestoreUDisk15 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewRestoreUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":       step.Scenario.GetVar("Zone"),
			"UDiskId":    step.Scenario.GetVar("udisk_id"),
			"SnapshotId": step.Scenario.GetVar("snapshot_id"),
			"Region":     step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.RestoreUDisk(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "RestoreUDiskResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(120) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "从备份恢复数据至UDisk",
	FastFail:      false,
}

var testStep526DescribeUDisk16 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewDescribeUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("udisk_id"),
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
			validation.Builtins.NewValidator("DataSet.0.Name", step.Must(functions.Concat("auto_", step.Scenario.GetVar("Name"))), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Status", "Available", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(120) * time.Second,
	MaxRetries:    30,
	RetryInterval: 10 * time.Second,
	Title:         "获取云硬盘列表",
	FastFail:      false,
}

var testStep526DeleteSnapshot17 = &driver.Step{
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
			"SnapshotId": step.Scenario.GetVar("snapshot_id"),
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
			validation.Builtins.NewValidator("SnapshotId", step.Scenario.GetVar("snapshot_id"), "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    0,
	RetryInterval: 0 * time.Second,
	Title:         "删除快照",
	FastFail:      false,
}

var testStep526DescribeSnapshot18 = &driver.Step{
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
				step.Scenario.GetVar("snapshot_id"),
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
			validation.Builtins.NewValidator("TotalCount", 0, "str_eq"),
			validation.Builtins.NewValidator("PerDiskQuota", 3, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(1) * time.Second,
	MaxRetries:    0,
	RetryInterval: 0 * time.Second,
	Title:         "描述快照",
	FastFail:      false,
}

var testStep526DeleteSnapshot19 = &driver.Step{
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
			"SnapshotId": step.Scenario.GetVar("snapshot_id1"),
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
			validation.Builtins.NewValidator("SnapshotId", step.Scenario.GetVar("snapshot_id1"), "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除快照",
	FastFail:      false,
}

var testStep526DescribeSnapshot20 = &driver.Step{
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
				step.Scenario.GetVar("snapshot_id1"),
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
			validation.Builtins.NewValidator("TotalCount", 0, "str_eq"),
			validation.Builtins.NewValidator("PerDiskQuota", 3, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "描述快照",
	FastFail:      false,
}

var testStep526DeleteSnapshot21 = &driver.Step{
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
			"SnapshotId": step.Scenario.GetVar("snapshot_id2"),
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
			validation.Builtins.NewValidator("SnapshotId", step.Scenario.GetVar("snapshot_id2"), "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除快照",
	FastFail:      false,
}

var testStep526DescribeSnapshot22 = &driver.Step{
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
				step.Scenario.GetVar("snapshot_id2"),
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
			validation.Builtins.NewValidator("TotalCount", 0, "str_eq"),
			validation.Builtins.NewValidator("PerDiskQuota", 3, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "描述快照",
	FastFail:      false,
}

var testStep526DeleteUDisk23 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewDeleteUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("udisk_id"),
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
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    0,
	RetryInterval: 1 * time.Second,
	Title:         "删除云硬盘",
	FastFail:      false,
}

var testStep526DescribeUDisk24 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewDescribeUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("udisk_id_fromkz_nofz"),
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
			validation.Builtins.NewValidator("DataSet.0.Name", "clone_from_kz_nofz", "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Status", "Available", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(60) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取云硬盘列表",
	FastFail:      true,
}

var testStep526DeleteUDisk25 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewDeleteUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("udisk_id_fromkz_nofz"),
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
		}
	},
	StartupDelay:  time.Duration(5) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除云硬盘",
	FastFail:      false,
}

var testStep526DescribeUDisk26 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewDescribeUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("udisk_id_fromkz_fz"),
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
			validation.Builtins.NewValidator("DataSet.0.Name", "clone_from_kz_fz", "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Status", "Available", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(200) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取云硬盘列表",
	FastFail:      true,
}

var testStep526DeleteUDisk27 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewDeleteUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("udisk_id_fromkz_fz"),
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
		}
	},
	StartupDelay:  time.Duration(5) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除云硬盘",
	FastFail:      false,
}
