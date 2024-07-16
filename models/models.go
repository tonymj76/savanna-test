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

	Commit struct {
		Date   string `json:"date"`
		Url    string `json:"url"`
		Author Author `json:"author"`
	}

	Author struct {
		Login string `json:"login"`
		Id    int    `json:"id"`
		Url   string `json:"url"`
	}

	Commits []*Commit
)
