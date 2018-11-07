package cmd

import (
	"github.com/absfs/afero"
	"github.com/spf13/cobra"
)

var ensureCommand = &cobra.Command{
	Use:   "ensure",
	Short: "Ensure that the actual state matches the target state",
	Run:   makeRun(runEnsureCommand),
}

func init() {
	rootCommand.AddCommand(ensureCommand)
}

func runEnsureCommand(fs afero.Fs, command *cobra.Command, args []string) error {
	targetState, err := config.getTargetState(fs)
	if err != nil {
		return err
	}
	return targetState.Ensure(fs, config.TargetDir, getUmask(), config.getDefaultActuator(fs))
}