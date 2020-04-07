package pods

import (
	"context"
	"fmt"

	"github.com/containers/libpod/cmd/podmanV2/registry"
	"github.com/containers/libpod/cmd/podmanV2/utils"
	"github.com/containers/libpod/pkg/domain/entities"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	pruneOptions = entities.PodPruneOptions{}
)

var (
	pruneDescription = fmt.Sprintf(`podman pod prune Removes all exited pods`)

	pruneCommand = &cobra.Command{
		Use:     "prune [flags]",
		Short:   "Remove all stopped pods and their containers",
		Long:    pruneDescription,
		RunE:    prune,
		Example: `podman pod prune`,
	}
)

func init() {
	registry.Commands = append(registry.Commands, registry.CliCommand{
		Mode:    []entities.EngineMode{entities.ABIMode, entities.TunnelMode},
		Command: pruneCommand,
		Parent:  podCmd,
	})
	flags := pruneCommand.Flags()
	flags.BoolVarP(&pruneOptions.Force, "force", "f", false, "Force removal of all running pods.  The default is false")
}

func prune(cmd *cobra.Command, args []string) error {
	var (
		errs utils.OutputErrors
	)
	if len(args) > 0 {
		return errors.Errorf("`%s` takes no arguments", cmd.CommandPath())
	}
	responses, err := registry.ContainerEngine().PodPrune(context.Background(), pruneOptions)

	if err != nil {
		return err
	}
	for _, r := range responses {
		if r.Err == nil {
			fmt.Println(r.Id)
		} else {
			errs = append(errs, r.Err)
		}
	}
	return errs.PrintErrors()
}
