package enums

type RemindFrequency int

const (
	RemindFrequencyJustOneTime RemindFrequency = iota + 1
	RemindFrequencyEveryHour
	RemindFrequencyEveryDay
	RemindFrequencyEveryMonth
	RemindFrequencyEveryYear
)
