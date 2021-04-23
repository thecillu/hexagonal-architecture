package com.cillu.adapters;

import com.cillu.core.domain.Player;
import com.cillu.core.domain.Team;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import com.cillu.ports.TeamServicePort;

import java.util.List;

@RestController
public class HttpHandler {

    private TeamServicePort teamServicePort;

    public HttpHandler(TeamServicePort teamServicePort){
        this.teamServicePort = teamServicePort;
    }

    @GetMapping("/api/v1/teams")
    @ResponseBody
    public List<Team> listTeams() throws Exception {
        return teamServicePort.listTeams();
    }

    @PostMapping("/api/v1/teams")
    @ResponseBody
    public Team createTeam(@RequestBody Team team) throws Exception{
        return teamServicePort.createNewTeam(team);
    }

    @GetMapping("/api/v1/teams/{teamID}/players")
    @ResponseBody
    public List<Player> getTeamPlayers(@PathVariable int teamID) throws Exception{
        return teamServicePort.listTeamPlayers(teamID);
    }

    @PostMapping("/api/v1/teams/{teamID}/invite-player")
    @ResponseBody
    public Player invitePlayer(@RequestBody Player player, @PathVariable int teamID) throws Exception{
        return teamServicePort.invitePlayer(player);
    }

}
