package git

import (
	"fmt"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/storage/memory"

	"github.com/codecoins/codecoins/log"
)

type Repository struct {
	r *git.Repository
}

func (r *Repository) Branches() Branches {
	return Branches{}
}

func (r *Repository) Checkout(b string) *Reference {
	return &Reference{}
}

func (r *Repository) Head() *Reference {
	h, err := r.r.Head()
	log.DieFatal(err)
	return &Reference{r, h}
}

type Branches struct {
}

func (b *Branches) Next() Branch {
	return Branch{}
}

type Branch struct {
}

func (b *Branch) Head() *Reference {
	return &Reference{}
}

type Reference struct {
	r   *Repository
	tip *plumbing.Reference
}

func (r *Reference) Commits() *Commits {
	cIter, err := r.r.r.Log(&git.LogOptions{From: r.tip.Hash()})
	log.DieFatal(err)
	return &Commits{
		cIter,
	}
}

func Clone(repo string) *Repository {
	log.Info(fmt.Sprintf("git clone %s", repo))

	var err error
	R := new(Repository)
	R.r, err = git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: repo,
	})
	if log.PrintError(err) != nil {
		return R
	}

	return R
}
