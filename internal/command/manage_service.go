package command

import "strings"

type ManageService struct {
	Storage *Storage
}

func NewManageService(s *Storage) *ManageService {
	es := ManageService{Storage: s}
	return &es
}

func (ms *ManageService) GetCommandTemplates() map[string]Command {
	result := map[string]Command{}

	for str, command := range ms.Storage.GetAll() {
		result[str] = command
		params := getTemplateParameters(command.ExecTmp)
		if len(params) > 0 {
			extendCommand := str + "-{{." + strings.Join(params, "}}-{{.") + "}}"
			result[extendCommand] = command
		}
	}
	return result
}
