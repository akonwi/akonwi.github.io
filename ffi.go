package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"
)

// ============================================================
// ParseFrontmatter — splits YAML frontmatter from markdown body
// Returns JSON: { title, date_iso, categories[], body_md, excerpt }
// ============================================================

func ParseFrontmatter(content string) (string, error) {
	content = strings.TrimSpace(content)
	if !strings.HasPrefix(content, "---") {
		// No frontmatter — treat entire content as body
		excerpt := extractExcerpt(content)
		out := map[string]any{
			"title":      "",
			"date_iso":   "",
			"categories": []string{},
			"body_md":    content,
			"excerpt":    excerpt,
			"published":  true,
		}
		encoded, err := json.Marshal(out)
		if err != nil {
			return "", fmt.Errorf("encode content: %w", err)
		}
		return string(encoded), nil
	}

	// Find closing ---
	rest := content[3:]
	endIdx := strings.Index(rest, "\n---")
	if endIdx < 0 {
		return "", fmt.Errorf("unclosed frontmatter")
	}

	yamlBlock := strings.TrimSpace(rest[:endIdx])
	bodyMD := strings.TrimSpace(rest[endIdx+4:])

	front := parseYAMLSimple(yamlBlock)

	// Extract fields
	title := ""
	if v, ok := front["title"]; ok {
		title = v
	}

	dateISO := ""
	if v, ok := front["date"]; ok {
		if t, err := parseDateValue(v); err == nil {
			dateISO = t.Format(time.RFC3339)
		}
	}

	categories := []string{}
	if v, ok := front["categories"]; ok {
		categories = parseYAMLList(v)
	}

	excerpt := extractExcerpt(bodyMD)

	published := true
	if v, ok := front["published"]; ok {
		published = v == "true"
	}

	out := map[string]any{
		"title":      title,
		"date_iso":   dateISO,
		"categories": categories,
		"body_md":    bodyMD,
		"excerpt":    excerpt,
		"published":  published,
	}

	encoded, err := json.Marshal(out)
	if err != nil {
		return "", fmt.Errorf("encode frontmatter: %w", err)
	}

	return string(encoded), nil
}

// Simple YAML parser for frontmatter subset
func parseYAMLSimple(yaml string) map[string]string {
	out := map[string]string{}
	lines := strings.Split(yaml, "\n")
	var currentKey string
	var currentList []string

	flushList := func() {
		if currentKey != "" && len(currentList) > 0 {
			out[currentKey] = strings.Join(currentList, "\x00")
			currentList = nil
		}
	}

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}

		// List item
		if strings.HasPrefix(trimmed, "- ") {
			if currentKey != "" {
				currentList = append(currentList, strings.TrimSpace(trimmed[2:]))
			}
			continue
		}

		// Key: value
		colonIdx := strings.Index(trimmed, ":")
		if colonIdx > 0 {
			flushList()
			key := strings.TrimSpace(trimmed[:colonIdx])
			value := ""
			if colonIdx+1 < len(trimmed) {
				value = strings.TrimSpace(trimmed[colonIdx+1:])
				value = stripQuotes(value)
			}
			currentKey = key
			if value != "" {
				out[key] = value
				currentKey = ""
			} else {
				currentList = []string{}
			}
		}
	}
	flushList()
	return out
}

