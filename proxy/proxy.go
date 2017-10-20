package proxy

type Proxy struct {
	Config
}

func New(config Config) *Proxy {
	return &Proxy{
		Config: config,
	}
}

func (p *Proxy) Run() {

}
