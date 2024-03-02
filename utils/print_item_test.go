package utils

import (
	"fmt"
	"testing"

	"inodaf/todo/internal/models"
)

func TestRenderTemplateWhenItemIsComplete(t *testing.T) {
	var item *models.Item = models.NewItem()
	item.Title = "Hello World"
	item.MarkAsDone()

	var expected string = fmt.Sprintf("\n----\n[x] #0: Hello World\nDone at %s\nCreated at %s\n", item.DoneAt, item.CreatedAt)
	var output string = RenderTemplate(item, 0, false)

	if output != expected {
		t.Errorf("Expected: %s, Got: %s", expected, output)
	}
}

func TestRenderTemplateWhenItemIsPending(t *testing.T) {
	var item *models.Item = models.NewItem()
	item.Title = "Hello World"

	var expected string = fmt.Sprintf("\n----\n[] #0: Hello World\nCreated at %s\n", item.CreatedAt)
	var output string = RenderTemplate(item, 0, false)

	if output != expected {
		t.Errorf("Expected: %s, Got: %s", expected, output)
	}
}

func TestRenderTemplateWithDetails(t *testing.T) {
	var item *models.Item = models.NewItem()
	item.Title = "Learn Testing in Go"
	item.Description = "A complete description"

	var expected string = fmt.Sprintf("\n----\n[] #0: Learn Testing in Go\nCreated at %s\n\n%s\n", item.CreatedAt, item.Description)
	var output string = RenderTemplate(item, 0, true)

	if output != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, output)
	}
}