package com.cillu.ports;

import com.cillu.core.domain.Player;
import com.cillu.core.domain.Team;

import java.util.List;

public interface TeamServicePort {
    Team createNewTeam(Team team) throws Exception;
    List<Team> listTeams() throws Exception;
    Player invitePlayer(Player player) throws Exception;
    List<Player> listTeamPlayers(int teamID) throws Exception;
}
