package daemon

import (
	"slices"

	"github.com/nning/juist/ipc"
	"github.com/nning/juist/types"
)

// tunnelFigures is rs with each peer's last handshake and traffic, read from
// the data plane now: the tunnel check that fills the rest of rs runs only
// every 30s. A macOS patch.
func (d *Daemon) tunnelFigures(rs []ipc.PeerReach) []ipc.PeerReach {
	d.mu.Lock()
	peers := d.installed
	d.mu.Unlock()
	ids := make(map[types.Identity]uint64, len(peers))
	for _, p := range peers {
		ids[p.id] = p.peer.ID
	}
	tunnels := d.manager.Tunnels()
	out := slices.Clone(rs)
	for i := range out {
		id, ok := ids[out[i].Identity]
		if !ok {
			continue
		}
		t := tunnels[id]
		if !t.LastHandshake.IsZero() {
			out[i].Handshake = t.LastHandshake.Unix()
		}
		out[i].RxBytes, out[i].TxBytes = t.RxBytes, t.TxBytes
	}
	return out
}
