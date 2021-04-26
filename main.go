package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thecillu/exagonal-architecture/adapters"
	"github.com/thecillu/exagonal-architecture/core/services"
)

func main() {
	memoryRepository := adapters.NewMemoryRepository()
	teamService := services.New(memoryRepository)
	httpHandler := adapters.NewHttpHandler(teamService)

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		teams := v1.Group("/teams")
		{
			teams.GET("", httpHandler.RetrieveTeams)
			teams.POST("", httpHandler.CreateNewTeam)
			teams.POST(":team_id/invite-player", httpHandler.InvitePlayer)
			teams.GET(":team_id/players", httpHandler.ListTeamPlayers)
		}
	}
	r.Run()
}
