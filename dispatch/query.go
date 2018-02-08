package dispatch

type (
	Query struct {
		QMap      map[string]string
		QKeys     []string
		IsOrdered bool
	}
)
