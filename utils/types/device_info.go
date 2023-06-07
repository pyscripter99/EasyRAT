package types

type DeviceInfoStruct struct {
	OSName      string `json:"os_name"`
	Arch        string `json:"architecture"`
	CPUCores    uint32 `json:"cpu_cores"`
	CPUThreads  uint32 `json:"cpu_threads"`
	RAMUsable   int64  `json:"ram_usable"`
	RAMPhysical int64  `json:"ram_physical"`
}
