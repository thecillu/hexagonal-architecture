package com.cillu.ports;

import com.cillu.core.domain.Player;
import com.cillu.core.domain.Team;

import java.util.List;

public interface RepositoryPort {
    Team saveTeam(Team team) throws Exception;
    List<Team> getTeams() throws Exception;
    Player savePlayer(Player player) throws Exception;
    List<Player> getPlayersByTeam(int teamID) throws Exception;
}
