package jellyfin

import (
	"errors"
	"net/url"
	"os"
	"strings"

	"github.com/gofiber/fiber/v3"
)

// BackendURL returns the Jellyfin URL the Pelagica backend should use for
// server-to-server requests.
//
// Resolution order:
//  1. jellyfin_backend_url query parameter
//  2. JELLYFIN_BACKEND_URL environment variable
//  3. jellyfin_url query parameter

func BackendURL(c fiber.Ctx) (string, error) {
	jellyfinURL := strings.TrimSpace(c.Query("jellyfin_backend_url"))
	if jellyfinURL == "" {
		jellyfinURL = strings.TrimSpace(os.Getenv("JELLYFIN_BACKEND_URL"))
	}
	if jellyfinURL == "" {
		jellyfinURL = strings.TrimSpace(c.Query("jellyfin_url"))
	}

	if jellyfinURL == "" {
		return "", errors.New("missing jellyfin_url query parameter")
	}

	if _, err := url.ParseRequestURI(jellyfinURL); err != nil {
		return "", errors.New("invalid Jellyfin URL")
	}

	return strings.TrimRight(jellyfinURL, "/"), nil
}
