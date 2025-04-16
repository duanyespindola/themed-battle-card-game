package waitingRoom

/**
* REQ-B85E
**/
type Status int

const (
	WaitingForAnotherPlayerStatus Status = iota + 1
	WaitingForMatchToStartStatus
)
