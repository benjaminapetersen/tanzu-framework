// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"github.com/spf13/cobra"

	cliapi "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/cli/v1alpha1"
)

func newPostInstallCmd(desc *cliapi.PluginDescriptor) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "post-install",
		Short:        "Run post install configuration for a plugin",
		Long:         "Run post install configuration for a plugin",
		Hidden:       true,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			// invoke postInstall for the plugin
			return desc.PostInstallHook()
		},
	}

	return cmd
}
