package state

import (
	"fmt"

	libcontext "github.com/Diaphteiros/kw/pluginlib/pkg/context"
	"github.com/Diaphteiros/kw/pluginlib/pkg/debug"
	liberrors "github.com/Diaphteiros/kw/pluginlib/pkg/errors"
	libstate "github.com/Diaphteiros/kw/pluginlib/pkg/state"
)

type KindState struct {
	ClusterName string `json:"clusterName"`
}

func (gs *KindState) String() string {
	// faster than yaml.Marshal, needs to be adapted if fields are added
	return fmt.Sprintf("clusterName: %s", gs.ClusterName)
}

func (gs *KindState) Id(pluginName string) string {
	return fmt.Sprintf("%s:%s", pluginName, gs.ClusterName)
}

func (gs *KindState) Notification() string {
	return fmt.Sprintf("Switched to kind cluster '%s'.", gs.ClusterName)
}

// Load fills the receiver state object with the data from the kubeswitcher state.
// The first return value is true if any state was actually loaded, false otherwise.
func (gs *KindState) Load(con *libcontext.Context) (bool, error) {
	debug.Debug("Loading gardenctl state from kubeswitcher state")
	ts, err := libstate.LoadTypedState[*KindState](con.GenericStatePath, con.PluginStatePath, con.CurrentPluginName)
	if err != nil {
		return false, liberrors.IgnoreStateFromAnotherPluginError(fmt.Errorf("error loading kubeswitcher state: %w", err))
	}
	if ts == nil || ts.PluginState == nil {
		return false, nil
	}
	gs.ClusterName = ts.PluginState.ClusterName
	return true, nil
}
