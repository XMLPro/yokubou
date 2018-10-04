package main

type User struct {
	ID       int
	Name     string
	Password string
}

type UserService interface {
	Auth() bool
	AllReactions() []Reaction
}

type Reaction struct {
	From id
	To   id
	Type ReactionType
}

type ReactionService interface {
	All() []Reaction
	Send()
}

type ReactionType int

const (
	Fuck ReactionType = iota
	God
	Cry
	Pray
)

func (t *ReactionType) String() string {
	switch t {
	case Fuck:
		return "糞"
	case God:
		return "神"
	case Cry:
		return "泣"
	case Pray:
		return "祈"
	}
}
