package replacer

type Replacer interface {
	Replace(marker string, newContent string) (string, error)
}
