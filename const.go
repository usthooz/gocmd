package gocmd

const (
	// DefaultSshPort
	DefaultSshPort string = "22"
)

const (
	// AuthByPassword use password
	AuthByPassword int = 1
	// AuthBySshKey use ssh key
	AuthBySshKey int = 2

	// DefaultAuthType default by password
	DefaultAuthType = AuthByPassword
)

const (
	// SshAuthTimeout ssh auth timeout
	SshAuthTimeout int64 = 5
)
