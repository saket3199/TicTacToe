package player

type Player struct {
	PlayerName string
	PlayerTurn bool
}

func New(name string, state bool) *Player {
	player := Player{
		PlayerName: name,
		PlayerTurn: state,
	}
	return &player
}
func (p *Player) SetPlayerName(Name string) {
	p.PlayerName = Name
}
func (p *Player) GetPlayerName() string {
	return p.PlayerName
}
func (p *Player) SetPlayerState(state bool) {
	p.PlayerTurn = state
}
func (p *Player) GetPlayerState() bool {
	return p.PlayerTurn
}
