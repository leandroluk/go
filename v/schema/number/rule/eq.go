// schema/number/rule/eq.go
package rule

import (
	"github.com/leandroluk/go/v/internal/engine"
	"github.com/leandroluk/go/v/internal/ruleset"
	"github.com/leandroluk/go/v/internal/types"
	"github.com/leandroluk/go/v/schema/number/util"
)

func Eq[N types.Number](code string, expected N) ruleset.Rule[N] {
	return ruleset.New("", func(actual N, context *engine.Context) (N, bool) {
		if util.NumberEqual[N](actual, expected) {
			return actual, false
		}

		stop := context.AddIssueWithMeta(code, "must be equal", map[string]any{
			"expected": expected,
			"actual":   actual,
		})
		return actual, stop
	})
}
