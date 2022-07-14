package routing

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"url_shortener_main/db"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"url"`
}

type handler struct {
	host string
	db   *db.Database
}

func New(host string, db *db.Database) *router.Router {
	router := router.New()

	h := handler{host, db}
	router.POST("/encode/", responseHandler(h.encode))
	router.GET("/{url}", h.decode)
	return router
}

func responseHandler(h func(ctx *fasthttp.RequestCtx) (interface{}, int, error)) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		data, code, err := h(ctx)
		if err != nil {
			log.Printf("Processing error")
			data = err.Error()
		}
		ctx.Response.Header.Set("Content-type", "application/json")
		ctx.Response.SetStatusCode(code)
		err = json.NewEncoder(ctx.Response.BodyWriter()).Encode(response{Data: data, Success: err == nil})
		if err != nil {
			log.Printf("Encoding error")
		}
	}
}

func (h handler) encode(ctx *fasthttp.RequestCtx) (interface{}, int, error) {
	rurl := string(ctx.PostArgs().Peek("url"))
	log.Println(rurl)
	uri, err := url.ParseRequestURI(rurl)
	if err != nil {
		return nil, fasthttp.StatusBadRequest, err
	}
	log.Println(uri)
	path, err := h.db.EncodeAndSaveLink(uri.String())
	if err != nil {
		return nil, fasthttp.StatusInternalServerError, err
	}
	res := url.URL{
		Scheme: "http",
		Host:   h.host,
		Path:   path,
	}
	log.Println("Encoded LINK:", res.String())
	return res.String(), fasthttp.StatusCreated, nil
}

func (h handler) decode(ctx *fasthttp.RequestCtx) { // decode and redirect to decoded link
	rurl := ctx.UserValue("url").(string)
	log.Println("URL: " + rurl)

	res, err := h.db.DecodeLink(rurl)
	if err != nil {
		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.Response.SetStatusCode(fasthttp.StatusNotFound)
	}

	ctx.Redirect(res, http.StatusMovedPermanently)
}
