module github.com/osbuild/blueprint-schema/cmd/oas-bundler

go 1.23.0

toolchain go1.23.6

replace github.com/osbuild/blueprint-schema => ../..

require github.com/pb33f/libopenapi v0.21.8

require (
	github.com/bahlo/generic-list-go v0.2.0 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/speakeasy-api/jsonpath v0.6.1 // indirect
	github.com/wk8/go-ordered-map/v2 v2.1.9-0.20240815153524-6ea36470d1bd // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
