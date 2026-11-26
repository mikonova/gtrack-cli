package types

const (
	NONE = iota
	HASH
	HEADER
	FILE
	NAME
	WHOLEFILE // name and then file again
	EOD       // end of data segment (name or file)
	EOF
	ABORT_CONN
	ACCEPT
	OP_SUCCESS
	OP_FAILURE
	OPERATION // repeat the whole operation
)
