package gitStatus

import (
	"fmt"
)

// NotGitRepositoryError : The specified path is not a Git repository
type NotGitRepositoryError struct {
	Path string
}

func (e NotGitRepositoryError) Error() string {
	return fmt.Sprintf("directory %v is not a git repository", e.Path)
}
