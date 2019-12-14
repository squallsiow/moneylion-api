package handler

// Handler : Singleton Handler
type Handler struct {
}

// New :
func New() (*Handler, error) {
	ctrl := &Handler{}
	return ctrl, nil
}
