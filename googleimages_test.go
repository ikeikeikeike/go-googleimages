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
		t.Errorf("Unexpected Googleimages.ResponseData.Results: %d length", len(data.ResponseData.Results))
	}

	if data.ResponseStatus <= 0 || data.ResponseStatus == 500 {
		t.Errorf("Unexpected Googleimages.ResponseStatus: %s", data.ResponseStatus)
	}

	if data.ResponseData.Cursor.EstimatedResultCount == "" {
		t.Errorf(
			"Unexpected Googleimages.Cursor.ResponseData.EstimatedResultCount: %s length",
			data.ResponseData.Cursor.EstimatedResultCount,
		)
	}

	if len(data.ResponseData.Cursor.Pages) <= 0 {
		t.Errorf(
			"Unexpected Googleimages.Cursor.ResponseData.Cursor.Pages: %d length",
			len(data.ResponseData.Cursor.Pages),
		)
	}

	if len(data.ResponseData.Results) <= 0 {
		t.Errorf(
			"Unexpected Googleimages.ResponseData.Results: %d length",
			len(data.ResponseData.Results),
		)
	}

	r := data.ResponseData.Results[0]

	if r.Width == "" {
		t.Errorf("Unexpected Result.Width: %s", r.Width)
	}
	if r.Height == "" {
		t.Errorf("Unexpected Result.Height: %s", r.Height)
	}
	if r.TbWidth == "" {
		t.Errorf("Unexpected Result.TbWidth: %s", r.TbWidth)
	}
	if r.TbHeight == "" {
		t.Errorf("Unexpected Result.TbHeight: %s", r.TbHeight)
	}
	if r.UnescapedUrl == "" {
		t.Errorf("Unexpected Result.UnescapedUrl: %s", r.UnescapedUrl)
	}
	if r.Url == "" {
		t.Errorf("Unexpected Result.Url: %s", r.Url)
	}
	if r.TitleNoFormatting == "" {
		t.Errorf("Unexpected Result.TitleNoFormatting: %s", r.TitleNoFormatting)
	}
	if r.OriginalContextUrl == "" {
		t.Errorf("Unexpected Result.OriginalContextUrl: %s", r.OriginalContextUrl)
	}
	if r.ContentNoFormatting == "" {
		t.Errorf("Unexpected Result.ContentNoFormatting: %s", r.ContentNoFormatting)
	}
	if r.TbUrl == "" {
		t.Errorf("Unexpected Result.TbUrl: %s", r.TbUrl)
	}
}
