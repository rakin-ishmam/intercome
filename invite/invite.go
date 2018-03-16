package invite

// Inviter wraps custome invite functionality
type Inviter interface {
	Read(fileLoc string)
	List(Filter) []Customer
	Err() error
}
