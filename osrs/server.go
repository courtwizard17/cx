func handleRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request for: %s", r.URL.Path)

	path := strings.TrimPrefix(r.URL.Path, "/osrs")

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	log.Printf("Adjusted path: %s", path)

	ext := filepath.Ext(path)
	switch ext {
	case ".html":
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
	case ".css":
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
	case ".js":
		w.Header().Set("Content-Type", "application/javascript")
	case ".png":
		w.Header().Set("Content-Type", "image/png")
	case ".jpg", ".jpeg":
		w.Header().Set("Content-Type", "image/jpeg")
	case ".gif":
		w.Header().Set("Content-Type", "image/gif")
	default:
		w.Header().Set("Content-Type", "application/octet-stream")
	}

	if path == "/" {
		log.Printf("Serving osrs.html")
		http.ServeFile(w, r, "osrs.html")
		return
	}

	filePath := filepath.Join(".", path)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Printf("File not found: %s", filePath)
		http.NotFound(w, r)
		return
	}

	log.Printf("Serving file: %s", filePath)
	http.ServeFile(w, r, filePath)
}
