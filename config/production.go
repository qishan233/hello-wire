package config

type Production struct {
}

func (p *Production) GetString(key string) string {
	return "production" + key
}

func NewProduction() *Production {
	return &Production{}
}
