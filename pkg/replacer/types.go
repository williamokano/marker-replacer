package replacer

type Replacer interface {
	Replace(marker string, newContent string) error
	GetString() string
}
