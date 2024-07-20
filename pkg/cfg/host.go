package cfg

type Host struct {
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	HostName string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Key      string `json:"key,omitempty"`
}
