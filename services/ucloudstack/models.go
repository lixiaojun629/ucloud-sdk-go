// Code is generated by ucloud-model, DO NOT EDIT IT.

package ucloudstack

/*
DiskInfo - 磁盘信息
*/
type DiskInfo struct {

	// 绑定资源ID
	AttachResourceID string

	// 硬盘计费模式。枚举值：Dynamic，表示小时；Month，表示月；Year，表示年；
	ChargeType string

	// 创建时间。时间戳
	CreateTime int

	// 硬盘ID
	DiskID string

	// 硬盘状态。Creating：创建中,BeingCloned：正在被克隆中,Unbound：已解绑,Unbounding：解绑中,Bounding：绑定中,Bound：已绑定,Upgrading：升级中,Deleting：删除中,Deleted：已删除,Releasing：销毁中,Released：已销毁
	DiskStatus string

	// 过期时间。时间戳
	ExpireTime int

	// 名称
	Name string

	// 地域
	Region string

	// 备注
	Remark string

	// 磁盘类型。例如：Normal,SSD
	SetType string

	// 大小。单位GB
	Size int

	// 可用区
	Zone string
}

/*
EIPInfo - 外网IP信息
*/
type EIPInfo struct {

	// 带宽大小
	Bandwidth int

	// 绑定资源ID
	BindResourceID string

	// 绑定资源类型
	BindResourceType string

	// 计费模式。枚举值：Dynamic，表示小时；Month，表示月；Year，表示年；
	ChargeType string

	// 创建时间。时间戳
	CreateTime int

	// ID
	EIPID string

	// 过期时间。时间戳
	ExpireTime int

	// 外网IP
	IP string

	// 名称
	Name string

	// 线路
	OperatorName string

	// 地域
	Region string

	// 备注
	Remark string

	// 状态。Allocating：申请中,Free：未绑定,Bounding：绑定中,Bound：已绑定,Unbounding：解绑中,Deleted：已删除,Releasing：销毁中,Released：已销毁,BandwidthChanging：带宽修改中
	Status string

	// 可用区
	Zone string
}

/*
MetricSet - 监控值
*/
type MetricSet struct {

	// 监控时间
	Timestamp int

	// 监控值
	Value float64
}

/*
MetricInfo - 监控信息
*/
type MetricInfo struct {

	// 监控值信息
	Infos []MetricSet

	// 监控指标。虚拟机的监控指标枚举值为：BlockProcessCount，表示阻塞进程数；CPUUtilization，表示CPU使用率；DiskReadOps，表示磁盘读次数；DiskWriteOps，表示磁盘写次数；IORead，表示磁盘读吞吐；IOWrite，表示磁盘写吞吐；LoadAvg，表示平均负载1分钟；MemUsage，表示内存使用率；NetPacketIn，表示网卡入包量；NetPacketOut，表示网卡出包量；NICIn，表示网卡入带宽；NICOut，表示网卡出带宽；SpaceUsage，表示空间使用率；TCPConnectCount，表示TCP连接数；
	MetricName string
}

/*
SGRuleInfo - 安全组规则信息
*/
type SGRuleInfo struct {

	// 端口号
	DstPort string

	// 方向。1：入，0：出
	IsIn string

	// 优先级。HIGH:高，MEDIUM:中，LOW:低
	Priority string

	// 协议
	ProtocolType string

	// 动作。ACCEPT：接受，DROP：拒绝
	RuleAction string

	// 规则ID
	RuleID string

	// IP或者掩码/段形式。10.0.0.2,10.0.10.10/16
	SrcIP string
}

/*
SGInfo - 安全组信息
*/
type SGInfo struct {

	// 创建时间，时间戳
	CreateTime int

	// 名称
	Name string

	// 地域
	Region string

	// 描述
	Remark string

	// 资源绑定数量
	ResourceCount int

	// 安全组规则。
	Rule []SGRuleInfo

	// 规则数量
	RuleCount int

	// 安全组ID
	SGID string

	// 状态。Creating：创建中,Updating：更新中,Available：有效,Deleted：已删除,Terminating：销毁中,Terminated：已销毁
	Status string

	// 更新时间，时间戳
	UpdateTime int

	// 可用区
	Zone string
}

/*
SubnetInfo - 子网信息
*/
type SubnetInfo struct {

	// 创建时间，时间戳
	CreateTime int

	// 名称
	Name string

	// 网段
	Network string

	// 地域
	Region string

	// 描述
	Remark string

	// 状态；Allocating：申请中,Available：有效,Deleting：删除中,Deleted：已删除
	State string

	// ID
	SubnetID string

	// 更新时间，时间戳
	UpdateTime int

	// 可用区
	Zone string
}

