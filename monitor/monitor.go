package monitor

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tonymj76/savannah/config"
	"github.com/tonymj76/savannah/ent"
	"github.com/tonymj76/savannah/ent/gitcommit"
	"github.com/tonymj76/savannah/ent/repository"
	"github.com/tonymj76/savannah/handlers"
	"time"
)

// MRepo uses a ticker to fetch data at regular intervals
func MRepo(h *handlers.RestService) {
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fetchAndSaveMonitoredRepoData(h)
		}
	}
}

func fetchAndSaveMonitoredRepoData(h *handlers.RestService) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fmt.Println("yes am working")
	// Select most recent entry in DB
	commitDate, err := h.DB.GitCommit.Query().
		Order(ent.Desc(gitcommit.FieldDate)).
		Limit(1).
		Select(gitcommit.FieldDate).
		String(ctx)
	if err != nil {
		log.Println("Error fetching repo info:", err)
		return
	}

	//// check if the selected commit is newer to the selected entry in DB
	commits, err := h.Service.FetchCommits(commitDate)
	if err != nil {
		log.Println("Error saving repo info:", err)
		return
	}
	var dataCommits []*ent.GitCommit

	for _, commit := range commits {
		saveCommit, err := h.DB.GitCommit.Create().
			SetURL(commit.Url).
			SetDate(time.Now()).
			SetGitcommit(commit.Commit).
			Save(ctx)

		if err != nil {
			log.Println("Error saving commits:", err)
			return
		}
		dataCommits = append(dataCommits, saveCommit)
	}
	entRepo, err := h.DB.Repository.Query().
		Where(repository.Name(config.GetEnv("repo_name", "savanna-test"))).
		Only(ctx)
	if err != nil {
		log.Println("Error fetching repo info:", err)
		return
	}
	if len(dataCommits) <= 0 {
		log.Println("Error no commit gotten yet")
	}
	entRepo.Update().AddGitCommits(dataCommits...)

	log.Println("Repository data saved successfully")
}
