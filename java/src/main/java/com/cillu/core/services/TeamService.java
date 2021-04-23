package com.cillu.core.services;

import com.cillu.core.domain.Player;
import com.cillu.core.domain.Team;
import com.cillu.ports.RepositoryPort;
import com.cillu.ports.TeamServicePort;

import java.util.List;

public class TeamService implements TeamServicePort {

    private RepositoryPort repositoryPort;

    public TeamService(RepositoryPort repositoryPort){
        this.repositoryPort = repositoryPort;
    }

    @Override
    public Team createNewTeam(Team team) throws Exception {
        return repositoryPort.saveTeam(team);
    }

    @Override
    public List<Team> listTeams() throws Exception {
        return repositoryPort.getTeams();
    }

    @Override
    public Player invitePlayer(Player player) throws Exception {
        return repositoryPort.savePlayer(player);
    }

    @Override
    public List<Player> listTeamPlayers(int teamID) throws Exception {
        return repositoryPort.getPlayersByTeam(teamID);
    }
}
