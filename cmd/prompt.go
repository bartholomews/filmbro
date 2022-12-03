package cmd

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
	"strconv"
)

// https://dev.to/divrhino/building-an-interactive-cli-app-with-go-cobra-promptui-346n
type promptContent struct {
	errorMsg string
	label    string
}

func promptInt(pc promptContent) int {
	validate := func(input string) error {
		_, err := strconv.Atoi(input)
		if len(input) <= 0 || err != nil {
			return errors.New(pc.errorMsg)
		}
		return nil
	}

	str := makePrompt(pc, validate)
	res, _ := strconv.Atoi(str)
	return res
}

func promptString(pc promptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}
	return makePrompt(pc, validate)
}

func makePrompt(pc promptContent, validate promptui.ValidateFunc) string {
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}

func promptGetSelect(label string, items []string) string {
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label: label,
			Items: items,
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}
