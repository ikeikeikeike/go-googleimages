package googleimages

import (
	"testing"
)

func TestSimple(t *testing.T) {
	c := NewClient()

	data, err := c.Fetch("simple")
	if err != nil {
		t.Fatal(err)
	}

	if len(data.ResponseData.Results) <= 0 {
		t.Fatalf("Unexpected Googleimages.ResponseData.Results: %d length", len(data.ResponseData.Results))
	}

	if data.ResponseStatus <= 0 || data.ResponseStatus == 500 {
		t.Fatalf("Unexpected Googleimages.ResponseStatus: %s", data.ResponseStatus)
	}

	if data.ResponseData.Cursor.EstimatedResultCount == "" {
		t.Fatalf(
			"Unexpected Googleimages.Cursor.ResponseData.EstimatedResultCount: %s length",
			data.ResponseData.Cursor.EstimatedResultCount,
		)
	}

	if len(data.ResponseData.Cursor.Pages) <= 0 {
		t.Fatalf(
			"Unexpected Googleimages.Cursor.ResponseData.Cursor.Pages: %d length",
			len(data.ResponseData.Cursor.Pages),
		)
	}

	if len(data.ResponseData.Results) <= 0 {
		t.Fatalf(
			"Unexpected Googleimages.ResponseData.Results: %d length",
			len(data.ResponseData.Results),
		)
	}

	r := data.ResponseData.Results[0]

	if r.Width == "" {
		t.Fatalf("Unexpected Result.Width: %s", r.Width)
	}
	if r.Height == "" {
		t.Fatalf("Unexpected Result.Height: %s", r.Height)
	}
	if r.TbWidth == "" {
		t.Fatalf("Unexpected Result.TbWidth: %s", r.TbWidth)
	}
	if r.TbHeight == "" {
		t.Fatalf("Unexpected Result.TbHeight: %s", r.TbHeight)
	}
	if r.UnescapedUrl == "" {
		t.Fatalf("Unexpected Result.UnescapedUrl: %s", r.UnescapedUrl)
	}
	if r.Url == "" {
		t.Fatalf("Unexpected Result.Url: %s", r.Url)
	}
	if r.TitleNoFormatting == "" {
		t.Fatalf("Unexpected Result.TitleNoFormatting: %s", r.TitleNoFormatting)
	}
	if r.OriginalContextUrl == "" {
		t.Fatalf("Unexpected Result.OriginalContextUrl: %s", r.OriginalContextUrl)
	}
	if r.ContentNoFormatting == "" {
		t.Fatalf("Unexpected Result.ContentNoFormatting: %s", r.ContentNoFormatting)
	}
	if r.TbUrl == "" {
		t.Fatalf("Unexpected Result.TbUrl: %s", r.TbUrl)
	}
}
