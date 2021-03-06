# Jessica
Last version: 1.5.1

## Installation
```
brew tap daskioff/jessica
brew install daskioff/jessica/jessica
```

## Upgrade
```
brew upgrade daskioff/jessica/jessica
```

## Reinstall and Uninstall
```
brew reinstall daskioff/jessica/jessica
brew uninstall --force daskioff/jessica/jessica
```

# Commands
`jessica help <command>` - Помощь по команде

|Command|Description|
|----|---|
|`version`|Выводит номер текущей версии приложения|
|[setup](#setup)|Конфигурация проекта|
|[readme](#readme-command)|Создание необходимых файлов и шаблонов для генерации `README.md` файла|
|[struct](#struct)|Создание и описание структуры проекта|
|[generator](#generator)|Генерация файлов для проекта|

# Usage

Формат `jessica <command> [action] [args]`

## Шаг 1
`jessica setup`

## Шаг 2 (Если нужна поддержка README.md)
`jessica readme`

## Шаг 3 (Если необходимо генерировать шаблоны)

```
jessica generator pull github.com/daskioff/jessica_templates newapp`
jessica generator gen newapp
```

## Example
Пример находится в папке `ExampleXcodeProj` 

Примеры шаблонов для команды `generator` можно посмотреть в разных ветках [репозитория](https://github.com/DaskiOFF/jessica_templates)

# Setup
Конфигурация проекта, чтобы использовать остальные команды кроме `version` и `setup`.

Запрашиваются только недостающие поля (для обновления всей конфигурации можно использовать параметры)

|Params|Description|
|----|---|
|`--force, --f`|Полное обновление конфигурации|

## Результатом команды являются два файла
### ~/.jessica.yml – глобальный файл конфигурации
  - `user_name` – Имя пользователя

### [project_path]/.jessica.yml – Файл конфигурации проекта
  - `company_name` – Имя компании (Для шаблонов)
  - `project_type` – Тип проекта [iOS|other]
  - `readme_template_filename` – Имя файла для шаблона README файла
  - `custom_project_struct_use` – Использовать или нет кастомную структуру проекта
  - `custom_project_struct_description` – Описание структуры проекта
  - `custom_project_struct_description_template_filename` – Имя файла с шаблоном описания структуры проекта
  - `templates_use` – Использовать шаблоны или нет
  - `templates_folder_name` – Имя папки, содержащей шаблоны

##### Для проекта типа `iOS`
  - `ios_dependencies_gemfile_use` – Использовать Gemfile или нет
  - `ios_dependencies_podfile_use` – Использовать Podfile или нет
  - `ios_xcodeproj_filename` – Имя xcodeproj файла проекта
  - `ios_target_name_code` – Название таргета кода проекта
  - `ios_folder_name_code` – Путь до папки с кодом проекта
  - `ios_target_name_unit_tests` – Название таргета unit тестов проекта
  - `ios_folder_name_unit_tests` – Путь до папки unit тестов проекта
  - `ios_target_name_ui_tests` – Название таргета ui тестов проекта
  - `ios_folder_name_ui_tests` – Путь до папки ui тестов проекта

##### Для проекта типа `other`
  - `other_project_name` – Название проекта
  - `other_project_folder_name` – Путь до папки с кодом проекта

# Readme command
Поддержка актуальности `README.md` файла

Пока есть недостающие данные – они будут запрашиваться, все последующие вызовы будут просто обновлять файл `README.md` по шаблону

Для iOS проекта
  - Будут запрошены:
    - Версия xcode, с которой проект последний раз собирался. Ответ будет сохранен в файл `.xcode-version`
    - Версия swift. Ответ будет сохранен в файл `.swift-version`
  - Будут созданы:
    - Файл `Gemfile` с первоначальными зависимостями
    - Файл `Podfile` с первоначальными зависимостями
    - Файл [readme_template_filename](#setup), описывающий шаблон результирующего файла `README.md`

Переменные, используемые при генерации `README.md` из шаблона [readme_template_filename](#setup):
  - Для всех типов проектов:
    - `projectName` – Имя проекта
  - Для iOS проекта:
    - `xcodeVersion` – Версия xcode из файла
    - `swiftVersion` – Версия swift из файла
    - `gemFileDependencies` – Список зависимостей Gemfile
    - `podFileDependencies` – Список зависимостей проекта Podfile

Если файл описания структуры [templates_folder_name](#setup) существует, он подключается в конец файла [readme_template_filename](#setup)

# Struct
Создание и описание структуры проекта.

|Action|Description|
|----|---|
|`gen`|Генерация структуры проекта, описанной в конфигурационном файле проекта|
|`example`|Пример описания структуры проекта в конфигурационном файле|

## gen
- Генерация структуры проекта
- Создание шаблона [custom_project_struct_description_template_filename](#setup) описания структуры проекта. В шаблоне доступны все те же переменные, что и для [readme_template_filename](#setup) файла

## example
Выводит пример описания структуры в конфигурационном файле

### Пример описания структуры проекта
```yml
custom_project_struct_description:
  - Data:
    - Entities
  - Domain:
    - Entities
  - Presentation:
    - Resources:
      - Fonts
    - Components:
      - Cells
    - Flows
    - ViewControllers
  - Support
```

# Generator
Генерация файлов из шаблонов для проекта.

|Action|Description|
|----|---|
|`list`|Список шаблонов|
|`pull git_url [branch]`|Клонирование репозитория в папку шаблонов|
|`gen TEMPLATE_NAME MODULE_NAME [PARAMS] [CUSTOM_KEYS]`|Генерация|

## list
Находит и выводит список всех доступных шаблонов из папки шаблонов, указанной в файле конфигурации по ключу [templates_folder_name](#setup), доступных для генерации с помощью действия `gen`

## pull
Пример использования
```
jessica generator pull github.com/daskioff/jessica_templates
```
или
```
jessica generator pull github.com/daskioff/jessica_templates my
```
или
```
jessica generator pull https://github.com/daskioff/jessica_templates.git
```

## gen
После указания действия `gen` необходимо указать имя шаблона и имя генерируемого модуля (необязательно). Далее перечисляются параметры, кастомные ключи и значения, которые доступны в шаблоне по ключу `{{.custom.имя_переданного_ключа}}`

### Пример
```
jessica generator gen repository User --nomock userCusomKey1:Value1 userCustom2:value2
```

### Описание файла, описывающего шаблон
Название шаблона – это имя папки с файлом `templates.yml`, которая находится в общей папке шаблонов проекта [templates_folder_name](#setup).

Структура файла `templates.yml`. Доступно 5 секций:
1. `variables` – optional
1. `questions` – optional
1. `code_files` – required
1. `test_files` – optional
1. `mock_files` – optional

#### Variables
Секция содержит переменные, которые можно использовать по ключу `{{.var.VariableName}}`. 

```yml
variables:
  key1: value1
  key2: "{{.var.key1}}value2"
```

#### Questions
Секция содержит вопросы, ответы на которые можно использовать в шаблоне по ключу `{{.answers.KeyName}}`. Формат описания вопроса в файле шаблона:

|Name|Type|Description|
|---|---|---|
|`key`|string|Ключ, по которому будет доступен ответ в генерируемом шаблоне|
|`text`|string|Текст вопроса|
|`required`|bool|Обязательно ли требуется непустой ответ на вопрос|

#### (Code|Test|Mock)_files
Каждая секция, описывающая файлы, может содержать список генерируемых файлов. В значениях можно использовать значения, описанные ниже в разделе "Шаблонные значения". Описание каждого файла содержит:

|Name|Type|Description|
|---|---|---|
|`name`|string|Суффикс генерируемого файла. Название модуля должно быть указано явно! (Например, `{{.moduleInfo.name}}`)|
|`template_path`|string|Путь внутри папки шаблона, относительно файла, описывающего шаблон|
|`output_path`|string|Выходной путь сгенерированного файла, возможно использование переменных|
|`rewrite`|bool|Значение true или false, означающее стоит ли перезаписывать генерируемый файл, если файл с таким именем по сохраняемому пути уже существует. Если ключ не указан, то значение будет запрошено во время выполнения|

По умолчанию генерируются все секции, `code_files` является обязательной всегда. Другие можно отключить передав `params`:
- `--notest` – для отключения генерации секции `test_files`
- `--nomock` – для отключения генерации секции `mock_files`

#### Шаблонные значения
- Все значения по ключу [Variables](#variables)
- Все значения по ключу [Custom](#custom)
- Все значения по ключу [Answers](#answers)
- Все значения по ключу [ModuleInfo](#moduleinfo)
- `projectName` – Имя проекта, для которого генерируется
- `projectCodeFolderPath` – Путь до папки с основным кодом проекта от корня
- `projectTestsName` – Имя таргета с тестами проекта из файла конфигурации (Для проектов типа iOS)
- `projectTestsCodeFolderPath` – Путь до папки с тестами проекта из файла конфигурации (Для проектов типа iOS)
- `projectUITestsName` – Имя таргета с ui тестами проекта из файла конфигурации (Для проектов типа iOS)
- `projectUITestsCodeFolderPath` – Путь до папки с ui тестами проекта из файла конфигурации (Для проектов типа iOS)

#### Пример файла, описывающего шаблон
```yml
variables:
  basePathData: Layers/DataLayer/Entities
  key2: "{{.var.basePathData}}/val2"

questions:
  - {key: versionApi,
    text: "Enter API version: ",
    required: true}

  - {key: entryPoint,
    text: "Enter Entry point: ",
    required: true}

  - {key: entityName,
    text: "Enter Entity name: ",
    required: true}

  - {key: suffix,
    text: "Enter suffix for module name: ",
    required: false}

code_files: 
  - {name: {{.moduleInfo.name}}BaseUseCase.swift, 
    template_path: code/baseUseCase.swift, 
    output_path: "{{.projectCodeFolderPath}}/{{.var.basePathData}}/base/{{.moduleInfo.name}}", 
    rewrite: true}

  - {name: "{{.moduleInfo.name}}{{.answers.suffix}}UseCase.swift", 
    template_path: code/usecase.swift, 
    output_path: "{{.projectCodeFolderPath}}/{{.var.basePathData}}/base/{{.moduleInfo.name}}/usecases", 
    rewrite: false}

test_files: 
  - {name: "{{.moduleInfo.name}}{{.answers.suffix}}UseCaseImplTests.swift",
    template_path: tests/useCaseImplTests.swift, 
    output_path: "{{.projectTestsName}}/Layers/DataLayer/Entities/{{.moduleInfo.name}}"}

mock_files:
  - {name: "PartialMock{{.moduleInfo.name}}Repository.swift", 
    template_path: mocks/partialMockUseCaseImpl.swift, 
    output_path: "{{.projectTestsName}}/Mocks/{{.moduleInfo.name}}"}
```

### Описание генерируемого файла
Переменные необходимо использовать с помощью конструкции `{{.VariableName}}`. Подробнее про используемый шаблонизатор можно прочитать [здесь](https://golang.org/pkg/text/template/)

#### Список доступных переменных, их типы и описания:

|VariableName|Type|Description|
|----|---|---|
|`date`|string|Текущая дата в формате dd.MM.yyyy|
|`year`|int|Текущий год|
|`fileName`|string|Имя сгенерированного файла|
|`projectName`|string|Имя проекта, для которого генерируется|
|`projectCodeFolderPath`|string|Путь до папки с основным кодом проекта от корня|

#### Дополнительные переменные для iOS проекта:

|VariableName|Type|Description|
|----|---|---|
|`projectTestsName`|string|Имя таргета с UNIT тестами, если задан в конфиге, иначе значение совпадает со значением по ключу projectName|
|`projectTestsCodeFolderPath`|string|Путь до папки с UNIT тестами, если задан в конфиге, иначе значение совпадает со значением по ключу projectCodeFolderPath|
|`projectUITestsName`|string|Имя таргета с UI тестами, если задан в конфиге, иначе значение совпадает со значением по ключу projectName|
|`projectUITestsCodeFolderPath`|string|Путь до папки с UI тестами, если задан в конфиге, иначе значение совпадает со значением по ключу projectCodeFolderPath|

#### Variables
Использовать `{{.var.VariableName}}`

Содержит ключи и значения, описанные в файле шаблона в секции `variables`, тип значения всегда `string`

#### Custom
Использовать `{{.custom.VariableName}}`

Содержит ключи и значения, переданные при запуске, тип значения всегда `string`

_Например:_
```
jessica generator gen usecase User key1:Value1
```

В шаблоне можно будет использовать как `{{.custom.key1}}`

#### Answers
Использовать `{{.answers.VariableName}}`

Содержит ключи, указанные при описании вопроса, и ответы, которые дал пользователь, тип значения всегда `string`

_Например, в шаблоне описать вопрос:_
```
questions:
  - {key: versionApi,
    text: "Enter API version: ",
    required: true}
```

В шаблоне можно будет использовать как `{{.answers.versionApi}}`

#### ModuleInfo
Использовать `{{.moduleInfo.VariableName}}`

|VariableName|Type|Description|
|----|---|---|
|`name`|string|Имя модуля, которое было передано при генерации (Например, UserModule)|
|`nameCapitalize`|string|Имя модуля, которое было передано при генерации, но с первой `Прописной` буквы (Например, UserModule)|
|`nameFirstLower`|string|Имя модуля, которое было передано при генерации, но с первой `строчной` буквы (Например, userModule)|
|`nameUppercase`|string|Имя модуля, которое было передано при генерации, но со всеми `ПРОПИСНЫМИ` буквами (Например, USERMODULE)|
|`nameLowercase`|string|Имя модуля, которое было передано при генерации, но со всеми `строчными` буквами (Например, usermodule)|

#### Developer
Использовать `{{.developer.VariableName}}`

|VariableName|Type|Description|
|----|---|---|
|`name`|string|Имя разработчика из глобального файла конфигурации|
|`companyName`|string|Имя компании из локального файла конфигурации|

## Autor

DaskiOFF, waydeveloper@gmail.com

## License

Jessica is available under the MIT license. See the LICENSE file for more info.

# Changelog
[Open Changelog](/CHANGELOG.md)