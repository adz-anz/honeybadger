package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"golang.org/x/exp/slices"
)

type payload struct {
	Method   string            `description:"The method to execute (PUT, POST, etc)."`
	Path     string            `description:"The path, following the host."`
	Headers  map[string]string `description:"Additional headers to be sent with the request."`
	Body     []byte            `description:"Anything that needs to be sent as the body with the request."`
	Response interface{}       `description:"The response from the request."`
}

var (
	client = &http.Client{Timeout: 10 * time.Second}

	// buildID is set by CI
	buildID = "dev" // TODO: set this to the actual build ID

	// userAgent is what gets included in all http requests to the api
	userAgent string = fmt.Sprintf("%s/%s", appName, buildID)

	status200Codes = []int{
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

	validMethods = []string{
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

func (p *payload) GetResponse(printResponse bool) error {
	// Ensure that a valid method has been specified.
	if !slices.Contains(validMethods, p.Method) {
		errMsg := fmt.Sprintf("Invalid method %s", p.Method)
		return errors.New(errMsg)
	}

	// Ensure that the defined API Host is valid. This is particularly important
	// in the event that a custom API Host is specified.
	postURL, err := url.Parse(apiHost)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to parse URL %s: %s", apiHost, err)
		return errors.New(errMsg)
	}

	postURL.Path = p.Path

	req, err := http.NewRequest(p.Method, postURL.String(), bytes.NewBuffer(p.Body))
	if err != nil {
		errMsg := fmt.Sprintf("Failed to create request: %s", err)
		return errors.New(errMsg)
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/json") // TODO: Do we move this to the individual functions?
	req.Header.Add("X-Honeycomb-Team", configKey)

	// If any additional headers are specified, add them to the request.
	for key, val := range p.Headers {
		req.Header.Add(key, val)
	}

	// Output on dry run, skip execution of command.
	if dryRun {
		fmt.Printf("Would have sent the following request:\n---\n")
		reqOut, err := httputil.DumpRequest(req, true)
		if err != nil {
			errMsg := fmt.Sprintf("Failed to dump the request to a byte sequence: %s", err)
			return errors.New(errMsg)
		}
		fmt.Printf(string(reqOut))
		return nil
	}

	// Execute the request.
	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to execute the request: %s", err)
		return errors.New(errMsg)
	}
	defer resp.Body.Close()

	// Error if the response code is not a 2XX code.
	if !slices.Contains(status200Codes, resp.StatusCode) {
		errMsg := fmt.Sprintf("Failed with %d and message: %s", resp.StatusCode, resp.Body)
		return errors.New(errMsg)
	}

	err = json.NewDecoder(resp.Body).Decode(p.Response)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to decode response: %s", err)
		return errors.New(errMsg)
	}

	if printResponse {
		respMarshal, err := json.MarshalIndent(p.Response, "", "  ")
		if err != nil {
			return err
		}

		fmt.Printf(string(respMarshal))
	}

	return nil
}
