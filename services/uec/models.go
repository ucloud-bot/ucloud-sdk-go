// Code is generated by ucloud-model, DO NOT EDIT IT.

package uec

/*
NodeList - 虚拟机资源id列表
*/
type NodeList struct {

	// 虚拟机资源id
	NodeId string
}

/*
RuleInfo - 防火墙规则
*/
type RuleInfo struct {

	// ACCEPT（接受）和DROP（拒绝）
	Action string

	// 端口，范围用"-"符号分隔，如：1-65535
	Port string

	// 优先级：HIGH（高），MEDIUM（中），LOW（低）
	Priority string

	// 协议，可选值：TCP，UDP，ICMP
	ProtocolType string

	// 备注
	Remark string

	// 源ip
	SrcIp string
}

/*
FirewallInfo - 防火墙信息
*/
type FirewallInfo struct {

	// 创建时间
	CreateTime int

	// 防火墙Id
	FirewallId string

	// 防火墙名称
	Name string

	// 描述
	Remark string

	// 防火墙绑定资源数量
	ResourceCount int

	// 防火墙规则组，详情参见RuleInfo
	Rule []RuleInfo

	// 防火墙组类型，枚举值为： "user defined", 用户自定义防火墙； "recommend web", 默认Web防火墙； "recommend non web", 默认非Web防火墙
	Type string
}

/*
ResourceInfo - 绑定防火墙的资源信息
*/
type ResourceInfo struct {

	// 节点名称
	Name string

	// 节点公网Ip列表
	PublicIpList []string

	// 节点备注
	Remark string

	// 资源Id
	ResourceId string

	// 节点状态，1部署中，2待启动，3启动中，4运行中，5正在停止，6已停止，7正在更新，8正在重启，9正在删除， 10已经删除,11异常
	State int
}

/*
EnvList - 容器环境变量列表
*/
type EnvList struct {

	// 环境变量key值
	Key string

	// 环境变量Value值
	Value string
}

/*
CfgDictList - 容器配置字典列表
*/
type CfgDictList struct {

	// 挂载路径
	MountPath string

	// 名称
	Name string

	// 资源id
	ResourceId string
}

/*
DockerInfo - 容器信息
*/
type DockerInfo struct {

	// 参数
	Args string

	// 容器配置字典（详情参考CfgDictList）
	CfgDictList []CfgDictList

	// 命令
	Command string

	// CPU核数（/核）精度0.1核
	CpuCores float64

	// 环境变量（详情参考EnvList）
	EnvList []EnvList

	// 镜像名称
	ImageName string

	// 内存大小（Gi）
	MemSize float64

	// 容器名称
	Name string

	// 容器状态，0：初始化；1：拉取镜像；2：拉取镜像失败；3：启动中；4：运行中；5：正在停止；6：已停止；7：已删除；8：镜像拉取成功；9：启动失败；99：异常
	State int

	// 工作目录
	WorkDir string
}

/*
StorVolumeInfo - 容器组存储卷信息
*/
type StorVolumeInfo struct {

	// 容量（单位GB）
	DiskSize int

	// 挂载点
	MountPoint string

	// 名称
	Name string

	// 资源id
	ResourceId string
}

/*
ImageList - 容器组镜像密钥列表
*/
type ImageList struct {

	// 镜像密钥
	ImageKey string

	// 仓库地址
	StoreAddr string

	// 用户名称
	UserName string
}

/*
IpList - 容器组外网ip列表
*/
type IpList struct {

	// 外网ip
	Ip string

	// 运营商
	Isp string
}

/*
HolderList - 容器组信息
*/
type HolderList struct {

	// 城市名称
	City string

	// 创建时间
	CreateTime int

	// 容器数量
	DockerCount int

	// 容器信息（详情参考DockerInfo）
	DockerInfo []DockerInfo

	// 过期时间
	ExpireTime int

	// 外网防火墙id
	FirewallId string

	// 容器组名称
	HolderName string

	// 机房id
	IdcId string

	// 容器组镜像密钥列表（详情参考ImageList）
	ImageList []ImageList

	// 容器组内网ip
	InnerIp string

	// 容器组外网ip集合（详情参考IpList）
	IpList []IpList

	// 外网绑定的带宽
	NetLimit int

	// 机房名称
	OcName string

	// 机器类型（normal通用型，hf高性能型）
	ProductType string

	// 省份名称
	Province string

	// 容器组资源id
	ResourceId string

	// 0：总是；1：失败是；2：永不
	RestartStrategy int

	// 容器组运行状态0：初始化；1：拉取镜像；2：启动中；3：运行中；4：错误；5：正在重启；6：正在删除；7：已经删除；8：容器运行错误；9：启动失败；99：异常
	State int

	// 存储卷数量
	StorVolumeCount int

	// 存储卷信息（详情参考StorVolumeInfo）
	StorVolumeInfo []StorVolumeInfo

	// 容器组子网id
	SubnetId string

	// 线路类型（运营商类型： 0-其它, 1-一线城市单线,2-二线城市单线, 3-全国教育网, 4-全国三通）
	Type int
}

/*
IdcInfo - 机房信息
*/
type IdcInfo struct {

	// 城市
	City string

	// 机房ID
	IdcId string

	// 运营商
	Isp string

	// 机房可创建节点最大数量
	MaxNodeCnt int

	// 机房名称
	Name string

	// 省份
	Province string

	// 运营商类型：0-其它, 1-一线城市单线,2-二线城市单线, 3-全国教育网, 4-全国三通
	Type int
}

