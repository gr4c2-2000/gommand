package command

import (
	"strings"
	"text/template"
	"time"
)

type Command struct {
	ExecTmp     string        `json:"execTmp"`
	DefaultArgs []string      `json:"defaultArgs"`
	WrapperName string        `json:"wrapperName"`
	Timeout     time.Duration `json:"timeout"`
	Sync        bool          `json:"sync"`
	ActionName  string        `json:"actionName"`
	Shell       string        `json:"shell"`
	WorkDir     string        `json:"workDir"`
	exec        Exec
}

type Action interface {
	OnSuccess(*Command)
	OnError(*Command)
}
type Exec struct {
	CommandStr string
	TmpArgs    []string
	CmdArgs    string
}

func (c *Command) GetCommand() (string, error) {
	tmpl, err := template.New(c.ExecTmp).Parse(c.ExecTmp)
	if err != nil {
		return "", err
	}
	tmpParams := getTemplateParameters(c.ExecTmp)

	ma := map[string]string{}
	for key, value := range tmpParams {
		if (len(c.exec.TmpArgs) - 1) >= key {
			ma[value] = c.exec.TmpArgs[key]
		}
	}
	command := strings.Builder{}
	err = tmpl.Execute(&command, ma)
	if err != nil {
		return "", err
	}

	return command.String() + " " + c.exec.CmdArgs, nil
}
