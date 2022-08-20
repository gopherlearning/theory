//go:build ignore
// +build ignore

package main

import (
	"fmt"
)

func main() {
	orig := Config{
		Version: "1",
		Plugins: []string{"11", "22"},
		Stat:    map[string]int{"333": 3},
	}

	fmt.Printf("task3\nOrig: %+v\nOrig Plugins %p\nOrig Stat %p\n", orig, orig.Plugins, &orig.Stat)
	clone := orig.Clone()
	fmt.Printf("task3\nClone: %+v\nClone Plugins %p\nClone Stat %p\n", clone, clone.Plugins, &clone.Stat)

}

type Config struct {
	Version string
	Plugins []string
	Stat    map[string]int
}

func (cfg *Config) Clone() *Config {
	clone := &Config{
		Version: cfg.Version,
		Plugins: append([]string{}, cfg.Plugins...),
		Stat:    make(map[string]int, len(cfg.Stat)),
	}
	for i, v := range cfg.Stat {
		clone.Stat[i] = v
	}
	return clone
}
