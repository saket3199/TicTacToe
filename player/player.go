package player

type Player struct {
	PlayerName string
	PlayerTurn bool
	PlayerMark string
}

func New(name string, state bool) *Player {
	return &Player{
		PlayerName: name,
		PlayerTurn: state,
		// PlayerMark: mark,
	}
}
func (p *Player) SetPlayerMark(Mark string) {
	p.PlayerMark = Mark
}

func (p *Player) GetPlayerMark() string {
	return p.PlayerMark
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
