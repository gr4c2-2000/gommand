package command

import (
	"fmt"
	"strings"
	"time"

	"github.com/gr4c2-2000/gommand/internal/proto"
)

func StorageToCommandListResult(stored map[string]Command) *proto.CommandListResult {
	res := make([]*proto.CommandTmp, 0, len(stored))

	for name, command := range stored {
		res = append(res, mapCommandToCommandTmp(command, name))
	}

	fmt.Printf("time : %s\n", time.Now().String())
	return &proto.CommandListResult{Items: res}

}

func mapCommandToCommandTmp(c Command, name string) *proto.CommandTmp {
	res := new(proto.CommandTmp)
	res.ActionName = c.ActionName
	res.DefaultArgs = c.DefaultArgs
	res.ExecTmp = c.ExecTmp
	res.Shell = c.Shell
	res.Sync = c.Sync
	res.WrapperName = c.WrapperName
	res.Name = name
	res.WorkDir = c.WorkDir
	res.Timeout = c.Timeout.String()

	return res
}

func mapCommandTmpToCommand(c *proto.CommandTmp) *Command {
	res := new(Command)
	res.ActionName = c.ActionName
	res.DefaultArgs = c.DefaultArgs
	res.ExecTmp = c.ExecTmp
	res.Shell = c.Shell
	res.Sync = c.Sync
	res.WrapperName = c.WrapperName
	res.WorkDir = c.WorkDir
	timeout, err := time.ParseDuration(c.Timeout)
	if err == nil {
		res.Timeout = timeout
	}
	return res
}

func SplitSpecial(s string) []string {
	specialRuns := []rune{'\'', '"'}
	var specialCharLock bool
	var usedSpecialCharacter rune
	spacePosition := 0
	for key, run := range s {
		if specialCharLock {
			if run == usedSpecialCharacter {
				specialCharLock = false
			}
			continue
		}
		if run == ' ' && !specialCharLock {
			spacePosition = key
			break
		}
		for _, val := range specialRuns {
			if run == val {
				specialCharLock = true
				usedSpecialCharacter = val
			}

		}

	}
	if spacePosition == 0 {
		return []string{s}
	} else {
		return []string{s[:spacePosition], s[spacePosition:]}
	}
}

func getTemplateParameters(tmplStr string) []string {
	var params []string

	// Split the template string by "{{" and "}}"
	parts := strings.Split(tmplStr, "{{")
	for _, part := range parts {
		// If it contains "}}" then it may be a parameter
		if strings.Contains(part, "}}") {
			// Get the content within "{{" and "}}"
			param := strings.TrimSpace(strings.Split(part, "}}")[0])
			// Ensure it starts with "." to be a template parameter
			if strings.HasPrefix(param, ".") {
				param = strings.Replace(param, ".", "", 1)
				params = append(params, param)
			}
		}
	}
	return params
}
