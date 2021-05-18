package cor

type baseHandler struct {
	isBreak bool
	nexts   []IHandler
}

func (b *baseHandler) Break() {
	b.isBreak = true
}

func (b baseHandler) IsBreak() bool {
	return b.isBreak
}

func (b *baseHandler) Next(handler IHandler) IHandler {
	b.nexts = append(b.nexts, handler)
	return b
}

func (b *baseHandler) Execute() {
	if b.isBreak {
		return
	}

	for _, handler := range b.nexts {
		if handler.IsBreak() {
			return
		}

		handler.Execute()
	}
}

func New() IHandler {
	return &baseHandler{}
}
