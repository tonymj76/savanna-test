package handlers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/tonymj76/savannah/config"
	"github.com/tonymj76/savannah/ent"
	"net/http"
	"time"
)

func (h *RestService) HandleWebhook(c *gin.Context) {
	var payload map[string]interface{}
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.JSON(c, "Repo data saved successfully", http.StatusOK, payload)
}

func (h *RestService) FetchAndSaveRepoData(c *gin.Context) {

	var dataCommits []*ent.GitCommit

	commits, err := h.Service.FetchCommits("")
	for _, commit := range commits {
		saveCommit, err := h.DB.GitCommit.Create().
			SetURL(commit.Url).
			SetDate(time.Now()).
			SetGitcommit(commit.Commit).
			Save(c)

		if err != nil {
			log.Println("Error saving commits:", err)
			config.JSON(c, "Failed to save commits", http.StatusInternalServerError, err)
			return
		}
		dataCommits = append(dataCommits, saveCommit)
	}
	if err != nil {
		log.Println("Error fetching commits:", err)
		config.JSON(c, "Failed to fetch commits", http.StatusInternalServerError, err)
		return
	}

	repoInfo, err := h.Service.FetchRepoInfo()
	if err != nil {
		log.Println("Error fetching repo info:", err)
		config.JSON(c, "Failed to fetch repo info", http.StatusInternalServerError, err)
		return
	}
	createdAt, _ := time.Parse(time.RFC1123, repoInfo.CreatedAt)
	updateAt, _ := time.Parse(time.RFC1123, repoInfo.UpdatedAt)
	info, err := h.DB.Repository.
		Create().
		SetName(repoInfo.Name).
		SetDescription(repoInfo.Description).
		SetCreatedAt(createdAt).
		SetForksCount(repoInfo.ForksCount).
		SetLanguage(repoInfo.Language).
		SetOpenIssuesCount(repoInfo.OpenIssuesCount).
		SetURL(repoInfo.URL).
		SetUpdatedAt(updateAt).
		SetWatchersCount(repoInfo.WatchersCount).
		SetStarCount(repoInfo.StarCount).
		AddGitCommits(dataCommits...).
		Save(c)

	if err != nil {
		log.Println("Error saving repo info:", err)
		config.JSON(c, "Failed to save repo info", http.StatusInternalServerError, err)
		return
	}

	config.JSON(c, "Repo data saved successfully", http.StatusOK, info)
}
