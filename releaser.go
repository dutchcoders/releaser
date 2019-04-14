package releaser

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/coreos/go-semver/semver"
)

type Releaser struct {
	*Client

	cache    bool
	cacheMap map[string]interface{}

	owner string
	repo  string

	m sync.Mutex
}

func WithCache() func(*Releaser) {
	return func(r *Releaser) {
		r.cache = true
	}
}

func WithRepo(s string) func(*Releaser) {
	return func(r *Releaser) {
		r.repo = s
	}
}

func WithOwner(s string) func(*Releaser) {
	return func(r *Releaser) {
		r.owner = s
	}
}

func New(options ...func(*Releaser)) (*Releaser, error) {
	client, _ := newClient()

	r := &Releaser{
		Client:   client,
		cacheMap: map[string]interface{}{},
	}

	for _, fn := range options {
		fn(r)
	}

	if r.owner == "" {
		return nil, fmt.Errorf("Owner not set")
	}

	if r.repo == "" {
		return nil, fmt.Errorf("Repo not set")
	}

	return r, nil
}

func SinceCommit(hash string) func(Commit) bool {
	found := false
	return func(c Commit) bool {
		if c.Sha == hash {
			found = true
			return true
		} else if strings.HasPrefix(c.Sha, hash) {
			found = true
			return true
		} else if found {
			return true
		}

		return false
	}
}

func (u *Releaser) Commits(filterOpts ...func(Commit) bool) ([]Commit, error) {
	u.m.Lock()
	defer u.m.Unlock()

	path := fmt.Sprintf("/repos/%s/%s/commits", u.owner, u.repo)

	commits := []Commit{}

	if !u.cache {
		// check in cache
	} else if val, ok := u.cacheMap[path]; ok {
		commits = val.([]Commit)
	}

	if len(commits) == 0 {
		req, err := u.NewRequest(http.MethodGet, path, nil)
		if err != nil {
			return nil, err
		}

		if err := u.Do(req, &commits); err != nil {
			return nil, fmt.Errorf("Could not check for update: %s", err.Error())
		}

		u.cacheMap[path] = commits
	}

	rCommits := []Commit{}
	for _, commit := range commits {
		filtered := false
		for _, fn := range filterOpts {
			filtered = filtered || fn(commit)
		}

		if filtered {
			continue
		}

		rCommits = append(rCommits, commit)
	}

	return rCommits, nil
}

func (u *Releaser) Available(version string) (*Release, error) {
	currentVersion, err := semver.NewVersion(version)
	if err != nil {
		return nil, err
	}

	req, err := u.NewRequest(http.MethodGet, fmt.Sprintf("/repos/%s/%s/releases", u.owner, u.repo), nil)
	if err != nil {
		return nil, err
	}

	releases := []Release{}
	if err := u.Do(req, &releases); err != nil {
		return nil, fmt.Errorf("Could not check for update: %s", err.Error())
	}

	if len(releases) == 0 {
		return nil, nil
	}

	name := releases[0].Name
	if strings.HasPrefix(name, "v") {
		name = name[1:]
	}

	if releaseVer, err := semver.NewVersion(name); err != nil {
		// could not parse version
		return nil, err
	} else if releaseVer.LessThan(*currentVersion) {
		return nil, nil
	} else if releaseVer.Equal(*currentVersion) {
		return nil, nil
	}

	return &releases[0], nil
}
