package gitStatus

import (
	"errors"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	version = "1.0.0"
	statusCommand = "git status --porcelain=v2 -z --branch --untracked-files=all"
	nilCharacter  = "\x00"
)

// GitStatus : Checks the status of a Git repository
type GitStatus struct {
}

// GetStatus : Gets the status of a Git repository
func (g *GitStatus) GetStatus(path string) (*GitStatusInfo, error) {
	if path != "" {
		// Check if path exists
		exists, err := g.dirExists(path)
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, errors.New("path does not exist")
		}
	}

	status, err := g.getGitStatus(path)
	if err != nil {
		if _, ok := err.(NotGitRepositoryError); ok {
			return &GitStatusInfo{path: path, isGit: false}, nil
		}
		return nil, err
	}

	info, err := g.parseGitStatus(status, path)
	if err != nil {
		return nil, err
	}

	return info, nil
}

//
// Private methods
//

func (g *GitStatus) getGitStatus(path string) (string, error) {
	prevPath, err := os.Getwd()
	if err != nil {
		return "", err
	}
	if path != "" {
		err = os.Chdir(path)
		if err != nil {
			return "", err
		}
	}

	command := exec.Command("/bin/bash", "-c", statusCommand)
	out, err := command.Output()
	if err != nil {
		_ = os.Chdir(prevPath)
		if exitErr, ok := err.(*exec.ExitError); ok {
			if strings.Contains(string(exitErr.Stderr), "not a git repository") {
				return "", NotGitRepositoryError{path}
			} else {
				return "", err
			}
		}
		return "", err
	}

	_ = os.Chdir(prevPath)
	return string(out), nil
}

func (g *GitStatus) parseGitStatus(status, path string) (*GitStatusInfo, error) {
	info := &GitStatusInfo{path: path,isGit: true}
	items := strings.Split(status, nilCharacter)
	var err error
	for _, value := range items {
		switch {
		case strings.HasPrefix(value, "#"):
			err = g.parseBranch(value, info)
			if err != nil {
				return nil, err
			}
		case strings.HasPrefix(value, "1"), strings.HasPrefix(value, "2"):
			g.parseFile(value, info)
		case strings.HasPrefix(value, "u"):
			info.unmerged += 1
		case strings.HasPrefix(value, "?"):
			info.untracked += 1
		default:
		}
	}

	return info, nil
}

func (g *GitStatus) parseBranch(branch string, info *GitStatusInfo) error {
	items := strings.Split(branch, " ")
	switch items[1] {
	case "branch.head":
		if items[2] != "(detached)" {
			info.branch = items[2]
		}
	case "branch.ab":
		ahead, err := strconv.Atoi(items[2])
		if err != nil {
			return err
		}
		info.ahead = ahead

		behind, err := strconv.Atoi(items[3])
		if err != nil {
			return err
		}
		info.behind = behind
	}

	return nil
}

func (g *GitStatus) parseFile(file string, info *GitStatusInfo) *GitStatusInfo {
	items := strings.Split(file, " ")
	fileStatus := items[1]

	if fileStatus[0] != '.' {
		info.staged += 1
	}

	switch {
	case fileStatus[1] == 'M':
		info.modified += 1
	case fileStatus[1] == 'D':
		info.deleted += 1
	}

	return info
}

// https://stackoverflow.com/questions/10510691/how-to-check-whether-a-file-or-directory-exists
func (g *GitStatus) dirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
