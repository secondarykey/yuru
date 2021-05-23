package window

//Focus
//Active

type Tile struct {
}

type Focuser interface {
	Focus()
}

type Pusher interface {
	Push()
}
