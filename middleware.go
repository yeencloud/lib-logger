package Logger

type LoggerMiddleware interface {
	Log(message StandardMessage)
}

var middlewares []LoggerMiddleware = make([]LoggerMiddleware, 0)

func AddMiddleware(middleware LoggerMiddleware) {
	middlewares = append(middlewares, middleware)
}
