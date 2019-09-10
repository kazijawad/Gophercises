// Package fpath implements HTTP handler functions for parsing
// paths and urls for maps and yaml
package fpath

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathURLs, err := parseYAML(yamlBytes)
	if err != nil {
		return nil, err
	}
	pathsToURLs := buildMap(pathURLs)
	return MapHandler(pathsToURLs, fallback), nil
}

func buildMap(pathURLs []pathURL) map[string]string {
	pathsToURLs := make(map[string]string)
	for _, pu := range pathURLs {
		pathsToURLs[pu.Path] = pu.URL
	}
	return pathsToURLs
}

func parseYAML(data []byte) ([]pathURL, error) {
	var pathURLs []pathURL
	err := yaml.Unmarshal(data, &pathURLs)
	if err != nil {
		return nil, err
	}
	return pathURLs, nil
}

type pathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
