package main

// Alias struct holds the information for each alias
type Alias struct {
	Name        string            `json:"alias"`       // Name is the parent alias
	Command     string            `json:"command"`     // Command is the command to execute
	Description string            `json:"description"` // Description of the commands and subcommands
	Subalias    map[string]string `json:"subalias"`    // Subalias are the subaliases of the parent alias
	Platform    string            `json:"platform"`
}

// Plugin struct hold all of the information for each plugin and their aliases
type Plugin struct {
	Name        string  `json:"plugin_name"` // Name of the plugin
	Author      string  `json:"author"`      // Author of the plugin
	URL         string  `json:"url"`         // Repository of plugin
	Version     string  `json:"version"`     // Version of plugin
	AliasMap    []Alias `json:"alias_map"`   // AliasMap of the aliases for plugin
	Description string  `json:"description"` // Description of the plugin
}

// Config for aly
type Config struct {
	Version         string   `json:"aly_version"`
	EnabledPlugins  []Plugin `json:"enabled_plugins"`
	DisabledPlugins []Plugin `json:"disabled_plugins"`
}
