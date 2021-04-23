package adapters

import (
	"fmt"

	"github.com/thecillu/exagonal-architecture/golang/core/domain"
)

var (
	team_counter   int64 = 1
	player_counter int64 = 1
)

type memoryRepository struct {
	inMemoryTeams   map[int64]domain.Team
	inMemoryPlayers map[int64][]domain.Player
}

func NewMemoryRepository() *memoryRepository {
	inMemoryTeams := map[int64]domain.Team{}
	inMemoryPlayers := map[int64][]domain.Player{}
	return &memoryRepository{inMemoryTeams: inMemoryTeams, inMemoryPlayers: inMemoryPlayers}
}

func (repo *memoryRepository) SaveTeam(team domain.Team) (domain.Team, error) {
	if team.ID == 0 {
		team.ID = team_counter
		team_counter++
	}
	repo.inMemoryTeams[team.ID] = team
	fmt.Println(repo.inMemoryTeams)
	return team, nil
}

func (repo *memoryRepository) GetTeams() ([]domain.Team, error) {
	var teams []domain.Team = make([]domain.Team, 0)
	for _, team := range repo.inMemoryTeams {
		teams = append(teams, team)
	}
	return teams, nil
}

func (repo *memoryRepository) SavePlayer(player domain.Player) (domain.Player, error) {
	if player.ID == 0 {
		player.ID = player_counter
		player_counter++
	}

	teamPlayers := repo.inMemoryPlayers[player.TeamID]
	fmt.Println("teamPlayers", teamPlayers)
	teamPlayers = append(teamPlayers, player)
	repo.inMemoryPlayers[player.TeamID] = teamPlayers
	fmt.Println(repo.inMemoryPlayers)
	return player, nil
}

func (repo *memoryRepository) GetPlayersByTeam(team_id int64) ([]domain.Player, error) {
	players := repo.inMemoryPlayers[team_id]
	if players == nil {
		players = []domain.Player{}
	}
	return players, nil
}
