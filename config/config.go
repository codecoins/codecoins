package config

import (
	"os"
	"regexp"
	"strings"

	"github.com/codecoins/codecoins/log"
)

type NoConfError struct{}

func (e NoConfError) Error() string {
	return "no config property found"
}

var (
	envRegEx = `\$[A-Z]+`
)

var confString = map[string]map[string]string{

	"log": {
		"level": "info|error|warning|debug|green",
	},

	"repo": {
		"url": "https://github.com/codecoins/codecoins",
	},

	"storage": {
		"path": "$APP/storage",
	},
}

func GetString(prop string) (string, error) {
	props := strings.Split(prop, ".")
	if len(props) == 2 {
		r, ok := confString[props[0]][props[1]]
		if ok {
			return env(r), nil
		}
	}
	return "", NoConfError{}
}

func env(s string) string {
	e, err := regexp.Compile(envRegEx)
	log.PrintError(err)

	if len(e.FindString(s)) > 0 {
		eStrings := e.Split(s, -1)
		var eStringsMap map[string]string
		for _, s := range eStrings {
			eStringsMap[s] = os.Getenv(s[1:])
		}
		for k, v := range eStringsMap {
			s = strings.Replace(s, k, v, -1)
		}
	}

	return s
}
