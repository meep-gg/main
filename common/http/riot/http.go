package http

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	. "meep.gg/log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Get(ctx context.Context, region, endpoint string, customParams ...map[string]interface{}) ([]byte, error) {
	baseURL := "https://" + region + ".api.riotgames.com"

	// Parse the base URL
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, status.Error(codes.Internal, "error parsing base URL")
	}

	// Parse the endpoint URL
	endpointURL, err := u.Parse(endpoint)
	if err != nil {
		return nil, status.Error(codes.Internal, "error parsing endpoint URL")
	}

	// Add the API key to the query parameters
	query := endpointURL.Query()
	if len(customParams) > 0 {
		for key, value := range customParams[0] {
			query.Set(key, fmt.Sprintf("%v", value))
		}
	}
	query.Set("api_key", API_KEY)
	endpointURL.RawQuery = query.Encode()

	// Make the HTTP request
	log.Printf("Requesting url " + Info(endpointURL.String()))
	req, _ := http.NewRequestWithContext(ctx, "GET", endpointURL.String(), nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer res.Body.Close()

	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		log.Printf("Response status code %v", Success(res.StatusCode))
	} else if res.StatusCode >= 300 && res.StatusCode <= 499 {
		log.Printf("Response status code %v", Warn(res.StatusCode))
	} else {
		log.Printf("Response status code %v", Error(res.StatusCode))
	}

	// Handle the response status code
	if res.StatusCode != 200 {
		return nil, handleStatusCode(res.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return body, nil
}

func handleStatusCode(code int) error {
	if err, ok := resCodes[code]; ok {
		return err
	}
	return status.Error(codes.Internal, "unknown error")
}

var resCodes map[int]error = map[int]error{
	400: status.Error(codes.NotFound, "not found"),
	401: status.Error(codes.Unauthenticated, "unauthorized"),
	403: status.Error(codes.PermissionDenied, "forbidden"),
	404: status.Error(codes.NotFound, "not found"),
	405: status.Error(codes.Unimplemented, "method not allowed"),
	415: status.Error(codes.InvalidArgument, "unsupported media type"),
	429: status.Error(codes.ResourceExhausted, "rate limit exceeded"),
	500: status.Error(codes.Internal, "internal server error"),
	502: status.Error(codes.Unavailable, "bad gateway"),
	503: status.Error(codes.Unavailable, "service unavailable"),
	504: status.Error(codes.Unavailable, "gateway timeout"),
}
