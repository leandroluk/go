// schema/text/rule/sha384.go
package rule

import "github.com/leandroluk/go/v/internal/ruleset"

func SHA384(code string) ruleset.Rule[string] {
	return digestRule(code, "invalid sha384", 48)
}
