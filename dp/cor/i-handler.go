package cor

type IHandler interface {
	Break()
	Execute()
	IsBreak() bool
	Next(mvc IHandler) IHandler
}
