// Code is generated by ucloud-model, DO NOT EDIT IT.

package udisk

/*
UDiskDataSet - DescribeUDisk

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn/api for detail.
*/
type UDiskDataSet struct {

	// 是否支持开启方舟，1支持 ，0不支持
	ArkSwitchEnable int

	// Year,Month,Dynamic,Trial,Postpay
	ChargeType string

	// 是否支持克隆，1支持 ，0不支持
	CloneEnable int

	// 该盘的cmk id
	CmkId string

	// cmk id 别名
	CmkIdAlias string

	// 该盘cmk的状态, Enabled(正常)，Disabled(失效)，Deleted(删除)，NoCmkId(非加密盘)
	CmkIdStatus string

	// 创建时间
	CreateTime int

	// 该盘的密文密钥
	DataKey string

	// 挂载的设备名称
	DeviceName string

	// 请求中的ProtocolVersion字段为1时，需结合IsBoot确定具体磁盘类型:普通数据盘：DiskType:"CLOUD_NORMAL",IsBoot:"False"； 普通系统盘：DiskType:"CLOUD_NORMAL",IsBoot:"True"；SSD数据盘：DiskType:"CLOUD_SSD",IsBoot:"False"；SSD系统盘：DiskType:"CLOUD_SSD",IsBoot:"True"；RSSD数据盘：DiskType:"CLOUD_RSSD",IsBoot:"False"。请求中的ProtocolVersion字段为0或没有该字段时，云硬盘类型参照如下:普通数据盘：DataDisk；普通系统盘：SystemDisk；SSD数据盘：SSDDataDisk；SSD系统盘：SSDSystemDisk；RSSD数据盘：RSSDDataDisk。
	DiskType string

	// 过期时间
	ExpiredTime int

	// 是否是系统盘，是："True", 否："False"
	IsBoot string

	// 资源是否过期，过期:"Yes", 未过期:"No"
	IsExpire string

	// 实例名称
	Name string

	// 容量单位GB
	Size int

	// 是否支持快照，1支持 ，0不支持
	SnapEnable int

	// 该盘快照个数
	SnapshotCount int

	// 该盘快照上限
	SnapshotLimit int

	// 状态:Available(可用),Attaching(挂载中), InUse(已挂载), Detaching(卸载中), Initializating(分配中), Failed(创建失败),Cloning(克隆中),Restoring(恢复中),RestoreFailed(恢复失败),
	Status string

	// 业务组名称
	Tag string

	// 是否开启数据方舟，开启:"Yes", 不支持:"No"
	UDataArkMode string

	// UDisk实例Id
	UDiskId string

	// 挂载的UHost的IP
	UHostIP string

	// 挂载的UHost的Id
	UHostId string

	// 挂载的UHost的Name
	UHostName string

	// 是否是加密盘，是:"Yes", 否:"No"
	UKmsMode string

	// 是否支持数据方舟，支持:"2.0", 不支持:"1.0"
	Version string

	// 可用区
	Zone string
}
