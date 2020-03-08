package docs

import (
	"google.golang.org/api/docs/v1"
	"net/http"
)

type Client struct {
	HTTPClient *http.Client
}

func (c *Client) NewChapter(docID string, texts []string) error {
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

		ind = ind + 2
	}

	requests = append(requests, paragraph(ind, texts[0]))
	requests = append(requests, &docs.Request{
		UpdateParagraphStyle:        &docs.UpdateParagraphStyleRequest{
			Fields: "*",
			ParagraphStyle:  &docs.ParagraphStyle{
				NamedStyleType:      "HEADING_1",
			},
			Range:           &docs.Range{
				EndIndex:        ind + int64(len(texts[0])) + 1,
				SegmentId:       "",
				StartIndex:      ind,
			},
		},
	})

	ind = ind + int64(len(texts[0])) + 1

	requests = append(requests, &docs.Request{
		UpdateParagraphStyle:        &docs.UpdateParagraphStyleRequest{
			Fields: "*",
			ParagraphStyle:  &docs.ParagraphStyle{
				NamedStyleType:      "NORMAL_TEXT",
			},
			Range:           &docs.Range{
				EndIndex:        ind + 1,
				SegmentId:       "",
				StartIndex:      ind,
			},
		},
	})

	ind = ind + 1

	for _, t := range texts[1:] {
		requests = append(requests, paragraph(ind, t))
		if t == "***" {
			requests = append(requests, &docs.Request{
				UpdateParagraphStyle:        &docs.UpdateParagraphStyleRequest{
					Fields: "*",
					ParagraphStyle:  &docs.ParagraphStyle{
						Alignment:      "CENTER",
						NamedStyleType: "NORMAL_TEXT",
					},
					Range:           &docs.Range{
						EndIndex:        ind + int64(len(t)) + 1,
						StartIndex:      ind,
					},
				},
			})
		}
		ind = ind + int64(len(t)) + 2
	}

	uReq := &docs.BatchUpdateDocumentRequest{
		Requests: requests,
	}
	_, err = srv.Documents.BatchUpdate(docID, uReq).Do()

	return err
}

func paragraph(index int64, text string) *docs.Request {
	return &docs.Request{
		InsertText: &docs.InsertTextRequest{
			Location: &docs.Location{
				Index: index,
			},
			Text: text + "\n\n",
		},
	}
}
