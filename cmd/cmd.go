package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"

	libcontext "github.com/Diaphteiros/kw/pluginlib/pkg/context"
	"github.com/Diaphteiros/kw/pluginlib/pkg/debug"
	libutils "github.com/Diaphteiros/kw/pluginlib/pkg/utils"

	"github.com/Diaphteiros/kw_kind/cmd/version"
	"github.com/Diaphteiros/kw_kind/pkg/config"
	"github.com/Diaphteiros/kw_kind/pkg/state"
)

var reload bool

var RootCmd = &cobra.Command{
	Use:               "kw_kind [<name>]",
	DisableAutoGenTag: true,
	Args:              cobra.RangeArgs(0, 1),
	Short:             "Switch to a local kind cluster",
	Long: `Switch to a local kind cluster.

This command sets the kubeconfig to the output of 'kind get kubeconfg --name <name>'.
No-op if the current kubeconfig is already set to the kubeconfig of the specified kind cluster.

The '--reload' flag causes the kubeconfig to be loaded even if it is the currently targeted cluster.
If the flag is set, the cluster name can be omitted to reload the current cluster's kubeconfig. This results in an error if the currently targeted kubeconfig was not set via this command.`,
	Run: func(cmd *cobra.Command, args []string) {
		// validate arguments
		clusterName := ""
		if len(args) == 0 {
			if !reload {
				libutils.Fatal(1, "cluster name needs to be provided if --reload is not set\n")
			}
		} else {
			clusterName = args[0]
		}

		// load context and config
		debug.Debug("Loading kubeswitcher context from environment")
		con, err := libcontext.NewContextFromEnv()
		if err != nil {
			libutils.Fatal(1, "error creating kubeswitcher context from environment (this is a plugin, did you run it as standalone?): %w\n", err)
		}
		debug.Debug("Kubeswitcher context loaded:\n%s", con.String())
		debug.Debug("Loading plugin configuration")
		cfg, err := config.LoadFromBytes([]byte(con.PluginConfig))
		if err != nil {
			libutils.Fatal(1, "error loading plugin configuration: %w\n", err)
		}
		debug.Debug("Plugin configuration loaded:\n%s", cfg.String())

		// load kind state
		kState := &state.KindState{}
		ok, err := kState.Load(con)
		if err != nil {
			libutils.Fatal(1, "error loading plugin state: %w\n", err)
		}
		if ok {
			debug.Debug("Plugin state loaded:\n%s", kState.String())
			if clusterName == "" {
				clusterName = kState.ClusterName
			} else if clusterName == kState.ClusterName && !reload {
				debug.Debug("Cluster name matches current cluster name and --reload flag is not set. Writing notification and exiting.")
				if err := con.WriteNotificationMessage(kState.Notification()); err != nil {
					libutils.Fatal(1, "%w\n", err)
				}
				return
			}
		} else {
			debug.Debug("Unable to load plugin state from kubeswitcher (either not found or current state is from a different plugin)")
			if clusterName == "" {
				libutils.Fatal(1, "Unable to reload kind cluster kubeconfig, because the current kubeconfig was not set via this subcommand.\nEither provide a kind cluster name or switch to a kind cluster first.\n")
			}
		}

		// prepare kind execution
		kindArgs := []string{"get", "kubeconfig", "--name", clusterName}
		bin := exec.Command(cfg.Binary, kindArgs...)
		// build command environment
		if bin.Env == nil {
			bin.Env = []string{}
		}
		bin.Env = append(bin.Env, os.Environ()...) // add current env vars

		// set channels
		errBuffer := libutils.NewWriteBuffer()
		outBuffer := libutils.NewWriteBuffer()
		bin.Stderr = errBuffer
		bin.Stdout = outBuffer
		bin.Stdin = cmd.InOrStdin()

		// run command
		debug.Debug("starting kind execution")
		if err := bin.Run(); err != nil {
			_ = outBuffer.Flush(cmd.OutOrStdout())
			_ = errBuffer.Flush(cmd.ErrOrStderr())
			libutils.Fatal(1, "error running kind: %w\n", err)
		}
		debug.Debug("finished kind execution")

		kcfgData := outBuffer.Data()
		// update state
		kState.ClusterName = clusterName
		if err := con.WriteKubeconfig(kcfgData, kState.Notification()); err != nil {
			libutils.Fatal(1, "%w\n", err)
		}
		if err := con.WriteId(kState.Id(con.CurrentPluginName)); err != nil {
			libutils.Fatal(1, "%w\n", err)
		}
		if err := con.WritePluginState(kState); err != nil {
			libutils.Fatal(1, "%w\n", err)
		}
	},
}

func init() {
	RootCmd.SetOut(os.Stdout)
	RootCmd.SetErr(os.Stderr)
	RootCmd.SetIn(os.Stdin)

	RootCmd.Flags().BoolVarP(&reload, "reload", "r", false, "Load the cluster kubeconfig even if it is the same cluster as the currently selected one. Can be used without cluster name to reload the current cluster's kubeconfig.")

	RootCmd.AddCommand(version.VersionCmd)
}
