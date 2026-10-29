package content

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadFullMCXDocumentURL(t *testing.T) {
	const url = "https://www.mcxindia.com/docs/default-source/circulars/english/2026/october/circular---594---2026.pdf"
	const body = "Circular summary\n\n---\n\nMore details."
	for _, newline := range []string{"\n", "\r\n"} {
		t.Run(strings.ReplaceAll(newline, "\n", "LF"), func(t *testing.T) {
			dir := t.TempDir()
			path := filepath.Join(dir, "mcx-2026-10-20-test.md")
			metadata := "---\ncircular_id: mcx-test\ndate: '2026-10-20'\nguid: " + url + "\npdf_url: " + url + "\nsource: mcx\ntitle: MCX circular\n---\n"
			if err := os.WriteFile(path, []byte(strings.ReplaceAll(metadata+body, "\n", newline)), 0600); err != nil {
				t.Fatal(err)
			}
			loader := NewLoader(dir)
			full, err := loader.LoadFull(path)
			if err != nil {
				t.Fatal(err)
			}
			if full.CircularID != "mcx-test" || full.Title != "MCX circular" || full.Source != "mcx" || full.PDFURL != url || full.GUID != url {
				t.Fatalf("metadata was corrupted: %+v", full)
			}
			if full.RawContent != strings.ReplaceAll(body, "\n", newline) {
				t.Fatalf("body = %q", full.RawContent)
			}
			summary, err := loader.LoadSummary(path)
			if err != nil {
				t.Fatal(err)
			}
			if summary.Title != full.Title || summary.Source != full.Source || summary.CircularID != full.CircularID {
				t.Fatal("summary and full metadata differ")
			}
		})
	}
}

func TestLoadFullRequiresFrontmatterDelimiters(t *testing.T) {
	for name, content := range map[string]string{
		"empty":            "",
		"missing opening":  "title: Circular\n---\nBody",
		"missing closing":  "---\ntitle: Circular---Title\n",
		"delimiter prefix": "---\ntitle: Circular\n---suffix\nBody",
	} {
		t.Run(name, func(t *testing.T) {
			dir := t.TempDir()
			path := filepath.Join(dir, "circular.md")
			if err := os.WriteFile(path, []byte(content), 0600); err != nil {
				t.Fatal(err)
			}
			if _, err := NewLoader(dir).LoadFull(path); err == nil {
				t.Fatal("expected invalid frontmatter to be rejected")
			}
		})
	}
}
