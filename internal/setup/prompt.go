package setup

import (
	"fmt"

	"github.com/manifoldco/promptui"
)


func PromptInstall() bool {
	stub := promptStub()
	s, err := stub.Run()
	if err != nil {
		fmt.Printf("prompt failed for %v", err)
	}

	if s == "y" {
		fmt.Printf("\nThat's it. Now run `eddie --help` to get started.\n")
		return true
	}

	fmt.Printf("\nThat's it. Now run `eddie --help` to get started.\n")
	return false
}


func promptStub() *promptui.Prompt {

	templates := &promptui.PromptTemplates{
		Prompt: "{{ . }}",
		Valid: "{{ . | bold }}",
		Invalid: "{{ . | bold }}",
		Success: "{{ . | bold }}",
	}

	return &promptui.Prompt{
		Label: " [1] Would you like to stub the database?",
		Templates: templates,
		IsConfirm: true,
	}
}

func promptWeight() *promptui.Prompt {

	return &promptui.Prompt{}
}

func promptHeight() *promptui.Prompt {

	return &promptui.Prompt{}
}

func promptSex() *promptui.Prompt {

	return &promptui.Prompt{}
}