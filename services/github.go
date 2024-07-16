package services

import (
	"encoding/json"
	"fmt"
	"github.com/tonymj76/savannah/config"
	"net/http"
	"time"

	"github.com/tonymj76/savannah/models"
)

var baseURL = config.GetEnv("github_baseURL", "https://api.github.com/repos")

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
func (s *GitHubService) FetchRepoInfo(owner, repo string) (*models.Repository, error) {
	url := fmt.Sprintf("%s/%s/%s", baseURL, owner, repo)
	resp, err := s.Client.Get(url)
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

func (s *GitHubService) FetchCommits(owner, repo string) (models.Commits, error) {
	url := fmt.Sprintf("%s/%s/%s/commits", baseURL, owner, repo)
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
