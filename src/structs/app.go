package structs

import (
	"github.com/gorilla/mux"
)

// App holds important attributes to use on service, like the Router
type App struct {
	Router *mux.Router
}
