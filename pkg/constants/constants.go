package constants

type EndpointAccessibility int

const (
	ADMIN_ONLY EndpointAccessibility = iota
	ADMIN_AND_OWNER
)

type AuthKeys string

const (
	CurrentUserID          AuthKeys = "cuID"
	CurrentUserUsername    AuthKeys = "cuUsername"
	CurrentUserEnabled     AuthKeys = "cuEnabled"
	CurrentUserIsSuperuser AuthKeys = "cuIsSuperuser"
)
