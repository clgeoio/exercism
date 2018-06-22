//Package twofer responds to a share!
package twofer

import (
	"fmt"
	"strings"
)

// ShareWith responds to a person and gives one to them or you!
func ShareWith(who string) string {
	name := strings.TrimSpace(who)
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
