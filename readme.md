# GTI

[![Build Status](https://travis-ci.org/ldez/gti.svg?branch=master)](https://travis-ci.org/ldez/gti)

A port of the original [gti](https://github.com/rwos/gti) in Go.

```
        ,---------------.
       /  /``````|``````\\
      /  /_______|_______\\________
     |]      GTI |'       |        |]
     =  .-:-.    |________|  .-:-.  =
      `  -+-  --------------  -+-  '
        '-:-'                '-:-'  
```


## Installation

### Binaries

To get the binary just download the latest release for your OS/Arch from [the release page](https://github.com/ldez/gti/releases) and put the binary somewhere convenient.

### From source

To install from source, just run:

```bash
go get -u github.com/ldez/gti
```

## Environment Variables

- `GTI_SPEED` [`int`]: display speed (default: `"1000"`)
- `GTI_VERBOSE` [`bool`]: display GTI version (default: `"false"`)
