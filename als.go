package main

import (
	"bufio"
	"compress/gzip"
	"os"

	"github.com/MrSplidge/go-xmldom"
)

// Given a plugin DOM node, find the name of the track within which it appears.
func findTrackNameForNode(node *xmldom.Node) string {
	for {
		// Work up through the node hierarchy to the root.
		if node = node.Parent; node == nil {
			break
		} else {
			// Look for a Name element below a track element.
			const nameElementQuery = "Name[parent::MidiTrack|parent::AudioTrack|parent::MainTrack|parent::MasterTrack|parent::GroupTrack|parent::ReturnTrack]/EffectiveName"
			for _, match := range node.Query(nameElementQuery) {
				return match.GetAttributeValue("Value")
			}
		}
	}
	return ""
}

// Examine the contents of a CPR file to obtain version information and a mapping of track names to plugin names.
func examineALS(path string) *projectInformation {
	info := newProjectInformation(path)

	// Open the project file
	file, err := os.Open(path)
	if err != nil {
		info.logError(err.Error())
		return &info
	}
	defer file.Close()

	// Decompress the file
	gzipReader, err := gzip.NewReader(bufio.NewReader(file))
	if err != nil {
		info.logError(err.Error())
		return &info
	}
	defer gzipReader.Close()

	// Parse the decompressed file content into an XML DOM.
	dom, err := xmldom.Parse(gzipReader)
	if err != nil {
		info.logError(err.Error())
		return &info
	}

	// Extract the version number of the project.
	if version := dom.Root.GetAttributeValue("MinorVersion"); len(version) != 0 {
		info.version = version
	}

	// Extract a mapping of track names to plugins.
	processPluginInfo := func(queryPath string) {
		for _, node := range dom.Root.Query(queryPath) {
			if track := findTrackNameForNode(node); len(track) != 0 {
				if plugin := node.GetAttributeValue("Value"); len(plugin) != 0 {
					info.mapTrackToPlugin(plugin, track)
				}
			}
		}
	}

	// For both VST2 and VST3 plugins.
	processPluginInfo("//VstPluginInfo/PlugName")
	processPluginInfo("//Vst3PluginInfo/Name")

	return &info
}
