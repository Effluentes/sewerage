package server

import "net/http"

type HTTPMethodsCallback struct {
	getHandler    HTTPHandler
	postHandler   HTTPHandler
	putHandler    HTTPHandler
	deleteHandler HTTPHandler
}

type HandleMethod func(*HTTPMethodsCallback)

func NewMethodHandler(opts ...HandleMethod) HTTPHandler {
	cfg := &HTTPMethodsCallback{}

	for _, opt := range opts {
		opt(cfg)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			if cfg.getHandler != nil {
				cfg.getHandler(w, r)
			} else {
				http.Error(w, "GET not implemented", http.StatusNotImplemented)
			}
		case http.MethodPost:
			if cfg.postHandler != nil {
				cfg.postHandler(w, r)
			} else {
				http.Error(w, "POST not implemented", http.StatusNotImplemented)
			}
		case http.MethodPut:
			if cfg.putHandler != nil {
				cfg.putHandler(w, r)
			} else {
				http.Error(w, "PUT not implemented", http.StatusNotImplemented)
			}
		case http.MethodDelete:
			if cfg.deleteHandler != nil {
				cfg.deleteHandler(w, r)
			} else {
				http.Error(w, "DELETE not implemented", http.StatusNotImplemented)
			}
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func WithGet(handler HTTPHandler) HandleMethod {
	return func(cfg *HTTPMethodsCallback) {
		cfg.getHandler = handler
	}
}

func WithPost(handler HTTPHandler) HandleMethod {
	return func(cfg *HTTPMethodsCallback) {
		cfg.postHandler = handler
	}
}

func WithPut(handler HTTPHandler) HandleMethod {
	return func(cfg *HTTPMethodsCallback) {
		cfg.putHandler = handler
	}
}

func WithDelete(handler HTTPHandler) HandleMethod {
	return func(cfg *HTTPMethodsCallback) {
		cfg.deleteHandler = handler
	}
}
