package pillow

type CreateDatabaseResponse struct {
	OK     bool   `json:"ok,omitempty"`
	Error  string `json:"error,omitempty"`
	Reason string `json:"reason,omitempty"`
}

type DeleteDatabaseResponse struct {
	OK     bool   `json:"ok,omitempty"`
	Error  string `json:"error,omitempty"`
	Reason string `json:"reason,omitempty"`
}
