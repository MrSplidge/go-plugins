# About

The go-plugins utility can examine ALS (11.0_433, 11.0_436, 11.0_11300) and CPR (12.0.70, 13.0.21) project files and generate tables showing the plugins used on each track, and vice versa.

# Building

* ```go get```
* ```go build```

# Usage

```
.\go-plugins -h

Usage of C:\Git\go\go-plugins\go-plugins.exe:
  -extensions value
        A semicolon-separated list of project file extensions to include when traversing the hierarchy (default .als;.cpr).
  -ignore-folders value
        A semicolon-separated list of folders to ignore when traversing the hierarchy.
  -num-threads int
        The number of worker threads to use. (default 64)
```

# Examples

1. Examine a single project file:

```
.\go-plugins C:\Music\Sets\Project1\Project1.als
```
```
.\go-plugins C:\Music\Sets\Project2\Project2.cpr
```

2. Examine all the projects within a folder heirarchy:

```
.\go-plugins C:\Music\Sets
```

3. Examine all the projects within a folder hierarchy, ignoring the contents of ```Backup``` and ```Experimental``` folders.

```
.\go-plugins -ignore-folders Backup;Experimental C:\Music\Sets
```

4. Examine only the project files with a ```.als``` file extension within a folder hierarchy.

```
.\go-plugins -extensions .als C:\Music\Sets
```

# Example Output

Note: Output will appear multicoloured in a Terminal, but monochrome if redirected to a file.

```
Project: C:\Music\Sets\43\43.als
Version: 11.0_11300

Plugin followed by a list of the tracks within which it appears:
  Chromaphone 3‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ 15-Chromaphone 3 ]
  DSEQ3‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ Master ]
  DUNE 3‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ 12 D3.Sqr, 19-DUNE 3, 22-DUNE 3, 23 D3.Pluck Arp Reverb, 24 D3.Pluck Arp Reverb, 5 D3.RBass1, 6 D3.RBass2, 7 D3.Drone ]
  Kick 2 x64‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ 4 Kick ]
  Newfangled Elevate‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ Master ]
  Newfangled Saturate‐‐‐‐‐‐‐‐‐‐‐‐‐[ Master ]
  Rift Feedback Lite‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ 11-VPS Avenger ]
  StandardCLIP‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ 3 Perc, 4 Kick, 5 D3.RBass1, 6 D3.RBass2 ]
  Surge XT‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ 20-Surge XT ]
  trueBalance‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ Master ]
  trueLevel‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ Master ]
  VPS Avenger‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ 10-VPS Avenger, 11-VPS Avenger, 14-VPS Avenger, 16-VPS Avenger, 17-VPS Avenger, 18-VPS Avenger, 21 VA.Riser, 8-VPS Avenger, 9-VPS Avenger ]
  VPS Avenger_x64‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ 13 V.PWM, 25-VPS Avenger_x64 ]

Track followed by a list of the plugins that it uses:
  10-VPS Avenger‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ VPS Avenger ]
  11-VPS Avenger‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ Rift Feedback Lite, VPS Avenger ]
  12 D3.Sqr‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ DUNE 3 ]
  13 V.PWM‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ VPS Avenger_x64 ]
  14-VPS Avenger‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ VPS Avenger ]
  15-Chromaphone 3‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ Chromaphone 3 ]
  16-VPS Avenger‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ VPS Avenger ]
  17-VPS Avenger‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ VPS Avenger ]
  18-VPS Avenger‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ VPS Avenger ]
  19-DUNE 3‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ DUNE 3 ]
  20-Surge XT‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ Surge XT ]
  21 VA.Riser‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ VPS Avenger ]
  22-DUNE 3‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ DUNE 3 ]
  23 D3.Pluck Arp Reverb‐‐‐‐‐‐‐‐‐‐[ DUNE 3 ]
  24 D3.Pluck Arp Reverb‐‐‐‐‐‐‐‐‐‐[ DUNE 3 ]
  25-VPS Avenger_x64‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ VPS Avenger_x64 ]
  3 Perc‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ StandardCLIP ]
  4 Kick‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ Kick 2 x64, StandardCLIP ]
  5 D3.RBass1‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ DUNE 3, StandardCLIP ]
  6 D3.RBass2‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ DUNE 3, StandardCLIP ]
  7 D3.Drone‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ DUNE 3 ]
  8-VPS Avenger‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ VPS Avenger ]
  9-VPS Avenger‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ VPS Avenger ]
  Master‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ DSEQ3, Newfangled Elevate, Newfangled Saturate, trueBalance, trueLevel ]

Project: C:\Music\Sets\92\92.cpr
Version: Version 12.0.70

Plugin followed by a list of the tracks within which it appears:
  Blackhole‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ Group 01 ]
  Chromaphone 3‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ Chime, Perc, Triangle ]
  DSEQ3‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ Stereo Out ]
  Newfangled Elevate‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ Stereo Out ]
  Newfangled Saturate‐‐‐‐‐‐‐‐‐‐‐‐‐[ Stereo Out ]
  ValhallaShimmer‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ Group 01 ]

Track followed by a list of the plugins that it uses:
  Chime‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ Chromaphone 3 ]
  Group 01‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ Blackhole, ValhallaShimmer ]
  Perc‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ Chromaphone 3 ]
  Stereo Out‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ DSEQ3, Newfangled Elevate, Newfangled Saturate ]
  Triangle‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐‐[ Chromaphone 3 ]

```

# Licence

Please see the included BSD 3-clause LICENSE file.

This program makes use of the following third-party packages. Please refer to these projects for additional licensing information.

* github.com/MrSplidge/go-coutil v1.0.0 // BSD 3-clause License
* github.com/MrSplidge/go-xmldom v1.1.4 // based on https://github.com/subchen/go-xmldom (Apache License 2.0 Jan 2004)
* github.com/antchfx/xpath v1.2.5 // (MIT License)
* github.com/mattn/go-isatty v0.0.20 // (MIT License (Expat))
