package fx

var (
	// 进程名
	ProcName string

	// 服务分组, 同一台机器(IP), 分组相同
	SvcGroup string

	// 同类服务区分, 进程ID
	SvcIndex int

	// 本进程对应的SvcID
	LocalSvcID string

	// 公网IP
	WANIP string

	// 服务发现地址
	DiscoveryAddress string
)
