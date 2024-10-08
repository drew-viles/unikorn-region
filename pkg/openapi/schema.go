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
