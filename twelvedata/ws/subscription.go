package ws

import (
	"encoding/json"
	"fmt"
	"strings"
)

// canonicalKey produces a stable key for an input so semantically equal
// subscriptions (a bare ticker vs the same ticker as a Symbol) deduplicate.
// The accepted concrete types are the same ones normalize emits: string and
// Symbol. Other types yield "".
func canonicalKey(item any) string {
	switch v := item.(type) {
	case string:
		return "s:" + strings.TrimSpace(v)
	case Symbol:
		b, _ := json.Marshal(v)
		return "o:" + string(b)
	}
	return ""
}

// normalize accepts the variadic forms of Subscribe / Unsubscribe input,
// splits comma-separated strings, validates the types, and deduplicates by
// canonical key. The result contains only string and Symbol values.
//
// Accepted types: string, Symbol, []string, []Symbol, []any (recursively).
// Anything else returns an error.
func normalize(inputs []any) ([]any, error) {
	out := make([]any, 0, len(inputs))
	seen := make(map[string]bool, len(inputs))

	var visit func(item any) error
	visit = func(item any) error {
		switch v := item.(type) {
		case string:
			for _, s := range strings.Split(v, ",") {
				s = strings.TrimSpace(s)
				if s == "" {
					continue
				}
				key := "s:" + s
				if !seen[key] {
					seen[key] = true
					out = append(out, s)
				}
			}
			return nil
		case Symbol:
			v.Symbol = strings.TrimSpace(v.Symbol)
			if v.Symbol == "" {
				return nil
			}
			key := canonicalKey(v)
			if !seen[key] {
				seen[key] = true
				out = append(out, v)
			}
			return nil
		case []string:
			for _, s := range v {
				if err := visit(s); err != nil {
					return err
				}
			}
			return nil
		case []Symbol:
			for _, s := range v {
				if err := visit(s); err != nil {
					return err
				}
			}
			return nil
		case []any:
			for _, s := range v {
				if err := visit(s); err != nil {
					return err
				}
			}
			return nil
		default:
			return fmt.Errorf("ws: unsupported subscribe input type %T (accepted: string, Symbol, []string, []Symbol, []any)", item)
		}
	}

	for _, item := range inputs {
		if err := visit(item); err != nil {
			return nil, err
		}
	}
	return out, nil
}
