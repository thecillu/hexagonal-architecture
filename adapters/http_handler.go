package adapters

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thecillu/exagonal-architecture/core/domain"
	"github.com/thecillu/exagonal-architecture/ports"
)

type httpHandler struct {
	service ports.TeamServicePort
}

func NewHttpHandler(service ports.TeamServicePort) *httpHandler {
	return &httpHandler{
		service: service,
	}
}

func (hdl *httpHandler) RetrieveTeams(c *gin.Context) {
	teams, err := hdl.service.ListTeams()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, teams)
}

func (hdl *httpHandler) CreateNewTeam(c *gin.Context) {
	var input domain.Team
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	team, err := hdl.service.CreateNewTeam(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, team)
}

func (hdl *httpHandler) ListTeamPlayers(c *gin.Context) {
	team_id, err := strconv.ParseInt(c.Param("team_id"), 10, 64)
	players, err := hdl.service.ListTeamPlayers(team_id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, players)
}

func (hdl *httpHandler) InvitePlayer(c *gin.Context) {
	var input domain.Player
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	player, err := hdl.service.InvitePlayer(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, player)
}
