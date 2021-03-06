package question

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/daskioff/jessica/utils/slices"
)

func AskQuestion(question string, answerIsRequired bool) (answer string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("%s ", question)
	for {
		answer, _ = reader.ReadString('\n')
		answer = strings.TrimSpace(answer)

		if !answerIsRequired || len(answer) > 0 {
			return
		}
	}
}

func AskQuestionWithAnswers(question string, answers []string) (answer string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("%s ", question)
	for {
		answer, _ = reader.ReadString('\n')
		answer = strings.TrimSpace(answer)

		if len(answer) > 0 && slices.Contains(answers, answer) {
			return
		}
	}
}

func AskQuestionWithBoolAnswer(question string) bool {
	reader := bufio.NewReader(os.Stdin)

	positiveAnswers := []string{"y", "yes", "true"}
	negativeAnswers := []string{"n", "no", "false"}

	question = question + "(" + strings.Join(positiveAnswers, "/") + "|" + strings.Join(negativeAnswers, "/") + ")"

	fmt.Printf("%s ", question)
	for {
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)
		answer = strings.ToLower(answer)

		if slices.Contains(positiveAnswers, answer) {
			return true
		} else if slices.Contains(negativeAnswers, answer) {
			return false
		}
	}
}

func AskQuestionWithChooseFileAnswer(question string, fileExt string) string {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	xcodeprojFiles := []string{}
	for _, f := range files {
		extension := filepath.Ext(f.Name())
		if extension == fileExt {
			xcodeprojFiles = append(xcodeprojFiles, f.Name())
		}
	}

	count := len(xcodeprojFiles)
	if count == 0 {
		return ""
	}
	if count == 1 {
		return xcodeprojFiles[0]
	}

	answers := []string{}
	variants := []string{}
	for i, fileName := range xcodeprojFiles {
		answers = append(answers, strconv.Itoa(i))
		variants = append(variants, strconv.Itoa(i)+". "+fileName)
	}

	question = question + "\n" + strings.Join(variants, "\n") + "\nIndex: "
	answer := AskQuestionWithAnswers(question, answers)

	index, err := strconv.Atoi(answer)
	if err != nil {
		index = 0
	}

	return xcodeprojFiles[index]
}

func AskQuestionWithChooseFolderAnswer(question string) string {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	folders := []string{".", "..Skip (command)"}
	for _, f := range files {
		if f.IsDir() {
			folders = append(folders, f.Name())
		}
	}

	answers := []string{}
	variants := []string{}
	for i, folderName := range folders {
		answers = append(answers, strconv.Itoa(i))
		variants = append(variants, strconv.Itoa(i)+". "+folderName)
	}

	question = question + "\n" + strings.Join(variants, "\n") + "\nIndex: "
	answer := AskQuestionWithAnswers(question, answers)

	index, err := strconv.Atoi(answer)
	if err != nil {
		index = 0
	}

	if index == 1 {
		return ""
	}

	return folders[index]
}
