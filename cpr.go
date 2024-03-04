package main

import (
	"fmt"
	"os"
)

// Represents a position within a slice. This could almost be a simple slice, but the position enables spans to be ordered.
type span struct {
	position int
	bytes    []byte
}

// Advance a span by [count] bytes. The span cannot be advanced beyond the end of its slice, at which point the [empty] method will return true.
func (s span) advance(count int) span {
	return span{
		position: min(s.position+count, len(s.bytes)),
		bytes:    s.bytes,
	}
}

// Return the byte at the specified offset from the span's current position.
func (s span) at(offset int) byte {
	return s.bytes[s.position+offset]
}

// Return a span that begins [count] bytes from the current position.
func (s span) subspan(count int) span {
	return span{
		position: s.position,
		bytes:    s.bytes[:s.position+count],
	}
}

// Return a byte slice that contains [count] bytes from the current position. This may be useful with dumpHex.
func (s span) subslice(count int) []byte {
	begin := s.position
	end := min(len(s.bytes), s.position+count)
	return s.bytes[begin:end]
}

// Return a string formed from [count] bytes that follow the current position.
func (s span) substring(count int) string {
	return string(s.bytes[s.position : s.position+count])
}

// Check whether a span's position has reached the end of its slice.
func (s span) empty() bool {
	return s.position == len(s.bytes)
}

// Check whether a span has at least [count] bytes remaining between its current position and the end of the slice.
func (s span) hasBytes(count int) bool {
	return len(s.bytes) >= s.position+count
}

// Read a FOURCC (four byte) marker from a span's current position. Returns a span following the marker, the FOURCC value, and any error encountered.
func readFOURCC(s span) (span, int, error) {
	return readDWORD(s)
}

// Read a DWORD (four byte) value from a span's current position. Returns a span following the value, the value, and any error encountered.
func readDWORD(s span) (span, int, error) {
	if !s.hasBytes(4) {
		return s, 0, fmt.Errorf("can't read DWORD")
	} else {
		v := uint(s.at(0)) << 24
		v |= uint(s.at(1)) << 16
		v |= uint(s.at(2)) << 8
		v |= uint(s.at(3))
		return s.advance(4), int(v), nil
	}
}

// Read a WORD (two byte) value from a span's current position. Returns a span following the value, the value, and any error encountered.
func readWORD(s span) (span, int, error) {
	if !s.hasBytes(2) {
		return s, 0, fmt.Errorf("can't read WORD")
	} else {
		v := uint(s.at(0)) << 8
		v |= uint(s.at(1))
		return s.advance(2), int(v), nil
	}
}

// Read a string from a span's current position. Returns a span following the string, the string, and any error encountered.
// The string may include a null terminator and an encoding byte order marker. Use encodedString to decode the string.
func readString(s span) (span, string, error) {
	textSpan, length, error := readDWORD(s)
	if error != nil {
		return s, "", fmt.Errorf("reading string: %s", error)
	}
	text := textSpan.substring(length)
	return textSpan.advance(length), text, nil
}

// Read a string from a span's current position. Returns a span following the string, the string, and any error encountered.
// The string is assumed to be null-terminated, but the returned string doesn't include the null terminator.
func readNullTerminatedString(s span) (span, string, error) {
	textSpan, length, error := readDWORD(s)
	if error != nil {
		return s, "", fmt.Errorf("reading string: %s", error)
	}
	text := textSpan.substring(length - 1)
	return textSpan.advance(length), text, nil
}

// Find a string within a span. Returns a span following the string. If the string is not found, an empty span is returned.
func findString(s span, searchText string) span {
	searchTextLength := len(searchText) + 1
	for !s.empty() {
		_, foundLength, error := readDWORD(s)
		if error != nil {
			break
		}
		if foundLength == searchTextLength {
			s3, foundText, error := readNullTerminatedString(s)
			if error != nil {
				break
			}
			if foundText == searchText {
				return s3
			}
		}
		s = s.advance(1)
	}
	return span{}
}

// Converts a FOURCC value to a string.
func fourccToString(fourcc int) string {
	return fmt.Sprintf("%c%c%c%c", 0xff&(fourcc>>24), 0xff&(fourcc>>16), 0xff&(fourcc>>8), 0xff&(fourcc>>0))
}

// Chunk types that can be mentioned in the ROOT chunk.
type chunkType int

