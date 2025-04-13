package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/phnthnhnm/go-pokedex/internal/pokecache"
)

var cache = pokecache.NewCache(5 * time.Minute)

func FetchLocationAreas(url string) (*LocationAreaResponse, error) {
	if cachedData, found := cache.Get(url); found {
		var data LocationAreaResponse
		if err := json.Unmarshal(cachedData, &data); err != nil {
			return nil, fmt.Errorf("failed to unmarshal cached data: %w", err)
		}
		return &data, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch location areas: %w", err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var data LocationAreaResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	cache.Add(url, body)
	return &data, nil
}

func FetchLocationAreaDetails(url string) (*LocationAreaDetailsResponse, error) {
	if cachedData, found := cache.Get(url); found {
		var data LocationAreaDetailsResponse
		if err := json.Unmarshal(cachedData, &data); err != nil {
			return nil, fmt.Errorf("failed to unmarshal cached data: %w", err)
		}
		return &data, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch location area details: %w", err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var data LocationAreaDetailsResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	cache.Add(url, body)
	return &data, nil
}
