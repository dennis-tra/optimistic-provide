package render

type ErrorCode string

const (
	ErrorCodeInternal        ErrorCode = "INTERNAL"
	ErrorCodeGetPeerFromPath ErrorCode = "GET_PEER_ID_FROM_PATH"
	ErrorCodeMalformedPeerID ErrorCode = "MALFORMED_PEER_ID"
	ErrorCodeHostNotFound    ErrorCode = "HOST_NOT_FOUND"
	ErrorCodeMalformedJSON   ErrorCode = "MALFORMED_JSON"
)
