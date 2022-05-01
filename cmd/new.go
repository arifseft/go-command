/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
    "errors"
    "os"

	"github.com/spf13/cobra"
	"github.com/manifoldco/promptui"
	"github.com/arifseft/go-command/data"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
        createNewNote()
	},
}

type promptContent struct {
    errorMsg string
    label    string
}

func init() {
	noteCmd.AddCommand(newCmd)
}

func promptGetInput(pc promptContent) string {
    validate := func(input string) error {
        if len(input) <= 0 {
            return errors.New(pc.errorMsg)
        }
        return nil
    }

    template := &promptui.PromptTemplates{
        Prompt: "{{ . }}",
        Valid: "{{ . | green }}",
        Invalid: "{{ . | red }}",
        Success: "{{ . | bold }}",
    }

    prompt := promptui.Prompt{
        Label: pc.label,
        Templates: template,
        Validate: validate,
    }

    result, err := prompt.Run()
    if err != nil {
        fmt.Println("Prompt failed %v\n", err)
        os.Exit(1)
    }

    fmt.Printf("Input: %s\n", result)
    return result
}

func promptGetSelect(pc promptContent) string {
    items := []string{"animal", "food", "person", "object"}
    index := -1

    var result string
    var err error

    for index < 0 {
        prompt := promptui.SelectWithAdd{
            Label: pc.label,
            Items: items,
            AddLabel: "Other",
        }

        index, result, err = prompt.Run()

        if index == -1 {
            items = append(items, result)
        }
    }

    if err != nil {
        fmt.Printf("Prompt failed %v\n", err)
    }

    fmt.Printf("Input: %s\n", result)
    return result
}

func createNewNote() {
    wordPromptContent := promptContent{
        "Please provide a word",
        "What word would you like to make note of?",
    }
    word := promptGetInput(wordPromptContent)

    definitionPromptContent := promptContent{
        "Please provide a definition",
        fmt.Sprintf("What is the definition of %s?", word),
    }
    definition := promptGetInput(definitionPromptContent)

    categoryPromptContent := promptContent{
        "Please provide a category",
        fmt.Sprintf("What category does %s belong to?", word),
    }
    category := promptGetSelect(categoryPromptContent)

    data.InsertNote(word, definition, category)
}
