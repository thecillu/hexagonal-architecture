package domain

type Player struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" binding:"required"`
	TeamID    int64  `json:"team_id"`
}
