package core

import "time"

type HostKeyPolicy string

const (
	HostKeyStrictKnownHosts HostKeyPolicy = "strict_known_hosts"
	HostKeyTOFUPin          HostKeyPolicy = "tofu_pin"
	HostKeyInsecureIgnore   HostKeyPolicy = "insecure_ignore" // dev only
)

type Auth struct {
	Password      string // may be encrypted string (handled by caller/decrypter)
	PrivateKeyPEM string // may be encrypted string
	Passphrase    string // may be encrypted string
}

type JumpHost struct {
	Enabled  bool
	Host     string
	Port     int
	Username string
	Auth     Auth
}

type Policy struct {
	ConnectTimeout   time.Duration
	HandshakeTimeout time.Duration
	CommandTimeout   time.Duration
	MaxOutputBytes   int64

	MaxConnsPerHost int
	MaxRetries      int
}

type Command struct {
	Cmd        string
	RunAs      string        // sudo -u RunAs
	Timeout    time.Duration // overrides Policy.CommandTimeout if >0
	WorkingDir string
	Env        map[string]string
}

type Meta struct {
	TraceID      string
	TenantID     string
	TargetID     string
	CheckID      string
	CheckVersion int
	SchemaVer    int
}

type RunSpec struct {
	Host     string
	Port     int
	Username string
	Auth     Auth
	Jump     JumpHost

	HostKeyPolicy  HostKeyPolicy
	KnownHostsPath string // used when HostKeyStrictKnownHosts

	Policy   Policy
	Commands []Command
	Meta     Meta
}

type CommandResult struct {
	Cmd        string
	RC         int
	Stdout     string
	Stderr     string
	Duration   time.Duration
	TimedOut   bool
	Truncated  bool
	TruncBytes int64
	RunAs      string
}

type Result struct {
	ObservedAt time.Time
	TotalTime  time.Duration
	Commands   []CommandResult

	// Summary: ok|warn|crit|error
	Status string
}
