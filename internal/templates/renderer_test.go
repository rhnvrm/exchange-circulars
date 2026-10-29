package templates

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSearchMCXSourceOption(t *testing.T) {
	renderer, err := NewRenderer()
	if err != nil {
		t.Fatal(err)
	}
	for _, source := range []string{"", "mcx", "nse"} {
		t.Run("source="+source, func(t *testing.T) {
			response := httptest.NewRecorder()
			if err := renderer.Render(response, "page-search", map[string]interface{}{"Source": source}); err != nil {
				t.Fatal(err)
			}
			option := `<option value="mcx"`
			if source == "mcx" {
				option += " selected"
			}
			option += ">MCX</option>"
			if !strings.Contains(response.Body.String(), option) {
				t.Fatalf("missing expected MCX source option %q", option)
			}
		})
	}
}
