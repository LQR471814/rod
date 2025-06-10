package rod_test

import (
	"testing"

	"github.com/go-rod/rod/lib/proto"
)

// This is the template to demonstrate how to test Rod.
func TestRod(t *testing.T) {
	g := setup(t)
	g.cancelTimeout()

	page := g.page

	err := proto.AccessibilityEnable{}.Call(page)
	if err != nil {
		panic(err)
	}

	wait := page.EachEvent(
		func(e *proto.AccessibilityLoadComplete) bool {
			fetchRelatives := false

			res, err := proto.AccessibilityGetPartialAXTree{
				BackendNodeID:  e.Root.BackendDOMNodeID,
				FetchRelatives: &fetchRelatives,
			}.Call(page)
			if err != nil {
				panic(err)
			}
			g.Len(res.Nodes, 1)
			return true
		},
	)

	page.MustNavigate(g.html(doc)).MustWaitLoad()

	wait()
}

const doc = `
<html>
  <body>ok</body>
</html>
`
