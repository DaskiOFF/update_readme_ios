package router

import (
	"errors"

	"github.com/daskioff/jessica/configs"
	"github.com/daskioff/jessica/flows"
	"github.com/daskioff/jessica/flows/hi"
	"github.com/daskioff/jessica/flows/projectstruct"
	"github.com/daskioff/jessica/flows/readme"
	"github.com/daskioff/jessica/flows/setup"
	"github.com/daskioff/jessica/flows/templategenerator"
	"github.com/daskioff/jessica/utils"
)

type Router struct {
	mapFlows map[string]flows.Flow
}

func NewRouter() *Router {
	mapFlows := make(map[string]flows.Flow)
	mapFlows["hi"] = hi.NewFlow()
	mapFlows["readme"] = readme.NewFlow()
	mapFlows["setup"] = setup.NewFlow()
	mapFlows["struct"] = projectstruct.NewFlow()
	mapFlows["generator"] = templategenerator.NewFlow()

	router := Router{mapFlows: mapFlows}

	return &router
}

// Handle Обрабатывает аргументы переданные программе и направляет на нужный flow
func (r *Router) Handle(args []string) error {
	if len(args) == 0 {
		return errNoArguments
	}

	command := args[0]
	isHelp := false
	if command == "help" {
		if len(args) < 2 {
			return errHelpNoArguments
		}
		isHelp = true
		command = args[1]
	}

	flow, ok := r.mapFlows[command]
	if !ok {
		return errRouteNotFound
	}

	if isHelp {
		utils.PrintlnInfoMessage(flow.Description())
	} else {
		if command != "setup" && command != "hi" {
			err := configs.ValidateProjectConfig()
			if err != nil {
				return errors.New(err.Error() + "\nДля начала необходимо настроить конфигурацию вызвав команду `jessica setup`")
			}
		}

		flow.Start(args[1:])
	}

	return nil
}