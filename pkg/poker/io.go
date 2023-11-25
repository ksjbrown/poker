package poker

import "time"

// PlayerIO defines the methods expected by the server logic,
// so that it can query info from the player, and report if info is invalid or acceptable, updates of state, etc.
type PlayerIO interface {
	onActionRequested(timeout time.Duration) PlayerAction
	onActionAccepted(pta PlayerAction)
	onActionRejected(timeout time.Duration, pta PlayerAction) PlayerAction
	onActionTimedOut()
	onGameStateUpdated(p *Game)
}
