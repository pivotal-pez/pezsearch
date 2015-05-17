package pezsearch

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/martini-contrib/render"
)

const (
	ScopeKeyword  = "scope:"
	TypeResource  = "type"
	ItemResource  = "item"
	LimitDefault  = 10
	OffsetDefault = 0
)

type searchController struct {
}

type queryParams struct {
	q      string
	scope  string
	limit  int
	offset int
}

// @Title search
// @Description Search API
// @Accept  json
// @Param   X-API-KEY       header  string  true        "APIKEY"
// @Param   q               query   string  true        "Retrieve items matching search terms"
// @Success 200 {object}  ResponseMessage
// @Failure 400 {object}  ResponseMessage
// @Resource /v1/search
// @Router /v1/search [get]
func (c *searchController) find(req *http.Request, render render.Render) {
	r := req.URL.Query()

	q := r.Get("q")
	if len(q) == 0 {
		render.JSON(400, errorMessage("You must supply search criteria"))
		return
	}

	scope, q := extractScope(q)

	limit := parseLimit(r)

	offset := parseOffset(r)

	params := &queryParams{
		q:      q,
		scope:  scope,
		limit:  limit,
		offset: offset,
	}

	db := initDB()
	count, results := db.find(params)

	reply := successMessage(results)

	meta := make(map[string]interface{})
	meta["count"] = count
	meta["uri"] = req.URL.RequestURI()
	reply.Meta = meta

	//TODO(dnem) populate links

	render.JSON(200, reply)
	return
}

func extractScope(in string) (scope string, out string) {
	f := strings.Fields(in)
	out = ""
	for i := range f {
		if strings.Contains(f[i], ScopeKeyword) {
			scope = strings.TrimPrefix(f[i], ScopeKeyword)
		} else {
			out = out + f[i] + " "
		}

	}
	out = strings.Trim(out, " ")
	return
}

func parseLimit(qs url.Values) (limit int) {
	limit = LimitDefault
	s := qs.Get("limit")
	if len(s) > 0 {
		l, err := strconv.Atoi(s)
		if err == nil {
			limit = l
		}
	}
	return
}

func parseOffset(qs url.Values) (offset int) {
	offset = OffsetDefault
	s := qs.Get("offset")
	if len(s) > 0 {
		o, err := strconv.Atoi(s)
		if err == nil {
			offset = o
		}
	}
	return
}

// TODO(dnem): partial response support
//   ?fields=comma,separated,list
//
//   What to do if field does not exist?
//     In a global search?
//     In a scoped search?
//     Should this not be supported?
