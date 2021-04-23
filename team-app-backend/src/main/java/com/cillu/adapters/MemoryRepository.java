package com.cillu.adapters;

import com.cillu.core.domain.Player;
import com.cillu.core.domain.Team;
import com.cillu.ports.RepositoryPort;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

public class MemoryRepository implements RepositoryPort {

    private int teamCounter = 1;
    private int playerCounter = 1;

    HashMap<String, Team> teams = new HashMap<String, Team>();;
    HashMap<String, List<Player>> players = new HashMap<>();

    @Override
    public Team saveTeam(Team team) throws Exception {
        team.setID(teamCounter);
        teams.put(String.valueOf(teamCounter), team);
        teamCounter++;
        return team;
    }

    @Override
    public List<Team> getTeams() throws Exception {
        return new ArrayList<>(teams.values());
    }

    @Override
    public Player savePlayer(Player player) throws Exception {
        List<Player> teamPlayers = players.get(String.valueOf(player.getTeamID()));
        if (teamPlayers == null) teamPlayers = new ArrayList<>();
        player.setID(playerCounter);
        teamPlayers.add(player);
        players.put(String.valueOf(player.getTeamID()), teamPlayers);
        playerCounter++;
        return player;
    }

    @Override
    public List<Player> getPlayersByTeam(int teamID) throws Exception {
        List<Player> teamPlayers  = players.get(String.valueOf(teamID));
        if (teamPlayers == null) teamPlayers = new ArrayList<>();
        return teamPlayers;
    }
}
