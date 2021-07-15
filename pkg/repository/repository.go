package repository

type Repository interface {
	GetFiles(options RepositoryOptions) ([]string, error)
	CreateFile(name string, commit string, data []byte, options RepositoryOptions) error
}

type RepositoryOptions struct {
	Organization string
	Repository   string
	Path         string
	Branch       string
}
