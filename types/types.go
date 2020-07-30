package types

type Credential struct {
	User        string `json:"user"`
	Password    string `json:"pass"`
	Description string `json:"description"`
	EnvType     string `json:"env_type"`
}
type CredentialTim struct {
	User     string `json:"user"`
	Password string `json:"pass"`
}

type HostnameTim struct {
	Credentials map[string]*CredentialTim
}

type Hostname struct {
	Credentials map[string]*Credential
}

type Terminal struct {
	Rows    string
	Columns string
}