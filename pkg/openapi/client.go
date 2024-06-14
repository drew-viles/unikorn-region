// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package openapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
	externalRef0 "github.com/unikorn-cloud/core/pkg/openapi"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegions request
	GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegions(ctx context.Context, organizationID OrganizationIDParameter, projectID ProjectIDParameter, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDExternalnetworks request
	GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDExternalnetworks(ctx context.Context, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDFlavors request
	GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDFlavors(ctx context.Context, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesWithBody request with any body
	PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesWithBody(ctx context.Context, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentities(ctx context.Context, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter, body PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDImages request
	GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDImages(ctx context.Context, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegions(ctx context.Context, organizationID OrganizationIDParameter, projectID ProjectIDParameter, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRequest(c.Server, organizationID, projectID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDExternalnetworks(ctx context.Context, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDExternalnetworksRequest(c.Server, organizationID, projectID, regionID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDFlavors(ctx context.Context, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDFlavorsRequest(c.Server, organizationID, projectID, regionID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesWithBody(ctx context.Context, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesRequestWithBody(c.Server, organizationID, projectID, regionID, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentities(ctx context.Context, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter, body PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesRequest(c.Server, organizationID, projectID, regionID, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDImages(ctx context.Context, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDImagesRequest(c.Server, organizationID, projectID, regionID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRequest generates requests for GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegions
func NewGetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRequest(server string, organizationID OrganizationIDParameter, projectID ProjectIDParameter) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "organizationID", runtime.ParamLocationPath, organizationID)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "projectID", runtime.ParamLocationPath, projectID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/organizations/%s/projects/%s/regions", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDExternalnetworksRequest generates requests for GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDExternalnetworks
func NewGetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDExternalnetworksRequest(server string, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "organizationID", runtime.ParamLocationPath, organizationID)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "projectID", runtime.ParamLocationPath, projectID)
	if err != nil {
		return nil, err
	}

	var pathParam2 string

	pathParam2, err = runtime.StyleParamWithLocation("simple", false, "regionID", runtime.ParamLocationPath, regionID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/organizations/%s/projects/%s/regions/%s/externalnetworks", pathParam0, pathParam1, pathParam2)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDFlavorsRequest generates requests for GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDFlavors
func NewGetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDFlavorsRequest(server string, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "organizationID", runtime.ParamLocationPath, organizationID)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "projectID", runtime.ParamLocationPath, projectID)
	if err != nil {
		return nil, err
	}

	var pathParam2 string

	pathParam2, err = runtime.StyleParamWithLocation("simple", false, "regionID", runtime.ParamLocationPath, regionID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/organizations/%s/projects/%s/regions/%s/flavors", pathParam0, pathParam1, pathParam2)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesRequest calls the generic PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentities builder with application/json body
func NewPostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesRequest(server string, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter, body PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesRequestWithBody(server, organizationID, projectID, regionID, "application/json", bodyReader)
}

// NewPostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesRequestWithBody generates requests for PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentities with any type of body
func NewPostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesRequestWithBody(server string, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "organizationID", runtime.ParamLocationPath, organizationID)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "projectID", runtime.ParamLocationPath, projectID)
	if err != nil {
		return nil, err
	}

	var pathParam2 string

	pathParam2, err = runtime.StyleParamWithLocation("simple", false, "regionID", runtime.ParamLocationPath, regionID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/organizations/%s/projects/%s/regions/%s/identities", pathParam0, pathParam1, pathParam2)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDImagesRequest generates requests for GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDImages
func NewGetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDImagesRequest(server string, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "organizationID", runtime.ParamLocationPath, organizationID)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "projectID", runtime.ParamLocationPath, projectID)
	if err != nil {
		return nil, err
	}

	var pathParam2 string

	pathParam2, err = runtime.StyleParamWithLocation("simple", false, "regionID", runtime.ParamLocationPath, regionID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/organizations/%s/projects/%s/regions/%s/images", pathParam0, pathParam1, pathParam2)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsWithResponse request
	GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsWithResponse(ctx context.Context, organizationID OrganizationIDParameter, projectID ProjectIDParameter, reqEditors ...RequestEditorFn) (*GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsResponse, error)

	// GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDExternalnetworksWithResponse request
	GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDExternalnetworksWithResponse(ctx context.Context, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter, reqEditors ...RequestEditorFn) (*GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDExternalnetworksResponse, error)

	// GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDFlavorsWithResponse request
	GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDFlavorsWithResponse(ctx context.Context, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter, reqEditors ...RequestEditorFn) (*GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDFlavorsResponse, error)

	// PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesWithBodyWithResponse request with any body
	PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesWithBodyWithResponse(ctx context.Context, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesResponse, error)

	PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesWithResponse(ctx context.Context, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter, body PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesResponse, error)

	// GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDImagesWithResponse request
	GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDImagesWithResponse(ctx context.Context, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter, reqEditors ...RequestEditorFn) (*GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDImagesResponse, error)
}

type GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *RegionsResponse
	JSON401      *externalRef0.UnauthorizedResponse
	JSON500      *externalRef0.InternalServerErrorResponse
}

// Status returns HTTPResponse.Status
func (r GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDExternalnetworksResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ExternalNetworksResponse
	JSON401      *externalRef0.UnauthorizedResponse
	JSON403      *externalRef0.ForbiddenResponse
	JSON404      *externalRef0.NotFoundResponse
	JSON500      *externalRef0.InternalServerErrorResponse
}

// Status returns HTTPResponse.Status
func (r GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDExternalnetworksResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDExternalnetworksResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDFlavorsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *FlavorsResponse
	JSON400      *externalRef0.BadRequestResponse
	JSON401      *externalRef0.UnauthorizedResponse
	JSON500      *externalRef0.InternalServerErrorResponse
}

// Status returns HTTPResponse.Status
func (r GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDFlavorsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDFlavorsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *IdentityResponse
	JSON400      *externalRef0.BadRequestResponse
	JSON401      *externalRef0.UnauthorizedResponse
	JSON403      *externalRef0.ForbiddenResponse
	JSON500      *externalRef0.InternalServerErrorResponse
}

// Status returns HTTPResponse.Status
func (r PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDImagesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ImagesResponse
	JSON400      *externalRef0.BadRequestResponse
	JSON401      *externalRef0.UnauthorizedResponse
	JSON500      *externalRef0.InternalServerErrorResponse
}

// Status returns HTTPResponse.Status
func (r GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDImagesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDImagesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsWithResponse request returning *GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsResponse
func (c *ClientWithResponses) GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsWithResponse(ctx context.Context, organizationID OrganizationIDParameter, projectID ProjectIDParameter, reqEditors ...RequestEditorFn) (*GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsResponse, error) {
	rsp, err := c.GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegions(ctx, organizationID, projectID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsResponse(rsp)
}

// GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDExternalnetworksWithResponse request returning *GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDExternalnetworksResponse
func (c *ClientWithResponses) GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDExternalnetworksWithResponse(ctx context.Context, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter, reqEditors ...RequestEditorFn) (*GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDExternalnetworksResponse, error) {
	rsp, err := c.GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDExternalnetworks(ctx, organizationID, projectID, regionID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDExternalnetworksResponse(rsp)
}

// GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDFlavorsWithResponse request returning *GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDFlavorsResponse
func (c *ClientWithResponses) GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDFlavorsWithResponse(ctx context.Context, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter, reqEditors ...RequestEditorFn) (*GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDFlavorsResponse, error) {
	rsp, err := c.GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDFlavors(ctx, organizationID, projectID, regionID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDFlavorsResponse(rsp)
}

// PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesWithBodyWithResponse request with arbitrary body returning *PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesResponse
func (c *ClientWithResponses) PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesWithBodyWithResponse(ctx context.Context, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesResponse, error) {
	rsp, err := c.PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesWithBody(ctx, organizationID, projectID, regionID, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesResponse(rsp)
}

func (c *ClientWithResponses) PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesWithResponse(ctx context.Context, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter, body PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesResponse, error) {
	rsp, err := c.PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentities(ctx, organizationID, projectID, regionID, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesResponse(rsp)
}

// GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDImagesWithResponse request returning *GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDImagesResponse
func (c *ClientWithResponses) GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDImagesWithResponse(ctx context.Context, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter, reqEditors ...RequestEditorFn) (*GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDImagesResponse, error) {
	rsp, err := c.GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDImages(ctx, organizationID, projectID, regionID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDImagesResponse(rsp)
}

// ParseGetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsResponse parses an HTTP response from a GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsWithResponse call
func ParseGetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsResponse(rsp *http.Response) (*GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest RegionsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest externalRef0.UnauthorizedResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef0.InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDExternalnetworksResponse parses an HTTP response from a GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDExternalnetworksWithResponse call
func ParseGetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDExternalnetworksResponse(rsp *http.Response) (*GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDExternalnetworksResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDExternalnetworksResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ExternalNetworksResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest externalRef0.UnauthorizedResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest externalRef0.ForbiddenResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest externalRef0.NotFoundResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef0.InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDFlavorsResponse parses an HTTP response from a GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDFlavorsWithResponse call
func ParseGetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDFlavorsResponse(rsp *http.Response) (*GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDFlavorsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDFlavorsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest FlavorsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest externalRef0.BadRequestResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest externalRef0.UnauthorizedResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef0.InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParsePostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesResponse parses an HTTP response from a PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesWithResponse call
func ParsePostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesResponse(rsp *http.Response) (*PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest IdentityResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest externalRef0.BadRequestResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest externalRef0.UnauthorizedResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest externalRef0.ForbiddenResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef0.InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDImagesResponse parses an HTTP response from a GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDImagesWithResponse call
func ParseGetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDImagesResponse(rsp *http.Response) (*GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDImagesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDImagesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ImagesResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest externalRef0.BadRequestResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest externalRef0.UnauthorizedResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef0.InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}
