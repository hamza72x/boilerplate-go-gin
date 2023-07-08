package server

func (s *Server) setRoutes() {
	s.setAdminRoutes()
	s.setAccountOwnerRoutes()
}

func (s *Server) setAdminRoutes() {
	admin := s.router.Group("/admin").Use(s.authAdmin())

	admin.POST("/create-account", s.admin_createAccount)
	admin.GET("/list-account", s.admin_listAccount)
	admin.GET("/get-account/:id", s.admin_getAccount)
}

func (s *Server) setAccountOwnerRoutes() {
	acc := s.router.Group("/account").Use(s.auth())

	acc.GET("/", s.accountOwner_getAccount)
	// acc.POST("/", s.accountOwner_updateAccount)
}
