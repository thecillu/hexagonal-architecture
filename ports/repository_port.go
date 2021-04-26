package ports

import "github.com/thecillu/exagonal-architecture/core/domain"

type RepositoryPort interface {
	GetTeams() ([]domain.Team, error)
	SaveTeam(team domain.Team) (domain.Team, error)
	SavePlayer(member domain.Player) (domain.Player, error)
	GetPlayersByTeam(team_id int64) ([]domain.Player, error)
}
