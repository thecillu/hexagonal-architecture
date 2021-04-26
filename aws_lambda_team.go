package main

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/aws/aws-lambda-go/events"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"

	"github.com/thecillu/exagonal-architecture/adapters"
	"github.com/thecillu/exagonal-architecture/core/domain"
	"github.com/thecillu/exagonal-architecture/core/services"
	"github.com/thecillu/exagonal-architecture/ports"
)

var teamService ports.TeamServicePort
var ginLambda *ginadapter.GinLambda

func init() {
	memoryRepository := adapters.NewMemoryRepository()
	teamService = services.New(memoryRepository)

	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		teams := v1.Group("/teams")
		{
			teams.GET("", retrieveTeams)
			teams.POST("", createNewTeam)
			teams.POST(":team_id/invite-player", invitePlayer)
			teams.GET(":team_id/players", listTeamPlayers)
		}
	}
	ginLambda = ginadapter.New(r)
}

func LambdaHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

/* COMMENTED TO RUN THE APPLICATION IN LOCAL. Uncomment if you want to dpeloy on AWS Lambda
func main() {
	lambda.Start(LambdaHandler)
}
*/

func retrieveTeams(c *gin.Context) {
	teams, err := teamService.ListTeams()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, teams)
}

func createNewTeam(c *gin.Context) {
	var input domain.Team
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	team, err := teamService.CreateNewTeam(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, team)
}

func listTeamPlayers(c *gin.Context) {
	team_id, err := strconv.ParseInt(c.Param("team_id"), 10, 64)
	players, err := teamService.ListTeamPlayers(team_id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, players)
}

func invitePlayer(c *gin.Context) {
	var input domain.Player
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	player, err := teamService.InvitePlayer(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, player)
}
