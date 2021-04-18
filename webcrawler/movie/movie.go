package movie

// Movie contains the basic information about a movie for enabling our searching API
// to find and return relevant movies to the user
type Movie struct {
	ID       string
	Name     string
	Director string
	Cast     []string
}
