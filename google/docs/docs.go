package docs

import (
	"chpunk/translation"
	"google.golang.org/api/docs/v1"
	"net/http"
)

type Client struct {
	HTTPClient *http.Client
}

func (c *Client) NewChapter(docID string, translations []*translation.Content) error {
	srv, err := docs.New(c.HTTPClient)
	if err != nil {
		return err
	}

	doc, err := srv.Documents.Get(docID).Do()
	if err != nil {
		return err
	}

	var requests []*docs.Request

	var ind int64
	ind = 1

	l := len(doc.Body.Content)
	if l > 2 {
		ind = doc.Body.Content[l-1].EndIndex - 1
		requests = append(requests, &docs.Request{
			InsertPageBreak: &docs.InsertPageBreakRequest{
				Location: &docs.Location{
					Index: ind,
				},
			},
		})

		ind += 2
	}

	request, nextIndex := paragraph(ind, translations[0].Yandex)
	requests = append(requests, request)
	requests = append(requests, &docs.Request{
		UpdateParagraphStyle: &docs.UpdateParagraphStyleRequest{
			Fields: "*",
			ParagraphStyle: &docs.ParagraphStyle{
				NamedStyleType: "HEADING_1",
			},
			Range: &docs.Range{
				EndIndex:   nextIndex - 1,
				SegmentId:  "",
				StartIndex: ind,
			},
		},
	})

	requests = append(requests, &docs.Request{
		UpdateParagraphStyle: &docs.UpdateParagraphStyleRequest{
			Fields: "*",
			ParagraphStyle: &docs.ParagraphStyle{
				NamedStyleType: "NORMAL_TEXT",
			},
			Range: &docs.Range{
				EndIndex:   nextIndex,
				SegmentId:  "",
				StartIndex: nextIndex - 1,
			},
		},
	})

	ind = nextIndex

	request, nextIndex = paragraph(ind, translations[0].Text)
	requests = append(requests, request)
	ind = nextIndex

	request, nextIndex = paragraph(ind, translations[0].Deepl)
	requests = append(requests, request)
	ind = nextIndex

	for _, t := range translations[1:] {
		request, nextIndex = paragraph(ind, t.Text)
		requests = append(requests, request)

		if t.IsNewParagraph() {
			requests = append(requests, &docs.Request{
				UpdateParagraphStyle: &docs.UpdateParagraphStyleRequest{
					Fields: "*",
					ParagraphStyle: &docs.ParagraphStyle{
						Alignment:      "CENTER",
						NamedStyleType: "NORMAL_TEXT",
					},
					Range: &docs.Range{
						EndIndex:   nextIndex - 1,
						StartIndex: ind,
					},
				},
			})
		} else {
			request, nextIndex = paragraph(nextIndex, t.Yandex)
			requests = append(requests, request)
			ind = nextIndex

			request, nextIndex = paragraph(ind, t.Deepl)
			requests = append(requests, request)
		}

		ind = nextIndex
	}

	uReq := &docs.BatchUpdateDocumentRequest{
		Requests: requests,
	}
	_, err = srv.Documents.BatchUpdate(docID, uReq).Do()

	return err
}

func paragraph(index int64, text string) (*docs.Request, int64) {
	request := &docs.Request{
		InsertText: &docs.InsertTextRequest{
			Location: &docs.Location{
				Index: index,
			},
			Text: text + "\n\n",
		},
	}

	endIndex := index + length(text) + 2

	return request, endIndex
}

func length(text string) int64 {
	return int64(len([]rune(text)))
}
