package types

type DeviceInfoStruct struct {
	UserName    string `json:"username"`
	HostName    string `json:"hostname"`
	MachineID   string `json:"machine_id"`
	OSName      string `json:"os_name"`
	Arch        string `json:"architecture"`
	CPUCores    uint32 `json:"cpu_cores"`
	CPUThreads  uint32 `json:"cpu_threads"`
	RAMUsable   int64  `json:"ram_usable"`
	RAMPhysical int64  `json:"ram_physical"`
}
