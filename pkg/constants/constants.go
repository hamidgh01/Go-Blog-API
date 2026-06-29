package constants

type EndpointAccessibility int

const (
	ADMIN_ONLY EndpointAccessibility = iota
	ADMIN_AND_OWNER
)

// Context Keys
const (
	CurrentUserID          string = "cuID"
	CurrentUserUsername    string = "cuUsername"
	CurrentUserEnabled     string = "cuEnabled"
	CurrentUserIsSuperuser string = "cuIsSuperuser"
)

// Http Header Keys
const (
	AuthorizationHeader string = "Authorization"
	RefreshTokenCookie  string = "refresh_token"
)