func parseYAMLList(value string) []string {
	parts := strings.Split(value, "\x00")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

func parseDateValue(value string) (time.Time, error) {
	formats := []string{
		"2006-01-02 15:04",
		"2006-01-02",
		time.RFC3339,
	}
	for _, f := range formats {
		if t, err := time.Parse(f, value); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("cannot parse date: %s", value)
}

func stripQuotes(s string) string {
	if len(s) >= 2 {
		if (s[0] == '"' && s[len(s)-1] == '"') || (s[0] == '\'' && s[len(s)-1] == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}

// ============================================================
// MarkdownToHTML — renders markdown to HTML (stdlib only)
// ============================================================

var (
	headingRegex  = regexp.MustCompile(`^(#{1,6})\s+(.+)$`)
	codeFence    = regexp.MustCompile("^```(\\w*)$")
	blockquoteRx = regexp.MustCompile(`^>\s?(.*)$`)
	listItemRx   = regexp.MustCompile(`^[-*]\s+(.+)$`)
	thematicBreak = regexp.MustCompile(`^[-*_]{3,}\s*$`)
)

func MarkdownToHTML(markdown string) (string, error) {
	lines := strings.Split(markdown, "\n")
	var out strings.Builder

	inCodeBlock := false
	var codeLang string
	var codeBuf strings.Builder
	inList := false

	flush := func() {
		if inCodeBlock {
			langAttr := ""
			if codeLang != "" {
				langAttr = fmt.Sprintf(` class="language-%s"`, codeLang)
			}
			out.WriteString(fmt.Sprintf("<pre><code%s>%s</code></pre>\n", langAttr, htmlEscape(codeBuf.String())))
			codeBuf.Reset()
			codeLang = ""
			inCodeBlock = false
		}
	}

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		if strings.HasPrefix(line, "```") {
			if inCodeBlock {
				flush()
			} else {
				flush()
				codeLang = strings.TrimSpace(line[3:])
				inCodeBlock = true
			}
			continue
		}

		if inCodeBlock {
			codeBuf.WriteString(line)
			codeBuf.WriteString("\n")
			continue
		}

		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			flush()
			if inList {
				out.WriteString("</ul>\n")
				inList = false
			}
			out.WriteString("\n")
			continue
		}

		// Thematic break
		if thematicBreak.MatchString(trimmed) {
			flush()
			if inList {
				inList = false
			}
			out.WriteString("<hr>\n")
			continue
		}

		// Headings
		if m := headingRegex.FindStringSubmatch(trimmed); m != nil {
			flush()
			if inList {
				out.WriteString("</ul>\n")
				inList = false
			}
			level := len(m[1])
			text := renderInline(m[2])
			out.WriteString(fmt.Sprintf("<h%[1]d>%s</h%[1]d>\n", level, text))
			continue
		}

		// Blockquote
		if m := blockquoteRx.FindStringSubmatch(trimmed); m != nil {
			flush()
			if inList {
				out.WriteString("</ul>\n")
				inList = false
			}
			text := renderInline(m[1])
			out.WriteString(fmt.Sprintf("<blockquote><p>%s</p></blockquote>\n", text))
			continue
		}

		// List item
		if m := listItemRx.FindStringSubmatch(trimmed); m != nil {
			flush()
			if !inList {
				out.WriteString("<ul>\n")
				inList = true
			}
			text := renderInline(m[1])
			out.WriteString(fmt.Sprintf("<li>%s</li>\n", text))
			continue
		} else if inList {
			inList = false
			out.WriteString("</ul>\n")
		}

		// Paragraph
		text := renderInline(trimmed)
		out.WriteString(fmt.Sprintf("<p>%s</p>\n", text))
	}

	flush()
	if inList {
		out.WriteString("</ul>\n")
	}

	return strings.TrimSpace(out.String()), nil
}

// Inline rendering: *italic*, **bold**, `code`, [text](url), ![alt](url)
func renderInline(text string) string {
	// Images first (so [ isn't consumed by links)
	imgRe := regexp.MustCompile(`!\[([^\]]*)\]\(([^)]+)\)`)
	text = imgRe.ReplaceAllString(text, `<img alt="$1" src="$2">`)

	// Links
	linkRe := regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`)
	text = linkRe.ReplaceAllString(text, `<a href="$2">$1</a>`)

	// Code (inline) - must handle backticks
	text = renderInlineCode(text)

	// Bold (** or __)
	boldRe := regexp.MustCompile(`\*\*(.+?)\*\*|__(.+?)__`)
	text = boldRe.ReplaceAllString(text, `<strong>$1$2</strong>`)

	// Italic (* or _)
	italicRe := regexp.MustCompile(`\*(.+?)\*|_(.+?)_`)
	text = italicRe.ReplaceAllString(text, `<em>$1$2</em>`)

	return text
}

func renderInlineCode(text string) string {
	// Replace backtick code: `code`  — simple single backtick
	codeRe := regexp.MustCompile("`([^`]+)`")
	return codeRe.ReplaceAllStringFunc(text, func(match string) string {
		inner := match[1 : len(match)-1]
		return "<code>" + htmlEscape(inner) + "</code>"
	})
}

func htmlEscape(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	s = strings.ReplaceAll(s, "\"", "&quot;")
	return s
}

// ============================================================
// Excerpt extraction
// ============================================================

func extractExcerpt(bodyMD string) string {
	// Take the first non-empty, non-heading paragraph, strip markdown
	blocks := strings.Split(bodyMD, "\n\n")
	for _, block := range blocks {
		block = strings.TrimSpace(block)
		if block == "" || strings.HasPrefix(block, "#") {
			continue
		}
		plain := stripMarkdown(block)
		if plain == "" {
			continue
		}
		if len(plain) <= 220 {
			return plain
		}
		return plain[:217] + "..."
	}
	return ""
}

func stripMarkdown(md string) string {
	// Remove code blocks
	md = regexp.MustCompile("```[\\s\\S]*?```").ReplaceAllString(md, "")
	// Remove inline code
	md = regexp.MustCompile("`([^`]+)`").ReplaceAllString(md, "$1")
	// Remove images
	md = regexp.MustCompile(`!\[[^\]]*\]\([^\)]*\)`).ReplaceAllString(md, "")
	// Remove links (keep text)
	md = regexp.MustCompile(`\[([^\]]+)\]\([^\)]*\)`).ReplaceAllString(md, "$1")
	// Remove heading markers
	md = regexp.MustCompile(`^#+\s+`).ReplaceAllString(md, "")
	// Remove bold/italic markers
	md = regexp.MustCompile(`[*_~]`).ReplaceAllString(md, "")
	// Collapse whitespace
	md = regexp.MustCompile(`\s+`).ReplaceAllString(md, " ")
	return strings.TrimSpace(md)
}
