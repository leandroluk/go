# Migração do go-playground/validator para golidator

No `go-playground/validator` você descreve validação em **tag**. No golidator você descreve em **Go**.
Não tem “tag mágica”. Tem schema explícito. É mais verboso no começo e muito menos misterioso depois.

## Diferenças que mais pegam gente
- `missing` e `null` **passam por padrão**. Use `.Required()` quando quiser presença.
- Coerção é **opt-in**. `WithCoerce(true)` liga o básico; o resto depende de flags.
- Você escolhe o schema certo para o tipo (text/number/date/etc), não joga tudo no mesmo saco.
- Em arrays/maps, `dive` não existe: você chama `.Items(...)` / `.Values(...)` / `.Keys(...)` e acabou.

## Mapa mental (como traduzir tags)
1) Ache o **tipo real**: string → `Text()`, int/float → `NumberSchemaOf[T]()`, bool → `Boolean()`, time → `Date()`, duration → `Duration()`,
slices → `ArraySchemaOf[T]()`, map → `RecordSchemaOf[T]()`, struct → `Object(...)`.
2) Converta **presence** (`required`/`omitempty`) e **limites** (`min/max/len`).
3) Converta **validações específicas** (email, uuid, ip, etc) para métodos do schema `text`.
4) Para coisa “esperta” (cross-field / condicionais), suba pro `object` (rule no nível do objeto) ou faça `Custom(rule)`.

---

## Tabela: tags “meta/controle”

| Tag (validator.v10) | Equivalente no golidator | Notas |
|---|---|---|
| `required` | `.Required()` | Presence (não é “não-zero”) |
| `omitempty` | `WithOmitZero(true)` | Só vale para entrada refletida (struct/map). Em JSON, “missing” já é missing |
| `omitnil` | (use ponteiro + ausência) | Se seu campo é `*T`, `nil` vira `null`/missing dependendo da origem |
| `omitzero` | `WithOmitZero(true)` | Igual ao `omitempty` na prática |
| `isdefault` | `.IsDefault()` | “se for zero-value, pula validações” (exceto `.Required()`) |
| `structonly` | `object.StructOnly()` | Só regras do objeto |
| `nostructlevel` | `object.NoStructLevel()` | Só fields |
| `dive` | `array.Items(...)` / `record.Values(...)` | Você decide o validator do item/valor |
| `keys,endkeys` | `record.Keys(...)` | Valida chave antes do valor |
| `|` (OR) | `combinator.AnyOf(...)` | Você monta os schemas e passa para o combinator |

## Tabela: tags comuns de string

| Tag | Equivalente no golidator (text) | Exemplo |
|---|---|---|
| `min=3` | `.Min(3)` | `Text().Min(3)` |
| `max=50` | `.Max(50)` | `Text().Max(50)` |
| `len=8` | `.Len(8)` | `Text().Len(8)` |
| `eq=foo` | `.Eq("foo")` | `Text().Eq("foo")` |
| `ne=foo` | `.Ne("foo")` | `Text().Ne("foo")` |
| `startswith=ab` | `.StartsWith("ab")` | `Text().StartsWith("ab")` |
| `endswith=ab` | `.EndsWith("ab")` | `Text().EndsWith("ab")` |
| `contains=ab` | `.Contains("ab")` | `Text().Contains("ab")` |
| `excludes=ab` | `.Excludes("ab")` | `Text().Excludes("ab")` |
| `lowercase` | `.Lowercase()` | `Text().Lowercase()` |
| `uppercase` | `.Uppercase()` | `Text().Uppercase()` |
| `oneof=a b c` | `.OneOf("a","b","c")` | `Text().OneOf(...)` |
| `email` | `.Email()` | `Text().Email()` |
| `url` | `.URL()` | `Text().URL()` |
| `http_url` | `.HTTPURL()` | `Text().HTTPURL()` |
| `uri` | `.URI()` | `Text().URI()` |
| `urn_rfc2141` | `.URNRFC2141()` | `Text().URNRFC2141()` |
| `uuid` / `uuid3` / `uuid4` / `uuid5` | `.UUID()` / `.UUID3()` / ... | `Text().UUID4()` |
| `ip` / `ipv4` / `ipv6` | `.IP()` / `.IPv4()` / `.IPv6()` | `Text().IPv4()` |
| `cidr` | `.CIDR()` | `Text().CIDR()` |
| `mac` | `.MAC()` | `Text().MAC()` |
| `hostname` / `fqdn` | `.Hostname()` / `.FQDN()` | `Text().FQDN()` |
| `port` | `.Port()` | `Text().Port()` |
| `numeric` | `.Numeric()` | só dígitos |
| `number` | `.Number()` | float/exp, sem `NaN/Inf` |
| `hexadecimal` | `.Hexadecimal()` | aceita `0x` |
| `hexcolor` | `.HexColor()` | `#fff` / `#ffffff` |
| `rgb` / `rgba` / `hsl` / `hsla` | `.RGB()` / `.RGBA()` / `.HSL()` / `.HSLA()` | CSS-like |
| `base64` | `.Base64()` | base64 padrão |
| `base64url` | `.Base64URL()` | base64 URL |
| `base64rawurl` | `.Base64RawURL()` | sem padding |
| `datauri` | `.DataURI()` | `data:...` |
| `ascii` / `printascii` / `multibyte` | `.ASCII()` / `.PrintASCII()` / `.Multibyte()` | |
| `credit_card` | `.CreditCard()` | Luhn + tamanho |
| `luhn_checksum` | `.LuhnChecksum()` | valida Luhn |
| `isbn` / `isbn10` / `isbn13` | `.ISBN()` / `.ISBN10()` / `.ISBN13()` | |
| `issn` | `.ISSN()` | |
| `e164` | `.E164()` | `+14155552671` |
| `semver` | `.SemVer()` | SemVer 2.0.0 |
| `cve` | `.CVE()` | `CVE-YYYY-NNNN...` |
| `file` / `dir` | `.File()` / `.Dir()` | exige existir |
| `filepath` / `dirpath` | `.FilePath()` / `.DirPath()` | formato; `DirPath` exige sep final quando não existe |
| `image` | `.Image()` | exige existir + decode de imagem (stdlib) |
| `md4` / `md5` / `sha1` / ... | `.MD5()` / `.SHA256()` / ... | aceita hex ou base64 (tamanho certo) |

