package init

//
//import (
//	"github.com/iods/go-standups/internal/config"
//	"github.com/iods/go-standups/internal/errors"
//	"github.com/iods/go-standups/internal/util/parseutil"
//	"github.com/spf13/cobra"
//)
//
//type initParams struct {
//	directory  string
//	projectkey string
//}
//
//// NewStandupInit is the default init command to generate the standup application and docs.
//func NewStandupInit() *cobra.Command {
//	cmd := cobra.Command{
//		Use:   "init",
//		Short: "Initialize the standups configuration.",
//		Long:  "`init` initializes the standups configuration required for the CLI tool to perform properly.",
//	}
//
//	cmd.Flags().SortFlags = false // what is this?
//
//	// if existing, use the current code folder to build note structure
//	// else create Developer/ folder in the users home directory
//	// if env file is using a custom option 1, separate from auto-create and allow for name of dir (default it ~/Developer)
//	cmd.Flags().String("directory", "", "Enter the path to install standups directory if using a custom option.")
//	// for me this is Corra the company i work for
//	cmd.Flags().String("project", "", "Default project to be used.")
//
//	return &cmd
//}
//
//func parseFlags(flags parseutil.ParseFlag) *initParams {
//
//	directory, err := flags.GetString("directory")
//	if err != nil {
//		errors.ExitOnError(err)
//	}
//
//	projectkey, err := flags.GetString("project")
//	if err != nil {
//		errors.ExitOnError(err)
//	}
//
//	return &initParams{
//		directory:  directory,
//		projectkey: projectkey,
//	}
//}
//
//func initStandup(cmd *cobra.Command, _ []string) {
//	params := parseutil.ParseFlag(cmd.Flags())
//
//	c := config.ConfigBuilder(
//		&config.C)
//}
