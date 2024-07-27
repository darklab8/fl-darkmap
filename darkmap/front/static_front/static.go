package static_front

import (
	_ "embed"
)

//go:embed custom.js
var CustomJS string

//go:embed common.css
var CommonCSS string

//go:embed custom.css
var CustomCSS string

//go:embed favicon.ico
var FaviconIco string
