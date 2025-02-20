// SPDX-FileCopyrightText: Copyright 2023 The Minder Authors
// SPDX-License-Identifier: Apache-2.0

// Package invite provides the auth invite command for the minder CLI.
package invite

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/mindersec/minder/cmd/cli/app"
	"github.com/mindersec/minder/internal/util"
	"github.com/mindersec/minder/internal/util/cli"
	"github.com/mindersec/minder/internal/util/cli/table"
	"github.com/mindersec/minder/internal/util/cli/table/layouts"
	minderv1 "github.com/mindersec/minder/pkg/api/protobuf/go/minder/v1"
)

// inviteListCmd represents the list command
var inviteListCmd = &cobra.Command{
	Use:   "list",
	Short: "List pending invitations",
	Long:  `List shows all pending invitations for the current minder user`,
	RunE:  cli.GRPCClientWrapRunE(inviteListCommand),
}

// inviteListCommand is the whoami subcommand
func inviteListCommand(ctx context.Context, cmd *cobra.Command, _ []string, conn *grpc.ClientConn) error {
	client := minderv1.NewUserServiceClient(conn)

	// No longer print usage on returned error, since we've parsed our inputs
	// See https://github.com/spf13/cobra/issues/340#issuecomment-374617413
	cmd.SilenceUsage = true
	format := viper.GetString("output")

	res, err := client.ListInvitations(ctx, &minderv1.ListInvitationsRequest{})
	if err != nil {
		return cli.MessageAndError("Error listing invitations", err)
	}

	switch format {
	case app.JSON:
		out, err := util.GetJsonFromProto(res)
		if err != nil {
			return cli.MessageAndError("Error getting json from proto", err)
		}
		cmd.Println(out)
	case app.YAML:
		out, err := util.GetYamlFromProto(res)
		if err != nil {
			return cli.MessageAndError("Error getting yaml from proto", err)
		}
		cmd.Println(out)
	case app.Table:
		if len(res.Invitations) == 0 {
			cmd.Println("No pending invitations")
			return nil
		}
		t := table.New(table.Simple, layouts.Default, []string{"Sponsor", "Project", "Role", "Expires", "Code"})
		for _, v := range res.Invitations {
			t.AddRow(v.SponsorDisplay, v.Project, v.Role, v.ExpiresAt.AsTime().Format(time.RFC3339), v.Code)
		}
		t.Render()
	default:
		return fmt.Errorf("unsupported output format: %s", format)
	}
	return nil
}

func init() {
	inviteCmd.AddCommand(inviteListCmd)
	inviteListCmd.Flags().StringP("output", "o", app.Table,
		fmt.Sprintf("Output format (one of %s)", strings.Join(app.SupportedOutputFormats(), ",")))
}
