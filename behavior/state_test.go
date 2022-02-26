package behavior

import "fmt"

type AudioPlayer struct {
	state ApState
	UI int64
	volume int64
	playlist []string
	currentSong string
}

func NewAudioPlayer(initState ApState) *AudioPlayer {
	return &AudioPlayer{
		state: initState,
	}
}
func (a AudioPlayer) clickLock()  {
	a.state.ClickLock()
}
func (a AudioPlayer) ClickPlay()  {
	a.state.ClickPlay()
}
func (a AudioPlayer) ClickNext()  {
	a.state.ClickNext()
}
func (a AudioPlayer) ClickPrevious()  {
	a.state.ClickPrevious()
}
func (a AudioPlayer) SetState(state ApState)  {
	a.state = state
}

type ApState interface {
	ClickLock()
	ClickPlay()
	ClickNext()
	ClickPrevious()
}

type LockedState struct {
	player *AudioPlayer
}
func (ls *LockedState) ClickLock() {
	fmt.Println("Nothing")
}
func (ls *LockedState) ClickPlay() {
	ls.player.SetState(ls)
}
func (ls *LockedState) ClickNext() {
	ls.player.SetState()
}
func (ls *LockedState) ClickPrevious() {

}



type PlayingState struct {
	ApState
}
func (ps *PlayingState) ClickLock() {

}
func (ps *PlayingState) ClickPlay() {

}
func (ps *PlayingState) ClickNext() {

}
func (ps *PlayingState) ClickPrevious() {

}

