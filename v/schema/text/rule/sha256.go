// schema/text/rule/sha256.go
package rule

import "github.com/leandroluk/go/v/internal/ruleset"

func SHA256(code string) ruleset.Rule[string] {
	return digestRule(code, "invalid sha256", 32)
}
