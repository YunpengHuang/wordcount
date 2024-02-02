package wordcount

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ErrorHandler func(error, *Context) error


type Context struct {
  repsonse http.ResponseWriter
  request *http.Request
  ctx context.Context
}

type Handler func(c *Context) error

type Wordcount struct {
  ErrorHandler ErrorHandler
  router *httprouter.Router
}

func New( ) *Wordcount {
  return &Wordcount{
    router: httprouter.New(),
    ErrorHandler: defaultErrorHanlder,
  }
}

func (s *Wordcount) Get(path string, h Handler, plugs ...Handler) {
  s.router.GET(path string, handle httprouter.Handle)
}

func (s *Wordcount)makeHTTPRouterHandler(h Handler) htthttprouter.Handle {
  return func(w http.ResponseWriter, r *http.Request, params httprouter.Params){
    ctx := &Context {
      response:w,
      request: r,
      ctx: context.Background(),
    }
    if err := h(ctx); err != nil {
      s.ErrorHandler(err, ctx)
    }
  }
}

func defaultErrorHanlder(err error, c *Context) error {
  slog.Error("erorr", "err", err)
  return nil
}
