package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dinuta/estuary-agent-go/src/command"
	"github.com/dinuta/estuary-agent-go/src/models"
	"github.com/dinuta/estuary-agent-go/src/utils"
	"github.com/dinuta/runcmd/constants"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var (
	// Used for flags.
	cid       string
	arguments string

	rootCmd = &cobra.Command{
		Use:     "runcmd",
		Short:   "A CLI runner which runs your system commands and outputs the results in JSON",
		Long:    `The CLI runs your system commands (linux/windows) sequentially and it outputs the results both in stdout and in JSON file`,
		Version: "v1.0.0",
		Args: func(cmd *cobra.Command, args []string) error {
			if viper.GetString("args") == "" {
				return errors.New("requires --args command argument. E.g. --args=\"dir;;echo 5\"")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			//viper args
			utils.AppendFile("command_info_logger.txt", cmd.Use+" --cid="+cid+" --args=\""+arguments+"\"\n")
			utils.CreateDir(constants.CMD_BACKGROUND_DIR)
			utils.CreateDir(constants.CMD_BACKGROUND_STREAMS_DIR)
			outputJsonFile := fmt.Sprintf(constants.CMD_BACKGROUND_OUTPUT, cid)
			cd := getCommandDescription()
			cdJson, err := json.Marshal(cd)
			if err != nil {
				panic(fmt.Sprintf(err.Error()))
			}
			utils.WriteFile(outputJsonFile, cdJson)

			command := command.NewCommand(cid, outputJsonFile)
			cd = command.RunCommands(strings.Split(arguments, ";;"))
			cdJson, err = json.Marshal(cd)
			if err != nil {
				panic(err.Error())
			}

			fmt.Print(fmt.Sprint(string(cdJson)))
		},
	}
)

func getCommandDescription() *models.CommandDescription {
	cd := models.NewCommandDescription()

	cd.SetPid(os.Getpid())
	cd.SetId(cid)
	cd.SetCommands(map[string]*models.CommandStatus{})

	return cd
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cid, "cid", "c", "none", "E.g. --cid 2 [Not mandatory]")
	rootCmd.PersistentFlags().StringVarP(&arguments, "args", "a", "",
		"E.g. --args \"ls;;pwd;;echo2\" [Mandatory]")
	rootCmd.MarkFlagRequired("args")

	viper.BindPFlag("cid", rootCmd.PersistentFlags().Lookup("cid"))
	viper.BindPFlag("args", rootCmd.PersistentFlags().Lookup("args"))
}
