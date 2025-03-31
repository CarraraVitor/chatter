package handlers

var PublicRoutes = make(map[string]bool)

func init() {
    PublicRoutes["/login"] = true
    PublicRoutes["/register"] = true
    PublicRoutes["/static/favicon.png"] = true
    PublicRoutes["/static/logonobg.png"] = true
    PublicRoutes["/static/atombg.png"] = true
}
