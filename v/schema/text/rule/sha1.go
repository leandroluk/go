// schema/text/rule/sha1.go
package rule

import "github.com/leandroluk/go/v/internal/ruleset"

func SHA1(code string) ruleset.Rule[string] {
	return digestRule(code, "invalid sha1", 20)
}
