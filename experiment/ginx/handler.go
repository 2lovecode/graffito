package ginx

type HandlerFunc func(ctx *XContext)

type HandlersChain []HandlerFunc

