package constants

type EndpointAccessibility int

const (
	ADMIN_ONLY EndpointAccessibility = iota
	ADMIN_AND_OWNER
)
