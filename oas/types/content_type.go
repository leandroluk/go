// oas/types/content_type.go
package types

type ContentType string

const (
	ContentType_ApplicationJson ContentType = "application/json"
	ContentType_TextPlain       ContentType = "text/plain"
	ContentType_TextHtml        ContentType = "text/html"
	ContentType_TextXml         ContentType = "text/xml"
	ContentType_TextCsv         ContentType = "text/csv"
	ContentType_ImageJpeg       ContentType = "image/jpeg"
	ContentType_ImagePng        ContentType = "image/png"
	ContentType_ImageGif        ContentType = "image/gif"
	ContentType_ImageSvg        ContentType = "image/svg+xml"
	ContentType_ImageWebp       ContentType = "image/webp"
)