## Tabela: tags comuns de número

| Tag | Equivalente | Notas |
|---|---|---|
| `gte=10` | `NumberSchemaOf[T]().Min(10)` | `T` define se é int/float |
| `lte=10` | `NumberSchemaOf[T]().Max(10)` | |
| `oneof=1 2 3` | `.OneOf(1,2,3)` | no schema numérico |
| `min=3`/`max=10` (slice/map) | `ArraySchemaOf[T]().Min(3)` / `RecordSchemaOf[V]().Max(10)` | depende do schema container |

---

## Exemplos

### 1) Antes: tags

```go
type User struct {
	Name string   `json:"name" validate:"required,min=3,max=50"`
	Age  int      `json:"age"  validate:"gte=0,lte=130"`
	Role string   `json:"role" validate:"oneof=admin user guest"`
	Tags []string `json:"tags" validate:"max=10,dive,min=2,max=20"`
}
```

### 1) Depois: schema

```go
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Role string `json:"role"`
	Tags []string `json:"tags"`
}

golidator.Object(func(u *User, s *golidator.ObjectSchema[User]) {
	s.Field(&u.Name, func(ctx *golidator.Context, v any) (any, bool) {
		return golidator.Text().Required().Min(3).Max(50).ValidateAny(v, ctx.Options)
	})

	s.Field(&u.Age, func(ctx *golidator.Context, v any) (any, bool) {
		return golidator.NumberSchemaOf[int]().Min(0).Max(130).ValidateAny(v, ctx.Options)
	})

	s.Field(&u.Role, func(ctx *golidator.Context, v any) (any, bool) {
		return golidator.Text().OneOf("admin", "user", "guest").ValidateAny(v, ctx.Options)
	})

	s.Field(&u.Tags, func(ctx *golidator.Context, v any) (any, bool) {
		itemSchema := golidator.Text().Min(2).Max(20)
		return golidator.ArraySchemaOf[string]().
			Max(10).
			Items(func(itemCtx *golidator.Context, item any) (any, bool) {
				return itemSchema.ValidateAny(item, itemCtx.Options)
			}).
			ValidateAny(v, ctx.Options)
	})
})
```

### 2) `dive` + `keys/endkeys` (map com validação de chave e valor)

```go
// validate:"min=1,max=20,keys,alphanum,endkeys,required,dive,required,email"
emailsByKey := golidator.RecordSchemaOf[string]().
	Min(1).
	Max(20).
	Keys(golidator.Text().Pattern("^[a-zA-Z0-9]+$")).
	Values(func(ctx *golidator.Context, v any) (any, bool) {
		return golidator.Text().Required().Email().ValidateAny(v, ctx.Options)
	})
```

### 3) OR (`|`) com combinator

```go
s := golidator.AnyOf(
	golidator.Text().Email(),
	golidator.Text().UUID(),
)

out, err := s.Validate("john@example.com")
```

---

## Quando não dá pra mapear 1:1 (e o que fazer)
Algumas tags são conveniência/opinião (ex: `required_without_all`, `excluded_unless`, `eqfield`, `gtcsfield`).
No golidator isso vira uma destas opções:

- **Regra no nível do objeto** (cross-field): você valida o struct inteiro e cria issue no path que quiser.
- **Condicionais** no `object` (RequiredIf/ExcludedIf/SkipUnless etc), quando disponíveis no seu build.
- `Custom(rule)` bem explícita (sem adivinhação).

A regra é simples: se a validação depende de **mais de um campo**, ela não pertence ao schema do campo isolado.
