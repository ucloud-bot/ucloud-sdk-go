// Code is generated by ucloud-model, DO NOT EDIT IT.



package uhost











/*
UHostImageSet - DescribeImage
*/
type UHostImageSet struct {
	
	// 创建时间，格式为Unix时间戳
	CreateTime int 
	
	// 特殊状态标识， 目前包含NetEnhnced（网络增强1.0）, NetEnhanced_Ultra]（网络增强2.0）,HotPlug(热升级),CloudInit
	Features []string 
	
	// 行业镜像类型（仅行业镜像将返回这个值）
	FuncType string 
	
	// 镜像描述
	ImageDescription string 
	
	// 镜像ID
	ImageId string 
	
	// 镜像名称
	ImageName string 
	
	// 镜像大小
	ImageSize int 
	
	// 镜像类型 标准镜像：Base， 行业镜像：Business，自定义镜像：Custom
	ImageType string 
	
	// 集成软件名称（仅行业镜像将返回这个值）
	IntegratedSoftware string 
	
	// 介绍链接（仅行业镜像将返回这个值）
	Links string 
	
	// 默认值为空'''。当CentOS 7.3/7.4/7.5等镜像会标记为“Broadwell”
	MinimalCPU string 
	
	// 操作系统名称
	OsName string 
	
	// 操作系统类型：Liunx，Windows
	OsType string 
	
	// 镜像状态， 可用：Available，制作中：Making， 不可用：Unavailable
	State string 
	
	// 供应商（仅行业镜像将返回这个值）
	Vendor string 
	
	// 可用区，参见 [可用区列表](../summary/regionlist.html)
	Zone string 
	
}



/*
SpreadInfo - 每个可用区中硬件隔离组信息
*/
type SpreadInfo struct {
	
	// 可用区中硬件隔离组中云主机的数量，不超过7。
	UHostCount int 
	
	// 可用区信息
	Zone string 
	
}



/*
IsolationGroup - 硬件隔离组信息
*/
type IsolationGroup struct {
	
	// 硬件隔离组id
	GroupId string 
	
	// 硬件隔离组名称
	GroupName string 
	
	// 备注
	Remark string 
	
	// 每个可用区中的机器数量。参见数据结构SpreadInfo。
	SpreadInfoSet []SpreadInfo 
	
}



/*
UHostIPSet - DescribeUHostInstance
*/
type UHostIPSet struct {
	
	// IP对应的带宽, 单位: Mb  (内网IP不显示带宽信息)
	Bandwidth int 
	
	// 【暂未支持】是否为默认网卡。true: 是默认网卡；其他值：不是。
	Default string 
	
	// IP地址
	IP string 
	
	// 外网IP资源ID 。(内网IP无对应的资源ID)
	IPId string 
	
	// IPv4/IPv6；
	IPMode string 
	
	// 当前网卡的Mac。
	Mac string 
	
	// IP地址对应的子网 ID。（北京一不支持，字段返回为空）
	SubnetId string 
	
	// 国际: Internation，BGP: Bgp，内网: Private
	Type string 
	
	// IP地址对应的VPC ID。（北京一不支持，字段返回为空）
	VPCId string 
	
	// 当前EIP的权重。权重最大的为当前的出口IP。
	Weight int 
	
}



/*
UHostDiskSet - DescribeUHostInstance
*/
type UHostDiskSet struct {
	
	// 备份方案。若开通了数据方舟，则为DataArk
	BackupType string 
	
	// 磁盘ID
	DiskId string 
	
	// 磁盘类型。请参考[[api:uhost-api:disk_type|磁盘类型]]。
	DiskType string 
	
	// 磁盘盘符
	Drive string 
	
	// "true": 加密盘 "false"：非加密盘
	Encrypted string 
	
	// 是否是系统盘。枚举值：\\ > True，是系统盘 \\ > False，是数据盘（默认）。Disks数组中有且只能有一块盘是系统盘。
	IsBoot string 
	
	// UDisk名字（仅当磁盘是UDisk时返回）
	Name string 
	
	// 磁盘大小，单位: GB
	Size int 
	
	// 【建议不再使用】磁盘类型。系统盘: Boot，数据盘: Data,网络盘：Udisk
	Type string 
	
}



