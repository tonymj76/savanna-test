package services

import (
	"encoding/json"
	"fmt"
	"github.com/tonymj76/savannah/config"
	"net/http"
	"time"

	"github.com/tonymj76/savannah/models"
)

var URL = config.GetEnv("github_baseURL", "https://api.github.com/repos/tonymj76/savanna-test")

type GitHubService struct {
	Client *http.Client
}

func NewGitHubService() *GitHubService {
	return &GitHubService{
		Client: &http.Client{Timeout: time.Second * 10},
	}
}

// FetchRepoInfo Implement the logic to fetch repository information from GitHub API
// and return the models.Repository struct
func (s *GitHubService) FetchRepoInfo() (*models.Repository, error) {
	resp, err := s.Client.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var repoInfo *models.Repository

	if err := json.NewDecoder(resp.Body).Decode(&repoInfo); err != nil {
		return nil, err
	}
	return repoInfo, nil
}

func (s *GitHubService) FetchCommits(since string) (models.Commits, error) {
	var url string
	if since == "" || since == " " {
		url = fmt.Sprintf("%s/commits", URL)
	} else {
		url = fmt.Sprintf("%s/commits?since=%s", URL, since)
	}
	resp, err := s.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var commits models.Commits
	if err := json.NewDecoder(resp.Body).Decode(&commits); err != nil {
		return nil, err
	}
	return commits, nil
}
