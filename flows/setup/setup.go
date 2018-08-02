package setup

import (
	"github.com/daskioff/jessica/configs/models"
	"github.com/daskioff/jessica/flows"
	"github.com/daskioff/jessica/utils/print"
	"github.com/daskioff/jessica/utils/slices"
)

type SetupFlow struct {
	globalConfig  *models.ConfigGlobal
	projectConfig *models.ConfigProject
	iosConfig     *models.ConfigIOS
	otherConfig   *models.ConfigOther
}

func (flow *SetupFlow) Start(args []string) {
	isForce := slices.Contains(args, "-force")

	flow.setup(isForce)
}

func (flow *SetupFlow) Description() string {
	return `
	--------------------------------------------------------------------------------
	Первичная настройка файла конфигурации
	Params:
		-force – Обновить всю конфигурацию
	--------------------------------------------------------------------------------
	`
}

// ----------------------------------------------------------------------------
func NewFlow(globalConfig *models.ConfigGlobal, projectConfig *models.ConfigProject, iosConfig *models.ConfigIOS, otherConfig *models.ConfigOther) flows.Flow {
	flow := SetupFlow{}
	flow.globalConfig = globalConfig
	flow.projectConfig = projectConfig
	flow.iosConfig = iosConfig

	return &flow
}

func (flow *SetupFlow) setup(isForce bool) {
	flow.setupGlobal(flow.globalConfig, isForce)
	flow.setupProject(flow.projectConfig, isForce)

	switch flow.projectConfig.GetProjectType() {
	case "iOS":
		flow.projectIOS(flow.iosConfig, isForce)
	case "Other":
		flow.projectOther(flow.otherConfig, isForce)
	}

	flow.globalConfig.Write()
	flow.projectConfig.Write()
	flow.iosConfig.Write()
	flow.otherConfig.Write()

	print.PrintlnSuccessMessage("Файл сконфигурирован")
}
