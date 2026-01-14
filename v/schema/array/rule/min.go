// schema/array/rule/min.go
package rule

import (
	"github.com/leandroluk/go/v/internal/engine"
	"github.com/leandroluk/go/v/internal/ruleset"
)

func Min(code string, minimum int) ruleset.Rule[int] {
	return ruleset.New("min", func(actual int, context *engine.Context) (int, bool) {
		if actual < minimum {
			stop := context.AddIssueWithMeta(code, "too short", map[string]any{
				"min":    minimum,
				"actual": actual,
			})
			return actual, stop
		}
		return actual, false
	})
}
