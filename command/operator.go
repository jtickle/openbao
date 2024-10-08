// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package command

import (
	"strings"

	"github.com/hashicorp/cli"
)

var _ cli.Command = (*OperatorCommand)(nil)

type OperatorCommand struct {
	*BaseCommand
}

func (c *OperatorCommand) Synopsis() string {
	return "Perform operator-specific tasks"
}

func (c *OperatorCommand) Help() string {
	helpText := `
Usage: bao operator <subcommand> [options] [args]

  This command groups subcommands for operators interacting with Vault. Most
  users will not need to interact with these commands. Here are a few examples
  of the operator commands:

  Initialize a new Vault cluster:

      $ bao operator init

  Force a Vault to resign leadership in a cluster:

      $ bao operator step-down

  Rotate Vault's underlying encryption key:

      $ bao operator rotate

  Please see the individual subcommand help for detailed usage information.
`

	return strings.TrimSpace(helpText)
}

func (c *OperatorCommand) Run(args []string) int {
	return cli.RunResultHelp
}
