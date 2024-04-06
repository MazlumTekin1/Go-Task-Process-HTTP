package middleware

/*
func AuthMiddleware(jwtService service.JWTService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader != "" {
				bearerToken := strings.Split(authHeader, " ")
				if len(bearerToken) == 2 {
					token, err := jwtService.ValidateToken(bearerToken[1])
					if token.Valid {
						next.ServeHTTP(w, r)
					} else {
						w.WriteHeader(http.StatusUnauthorized)
						w.Write([]byte(err.Error()))
						return
					}
				}
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		})
	}
}
*/
