package region

import "fmt"

type Region string

const (
	English Region = "https://en.wikipedia.org/w/api.php"
	Hebrew  Region = "https://he.wikipedia.org/w/api.php"
)

// CustomRegion returns a custom region given it's wikipedia prefix
func CustomRegion(prefix string) Region {
	return Region(fmt.Sprintf("https://%s.wikipedia.org/w/api.php", prefix))
}
