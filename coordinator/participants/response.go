package participants

// State
type CommitState int64

const (
	Ready CommitState = iota
	Committed
)

// Map of ip to CommitState
type CommitStateMap map[string]CommitState
