package behavior

type Iterator interface {
	GetNext() *User
	HasNext() bool
}

type Collection interface {
	CreateIterator() Iterator
}

type User struct {
	name string
	age int
}

type UserIterator struct {
	index int
	users []*User
}
func (ui *UserIterator) HasNext() bool {
	return ui.index < len(ui.users)
}
func (ui *UserIterator) GetNext() *User {
	if ui.HasNext() {
		user := ui.users[ui.index]
		ui.index ++
		return user
	}
	return nil
}

type UserCollection struct {
	users []*User
}
func (uc *UserCollection) CreateIterator() Iterator {
	return &UserIterator{
		users: uc.users,
	}
}