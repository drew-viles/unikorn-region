// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	externalRef0 "github.com/unikorn-cloud/core/pkg/openapi"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"AP//SBLGC8hFAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	pathPrefix := path.Dir(pathToFile)

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(pathPrefix, "https://raw.githubusercontent.com/unikorn-cloud/core/main/pkg/openapi/common.spec.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
