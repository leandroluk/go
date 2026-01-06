# Schemas

Referência rápida (com exemplos curtos).

## Princípios
- Por padrão: `missing` e `null` passam (não geram issue).
- `Required()` força presença.
- `Default(...)` aplica primeiro.
- Ordem (alto nível): default → missing/null → parse/coerce/type → constraints → custom.
- Coerção é **opt-in**: `WithCoerce(true)`.

---

## text

Usa `validator.Text()`.

### Meta
- `.Required()`
- `.IsDefault()` (se a string for zero-value, pula validações — exceto `Required`)
- `.Default(v)` / `.DefaultFunc(fn)`
- `.Custom(rule)`

### Tamanho e igualdade
- `.Len(n)`
- `.Min(n)` / `.Max(n)`
- `.Eq(v)` / `.Ne(v)`
- `.EqIgnoreCase(v)` / `.NeIgnoreCase(v)`
- `.OneOf(v1, v2, ...)`

### Conteúdo (substring/prefix/suffix)
- `.Contains(v)` / `.Excludes(v)`
- `.StartsWith(v)` / `.NotStartsWith(v)`
- `.EndsWith(v)` / `.NotEndsWith(v)`
- `.Lowercase()` / `.Uppercase()`

### Regex e formatos “web”
- `.Pattern(regex)` / `.PatternRegexp(*regexp.Regexp)`
- `.Email()`
- `.URL()`
- `.HTTPURL()`
- `.URI()`
- `.URNRFC2141()`

### IDs e rede
- `.UUID()` / `.UUID3()` / `.UUID4()` / `.UUID5()`
- `.IP()` / `.IPv4()` / `.IPv6()`
- `.CIDR()`
- `.MAC()`
- `.Hostname()` / `.FQDN()`
- `.Port()`

### Números e cores (string)
- `.Numeric()` (só dígitos)
- `.Number()` (float/exp, sem `NaN/Inf`)
- `.Hexadecimal()` (aceita `0x`)
- `.HexColor()` (`#fff` / `#ffffff`)
- `.RGB()` / `.RGBA()` / `.HSL()` / `.HSLA()` (CSS-like)

### Encoding e charset
- `.Base64()`
- `.Base64URL()` / `.Base64RawURL()`
- `.DataURI()`
- `.ASCII()` / `.PrintASCII()` / `.Multibyte()`

### “Fintech/docs”
- `.CreditCard()` / `.LuhnChecksum()`
- `.ISBN()` / `.ISBN10()` / `.ISBN13()`
- `.ISSN()`
- `.E164()`
- `.SemVer()`
- `.CVE()`

### Filesystem
- `.File()` / `.Dir()` (exigem existir)
- `.FilePath()` / `.DirPath()` (formato; `DirPath` exige separador final quando não existe)
- `.Image()` (arquivo existente + decode de imagem via stdlib)

### Hash digests
Aceitam **hex** (tamanho fixo) ou **base64** (tamanho certo).
- `.MD4()` / `.MD5()`
- `.SHA1()` / `.SHA224()` / `.SHA256()` / `.SHA384()` / `.SHA512()`
- `.SHA512_224()` / `.SHA512_256()`
- `.SHA3_224()` / `.SHA3_256()` / `.SHA3_384()` / `.SHA3_512()`
- `.RIPEMD160()`
- `.BLAKE2B_256()` / `.BLAKE2B_384()` / `.BLAKE2B_512()`
- `.BLAKE2S_256()`

### Coerce
- number/bool → string com `WithCoerce(true)`.

Exemplo rápido:

```go
s := validator.Text().Required().Email()

out, err := s.Validate("john@example.com")
```

---

## number

Usa `validator.NumberSchemaOf[T]()`.

- `.Required()`
- `.Min(n)` / `.Max(n)`
- `.OneOf(v1, v2, ...)`
- `.Default(v)` / `.DefaultFunc(fn)`
- `.Custom(rule)`
- Coerce base: string → number com `WithCoerce(true)`
- Flags comuns: trim space, underscore

```go
s := validator.NumberSchemaOf[int]().Min(0).Max(130)
```

---

## boolean

Usa `validator.Boolean()` (ou equivalente no seu build).

- `.Required()`
- `.Default(v)` / `.DefaultFunc(fn)`
- `.Custom(rule)`
- Coerce: string / 0 / 1 com `WithCoerce(true)` (conforme regras do schema)

---

## date

Usa `validator.Date()` (ou `DateSchemaOf[time.Time]()` dependendo do seu build).

- `.Required()`
- `.Min(t)` / `.Max(t)`
- `.Default(v)` / `.DefaultFunc(fn)`
- `.Custom(rule)`
- Parse usa `WithTimeLocation(...)` e `WithDateLayouts(...)`

---

## duration

Usa `validator.Duration()` (ou equivalente).

- `.Required()`
- `.Min(d)` / `.Max(d)`
- `.Default(v)` / `.DefaultFunc(fn)`
- `.Custom(rule)`
- String: `time.ParseDuration`
- AST number: nanosegundos
- Go number: seconds/millis só com `WithCoerce(true)` + flags

---

## array

Usa `validator.ArraySchemaOf[T]()`.

- `.Required()`
- `.Min(n)` / `.Max(n)`
- `.Default(v)` / `.DefaultFunc(fn)`
- `.Items(validator)`
- `.Unique()` / `.UniqueByHash(fn)` / `.UniqueByEqual(fn)`
- `.Custom(rule)`
- Coerce: singleton → array com `WithCoerce(true)` (quando suportado)

---

## record (map[string]V)

Usa `validator.RecordSchemaOf[V]()`.

- `.Required()`
- `.Min(n)` / `.Max(n)`
- `.Default(v)` / `.DefaultFunc(fn)`
- `.Keys(schema)` / `.KeysFunc(fn)`
- `.Values(validator)`
- `.Unique()` / `.UniqueByHash(fn)` / `.UniqueByEqual(fn)`
- `.Custom(rule)`

---

## object (struct)

Usa `validator.Object(...)`.

- `.Required()`
- `.Default(v)` / `.DefaultFunc(fn)`
- `.Field(&u.X, validator)` (nome via tag json primeiro)
- Rules no nível do objeto
- `.StructOnly()` / `.NoStructLevel()`
- Condicionais cross-field (required/excluded/skip), quando configuradas