/*
UHostInstanceSet - DescribeUHostInstance
*/
type UHostInstanceSet struct {
	
	// 是否自动续费，自动续费：“Yes”，不自动续费：“No”
	AutoRenew string 
	
	// 基础镜像ID（指当前自定义镜像的来源镜像）
	BasicImageId string 
	
	// 基础镜像名称（指当前自定义镜像的来源镜像）
	BasicImageName string 
	
	// 系统盘状态 Normal表示初始化完成；Initializing表示在初始化。仍在初始化的系统盘无法制作镜像。
	BootDiskState string 
	
	// 虚拟CPU核数，单位: 个
	CPU int 
	
	// 计费模式，枚举值为： Year，按年付费； Month，按月付费； Dynamic，按需付费（需开启权限）；
	ChargeType string 
	
	// true，支持cloutinit方式初始化；false,不支持
	CloudInitFeature bool 
	
	// 创建时间，格式为Unix时间戳
	CreateTime int 
	
	// 
	DeleteTime int `deprecated:"true"`
	
	// 磁盘信息见 UHostDiskSet
	DiskSet []UHostDiskSet 
	
	// 到期时间，格式为Unix时间戳
	ExpireTime int 
	
	// GPU个数
	GPU int 
	
	// 【建议不再使用】主机系列：N2，表示系列2；N1，表示系列1
	HostType string 
	
	// true: 开启热升级； false，未开启热升级
	HotplugFeature bool 
	
	// 详细信息见 UHostIPSet
	IPSet []UHostIPSet 
	
	// 
	IPs []string `deprecated:"true"`
	
	// true:有ipv6特性；false，没有ipv6特性
	IPv6Feature bool 
	
	// 【建议不再使用】主机的系统盘ID。
	ImageId string 
	
	// 隔离组id，不在隔离组则返回""
	IsolationGroup string 
	
	// 主机的生命周期类型。目前仅支持Normal：普通；
	LifeCycle string 
	
	// 云主机机型（新）。参考[[api:uhost-api:uhost_type#主机概念20版本|云主机机型说明]]。
	MachineType string 
	
	// 内存大小，单位: MB
	Memory int 
	
	// UHost实例名称
	Name string 
	
	// 网络增强。Normal: 无；Super： 网络增强1.0； Ultra: 网络增强2.0
	NetCapability string 
	
	// 【建议不再使用】网络状态。 连接：Connected， 断开：NotConnected
	NetworkState string 
	
	// 创建主机的最初来源镜像的操作系统名称（若直接通过基础镜像创建，此处返回和BasicImageName一致）
	OsName string 
	
	// 操作系统类别。返回"Linux"或者"Windows"
	OsType string 
	
	// 备注
	Remark string 
	
	// 实例状态，枚举值：\\ >初始化: Initializing; \\ >启动中: Starting; \\> 运行中: Running; \\> 关机中: Stopping; \\ >关机: Stopped \\ >安装失败: Install Fail; \\ >重启中: Rebooting
	State string 
	
	// 【建议不再使用】主机磁盘类型。 枚举值为：\\ > LocalDisk，本地磁盘; \\ > UDisk 云盘。\\只要有一块磁盘为本地盘，即返回LocalDisk。
	StorageType string 
	
	// 【建议不再使用】仅北京A的云主机会返回此字段。基础网络模式：Default；子网模式：Private
	SubnetType string 
	
	// 业务组名称
	Tag string 
	
	// 【建议不再使用】数据方舟模式。枚举值：\\ > Yes: 开启方舟； \\ > no，未开启方舟
	TimemachineFeature string 
	
	// 总的数据盘存储空间。
	TotalDiskSpace int 
	
	// UHost实例ID
	UHostId string 
	
	// 【建议不再使用】云主机机型（旧）。参考[[api:uhost-api:uhost_type|云主机机型说明]]。
	UHostType string 
	
	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone string 
	
}



/*
UHostTagSet - DescribeUHostTags
*/
type UHostTagSet struct {
	
	// 业务组名称
	Tag string 
	
	// 该业务组中包含的主机个数
	TotalCount int 
	
	// 可用区
	Zone string 
	
}



/*
UHostPriceSet - 主机价格
*/
type UHostPriceSet struct {
	
	// 计费类型。Year，Month，Dynamic
	ChargeType string 
	
	// 产品列表价。
	ListPrice float64 
	
	// 限时优惠的折前原价（即列表价乘以商务折扣后的单价）。
	OriginalPrice float64 
	
	// 价格，单位: 元，保留小数点后两位有效数字
	Price float64 
	
}


