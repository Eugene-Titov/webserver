package ip

type ServerIP interface {
	GetIP() string
	GetPort() string
}