const (
	CT_Version chunkType = iota
	CT_Arrangement
	CT_ComputerGuid
	CT_Devices
	CT_WindowLayouts
	CT_ProjectLayouts
	CT_Unknown
)

// Examines the content of a ROOT chunk to discover the type of the following ARCH chunk.
// Returns the ARCH chunk type and any error encountered.
func scanRootChunk(s span) (chunkType, error) {
	_, text, err := readString(s)
	if err != nil {
		return CT_Unknown, fmt.Errorf("error reading ROOT chunk: %s", err)
	}
	switch text {
	case "Version":
		return CT_Version, nil
	case "Arrangement1":
		return CT_Arrangement, nil
	case "ComputerGuid":
		return CT_ComputerGuid, nil
	case "Devices":
		return CT_Devices, nil
	case "WindowLayouts":
		return CT_WindowLayouts, nil
	case "ProjectLayouts":
		return CT_ProjectLayouts, nil
	default:
		return CT_Unknown, nil
	}
}

// Checks whether a value is one that can appear before a serialized object name.
func isObjectIntro(value int) bool {
	return value == 0xfffffffe || value == 0xffffffff
}

// Scans the contents of a Version ARCH chunk, and returns the project version string and any error encountered. If no version is found, an empty string is returned.
func scanArchChunk_Version(s span) string {
	//fmt.Printf("scanArchChunk_Version\n")
	//dumpHex(s.subslice(256))

	for !s.empty() {
		s2, objectIntro, _ := readDWORD(s)
		if isObjectIntro(objectIntro) {
			var objectType string
			s2, objectType, _ = readNullTerminatedString(s2)

			switch objectType {
			case "CmObject":
				s2, _, _ = readWORD(s2) // ignore
			case "PAppVersion":
				s2, _, _ = readWORD(s2)                 // ignore
				s2, _, _ = readDWORD(s2)                // ignore
				s2, _, _ = readNullTerminatedString(s2) // ignore
				var version string
				_, version, _ = readNullTerminatedString(s2)
				return version
			}
		}
		s = s2
	}
	return ""
}

// Associates a named thing (a plugin or track) with a location.
type namedLocation struct {
	name     string
	location span
}

// Scans an Arrangement or Devices ARCH chunk for information about plugins and the tracks on which they appear.
// Results, if any, are stored in the projectInformation object passed in.
func scanArchChunk(s span, pi *projectInformation) {
	//fmt.Printf("scanArchChunk\n")
	//dumpHex(s.bytes[:256])

	// Find tracks
	tracks := findTracks(s)
	maybeUnused(tracks)

	// Find plugins
	plugins := findPlugins(s)
	maybeUnused(plugins)

	// Associate tracks with plugins
	for _, pluginLocation := range plugins {
		// Find the name of the track immediately prior to the plugin location.
		trackName := ""
		for _, trackLocation := range tracks {
			if trackLocation.location.position < pluginLocation.location.position {
				trackName = trackLocation.name
			} else {
				break
			}
		}

		if len(trackName) != 0 {
			pi.mapTrackToPlugin(pluginLocation.name, trackName)
		}
	}
}

// Finds and returns the named locations of specified tracks in the given span.
func findTracks(s span) []namedLocation {
	tracks := []namedLocation{}

	for _, trackType := range []string{"VST Multitrack", "Output Channels"} {
		for track := findString(s, trackType); !track.empty(); track = findString(track, trackType) {
			next, _, _ := readDWORD(track) // ignore
			next, _, _ = readDWORD(next)   // ignore
			next, _, _ = readDWORD(next)   // ignore
			next, text, _ := readNullTerminatedString(next)
			if text == "RuntimeID" {
				next, _, _ = readWORD(next)  // ignore
				next, _, _ = readDWORD(next) // ignore
				next, _, _ = readDWORD(next) // ignore
				next, text, _ = readNullTerminatedString(next)
				if text == "Name" {
					next, _, _ = readWORD(next)  // ignore
					next, _, _ = readWORD(next)  // ignore
					next, _, _ = readDWORD(next) // ignore
					next, text, _ = readNullTerminatedString(next)
					if text == "String" {
						next, _, _ = readWORD(next) // ignore
						_, text, _ = readString(next)
						tracks = append(tracks, namedLocation{name: decodeString(text), location: track})
						//fmt.Printf("Track: %s location %d\n", decodeString(text), track.position)
					}
				}
			}
		}
	}
	return tracks
}

