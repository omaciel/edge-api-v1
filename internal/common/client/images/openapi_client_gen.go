// Package images provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version (devel) DO NOT EDIT.
package images

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
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
	// GetImages request
	GetImages(ctx context.Context, params *GetImagesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateImage request with any body
	CreateImageWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateImage(ctx context.Context, body CreateImageJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteImage request
	DeleteImage(ctx context.Context, imageId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetImage request
	GetImage(ctx context.Context, imageId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateImage request with any body
	UpdateImageWithBody(ctx context.Context, imageId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateImage(ctx context.Context, imageId string, body UpdateImageJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteImagesImageIdUpdate request
	DeleteImagesImageIdUpdate(ctx context.Context, imageId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateNewVersion request with any body
	CreateNewVersionWithBody(ctx context.Context, imageId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateNewVersion(ctx context.Context, imageId string, body CreateNewVersionJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetImages(ctx context.Context, params *GetImagesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetImagesRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateImageWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateImageRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateImage(ctx context.Context, body CreateImageJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateImageRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteImage(ctx context.Context, imageId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteImageRequest(c.Server, imageId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetImage(ctx context.Context, imageId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetImageRequest(c.Server, imageId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateImageWithBody(ctx context.Context, imageId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateImageRequestWithBody(c.Server, imageId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateImage(ctx context.Context, imageId string, body UpdateImageJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateImageRequest(c.Server, imageId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteImagesImageIdUpdate(ctx context.Context, imageId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteImagesImageIdUpdateRequest(c.Server, imageId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateNewVersionWithBody(ctx context.Context, imageId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateNewVersionRequestWithBody(c.Server, imageId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateNewVersion(ctx context.Context, imageId string, body CreateNewVersionJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateNewVersionRequest(c.Server, imageId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetImagesRequest generates requests for GetImages
func NewGetImagesRequest(server string, params *GetImagesParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/images")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.SortBy != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "sort_by", runtime.ParamLocationQuery, *params.SortBy); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Name != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "name", runtime.ParamLocationQuery, *params.Name); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Status != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "status", runtime.ParamLocationQuery, *params.Status); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreateImageRequest calls the generic CreateImage builder with application/json body
func NewCreateImageRequest(server string, body CreateImageJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateImageRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateImageRequestWithBody generates requests for CreateImage with any type of body
func NewCreateImageRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/images")
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

// NewDeleteImageRequest generates requests for DeleteImage
func NewDeleteImageRequest(server string, imageId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "imageId", runtime.ParamLocationPath, imageId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/images/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetImageRequest generates requests for GetImage
func NewGetImageRequest(server string, imageId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "imageId", runtime.ParamLocationPath, imageId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/images/%s", pathParam0)
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

// NewUpdateImageRequest calls the generic UpdateImage builder with application/json body
func NewUpdateImageRequest(server string, imageId string, body UpdateImageJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateImageRequestWithBody(server, imageId, "application/json", bodyReader)
}

// NewUpdateImageRequestWithBody generates requests for UpdateImage with any type of body
func NewUpdateImageRequestWithBody(server string, imageId string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "imageId", runtime.ParamLocationPath, imageId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/images/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewDeleteImagesImageIdUpdateRequest generates requests for DeleteImagesImageIdUpdate
func NewDeleteImagesImageIdUpdateRequest(server string, imageId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "imageId", runtime.ParamLocationPath, imageId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/images/%s/update", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreateNewVersionRequest calls the generic CreateNewVersion builder with application/json body
func NewCreateNewVersionRequest(server string, imageId string, body CreateNewVersionJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateNewVersionRequestWithBody(server, imageId, "application/json", bodyReader)
}

// NewCreateNewVersionRequestWithBody generates requests for CreateNewVersion with any type of body
func NewCreateNewVersionRequestWithBody(server string, imageId string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "imageId", runtime.ParamLocationPath, imageId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/images/%s/update", pathParam0)
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
	// GetImages request
	GetImagesWithResponse(ctx context.Context, params *GetImagesParams, reqEditors ...RequestEditorFn) (*GetImagesResponse, error)

	// CreateImage request with any body
	CreateImageWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateImageResponse, error)

	CreateImageWithResponse(ctx context.Context, body CreateImageJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateImageResponse, error)

	// DeleteImage request
	DeleteImageWithResponse(ctx context.Context, imageId string, reqEditors ...RequestEditorFn) (*DeleteImageResponse, error)

	// GetImage request
	GetImageWithResponse(ctx context.Context, imageId string, reqEditors ...RequestEditorFn) (*GetImageResponse, error)

	// UpdateImage request with any body
	UpdateImageWithBodyWithResponse(ctx context.Context, imageId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateImageResponse, error)

	UpdateImageWithResponse(ctx context.Context, imageId string, body UpdateImageJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateImageResponse, error)

	// DeleteImagesImageIdUpdate request
	DeleteImagesImageIdUpdateWithResponse(ctx context.Context, imageId string, reqEditors ...RequestEditorFn) (*DeleteImagesImageIdUpdateResponse, error)

	// CreateNewVersion request with any body
	CreateNewVersionWithBodyWithResponse(ctx context.Context, imageId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateNewVersionResponse, error)

	CreateNewVersionWithResponse(ctx context.Context, imageId string, body CreateNewVersionJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateNewVersionResponse, error)
}

type GetImagesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Count *int             `json:"count,omitempty"`
		Items *[]ImageResponse `json:"items,omitempty"`
	}
	JSON404 *Error
}

// Status returns HTTPResponse.Status
func (r GetImagesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetImagesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateImageResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *ImageResponse
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r CreateImageResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateImageResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteImageResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r DeleteImageResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteImageResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetImageResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ImageResponse
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r GetImageResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetImageResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateImageResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r UpdateImageResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateImageResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteImagesImageIdUpdateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r DeleteImagesImageIdUpdateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteImagesImageIdUpdateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateNewVersionResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r CreateNewVersionResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateNewVersionResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetImagesWithResponse request returning *GetImagesResponse
func (c *ClientWithResponses) GetImagesWithResponse(ctx context.Context, params *GetImagesParams, reqEditors ...RequestEditorFn) (*GetImagesResponse, error) {
	rsp, err := c.GetImages(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetImagesResponse(rsp)
}

// CreateImageWithBodyWithResponse request with arbitrary body returning *CreateImageResponse
func (c *ClientWithResponses) CreateImageWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateImageResponse, error) {
	rsp, err := c.CreateImageWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateImageResponse(rsp)
}

func (c *ClientWithResponses) CreateImageWithResponse(ctx context.Context, body CreateImageJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateImageResponse, error) {
	rsp, err := c.CreateImage(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateImageResponse(rsp)
}

// DeleteImageWithResponse request returning *DeleteImageResponse
func (c *ClientWithResponses) DeleteImageWithResponse(ctx context.Context, imageId string, reqEditors ...RequestEditorFn) (*DeleteImageResponse, error) {
	rsp, err := c.DeleteImage(ctx, imageId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteImageResponse(rsp)
}

// GetImageWithResponse request returning *GetImageResponse
func (c *ClientWithResponses) GetImageWithResponse(ctx context.Context, imageId string, reqEditors ...RequestEditorFn) (*GetImageResponse, error) {
	rsp, err := c.GetImage(ctx, imageId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetImageResponse(rsp)
}

// UpdateImageWithBodyWithResponse request with arbitrary body returning *UpdateImageResponse
func (c *ClientWithResponses) UpdateImageWithBodyWithResponse(ctx context.Context, imageId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateImageResponse, error) {
	rsp, err := c.UpdateImageWithBody(ctx, imageId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateImageResponse(rsp)
}

func (c *ClientWithResponses) UpdateImageWithResponse(ctx context.Context, imageId string, body UpdateImageJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateImageResponse, error) {
	rsp, err := c.UpdateImage(ctx, imageId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateImageResponse(rsp)
}

// DeleteImagesImageIdUpdateWithResponse request returning *DeleteImagesImageIdUpdateResponse
func (c *ClientWithResponses) DeleteImagesImageIdUpdateWithResponse(ctx context.Context, imageId string, reqEditors ...RequestEditorFn) (*DeleteImagesImageIdUpdateResponse, error) {
	rsp, err := c.DeleteImagesImageIdUpdate(ctx, imageId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteImagesImageIdUpdateResponse(rsp)
}

// CreateNewVersionWithBodyWithResponse request with arbitrary body returning *CreateNewVersionResponse
func (c *ClientWithResponses) CreateNewVersionWithBodyWithResponse(ctx context.Context, imageId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateNewVersionResponse, error) {
	rsp, err := c.CreateNewVersionWithBody(ctx, imageId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateNewVersionResponse(rsp)
}

func (c *ClientWithResponses) CreateNewVersionWithResponse(ctx context.Context, imageId string, body CreateNewVersionJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateNewVersionResponse, error) {
	rsp, err := c.CreateNewVersion(ctx, imageId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateNewVersionResponse(rsp)
}

// ParseGetImagesResponse parses an HTTP response from a GetImagesWithResponse call
func ParseGetImagesResponse(rsp *http.Response) (*GetImagesResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetImagesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Count *int             `json:"count,omitempty"`
			Items *[]ImageResponse `json:"items,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParseCreateImageResponse parses an HTTP response from a CreateImageWithResponse call
func ParseCreateImageResponse(rsp *http.Response) (*CreateImageResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateImageResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest ImageResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseDeleteImageResponse parses an HTTP response from a DeleteImageWithResponse call
func ParseDeleteImageResponse(rsp *http.Response) (*DeleteImageResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteImageResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseGetImageResponse parses an HTTP response from a GetImageWithResponse call
func ParseGetImageResponse(rsp *http.Response) (*GetImageResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetImageResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ImageResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseUpdateImageResponse parses an HTTP response from a UpdateImageWithResponse call
func ParseUpdateImageResponse(rsp *http.Response) (*UpdateImageResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateImageResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseDeleteImagesImageIdUpdateResponse parses an HTTP response from a DeleteImagesImageIdUpdateWithResponse call
func ParseDeleteImagesImageIdUpdateResponse(rsp *http.Response) (*DeleteImagesImageIdUpdateResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteImagesImageIdUpdateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseCreateNewVersionResponse parses an HTTP response from a CreateNewVersionWithResponse call
func ParseCreateNewVersionResponse(rsp *http.Response) (*CreateNewVersionResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateNewVersionResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}
