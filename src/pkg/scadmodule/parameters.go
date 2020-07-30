package scadmodule

type twoDimensions struct {
	x, y int
}

type threeDimensions struct {
	twoDimensions
	z int
}

type vector struct { threeDimensions }

type cubeSize struct { threeDimensions }