// Finds and returns the named locations of specified plugins in the given span.
func findPlugins(s span) []namedLocation {
	var plugins []namedLocation

	for _, pluginType := range []string{"VstCtrlInternalEffect"} {
		for plugin := findString(s, pluginType); !plugin.empty(); plugin = findString(plugin, pluginType) {
			next, text, _ := readNullTerminatedString(plugin)
			if text == "Plugin" {
				next, _, _ = readDWORD(next) // ignore
				next, _, _ = readWORD(next)  // ignore
				next, _, _ = readWORD(next)  // ignore
				next, text, _ = readNullTerminatedString(next)
				if text == "Plugin UID" {
					next, _, _ = readDWORD(next) // ignore
					next, _, _ = readDWORD(next) // ignore
					next, text, _ = readNullTerminatedString(next)
					if text == "GUID" {
						next, _, _ = readWORD(next)                 // ignore
						next, _, _ = readNullTerminatedString(next) // ignore
						next, text, _ = readNullTerminatedString(next)

						var pluginName string
						if text == "Plugin Name" {
							next, _, _ = readWORD(next) // ignore
							next, pluginName, _ = readString(next)
						}
						next, text, _ = readNullTerminatedString(next)
						if text == "Original Plugin Name" {
							next, _, _ = readWORD(next) // ignore
							_, pluginName, _ = readString(next)
						}
						plugins = append(plugins, namedLocation{name: decodeString(pluginName), location: plugin})
						//fmt.Printf("plugin: %s location %d\n", decodeString(pluginName), plugin.position)
					}
				}
			}
		}
	}
	return plugins
}

// Convert a string that might contain a null terminator followed by an encoding byte order marker into a string.
func decodeString(text string) string {
	for index, ch := range text {
		if ch == 0 {
			// Chop off the null terminator and anything following it
			return text[:index]
		}
	}
	return text
}

const (
	riffFourcc = 0x52494646 // 'RIFF'
	rootFourcc = 0x524f4f54 // 'ROOT'
	archFourcc = 0x41524348 // 'ARCH'
)

// Examine the contents of a CPR file to obtain version information and a mapping of track names to plugin names.
func examineCPR(projectPath string) *projectInformation {
	info := newProjectInformation(projectPath)

	content, error := os.ReadFile(projectPath)
	if error != nil {
		info.logError(fmt.Sprintf("opening file %s", error.Error()))
	} else {
		s := span{bytes: content}
		//dumpHex(s.bytes[:256])

		s, riff, error := readFOURCC(s)
		if error != nil || riff != riffFourcc {
			info.logError(fmt.Sprintf("reading RIFF FOURCC: %s", error))
			return &info
		}

		s, riffSize, error := readDWORD(s)
		maybeUnused(riffSize)
		if error != nil {
			info.logError(fmt.Sprintf("reading RIFF chunk size: %s", error))
			return &info
		}

		s, formType, error := readFOURCC(s)
		maybeUnused(formType)
		if error != nil {
			info.logError(fmt.Sprintf("reading FORM FOURCC: %s", error))
			return &info
		}

		lastRootChunkType := CT_Unknown

		for !s.empty() {
			s2, chunkFourcc, error := readFOURCC(s)
			if error != nil {
				info.logError(fmt.Sprintf("reading chunk FOURCC: %s", error))
				return &info
			}
			s2, chunkSize, error := readDWORD(s2)
			if error != nil {
				info.logError(fmt.Sprintf("reading chunk size: %s", error))
				return &info
			}
			//fmt.Printf("Chunk: %s %d\n", fourccToString(chunkFourcc), chunkSize)

			switch chunkFourcc {
			case rootFourcc:
				// Work out what type the ARCH chunk that follows this ROOT chunk will be.
				lastRootChunkType, error = scanRootChunk(s2.subspan(chunkSize))
				if error != nil {
					info.logError(fmt.Sprintf("reading ROOT chunk: %s", error))
					return &info
				}
			case archFourcc:
				// Process the ARCH chunk based on the type discovered in the preceeding ROOT chunk.
				switch lastRootChunkType {
				case CT_Version:
					info.version = scanArchChunk_Version(s2.subspan(chunkSize))
				case CT_Arrangement:
					cs := s.subspan(chunkSize)
					scanArchChunk(cs, &info)
				case CT_Devices:
					cs := s.subspan(chunkSize)
					scanArchChunk(cs, &info)
				}
			}

			s = s2.advance(chunkSize)
		}
	}
	return &info
}
