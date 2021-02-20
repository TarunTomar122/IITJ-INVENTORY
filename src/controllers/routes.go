package controllers

//InitializeRoutes ...
func (s *Server) InitializeRoutes() {
	r := s.Router

	api := r.Group("/api")
	{
		api.GET("/getHome", s.GetHome)
		api.POST("/checkEmailExists", s.CheckIfEmailExists)
	}

	auth := r.Group(("/auth"))
	{
		auth.GET("/login", s.Login)
		auth.GET("/callback", s.Callback)
		auth.GET("error", s.Error)
	}
}
