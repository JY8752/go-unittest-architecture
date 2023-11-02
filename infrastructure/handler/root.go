package handler

type Handler interface {
	Register()
}

type Root struct {
	handlers []Handler
}

func NewRoot(g *Gacha) *Root {
	return &Root{[]Handler{g}}
}

func (r *Root) RegisterAll() {
	for _, h := range r.handlers {
		h.Register()
	}
}
