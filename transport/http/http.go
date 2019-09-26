package http

import (
	"context"
	"encoding/json"
	"mime"
	"net/http"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

var schemaDecoder = schema.NewDecoder()

func NewHttpServer(dec DecodeReq, handler Handler, middles ...MiddleWareOption) *Server {
	var s = &Server{dec: dec, handler: handler}
	for _, option := range middles {
		option(s)
	}
	return s
}

func (server *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 没有统一err处理，简单返回相应code
	var ctx context.Context
	for _, before := range server.before {
		ctx = before(ctx, r)
	}

	structReq, err := server.dec(ctx, r)
	if err != nil {
		enCodeErr(ctx, w, ErrorMsg{err.Error(), 400})
		return
	}
	resp, err := server.handler(ctx, structReq)
	if err != nil {
		enCodeErr(ctx, w, ErrorMsg{err.Error(), 500})
		return
	}
	if err := EncodeJSONResponse(ctx, w, resp); err != nil {
		enCodeErr(ctx, w, ErrorMsg{err.Error(), 500})
		return
	}
	return
}

func NewMuxServer() *mux.Router {
	r := mux.NewRouter()
	r.Use(handlers.CORS(
		handlers.AllowedMethods([]string{"PUT", "GET", "POST", "DELETE", "HEAD"}),
		handlers.AllowedHeaders([]string{"Origins", "X-Requested-With", "Content-Type", "Accept"}),
		handlers.AllowedOrigins([]string{"http://*", "https://*"}),
	))

	//添加middleware  r.use()
	return r
}

// EncodeJSONResponse ...
func EncodeJSONResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	// status code
	code := http.StatusOK
	if response == nil {
		code = http.StatusNoContent
	}
	w.WriteHeader(code)
	// body
	if code == http.StatusNoContent {
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

// EncodeJSONResponse ...
func enCodeErr(ctx context.Context, w http.ResponseWriter, msg ErrorMsg) {
	// status code
	w.WriteHeader(msg.code)
	_ = json.NewEncoder(w).Encode(msg.Reason)
}

// DecodePostForm loads form data from r.PostForm (which includes query params)
func DecodePostForm(r *http.Request, req interface{}) error {
	// This decodes urlencoded forms only if the content-type is correct
	if err := r.ParseForm(); err != nil {
		return err
	}
	// query params
	for k, v := range r.Form {
		// Support ?id[]=1&id[]=2 query string format
		if strings.HasSuffix(k, "[]") {
			newKey := strings.TrimSuffix(k, "[]")
			for _, i := range v {
				r.Form.Add(newKey, i)
			}
			r.Form.Del(k)
		}
	}
	if r.Form != nil {
		if err := schemaDecoder.Decode(req, r.Form); err != nil {
			return nil
		}
	}
	// post data
	if r.PostForm != nil {
		return schemaDecoder.Decode(req, r.PostForm)
	}
	return nil
}

// DecodeMuxVars loads URL vars into the request struct
func DecodeMuxVars(r *http.Request, req interface{}) error {
	vars := mux.Vars(r)
	if len(vars) == 0 {
		return nil
	}
	schemaInput := map[string][]string{}
	for k, v := range vars {
		schemaInput[k] = []string{v}
	}
	return schemaDecoder.Decode(req, schemaInput)
}

// IsFormRequest
func IsFormRequest(r *http.Request) bool {
	ctHeader := r.Header.Get("Content-Type")
	if ctHeader == "" {
		ctHeader = "application/json"
	}
	ct, _, err := mime.ParseMediaType(ctHeader)
	if err != nil {
		return false
	}
	return ct == "application/x-www-form-urlencoded" || ct == "multipart/form-data"
}
