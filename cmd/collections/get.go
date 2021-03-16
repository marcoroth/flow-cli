package collections

import (
	"github.com/onflow/flow-cli/cmd"
	"github.com/onflow/flow-cli/flow/lib"
	"github.com/onflow/flow-cli/flow/services"
	"github.com/psiemens/sconfig"
	"github.com/spf13/cobra"
)

type flagsCollections struct {
}

type cmdGet struct {
	cmd   *cobra.Command
	flags flagsCollections
}

// NewGetCmd creates new get command
func NewGetCmd() cmd.Command {
	return &cmdGet{
		cmd: &cobra.Command{
			Use:   "get <collection_id>",
			Short: "Get collection info",
			Args:  cobra.ExactArgs(1),
		},
	}
}

// Run collection command
func (s *cmdGet) Run(
	cmd *cobra.Command,
	args []string,
	project *lib.Project,
	services *services.Services,
) (cmd.Result, error) {
	collection, err := services.Collections.Get(args[0])
	return &CollectionResult{collection}, err
}

// GetFlags for collection
func (s *cmdGet) GetFlags() *sconfig.Config {
	return sconfig.New(&s.flags)
}

// GetCmd get command
func (s *cmdGet) GetCmd() *cobra.Command {
	return s.cmd
}
