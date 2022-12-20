package structs

type SysInfo struct {
	HostName string
	OsName   string
	CPU      string
}

type GPUInfo struct {
	GPU string
}

type MemoryInfo struct {
	TotalMemory     string
	AvailableMemory string
}

type PcInfo struct {
	SysInfo
	GPUInfo
	MemoryInfo
}

func NewSysInfo(hostName, osName, cpu string) SysInfo {
	return SysInfo{
		HostName: hostName,
		OsName:   osName,
		CPU:      cpu,
	}
}
func NewGpuInfo(gpu string) GPUInfo {
	return GPUInfo{
		GPU: gpu,
	}
}

func NewMemoryInfo(total, avail string) MemoryInfo {
	return MemoryInfo{
		TotalMemory:     total,
		AvailableMemory: avail,
	}
}

func NewPcInfo(SysInfo SysInfo, GPUInfo GPUInfo, MemInfo MemoryInfo) PcInfo {
	return PcInfo{
		SysInfo:    SysInfo,
		GPUInfo:    GPUInfo,
		MemoryInfo: MemInfo,
	}
}
