package main

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

	return func(wri http.ResponseWriter, req *http.Request) {
		urlpath := req.URL.Path
		if p, stat := pathsToUrls[urlpath]; stat {
			http.Redirect(wri, req, p, http.StatusFound)
			return
		}
		fallback.ServeHTTP(wri, req)
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
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.

func YAMLHandler(yamlbytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yamlbytes)
	if err != nil {
		return nil, err
	}
	pathMap := buildMapfromYaml(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

type UrlPath struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func parseYAML(yamlbytes []byte) ([]UrlPath, error) {
	var pathofurls []UrlPath
	err := yaml.Unmarshal(yamlbytes, &pathofurls)
	if err != nil {
		return nil, err
	}
	return pathofurls, nil
}

func buildMapfromYaml(parsedYaml []UrlPath) map[string]string {
	urlpaths := make(map[string]string)
	for _, eachone := range parsedYaml {
		urlpaths[eachone.Path] = eachone.URL
	}
	return urlpaths
}
