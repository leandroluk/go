// oas/builder/content_builder.go
package builder

import "github.com/leandroluk/go/oas/types"

type ContentBuilder struct {
	content map[string]*types.MediaType
}

func newContentBuilder(target map[string]*types.MediaType) *ContentBuilder {
	if target == nil {
		target = map[string]*types.MediaType{}
	}
	return &ContentBuilder{content: target}
}

func (b *ContentBuilder) Content() map[string]*types.MediaType {
	return b.content
}

func (b *ContentBuilder) Type(contentType string, build func(target *MediaTypeBuilder)) *ContentBuilder {
	media := &types.MediaType{}
	mb := &MediaTypeBuilder{mediaType: media}
	if build != nil {
		build(mb)
	}
	b.content[contentType] = media
	return b
}

func (b *ContentBuilder) JSON(build func(target *MediaTypeBuilder)) *ContentBuilder {
	return b.Type(string(types.ContentType_ApplicationJson), build)
}