/*
SubnetInfo - 子网信息
*/
type SubnetInfo struct {

	// 可用ip数
	AvailableIPCnt int

	// 子网cidr
	CIDR string

	// 备注
	Comment string

	// 创建时间
	CreateTime int

	// 机房ID
	IdcId string

	// 子网ID
	SubnetId string

	// 子网名称
	SubnetName string

	// 总ip数
	TotalIpCnt int
}

/*
NodeIpList - 虚拟机外网ip列表
*/
type NodeIpList struct {

	// 外网ip
	Ip string

	// 运营商
	Isp string

	// 运营商名称
	IspName string
}

/*
NodeInfo - 节点信息
*/
type NodeInfo struct {

	// 付费类型：1按时, 2按月,3按年
	ChargeType int

	// 城市
	City string

	// Cpu核数
	CoreNum int

	// 创建时间
	CreateTime int

	// 数据盘大小， 单位GB
	DiskSize int

	// 过期时间
	ExpiredTime int

	// 防火墙Id
	FirewallId string

	// 机房ID
	IdcId string

	// 镜像名称
	ImageName string

	// 节点内存大小，单位GB
	MemSize int

	// 节点带宽限制， 单位Mbs
	NetLimit int

	// 节点ID
	NodeId string

	// 外网ip集合（详情参考NodeIpList）
	NodeIpList []NodeIpList

	// 节点名称
	NodeName string

	// 机房名称
	OcName string

	// 机器类型
	ProductType string

	// 省份
	Province string

	// 节点状态，1部署中，2待启动，3启动中，4运行中，5正在停止，6已停止，7正在更新，8正在重启，9正在删除， 10已经删除,11异常
	State int

	// 系统盘大小， 单位GB
	SysDiskSize int

	// 运营商类型： 0-其它, 1-一线城市单线,2-二线城市单线, 3-全国教育网, 4-全国三通
	Type int
}

/*
NodeIspList - 节点运营商列表
*/
type NodeIspList struct {

	// 城市
	City string

	// 机房名称
	IdcName string

	// 机房运营商名称
	IspName string

	// 线路类型
	LineType string

	// 省份
	Province string
}

/*
MonitorInfo - 监控信息
*/
type MonitorInfo struct {

	// 时间戳
	TimeStamp int

	// 值
	Value int
}

/*
MetricisDataSet - 监控数据
*/
type MetricisDataSet struct {

	// cpu利用率（详情参考MonitorInfo）
	CPUUtilization []MonitorInfo

	// 内存使用率（详情参考MonitorInfo）
	MemUtilization []MonitorInfo

	// 网卡入带宽（详情参考MonitorInfo）
	NICIn []MonitorInfo

	// 网卡出带宽（详情参考MonitorInfo）
	NICOut []MonitorInfo

	// 网卡入包数（详情参考MonitorInfo）
	NetPacketIn []MonitorInfo

	// 网卡出包数（详情参考MonitorInfo）
	NetPacketOut []MonitorInfo
}

/*
ResourceSet - 受到影响的资源列表
*/
type ResourceSet struct {

	// 节点id
	NodeId string

	// 机器外网ip集合
	OuterIps []string
}

/*
IDCCutInfo - 机房割接信息
*/
type IDCCutInfo struct {

	// 城市
	City string

	// 割接类型（中断、抖动、断电）
	CutType string

	// 割接结束时间
	EndTime int

	// 机房名称
	IDCName string

	// 省份
	Province string

	// 受影响的资源信息列表
	ResourceSet []ResourceSet

	// 割接开始时间
	StartTime int
}

/*
DataSet - 监控信息集合
*/
type DataSet struct {

	// cpu使用率
	CPUUtilization []MonitorInfo

	// 磁盘读取次数
	DiskReadOps []MonitorInfo

	// 磁盘写入次数
	DiskWriteOps []MonitorInfo

	// 磁盘读取量
	IORead []MonitorInfo

	// 磁盘写入量
	IOWrite []MonitorInfo

	// 内存使用率
	MemUtilization []MonitorInfo

	// 网卡入带宽
	NICIn []MonitorInfo

	// 网卡出带宽
	NICOut []MonitorInfo

	// 网卡入包量
	NetPacketIn []MonitorInfo

	// 网卡出包量
	NetPacketOut []MonitorInfo
}

/*
DeployImageInfo - 镜像部署信息
*/
type DeployImageInfo struct {

	// 机房ID
	IdcId string

	// 镜像状态 1-可用, 2-不可用, 3-获取中, 4-转换中, 5-部署中
	State int
}

/*
ImageInfo - 镜像详情
*/
type ImageInfo struct {

	// 镜像创建时间戳
	CreateTime int

	// 部署详情列表
	DeployInfoList []DeployImageInfo

	// 镜像描述
	ImageDesc string

	// 镜像ID
	ImageId string

	// 镜像名称
	ImageName string

	// 镜像大小，单位GB
	ImageSize int

	// 镜像类型：1标准镜像，2行业镜像，3自定义镜像
	ImageType int

	// 系统类型：unix, windows
	OcType string

	// 镜像状态：镜像状态 1可用，2不可用
	State int
}
