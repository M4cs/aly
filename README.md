# aly - Command Line Alias Manager and Packager

Aly offers the simplest way to manage, share, and obtain command line aliases!

## Features

- Aliases shared, loaded, and stored as JSON files for easy customization and readability
- Configure w/ the tool or manually
- Download remote Alias plugins to load a bunch all at once!
- Update plugins remotely or from local files!

## Installation

First, install Go and then run:

```
go get -u github.com/M4cs/aly
go install github.com/M4cs/aly
```

Next, run `aly -h` to see available options.

## Usage

### Installing a Plugin

```
# From Remote URL
aly -a 'https://raw.githubusercontent.com/M4cs/aly/master/example_unix_plugin.json'

# From Local File
aly -a '/path/to/plugin.json' -f

# Enable the plugin
aly -e 'Plugin Name'
```

### Enabling/Disable Plugins

```
# Enable
aly -e 'Plugin Name'

# Disable
aly -d 'Plugin Name'
```

### Update Plugins

```
# Update A Single Plugin from their URL
aly -u 'Plugin Name'

# Update A Single Plugin from Local File
aly -u '/path/to/plugin.json' -f

# Update All Plugins
aly -t
```

### See All Installed Plugins

```
aly -i
```

### See Plugin Info

```
aly -p 'Plugin Name'
```

### Delete A Plugin Completely

```
aly -r 'Plugin Name'
```