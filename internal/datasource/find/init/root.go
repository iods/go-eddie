package init

//
///*
//Standups: A Go application for trying harder.
//Copyright (c) 2022 Rye Miller
//https://ryemiller.io
//*/
//
//import (
//	"fmt"
//	"github.com/AlecAivazis/survey/v2"
//	// "github.com/iods/go-standups/internal/util/common"
//	"github.com/spf13/cobra"
//)
//
//var ques = []*survey.Question{
//	{
//		Name: "directory",
//		Prompt: &survey.Confirm{
//			Message: "Do you want to use the default directory?:",
//		},
//		Validate: survey.Required,
//	},
//	{
//		Name: "projectkey",
//		Prompt: &survey.Input{
//			Message: "Add a key for the default project:",
//		},
//		Validate: survey.Required,
//	},
//}
//
//// StandupRootCmd replaces the Execute() function for the root command.
//func StandupRootCmd() *cobra.Command {
//	cmd := cobra.Command{
//		Use:           "projctl",
//		Short:         "A Go application for TL's wanting to try harder.",
//		SilenceErrors: true, // do i need this?
//		SilenceUsage:  true, // do i need this?
//		Long: `
//	 _____ _____ _____ _____ _____ _____ _____ _____
//	||s  |||t  |||a  |||n  |||d  |||u  |||p  |||s  ||
//	||___|||___|||___|||___|||___|||___|||___|||___||
//	|/___\|/___\|/___\|/___\|/___\|/___\|/___\|/___\|
//
//	       the cli tool for TLs 🖕🤬 v0.1.0
//
//`,
//		Run: func(cmd *cobra.Command, args []string) {
//			fmt.Println("Hello Standups CLI")
//		},
//	}
//
//	ans := struct {
//		Directory  bool
//		ProjectKey string
//	}{}
//
//	err := survey.Ask(ques, &ans)
//	if err != nil {
//		fmt.Println(err.Error())
//		return nil
//	}
//
//	fmt.Printf("Key %s will be using the %s directory.", ans.ProjectKey, ans.Directory)
//
//	return &cmd
//}
