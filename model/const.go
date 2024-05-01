package model

const QUOTES_BASE_URL = "https://api.quotable.io/random"

// Request Header keys
const (
	HeaderKeyUserAgent = "User-Agent"
	HeaderKeyUserIP    = "X-Forwarded-For"
	HeaderKeyRequestID = "Request_ID"
)

// Logger additional fields key
const (
	LoggerKeyRequestID = "REQUEST_ID"
	LoggerKeyOperation = "OPERATION"
	LoggerKeyUserIP    = "USER_IP"
	LoggerKeyUserAgent = "USER_AGENT"
)

const (
	SLOTS = "SLOTS"
)
