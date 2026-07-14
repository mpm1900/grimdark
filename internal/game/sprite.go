package game

type Sprite struct {
	SpriteURL string
}

func (s *Sprite) Current(g *Game, ctx Context) string {
	return s.SpriteURL
}
