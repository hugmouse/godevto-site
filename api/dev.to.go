package handler

import (
    "fmt"
    api "github.com/hugmouse/godevto"
    "log"
    "net/http"
    "strconv"
    "strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    wstring := strings.Builder{}
    response, err := api.GetPublishedArticles(api.QueryArticle{
        Top: 1,
    }); if err != nil {
        log.Fatalln(err)
    }

    var i int

    // god-gopher forgive me
    // HTML OUTPUT:
    // ------------
    // <div>
    // 	<div class="borderme"
    //		<img src="">
    //		<h2><a href=""> Title </a></h2>
    //		<p> Description </p>
    //		<div id="comments">
    //			<p class="left"> Comments <code> amount </code></p>
    //		</div>
    //		<div id="rating">
    //			<p class="right"> Rating <code> amount </code></p>
    // 		</div>
    //	</div>
    //	<br>
    // </div>
    // ------------
    wstring.WriteString("<div>")
    for i = range response  {
        wstring.WriteString("<div class=\"borderme\">")

        // image
        if response[i].CoverImage != "" {
            wstring.WriteString("<img src=\"")
            wstring.WriteString(response[i].CoverImage)
            wstring.WriteString("\">")
        }

        // title
        wstring.WriteString("<h2><a href=\"")
        wstring.WriteString(response[i].URL)
        wstring.WriteString("\">")
        wstring.WriteString(response[i].Title)
        wstring.WriteString("</a></h2>")
        // description
        wstring.WriteString("<p>")
        wstring.WriteString(response[i].Description)
        wstring.WriteString("</p>")

        // stats
        wstring.WriteString("<div id=\"comments\">")

        wstring.WriteString("<p class=\"left\">Comments: <code>")
        wstring.WriteString(strconv.Itoa(response[i].CommentsCount))
        wstring.WriteString("</code></p>")

        wstring.WriteString("</div>")

        wstring.WriteString("<div id=\"rating\">")

        wstring.WriteString("<p class=\"right\">Rating: <code>")
        wstring.WriteString(strconv.Itoa(response[i].PositiveReactionsCount))
        wstring.WriteString("</code></p>")

        wstring.WriteString("</div>")

        // end
        wstring.WriteString("</div>")

        wstring.WriteString("<br>")
    }
    wstring.WriteString("</div>")

    _, _ = fmt.Fprintln(w, wstring.String())
}
