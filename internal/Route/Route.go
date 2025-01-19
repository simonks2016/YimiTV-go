package Route

type Route int

const (
	ValidateToken Route = iota
)

func (token Route) Path() string {
	return "/app/validate_token"
}

func GetRoute(host string, r Route) string {
	return host + r.Path()
}
