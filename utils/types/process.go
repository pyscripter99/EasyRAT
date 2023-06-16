package types

type Process struct {
	PID       int    `json:"pid"`
	ParentPID int    `json:"parent_pid"`
	ProcName  string `json:"proc_name"`
}
