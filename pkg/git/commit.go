package git

import (
	"github.com/codecoins/codecoins/pkg/log"

	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

type Commits struct{
	i object.CommitIter
}

func(c *Commits) Next() (*Commit,error){
	cc, err := c.i.Next()
	return &Commit{cc ,CommitStat{}},err
}

type Commit struct{
	c *object.Commit
	stats CommitStat
}

func(c *Commit)Hash() string {
	hash := c.c.Hash
	return hash.String()
}

func(c *Commit)Author()Author{
	return Author{
		c.c.Author,
	}
}

func(c *Commit)Message()string {
	return c.c.Message
}

func(c *Commit)String()string {
	return c.c.String()
}

func(c *Commit)Stats()CommitStat{
	if c.stats == (CommitStat{}){
		c.setStats()
	}
	return c.stats
}

func (c *Commit)setStats(){
	fs, err := c.c.Stats()
	if log.PrintError(err) != nil {
		return
	}
	c.stats = parseStat(fs.String())
}

func parseStat(s string)CommitStat {
	var cs CommitStat
	for _,c := range s {
		if string(c) == "+"{
			cs.Added++
		} else if string(c) == "-"{
			cs.Removed++
		}
	}
	return cs
}

type CommitStat struct {
	Added int
	Removed int
}

func (c *CommitStat) Score() int {
	if c.Removed > 2 {
		return (c.Added + c.Removed) * 10 / c.Removed
	}

	return c.Added * 5 / 2
}

type Author struct {
	a object.Signature
}

func(a Author)Email()string {
	return a.a.Email
}

