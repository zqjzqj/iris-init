package config

type Web struct {
	port uint64
}

func (w Web) GetPort() uint64 {
	return w.port
}
