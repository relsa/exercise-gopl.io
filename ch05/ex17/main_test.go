package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestElementsByTagName(t *testing.T) {

	tests := []struct {
		htmlStr string
		names   []string
		want    []string
	}{
		{
			`
<html>
<head> </head>
<body>
	<h1 id="1">test</h1>
	<div id="2">
		<div id="3">
			<h2 id="5">hoge</h2>
			<p id="6">hogehoge</p>
		</div>
		<div id="4">
			<h2 id="7">fuga</h2>
			<p id="8">fugafuga</p>
		</div>
	</div>
</body>
</html>
			`,
			[]string{"div"},
			[]string{"2", "3", "4"},
		},
	}

	for _, test := range tests {
		doc, err := html.Parse(strings.NewReader(test.htmlStr))
		if err != nil {
			t.Error(err)
			continue
		}
		got := ElementsByTagName(doc, test.names...)

		if len(got) != len(test.want) {
			t.Errorf("got %d nodes, want %d nodes", len(got), len(test.want))
			continue
		}

		ws := make(map[string]bool)
		for _, w := range test.want {
			ws[w] = true
		}

		for _, g := range got {
			gid, _ := getAttr(g, "id")
			if _, ok := ws[gid]; ok {
				ws[gid] = false
			} else {
				t.Errorf("unexpected node id=%q is got", gid)
			}
		}

		for w, v := range ws {
			if v {
				t.Errorf("node id=%q is not found", w)
			}
		}
	}
}

func getAttr(n *html.Node, name string) (string, bool) {
	for _, attr := range n.Attr {
		if attr.Key == name {
			return attr.Val, true
		}
	}
	return "", false
}
