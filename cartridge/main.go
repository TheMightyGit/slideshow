package cartridge

import (
	"embed"
	"image"

	"github.com/TheMightyGit/marv/marvlib"
	"github.com/TheMightyGit/marv/marvtypes"
)

//go:embed "resources/*"
var Resources embed.FS

const (
	GfxBankFont = iota
	GfxBankGfx1
	GfxBankGfx2
	GfxBankGfx3
	GfxBankGfx4
	GfxBankGfx5
	GfxBankEnd = GfxBankGfx5
)
const (
	MapBankText = iota
	MapBankPics
)
const (
	SpriteText = iota
	Sprite1
	Sprite2
	Sprite3
	Sprite4
	Sprite5
)

func Start() {
	area := marvlib.API.MapBanksGet(MapBankPics).AllocArea(image.Point{1, 1}) // 1x1 at 256x256 tiles
	area.Clear(0, 0)
	for i := GfxBankGfx1; i <= GfxBankEnd; i++ {
		marv.Sprites[i].ChangePos(image.Rectangle{
			Min: image.Point{X: 256, Y: 0}, // off screen at the start
			Max: image.Point{X: 512, Y: 0},
		})
		// marv.Sprites[i].Mode7(true)
		marvlib.API.SpritesGet(i).Show(i, area)
	}

	marvlib.API.SfxBanks(0).PlayLooped()
}

var (
	xpos = -32
)

func Update() {
	rec := image.Rectangle{
		Max: image.Point{X: 256, Y: 240},
	}
	for i := GfxBankGfx1; i <= GfxBankEnd; i++ {
		rec.Min.X = (i * 256) - xpos
		marvlib.API.SpritesGet(i).ChangePos(rec)
	}
	if xpos++; xpos > (GfxBankEnd+1)*256 {
		xpos = 0
	}
}
