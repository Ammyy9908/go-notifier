package injector

import (
	"fmt"
	"strings"
)

func InjectPlaceholders(text string, placeholders map[string]string) string {
	for key, value := range placeholders {
		placeholder := fmt.Sprintf("{{%s}}", key)
		text = strings.ReplaceAll(text, placeholder, value)
	}
	return text
}
