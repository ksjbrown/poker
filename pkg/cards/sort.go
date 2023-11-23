package cards

// SortAlgorithm represents a method of sorting cards
// Implementations give the difference between left and right based on some sorting metric
type SortAlgorithm func(left Card, right Card) int

// StandardSort will sort the cards in this hand according to their face rank, i.e. Ace first, Two second, and so on, up til King
func StandardSort(left Card, right Card) int {
	return int(left.Rank - right.Rank)
}

// AceHighSort will sort cards, with Ace being considered the highest card, i.e. Two first, Three second, ... Queen, King, then Ace
func AceHighSort(left Card, right Card) int {
	return left.Score() - right.Score()
}
