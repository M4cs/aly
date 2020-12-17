<p align="center">
    <a align="center"><img src=""></a>
</p>

# aly - Command Line Alias Manager and Packager

Aly offers the simplest way to manage, share, and obtain command line aliases!

### Table of Contents

#### I. [Features](https://github.com/M4cs/aly#features)
#### II. [Installation](https://github.com/M4cs/aly#installation)
#### III. [Usage](https://github.com/M4cs/aly#usage)
  - [Installing A Plugin](https://github.com/M4cs/aly#installing-a-plugin)
  - [Enabling/Disabling Plugins](https://github.com/M4cs/aly#enablingdisabling-plugins)
  - [Updating Plugins](https://github.com/M4cs/aly#update-plugins)
  - [Listing Plugins](https://github.com/M4cs/aly#see-all-installed-plugins)
  - [See Plugin Info](https://github.com/M4cs/aly#see-plugin-info)
  - [Deleting Plugins](https://github.com/M4cs/aly#delete-a-plugin-completely)
#### IV. [Creating Plugins](https://github.com/M4cs/aly#creating-a-plugin)
#### V. [Adding an Official Plugin](https://github.com/M4cs/aly#adding-your-plugin-to-the-official-repository)

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

### Enabling/Disabling Plugins

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

## Creating a Plugin

All plugins are in the JSON format. You can create a plugin very easily using the below formatting.

```json
{
    "plugin_name": "Your Plugin Name",
    "author": "github/M4cs",
    "description": "A small description about your plugin",
    "version": "1.0.0",
    "aliasmap": [
        {
            "alias": "ec",
            "description": "Example Alias",
            "command": "echo",
            "subalias": {
                "t": "'testing'",
                "hw": "'hello, world!'",
                "a": "$1"
            } 
        }
    ]
}
```

### Top-Level

- **plugin_name** - The name of your plugin. This is what it will be referred to for updating, enabling, disabling, etc. Keep it short and unique!
- **author** - Your name. I recommend using your github/USERNAME
- **description** - A small description about your plugin
- **version** - Version number for plugin. This should be semantic syntax.
- **aliasmap** - Your map of aliase groups for the plugin.


#### aliasmap

An alias group includes a base alias, base command, and subaliases. If there are subaliases, they will be concatenated with the base alias.

- **alias** - The base alias for your command. This will run whatever is in `command`.
- **command** - The base command for your alias. This will run whenever the `alias` is run.
- **description** - Description of your alias group.
- **subaliases** - subalias:args dictionary. This will add the `subalias` to the `alias` specified and add the `args` string to your base `command` string.

Using the JSON example above:

- `ec` will run `echo`
- `ect` will run `echo 'testing'`
- `echw` will run `echo 'hello, world!'`
- `eca` will run `echo $1` using whatever argument you send to it (or Windows equivalent)

## Adding Your Plugin To The Official Repository

You should add a folder with your username into `./official_plugins/`. For example, I'd add `./official_plugins/M4cs/`. Inside of this folder, you will add your different plugins using `.json` files. You should include a `README.md` with some information about what your plugin(s) offer!

1. Make a Pull Request using the `Add a Plugin` Template [Here]()
2. Fill Out Pull Request Template w/ Your Fork
3. Submit Pull Request and Await Review!