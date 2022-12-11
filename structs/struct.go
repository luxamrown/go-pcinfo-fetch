package structs

type SysInfo struct {
	HostName string
	OsName   string
	CPU      string
}

type GPUInfo struct {
	GPU string
}

type PcInfo struct {
	SysInfo
	GPUInfo
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

func NewPcInfo(SysInfo SysInfo, GPUInfo GPUInfo) PcInfo {
	return PcInfo{
		SysInfo: SysInfo,
		GPUInfo: GPUInfo,
	}
}
