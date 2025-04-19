package router

type Player struct {
	PlayerID string
	Name string
	IsConnected bool
	Result TypingResult
}

type TypingResult struct {
	WPM int
	Accuracy int
}
