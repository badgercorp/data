package entities

type DatabaseComponent struct {
	ID                     int    `json:"id,omitempty"`
	Component              string `json:"component,omitempty"`
	Organization           string `json:"organization,omitempty"`
	RepositoryID           string `json:"repository_id,omitempty"`
	RepositoryName         string `json:"repository_name,omitempty"`
	RepositoryOrganization string `json:"repository_owner,omitempty"`
	RepositorySlug         string `json:"repository_slug,omitempty"`
	OnboardedAt            string `json:"onboarded_at,omitempty"`
	ArchivedAt             string `json:"archived_at,omitempty"`
	Status                 string `json:"status,omitempty"`
}
