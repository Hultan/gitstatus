package gitStatus

// GitStatusInfo : Contains information about the git status of a repository
type GitStatusInfo struct {
	path string
	isGit bool

	branch string
	ahead  int
	behind int

	staged    int
	modified  int
	deleted   int
	unmerged  int
	untracked int
}

// Path : Returns the path to the git repository
func (g *GitStatusInfo) Path() string {
	return g.path
}

// IsGit : Returns the true if the path is a git repository
func (g *GitStatusInfo) IsGit() bool {
	return g.isGit
}

// Branch : Returns the name of the current branch
func (g *GitStatusInfo) Branch() string {
	return g.branch
}

// Ahead : Returns the number of commits the repository is ahead of master
func (g *GitStatusInfo) Ahead() int {
	return g.ahead
}

// Behind : Returns the number of commits the repository is behind the master
func (g *GitStatusInfo) Behind() int {
	return g.behind
}

// Staged : Returns the number of staged files
func (g *GitStatusInfo) Staged() int {
	return g.staged
}

// Modified : Returns the number of modified files
func (g *GitStatusInfo) Modified() int {
	return g.modified
}

// Deleted : Returns the number of deleted files
func (g *GitStatusInfo) Deleted() int {
	return g.deleted
}

// Unmerged : Returns the number of unmerged files
func (g *GitStatusInfo) Unmerged() int {
	return g.unmerged
}

// Untracked : Returns the number of untracked files
func (g *GitStatusInfo) Untracked() int {
	return g.untracked
}
