package main
type APIserver struct {
	listenAddr string
}
func NewAPIServer(listenAddr string) *APIserver {
	return &APIserver{listenAddr: listenAddr}
}