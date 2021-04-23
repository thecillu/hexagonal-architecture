package services

import (
	"github.com/thecillu/exagonal-architecture/golang/core/domain"
	"github.com/thecillu/exagonal-architecture/golang/ports"
)

type teamService struct {
	repository ports.RepositoryPort
}

func New(repository ports.RepositoryPort) *teamService {
	return &teamService{
		repository: repository,
	}
}

func (srv *teamService) CreateNewTeam(team domain.Team) (domain.Team, error) {
	team, err := srv.repository.SaveTeam(team)
	if err != nil {
		return domain.Team{}, err
	}
	return team, nil
}

func (srv *teamService) ListTeams() ([]domain.Team, error) {
	teams, err := srv.repository.GetTeams()
	if err != nil {
		return nil, err
	}
	return teams, nil
}

func (srv *teamService) InvitePlayer(player domain.Player) (domain.Player, error) {
	player, err := srv.repository.SavePlayer(player)
	if err != nil {
		return domain.Player{}, err
	}
	return player, nil
}

func (srv *teamService) ListTeamPlayers(team_id int64) ([]domain.Player, error) {
	players, err := srv.repository.GetPlayersByTeam(team_id)
	if err != nil {
		return nil, err
	}
	return players, nil
}
