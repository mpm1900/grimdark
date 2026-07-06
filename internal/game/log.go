package game

import (
	"sort"
	"strings"
)

type Log struct {
	Depth    int               `json:"depth"`
	Template string            `json:"template"`
	Terms    map[string]string `json:"terms"`
	Type     string            `json:"type"`
}

func NewLog(template string, terms map[string]string) Log {
	return Log{
		Template: template,
		Terms:    terms,
		Type:     "message",
	}
}
func (g *Game) ResetLogDepth() {
	g.meta.log_depth = 0
}
func (g *Game) IncLogDepth() {
	g.meta.log_depth = g.meta.log_depth + 1
}

func (l Log) Bind(context Context) Bindable[Log] {
	return bind(l, context)
}

func (l Log) Resolve() string {
	if len(l.Terms) == 0 {
		return l.Template
	}

	keys := make([]string, 0, len(l.Terms))
	for key := range l.Terms {
		keys = append(keys, key)
	}

	// Replace longer keys first to avoid partial overlaps.
	sort.Slice(keys, func(i, j int) bool {
		return len(keys[i]) > len(keys[j])
	})

	parts := make([]string, 0, len(keys)*2)
	for _, key := range keys {
		parts = append(parts, key, l.Terms[key])
	}

	return strings.NewReplacer(parts...).Replace(l.Template)
}
