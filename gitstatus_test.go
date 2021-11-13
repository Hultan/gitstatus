package gitStatus

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGitStatus_NonGitFolder(t *testing.T) {
	git := &GitStatus{}

	path := "/home/per"
	info, err := git.GetStatus(path)
	assert.Nil(t, err, "Should not return an error")
	assert.NotNil(t, info, "Should return a GitStatusInfo object")
	assert.Equal(t, path, info.Path(), fmt.Sprintf("path is not '%s'", path))
	assert.False(t, info.IsGit(), "Should not be a Git folder")

	assert.Equal(t, "", info.Branch(), "Branch() is not empty string")
	assert.Equal(t, 0, info.Ahead(), "Ahead() is not 0")
	assert.Equal(t, 0, info.Behind(), "Behind() is not 0")

	assert.Equal(t, 0, info.Staged(), "staged is not 0")
	assert.Equal(t, 0, info.Modified(), "modified is not 0")
	assert.Equal(t, 0, info.Deleted(), "deleted is not 0")
	assert.Equal(t, 0, info.Unmerged(), "unmerged is not 0")
	assert.Equal(t, 0, info.Untracked(), "untracked is not 0")
}

func TestGitStatus_GitFolder(t *testing.T) {
	git := &GitStatus{}

	path := "/home/per/code/gitprompt-go-test"
	info, err := git.GetStatus(path)
	assert.Nil(t, err, "Should not return an error")
	assert.NotNil(t, info, "Should return a GitStatusInfo object")
	assert.Equal(t, path, info.Path(), fmt.Sprintf("path is not '%s'", path))
	assert.True(t, info.IsGit(), "Should be a Git folder")

	assert.Equal(t, "testBranch", info.Branch(), fmt.Sprintf("Branch() is not 'testBranch'"))
	assert.Equal(t, 0, info.Ahead(), "Ahead() is not 0")
	assert.Equal(t, 0, info.Behind(), "Behind() is not 0")

	assert.Equal(t, 1, info.Staged(), "staged is not 1")
	assert.Equal(t, 1, info.Modified(), "modified is not 1")
	assert.Equal(t, 1, info.Deleted(), "deleted is not 1")
	assert.Equal(t, 0, info.Unmerged(), "unmerged is not 0")
	assert.Equal(t, 1, info.Untracked(), "untracked is not 1")
}
