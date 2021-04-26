package ports

import domain "github.com/thecillu/exagonal-architecture/core/domain"

type TeamServicePort interface {
	CreateNewTeam(team domain.Team) (domain.Team, error)
	ListTeams() ([]domain.Team, error)
	InvitePlayer(player domain.Player) (domain.Player, error)
	ListTeamPlayers(team_id int64) ([]domain.Player, error)
}
