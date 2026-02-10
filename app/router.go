package app

import (
	"net/http"

	"github.com/IwanPlamboyan/contact-manajemen-golang/controller"
	"github.com/IwanPlamboyan/contact-manajemen-golang/middleware"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(
		jwtMiddleware *middleware.JWTMiddleware,
		authController *controller.AuthController,
		userController *controller.UserController,
		contactController *controller.ContactController,
		addressController *controller.AddressController,
	) *httprouter.Router {
	router := httprouter.New()
	
	router.POST("/api/auth/register", authController.Register)
	router.POST("/api/auth/login", authController.Login)
	router.POST("/api/auth/refresh", authController.RefreshToken)
	router.POST("/api/auth/logout", authController.Logout)
	router.PATCH("/api/users/current", protect(jwtMiddleware, userController.UpdateMyProfile))
	router.GET("/api/users/current", protect(jwtMiddleware, userController.GetMyProfile))

	router.GET("/api/contacts", protect(jwtMiddleware, contactController.Search))
	router.GET("/api/contacts/:contactId", protect(jwtMiddleware, contactController.GetByID))
	router.POST("/api/contacts", protect(jwtMiddleware, contactController.Create))
	router.PUT("/api/contacts/:contactId", protect(jwtMiddleware, contactController.Update))
	router.DELETE("/api/contacts/:contactId", protect(jwtMiddleware, contactController.Delete))

	router.GET("/api/contacts/:contactId/addresses", protect(jwtMiddleware, addressController.List))
	router.GET("/api/contacts/:contactId/addresses/:addressId", protect(jwtMiddleware, addressController.GetByID))
	router.POST("/api/contacts/:contactId/addresses", protect(jwtMiddleware, addressController.Create))
	router.PUT("/api/contacts/:contactId/addresses/:addressId", protect(jwtMiddleware, addressController.Update))
	router.DELETE("/api/contacts/:contactId/addresses/:addressId", protect(jwtMiddleware, addressController.Delete))
	return router
}

func protect(jwt *middleware.JWTMiddleware, handler httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
        wrapped := jwt.Handle(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            handler(w, r, p)
        }))
        wrapped.ServeHTTP(w, r)
    }
}