package entities

type DatabaseComponent struct {
	ID                     int    `json:"id"`
	Component              string `json:"component"`
	Organization           string `json:"organization"`
	RepositoryID           string `json:"repository_id "`
	RepositoryName         string `json:"repository_name"`
	RepositoryOrganization string `json:"repository_owner"`
	RepositorySlug         string `json:"repository_slug"`
	OnboardedAt            string `json:"onboarded_at"`
	ArchivedAt             string `json:"archived_at,omitempty"`
	Status                 string `json:"status"`
}
