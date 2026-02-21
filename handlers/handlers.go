package handlers

import (
	"math/rand"
	"net/http"
	"time"

	"fiction-turnament/db"
	"fiction-turnament/fictions"
	"fiction-turnament/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var input struct {
		Name string `json:"name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name обязателен"})
		return
	}

	user := models.User{Name: input.Name}
	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func ListUsers(c *gin.Context) {
	var users []models.User
	if err := db.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func CreateTournament(c *gin.Context) {
	var input struct {
		UserID int64 `json:"user_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id обязателен"})
		return
	}

	var user models.User
	if err := db.DB.First(&user, input.UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "пользователь не найден"})
		return
	}

	tournament := models.Tournament{
		UserID: input.UserID,
		Status: models.TournamentStatusInProgress,
	}
	if err := db.DB.Create(&tournament).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	shuffled := fictions.Shuffled(rng)

	matches := make([]models.TournamentMatch, 0, 31)
	for i := 0; i < 16; i++ {
		matches = append(matches, models.TournamentMatch{
			TournamentID:    tournament.ID,
			Round:           1,
			SlotInRound:     i,
			Contestant1Slug: shuffled[i*2].Slug,
			Contestant2Slug: shuffled[i*2+1].Slug,
		})
	}

	for round, count := range []int{8, 4, 2, 1} {
		for slot := 0; slot < count; slot++ {
			matches = append(matches, models.TournamentMatch{
				TournamentID: tournament.ID,
				Round:        round + 2,
				SlotInRound:  slot,
			})
		}
	}

	if err := db.DB.Create(&matches).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	db.DB.Preload("User").First(&tournament, tournament.ID)
	c.JSON(http.StatusCreated, tournament)
}

func ListTournaments(c *gin.Context) {
	var tournaments []models.Tournament
	if err := db.DB.Preload("User").Find(&tournaments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tournaments)
}

func GetFictions(c *gin.Context) {
	c.JSON(http.StatusOK, fictions.List)
}

func GetTournament(c *gin.Context) {
	id := c.Param("id")
	var tournament models.Tournament
	if err := db.DB.Preload("User").First(&tournament, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "турнир не найден"})
		return
	}

	var matches []models.TournamentMatch
	if err := db.DB.Where("tournament_id = ?", tournament.ID).Order("round, slot_in_round").Find(&matches).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tournament": tournament,
		"matches":    matches,
	})
}

func VoteMatch(c *gin.Context) {
	tournamentID := c.Param("id")
	matchID := c.Param("matchId")

	var input struct {
		WinnerSlug string `json:"winner_slug" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "winner_slug обязателен"})
		return
	}

	var match models.TournamentMatch
	if err := db.DB.Where("id = ? AND tournament_id = ?", matchID, tournamentID).First(&match).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "матч не найден"})
		return
	}

	if input.WinnerSlug != match.Contestant1Slug && input.WinnerSlug != match.Contestant2Slug {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный winner_slug"})
		return
	}

	match.WinnerSlug = input.WinnerSlug
	if err := db.DB.Save(&match).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	nextRound := match.Round + 1
	nextSlot := match.SlotInRound / 2

	if nextRound <= 5 {
		var nextMatch models.TournamentMatch
		if err := db.DB.Where("tournament_id = ? AND round = ? AND slot_in_round = ?", tournamentID, nextRound, nextSlot).First(&nextMatch).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "следующий матч не найден"})
			return
		}

		if match.SlotInRound%2 == 0 {
			nextMatch.Contestant1Slug = input.WinnerSlug
		} else {
			nextMatch.Contestant2Slug = input.WinnerSlug
		}
		if err := db.DB.Save(&nextMatch).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	if match.Round == 5 {
		db.DB.Model(&models.Tournament{}).Where("id = ?", tournamentID).Update("status", models.TournamentStatusCompleted)
	}

	var tournament models.Tournament
	db.DB.Preload("User").First(&tournament, tournamentID)
	var matches []models.TournamentMatch
	db.DB.Where("tournament_id = ?", tournamentID).Order("round, slot_in_round").Find(&matches)

	c.JSON(http.StatusOK, gin.H{
		"tournament": tournament,
		"matches":    matches,
	})
}
