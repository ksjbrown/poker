package hands

// Score is a comparison metric for Hand structs.
// Three points are compared:
//   - The Major score, or HandRank.
//     This is equivalent to the type of the hand.
//   - The Minor score, which compares equivalent hand types by their significant cards in the hand.
//   - The Minor score, which compares the remaining cards that do not make up the hand.
//     These are often referred to as "kicker" cards.
type Score struct {
	Major Rank
	Minor int
	Micro int
}

// CompareTo returns the result of a comparison of this hand to an argument hand.
// A positive value means that the calling HandScore is better.
// A negative value means that the argument HandScore is better.
// A magnitude of 3 indicates the HandRank of the winning hand is better than the losing hand.
// A magnitude of 2 indicates the HandRanks of the hands are equal, but the winning has a better variant of the HandRank.
// A magnitude of 2 indicates that the HandRanks are equal, but the winning hand has better kicker cards.
// A return value of 0 means that one hand is not better than the other.
func (hs Score) Compare(other Score) int {
	panic("not yet implemented")
}
