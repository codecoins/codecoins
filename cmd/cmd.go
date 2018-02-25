package main

import (
	"fmt"

	"github.com/codecoins/codecoins/pkg/config"
	"github.com/codecoins/codecoins/pkg/git"
	"github.com/codecoins/codecoins/pkg/log"
)

func main() {
	GetGit()
}

func GetGit() {

		repo,err := config.GetString("repo.url")
		logLevel,err := config.GetString("log.level")

		log.DieFatal(err)
		log.SetLevel(logLevel)

		r := git.Clone(repo)
		ref := r.Head()
		coms := ref.Commits()

		commitCounter := 0
		commitLimit := 10

		for {
			if commitCounter > commitLimit {
				break
			}
			c,err := coms.Next()
			if err != nil {
				log.PrintError(err)
				break
			}
			commitCounter++
			log.Info(fmt.Sprintf("Commit: %s", c.Hash()))
			log.Info(fmt.Sprintf("Commited by: %s", c.Author().Email()))
			stat := c.Stats()
			log.Info(fmt.Sprintf("Score: %d", stat.Score()))
		}
}