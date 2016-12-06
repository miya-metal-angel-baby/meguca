// This file is automatically generated by qtc from "thread.html".
// See https://github.com/valyala/quicktemplate for details.

//line thread.html:1
package templates

//line thread.html:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line thread.html:1
import "github.com/bakape/meguca/lang"

//line thread.html:2
import "github.com/bakape/meguca/common"

//line thread.html:3
import "github.com/bakape/meguca/config"

//line thread.html:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line thread.html:5
func streamrenderThread(qw422016 *qt422016.Writer, t common.Thread, json []byte, title string, omit, imageOmit int, ln lang.Pack) {
	//line thread.html:5
	qw422016.N().S(`<h1 id="page-title">`)
	//line thread.html:7
	qw422016.E().S(title)
	//line thread.html:7
	qw422016.N().S(`</h1><span class="aside-container"><span class="act" id="top"><a href="#bottom">`)
	//line thread.html:12
	qw422016.N().S(ln.UI["bottom"])
	//line thread.html:12
	qw422016.N().S(`</a></span><span id="expand-images" class="act"><a>`)
	//line thread.html:17
	qw422016.N().S(ln.Common.Posts["expand"])
	//line thread.html:17
	qw422016.N().S(`</a></span></span><noscript>TODO: Noscript post creation</noscript><hr><div id="thread-container">`)
	//line thread.html:26
	root := config.Get().RootURL

	//line thread.html:27
	streamrenderArticle(qw422016, t.Post, t.ID, omit, imageOmit, t.Subject, root)
	//line thread.html:28
	for _, p := range t.Posts {
		//line thread.html:29
		streamrenderArticle(qw422016, p, t.ID, 0, 0, "", root)
		//line thread.html:30
	}
	//line thread.html:30
	qw422016.N().S(`</div><script id="post-data" type="application/json">`)
	//line thread.html:33
	qw422016.N().Z(json)
	//line thread.html:33
	qw422016.N().S(`</script><script id="board-configs" type="application/json">`)
	//line thread.html:36
	qw422016.N().Z(config.GetBoardConfigs(t.Board).JSON)
	//line thread.html:36
	qw422016.N().S(`</script><div id="bottom-spacer"></div><aside class="act posting"><a>`)
	//line thread.html:41
	qw422016.N().S(ln.UI["reply"])
	//line thread.html:41
	qw422016.N().S(`</a></aside><hr><span class="aside-container"><span class="act" id="bottom"><a href="." class="history">`)
	//line thread.html:48
	qw422016.N().S(ln.UI["return"])
	//line thread.html:48
	qw422016.N().S(`</a></span><span class="act"><a href="#top">`)
	//line thread.html:53
	qw422016.N().S(ln.UI["top"])
	//line thread.html:53
	qw422016.N().S(`</a></span><span class="act"><a href="?last=100" class="history reload">`)
	//line thread.html:58
	qw422016.N().S(ln.UI["last"])
	//line thread.html:58
	qw422016.N().S(`100</a></span><span id="lock" style="visibilty: hidden;">`)
	//line thread.html:62
	qw422016.N().S(ln.UI["lockedToBottom"])
	//line thread.html:62
	qw422016.N().S(`</span></span>`)
//line thread.html:65
}

//line thread.html:65
func writerenderThread(qq422016 qtio422016.Writer, t common.Thread, json []byte, title string, omit, imageOmit int, ln lang.Pack) {
	//line thread.html:65
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line thread.html:65
	streamrenderThread(qw422016, t, json, title, omit, imageOmit, ln)
	//line thread.html:65
	qt422016.ReleaseWriter(qw422016)
//line thread.html:65
}

//line thread.html:65
func renderThread(t common.Thread, json []byte, title string, omit, imageOmit int, ln lang.Pack) string {
	//line thread.html:65
	qb422016 := qt422016.AcquireByteBuffer()
	//line thread.html:65
	writerenderThread(qb422016, t, json, title, omit, imageOmit, ln)
	//line thread.html:65
	qs422016 := string(qb422016.B)
	//line thread.html:65
	qt422016.ReleaseByteBuffer(qb422016)
	//line thread.html:65
	return qs422016
//line thread.html:65
}