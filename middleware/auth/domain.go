package auth

// ProjectsData Projects data struct
type ProjectsData struct {
	Projects []ProjectRole `json:"projects"`
}

// ProjectRole struct
type ProjectRole struct {
	RoleID int `json:"roleID"`
}

// UserData data struct
type UserData struct {
	UserID       string
	UserMobile   string
	ProfileID    *string
	RoleID       *int
	Meta         map[string]MetaData
	IsSuperAdmin bool
}

// MetaData User meta data struct
type MetaData struct {
	RoleID    *int   `json:"roleID,omitempty"`
	ProfileID string `json:"profileID"`
}
