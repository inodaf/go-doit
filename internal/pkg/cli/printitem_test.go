package cli_test

import (
	"fmt"
	"inodaf/todo/internal/pkg/cli"
	"inodaf/todo/internal/pkg/models"
	"testing"
)

func TestRenderTemplateWhenItemIsComplete(t *testing.T) {
	item, _ := models.NewItem("Hello World")
	item.MarkAsDone()

	var expected string = fmt.Sprintf("\n----\n[x] #0: Hello World\nDone at %s\nCreated at %s\n", item.DoneAt, item.CreatedAt)
	var output string = cli.RenderTemplate(item, false)

	if output != expected {
		t.Errorf("Expected: %s, Got: %s", expected, output)
	}
}

func TestRenderTemplateWhenItemIsPending(t *testing.T) {
	item, _ := models.NewItem("Hello World")

	var expected string = fmt.Sprintf("\n----\n[] #0: Hello World\nCreated at %s\n", item.CreatedAt)
	var output string = cli.RenderTemplate(item, false)

	if output != expected {
		t.Errorf("Expected: %s, Got: %s", expected, output)
	}
}

func TestRenderTemplateWithDetails(t *testing.T) {
	item, _ := models.NewItem("Learn Testing in Go")
	item.Description = "A complete description"

	var expected string = fmt.Sprintf("\n----\n[] #0: Learn Testing in Go\nCreated at %s\n\n%s\n", item.CreatedAt, item.Description)
	var output string = cli.RenderTemplate(item, true)

	if output != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, output)
	}
}