/*
UserInfo - 租户信息
*/
type UserInfo struct {

	// 账户余额
	Amount float64

	// 账户创建时间。时间戳
	CreateTime int

	// 租户名称
	Email string

	// 私钥
	PrivateKey string

	// 公钥
	PublicKey string

	// 用户状态。USER_STATUS_AVAILABLE：正常，USER_STATUS_FREEZE：冻结
	Status string

	// 更新时间。时间戳
	UpdateTime int

	// 租户ID.
	UserID int
}

/*
VMIPInfo - UCloudStack虚拟机IP信息
*/
type VMIPInfo struct {

	// IP 值
	IP string

	// 网卡 ID
	InterfaceID string

	// 是否是弹性网卡。枚举值：Y，表示是；N，表示否；
	IsElastic string

	// MAC 地址值
	MAC string

	// 安全组 ID
	SGID string

	// 安全组名称
	SGName string

	// 子网 ID
	SubnetID string

	// 子网名称
	SubnetName string

	// IP 类型。枚举值：Private，表示内网；Public，表示外网；Physical，表示物理网；
	Type string

	// VPC ID
	VPCID string

	// VPC 名称
	VPCName string
}

/*
VMDiskInfo - UCloudStack虚拟机磁盘信息
*/
type VMDiskInfo struct {

	// 磁盘 ID
	DiskID string

	// 磁盘盘符
	Drive string

	// 是否是弹性磁盘。枚举值为：Y，表示是；N，表示否；
	IsElastic string

	// 磁盘名称
	Name string

	// 磁盘大小，单位 GB
	Size int

	// 磁盘类型。枚举值：Boot，表示系统盘；Data，表示数据盘；
	Type string
}

/*
VMInstanceInfo - UCloudStack虚拟机信息
*/
type VMInstanceInfo struct {

	// CPU 个数
	CPU int

	// 虚拟机计费模式。枚举值：Dynamic，表示小时；Month，表示月；Year，表示年；
	ChargeType string

	// 虚拟机创建时间
	CreateTime int

	// 磁盘信息
	DiskInfos []VMDiskInfo

	// 虚拟机过期时间
	ExpireTime int

	// IP 信息
	IPInfos []VMIPInfo

	// 镜像 ID
	ImageID string

	// 内存大小，单位 M
	Memory int

	// 虚拟机名称
	Name string

	// 操作系统名称
	OSName string

	// 操作系统类型
	OSType string

	// Region
	Region string

	// Region 别名
	RegionAlias string

	// 备注
	Remark string

	// 虚拟机状态。枚举值：Initializing，表示初始化；Starting，表示启动中；Restarting，表示重启中；Running，表示运行；Stopping，表示关机中；Stopped，表示关机；Deleted，表示已删除；Resizing，表示修改配置中；Terminating，表示销毁中；Terminated，表示已销毁；Migrating，表示迁移中；WaitReinstall，表示重装中；Reinstalling，表示重装中；Poweroffing，表示断电中；ChangeSGing，表示修改防火墙中；
	State string

	// 子网 ID
	SubnetID string

	// 子网 名称
	SubnetName string

	// 虚拟机 ID
	VMID string

	// 虚拟机类型
	VMType string

	// 虚拟机类型别名
	VMTypeAlias string

	// VPC ID
	VPCID string

	// VPC 名称
	VPCName string

	// Zone
	Zone string

	// Zone 别名
	ZoneAlias string
}

/*
VPCInfo - VPC信息
*/
type VPCInfo struct {

	// 创建时间，时间戳
	CreateTime int

	// 名称
	Name string

	// 地域。
	Region string

	// 描述
	Remark string

	// 状态；Allocating：申请中,Available：有效,Terminating：销毁中,Terminated：已销毁
	State string

	// 该VPC下拥有的子网数目
	SubnetCount int

	// 该VPC下子网信息。
	SubnetInfos []SubnetInfo

	// 修改时间，时间戳
	UpdateTime int

	// VPC的ID
	VPCID string

	// 可用区
	Zone string
}

/*
PriceInfo - 价格信息
*/
type PriceInfo struct {

	// 计费模式。枚举值：Dynamic，表示小时；Month，表示月；Year，表示年；
	ChargeType string

	// 价格
	Price float64
}
