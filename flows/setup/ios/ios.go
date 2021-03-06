package ios

import (
	"strings"

	"github.com/daskioff/jessica/configs/models"
	"github.com/daskioff/jessica/flows/internal"
	"github.com/daskioff/jessica/utils/gemfile"
	"github.com/daskioff/jessica/utils/podfile"
	"github.com/daskioff/jessica/utils/print"
	"github.com/daskioff/jessica/utils/question"
	"github.com/daskioff/jessica/utils/xcodeproj"
)

func Setup(config *models.ConfigIOS, isForce bool) {
	err := config.Validate()
	if err == nil && !isForce {
		return
	}

	if !config.HasGemfileUse() || isForce {
		answer := question.AskQuestionWithBoolAnswer("Использовать Gemfile?")
		config.SetGemfileUse(answer)

		if answer {
			if err := gemfile.CreateDefaultIOS(); err == nil {
				print.PrintlnSuccessMessage(podfile.DefaultFilename + " создан")
			}
		}
	}

	if !config.HasPodfileUse() || isForce {
		answer := question.AskQuestionWithBoolAnswer("Использовать Podfile?")
		config.SetPodfileUse(answer)

		if answer {
			if err := podfile.CreateDefault(); err == nil {
				print.PrintlnSuccessMessage(podfile.DefaultFilename + " создан")
			}
		}
	}

	if !config.HasProjectName() || !config.HasXcodeprojFilename() || isForce {
		xcodeprojFilename := question.AskQuestionWithChooseFileAnswer("Выберите "+internal.IosProjectFileExtension+" файл:", internal.IosProjectFileExtension)
		if xcodeprojFilename == "" {
			print.PrintlnAttentionMessage("Пропущена настройка iOS проекта")
			return
		}
		config.SetXcodeprojFilename(xcodeprojFilename)
		config.SetProjectName(strings.Replace(xcodeprojFilename, internal.IosProjectFileExtension, "", 1))

		xcodeproj.Install()
	}

	if !config.HasTargetNameCode() || !config.HasFolderNameCode() || isForce {
		codeProjectFolderName := question.AskQuestionWithChooseFolderAnswer("\nВыберите папку с кодом проекта (обычно она называется так же как и xcodeproj файл): ")
		if codeProjectFolderName == "" {
			codeProjectFolderName = "."
		}
		config.SetFolderNameCode(codeProjectFolderName)
		config.SetTargetNameCode(codeProjectFolderName)
	}

	if !config.HasTargetNameUnitTests() || !config.HasFolderNameUnitTests() || isForce {
		unitTestsFolderName := question.AskQuestionWithChooseFolderAnswer("\nВыберите папку с UNIT тестами: ")
		if unitTestsFolderName != "" {
			config.SetFolderNameUnitTests(unitTestsFolderName)
			config.SetTargetNameUnitTests(unitTestsFolderName)
		}
	}

	if !config.HasTargetNameUITests() || !config.HasFolderNameUITests() || isForce {
		uiTestsFolderName := question.AskQuestionWithChooseFolderAnswer("\nВыберите папку с UI тестами: ")
		if uiTestsFolderName != "" {
			config.SetFolderNameUITests(uiTestsFolderName)
			config.SetTargetNameUITests(uiTestsFolderName)
		}
	}
}
