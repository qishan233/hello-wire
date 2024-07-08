package config

type Offline struct {
}

func (o *Offline) GetString(key string) string {
	return "offline" + key
}

func NewOffline() *Offline {
	return &Offline{}
}
