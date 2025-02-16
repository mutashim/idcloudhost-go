package idcloudhost

type Storage struct {
	CreatedAt string
	ID        int64
	Name      string
	Pool      string
	Primary   bool
	Replica   interface{}
	Shared    bool
	Size      int64
	Type      string
	UpdatedAt interface{}
	UserID    int64
	UUID      string
}
