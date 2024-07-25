/* cube.go */

package cube

type Dimensions struct {
	Width int
	Length int
	Height int
}

func (d *Dimensions) SetSize (W, L, H int) {
	d.Width = W
	d.Length = L
	d.Height = H
}

func (d *Dimensions) GetVolume () int {
	return d.Width * d.Length * d.Height
}

func (d *Dimensions) GetArea () int {
	return d.Width * d.Length
}
