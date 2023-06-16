package modules

import (
	"fmt"

	"github.com/mitchellh/go-ps"
)

func GetProcList() ([]ps.Process, error) {
	proc_list, err := ps.Processes()
	if err != nil {
		fmt.Printf("Error getting process list. %s\n", err.Error())
	}

	return proc_list, nil
}
