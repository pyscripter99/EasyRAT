package types

type ConnectResp struct {
	SessionID string `json:"session_id"`
	Connected bool   `json:"status"`
}

type ServerStatusResp struct {
	Online bool   `json:"online"`
	Status string `json:"status"`
}
