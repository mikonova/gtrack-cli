package category

const (
	REQ_TO_RECV = iota
	REQ_TO_RECV_AND_DEL
	REQ_TO_SEND
	REQ_TO_DEL
	REQ_TO_REPEAT
	DATA
	INFO
)
