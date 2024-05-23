package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"golang.org/x/exp/slices"
)

type Payload struct {
	Method  string            `description:"The method to execute (PUT, POST, etc)."`
	Path    string            `description:"The path, following the host."`
	Headers map[string]string `description:"Additional headers to be sent with the request."`
	Body    []byte            `description:"Anything that needs to be sent as the body with the request."`
}

var (
	client = http.Client{}

	// BuildID is set by CI
	BuildID string = "dev" // TODO: set this to the actual build ID

	// UserAgent is what gets included in all http requests to the api
	UserAgent string = fmt.Sprintf("%s/%s", appName, BuildID)

	Status200Codes = []int{
		http.StatusOK,
		http.StatusCreated,
		http.StatusAccepted,
		http.StatusNonAuthoritativeInfo,
		http.StatusNoContent,
		http.StatusResetContent,
		http.StatusPartialContent,
		http.StatusMultiStatus,
		http.StatusAlreadyReported,
		http.StatusIMUsed,
	}

	ValidMethods = []string{
		http.MethodGet,
		http.MethodHead,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodConnect,
		http.MethodOptions,
		http.MethodTrace,
	}
)

func (p *Payload) Execute() error {
	if !slices.Contains(ValidMethods, p.Method) {
		errMsg := fmt.Sprintf("Invalid method %s", p.Method)
		return errors.New(errMsg)
	}
	postURL, err := url.Parse(apiHost)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to parse URL %s", apiHost)
		return errors.New(errMsg)
	}
	postURL.Path = p.Path

	req, _ := http.NewRequest(p.Method, postURL.String(), bytes.NewBuffer(p.Body))
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Content-Type", "application/json") // TODO: Do we move this to the individual functions?
	req.Header.Add("X-Honeycomb-Team", configKey)
	for key, val := range p.Headers {
		req.Header.Add(key, val)
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if !slices.Contains(Status200Codes, resp.StatusCode) {
		errMsg := fmt.Sprintf("Failed with %d and message: %s", resp.StatusCode, body)
		return errors.New(errMsg)
	}
	fmt.Println(string(body))
	return nil
}
