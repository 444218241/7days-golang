package gee

import (
	"log"
	"net/http"
)

type HandlerFunc func(*Context)

type Engine struct {
	*RouteGroup
	router *router
	groups []*RouteGroup // store all groups
}
type RouteGroup struct {
	prefix      string
	middlewares []HandlerFunc // support middleware
	parent      *RouteGroup   // support nesting
	engine      *Engine       // all groups share a Engine instance
}

func (e *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	c := newContext(writer, request)
	e.router.handle(c)
}

// New is the constructor fof gee.Engine
func New() *Engine {
	engine := &Engine{
		router: newRouter(),
	}
	engine.RouteGroup = &RouteGroup{
		engine: engine,
	}
	engine.groups = []*RouteGroup{
		engine.RouteGroup,
	}
	return engine
}

// Group is defined to create a new RouteGroup
// remember all groups share the same Engine instance
func (group *RouteGroup) Group(prefix string) *RouteGroup {
	engine := group.engine
	newGroup := &RouteGroup{
		prefix:      group.prefix + prefix,
		middlewares: nil,
		parent:      group,
		engine:      engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouteGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

func (group *RouteGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouteGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}
