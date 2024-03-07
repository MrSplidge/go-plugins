package main

import (
	"os"
	"sort"
	"strings"

	"github.com/mattn/go-isatty"
)

// Information about a project.
type projectInformation struct {
	path             string
	version          string
	pluginToTrackMap map[string][]string
	trackToPluginMap map[string][]string
	errors           []string
}

// Create and initialize a new projectInformation instance.
func newProjectInformation(projectPath string) projectInformation {
	return projectInformation{
		path:             projectPath,
		version:          "<unknown version>",
		pluginToTrackMap: map[string][]string{},
		trackToPluginMap: map[string][]string{},
		errors:           []string{},
	}
}

// Log an error message against a project.
func (pi *projectInformation) logError(text string) {
	pi.errors = append(pi.errors, text)
}

// Map a plugin to a track, and the reverse (track to plugin).
func (pi *projectInformation) mapTrackToPlugin(plugin, track string) {
	pi.pluginToTrackMap[plugin] = append(pi.pluginToTrackMap[plugin], track)
	pi.trackToPluginMap[track] = append(pi.trackToPluginMap[track], plugin)
}

// Generate a coloured description of a project.
func (pi *projectInformation) ColouredString(projectColour, keyColour, valueColour, errorColour, resetColour string) string {
	type mapType int

	const (
		mapTracksToPlugins mapType = iota
		mapPluginsToTracks
	)

	var sb strings.Builder

	sb.WriteString(projectColour + "Project: " + pi.path + "\nVersion: " + pi.version + resetColour + "\n\n")

	for _, mt := range []mapType{mapPluginsToTracks, mapTracksToPlugins} {
		displayMap := pi.pluginToTrackMap
		switch mt {
		case mapTracksToPlugins:
			displayMap = pi.trackToPluginMap
			sb.WriteString("Track followed by a list of the plugins that it uses:\n")
		case mapPluginsToTracks:
			sb.WriteString("Plugin followed by a list of the tracks within which it appears:\n")
		}

		maximumKeyWidth := calculateMaximumKeyWidth(displayMap)

		iterateOverCISortedMap(displayMap, func(key string, value []string) {
			//iterateOverMap(displayMap, func(key string, value []string) {
			sb.WriteString("  " + keyColour + padStringToWidth(key, max(32, maximumKeyWidth+3), '.') + resetColour)
			if len(value) != 0 {
				// Sort the list of plugins or tracks
				sort.Slice(value, func(i, j int) bool {
					return strings.ToLower(value[i]) < strings.ToLower(value[j])
				})
				// Deduplicate them too.
				value = dedup(value)
				sb.WriteString("[ " + valueColour + strings.Join(value, resetColour+", "+valueColour) + resetColour + " ]\n")
			}
		})

		sb.WriteString("\n")
	}

	if len(pi.errors) != 0 {
		sb.WriteString(errorColour + "Errors:" + resetColour + "\n")
		for _, err := range pi.errors {
			sb.WriteString(" " + errorColour + err + resetColour + "\n")
		}
	}

	return sb.String()
}

// Generate a coloured or monochrome description for a project based on whether stdout is a terminal or a file.
func (pi *projectInformation) String() string {
	isAtty := isatty.IsTerminal(os.Stdout.Fd())
	if isAtty {
		return pi.ColouredString(yellow, green, cyan, red, reset)
	} else {
		return pi.ColouredString("", "", "", "", "")
	}
}
