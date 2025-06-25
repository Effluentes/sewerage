// package handlers

// import (
// 	"sewerage/internal/infrastructure/server"
// 	"log/slog"
// 	"net/http"
// )

// func CombineUserHandlers() *server.HTTPServer {
// 	userEndpoint := server.NewHTTPServer()
// 	userEndpoint.HandleFunc("/default", server.NewMethodHandler(userGet, userPost))
// 	return userEndpoint
// }

// func userGet(w http.ResponseWriter, r *http.Request) {
// 	slog.Info("Handle user")
// 	server.NewMethodHandler()
// 	w.Write([]byte("Hello from Snippetbox"))
// }

// func userPost(w http.ResponseWriter, r *http.Request) {
// 	slog.Info("Handle user")
// 	server.NewMethodHandler()
// 	w.Write([]byte("Hello from Snippetbox"))
// }

package handlers

import (
	"sewerage/internal/infrastructure/server"
	"sewerage/internal/domain/dto"
	"sewerage/internal/domain/controller"
	"sewerage/internal/domain/repositories"
	"sewerage/internal/domain/services"
	"log/slog"
	"net/http"
    "encoding/json"
)
type UserHandler struct {
	userController *controller.UserController
}

func NewUserHandler(userController *controller.UserController) *UserHandler {
	return &UserHandler{userController: userController}
}

func CombineUserHandlers() *server.HTTPServer {
	userEndpoint := server.NewHTTPServer()

	// Inicjalizacja zależności
	// userRepo := repositories.NewUserRepository(db)
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	userHandler := NewUserHandler(userController)

	// Mapowanie metod
	combinedHandler := server.NewMethodHandler(
		server.WithGet(userHandler.getUser),
		server.WithPost(userHandler.createUser),
	)

	userEndpoint.HandleFunc("/user", combinedHandler)
	return userEndpoint
}

func (userHandler *UserHandler) getUser(w http.ResponseWriter, r *http.Request) {
	slog.Debug("Handle getUser")
	w.Write([]byte("GET handler for /default"))
}

func (userHandler *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	slog.Debug("Handle createUser")

    var req dto.CreateUserRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Info("decoded failed")
		return
    }
	if err := userHandler.userController.CreateUser(&req); err != nil {
		return
	}
	slog.Debug("CreateUser")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Success"))
}

//==========================

// Wersja z dynamicznym routingiem

// func CombineUserHandlers() *server.HTTPServer {
//     userEndpoint := server.NewHTTPServer()

//     // Handler z parametrem ID
//     userByIDHandler := server.NewMethodHandler(
//         server.WithGet(func(w http.ResponseWriter, r *http.Request) {
//             // Pobierz ID z ścieżki (wymaga routera obsługującego parametry)
//             id := chi.URLParam(r, "id") // jeśli używasz chi
//             w.Write([]byte("GET user with ID: " + id))
//         }),
//     )

//     userEndpoint.HandleFunc("/default", server.NewMethodHandler(
//         server.WithGet(userGet),
//         server.WithPost(userPost),
//     ))

//     userEndpoint.HandleFunc("/users/{id}", userByIDHandler)

//     return userEndpoint
// }
//==========================
//Wersja z middleware

// func CombineUserHandlers() *server.HTTPServer {
//     userEndpoint := server.NewHTTPServer()

//     // Middleware dla GET
//     loggedGet := loggingMiddleware(userGet)

//     // Middleware dla POST
//     authPost := authMiddleware(userPost)

//     combinedHandler := server.NewMethodHandler(
//         server.WithGet(loggedGet),
//         server.WithPost(authPost),
//     )

//     userEndpoint.HandleFunc("/default", combinedHandler)
//     return userEndpoint
// }

// func loggingMiddleware(next server.HTTPHandler) server.HTTPHandler {
//     return func(w http.ResponseWriter, r *http.Request) {
//         log.Println("Before GET handler")
//         next(w, r)
//         log.Println("After GET handler")
//     }
// }

// func authMiddleware(next server.HTTPHandler) server.HTTPHandler {
//     return func(w http.ResponseWriter, r *http.Request) {
//         if r.Header.Get("Authorization") == "" {
//             http.Error(w, "Unauthorized", http.StatusUnauthorized)
//             return
//         }
//         next(w, r)
//     }
// }
//==========================
