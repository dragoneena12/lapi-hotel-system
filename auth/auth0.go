package auth

type Auth0UserInfo struct {
	UserID string   `json:"sub"`
	Roles  []string `json:"https://lapi.tokyo/claims/roles"`
}
