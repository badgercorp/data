package entities

type DatabaseComponent struct {
	ID                     int
	Component              string
	Organization           string
	RepositoryID           string
	RepositoryName         string
	RepositoryOrganization string
	RepositorySlug         string
	OnboardedAt            string
	ArchivedAt             string
	Status                 string
}
