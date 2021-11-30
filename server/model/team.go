package model

// Team is information global to a team
// swagger:model
type Team struct {
	// ID of the team
	// required: true
	ID string `json:"id"`

	// Title of the team
	// required: false
	Title string `json:"title"`

	// Token required to register new users
	// required: true
	SignupToken string `json:"signupToken"`

	// Team settings
	// required: false
	Settings map[string]interface{} `json:"settings"`

	// ID of user who last modified this
	// required: true
	ModifiedBy string `json:"modifiedBy"`

	// Updated time
	// required: true
	UpdateAt int64 `json:"updateAt"`
}

// ToDo: should we maintain this struct??
// UserTeam is a summary of a single association between
// a user and a team
// swagger:model
type UserTeam struct {
	// ID of the team
	// required: true
	ID string `json:"id"`

	// Title of the team
	// required: false
	Title string `json:"title"`

	// Number of boards in the team
	BoardCount int `json:"boardCount"`
}
