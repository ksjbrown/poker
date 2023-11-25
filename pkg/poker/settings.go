package poker

type SettingsItem struct {
	BigBlind   int
	SmallBlind int
	Credits    int
}

type Settings []SettingsItem
