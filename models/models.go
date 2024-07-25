package models

type (
	Repository struct {
		Name            string `json:"name"`
		URL             string `json:"url"`
		Description     string `json:"description"`
		Language        string `json:"languages_url"`
		CreatedAt       string `json:"created_at"`
		UpdatedAt       string `json:"updated_at"`
		OpenIssuesCount int    `json:"open_issues_count"`
		WatchersCount   int    `json:"watchers_count"`
		StarCount       int    `json:"stargazers_count"`
		ForksCount      int    `json:"forks_count"`
	}

	GitCommit struct {
		Date   string         `json:"date"`
		Url    string         `json:"url"`
		Commit map[string]any `json:"commit"`
	}

	Commits []*GitCommit
)
