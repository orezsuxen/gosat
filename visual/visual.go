package visual

import (
	vec "local/gosat/vector"

	"github.com/veandco/go-sdl2/sdl"
)

type TextLabel struct {
	text     string
	position vec.Vector
	scale    vec.Vector
	spacing  vec.Vector
}

type TextDrawer struct {
	symbols         map[byte]uint64 //arount 93 symbols right now
	rend            *sdl.Renderer
	texture         *sdl.Texture        //draw once then copy ?
	texturePosition map[byte]vec.Vector // where symbols are on the texture(x,y)
	textureScale    vec.Vector          // size of symbols in texture(w,h)
}

func (t *TextDrawer) Init() (ok bool) {

	tmp, err := t.rend.CreateTexture(
		sdl.PIXELFORMAT_RGBA8888,
		sdl.TEXTUREACCESS_TARGET,
		800, 600, //size
	)
	if err != nil {
		return false
	}
	t.texture = tmp
	//TODO: render symbols to texture

	return true

}

func NewTextDrawer(r *sdl.Renderer) TextDrawer {
	return TextDrawer{
		rend: r,
		symbols: map[byte]uint64{
			'A': 0x006666667E663C18,
			'B': 0x003E66663E66663E,
			'C': 0x003C66060606663C,
			'D': 0x001E36666666361E,
			'E': 0x007E06061E06067E,
			'F': 0x000606061E06067E,
			'G': 0x003C66667606663C,
			'H': 0x006666667E666666,
			'I': 0x003C18181818183C,
			'J': 0x001C363030303078,
			'K': 0x0066361E0E1E3666,
			'L': 0x007E060606060606,
			'M': 0x00C6C6C6D6FEEEC6,
			'N': 0x006666767E7E6E66,
			'O': 0x003C66666666663C,
			'P': 0x000606063E66663E,
			'Q': 0x00703C666666663C,
			'R': 0x0066361E3E66663E,
			'S': 0x003C66603C06663C,
			'T': 0x001818181818187E,
			'U': 0x003C666666666666,
			'V': 0x00183C6666666666,
			'W': 0x00C6EEFED6C6C6C6,
			'X': 0x0066663C183C6666,
			'Y': 0x001818183C666666,
			'Z': 0x007E060C1830607E,

			'a': 0x007C667C603C0000,
			'b': 0x003E66663E060600,
			'c': 0x003C0606063C0000,
			'd': 0x007C66667C606000,
			'e': 0x003C067E663C0000,
			'f': 0x001818187C187000,
			'g': 0x3E607C66667C0000,
			'h': 0x006666663E060600,
			'i': 0x003C18181C001800,
			'j': 0x3C60606060006000,
			'k': 0x0066361E36060600,
			'l': 0x003C181818181C00,
			'm': 0x00C6D6FEFE660000,
			'n': 0x00666666663E0000,
			'o': 0x003C6666663C0000,
			'p': 0x06063E66663E0000,
			'q': 0x60607C66667C0000,
			'r': 0x00060606663E0000,
			's': 0x003E603C067C0000,
			't': 0x00701818187E1800,
			'u': 0x007C666666660000,
			'v': 0x00183C6666660000,
			'w': 0x006C7CFED6C60000,
			'x': 0x00663C183C660000,
			'y': 0x1E307C6666660000,
			'z': 0x007E0C18307E0000,

			'0': 0x003C66666E76663C,
			'1': 0x007E1818181C1818,
			'2': 0x007E060C3060663C,
			'3': 0x003C66603860663C,
			'4': 0x006060FE66787060,
			'5': 0x003C6660603E067E,
			'6': 0x003C66663E06663C,
			'7': 0x001818181830667E,
			'8': 0x003C66663C66663C,
			'9': 0x003C66607C66663C,

			'!':  0x0018000018181818,
			'"':  0x0000000000666666,
			'#':  0x0066667E66FF6666,
			'$':  0x00183E603C067C18,
			'%':  0x0062660C18306646,
			'&':  0x00FC66E61C3C663C,
			'\'': 0x0000000000181818,
			'(':  0x0030180C0C0C1830,
			')':  0x000C18303030180C,
			'*':  0x0000663CFF3C6600,
			'+':  0x000018187E181800,
			',':  0x0C18180000000000,
			'-':  0x000000007E000000,
			'.':  0x0018180000000000,
			'/':  0x00060C183060C000,

			':': 0x0000180000180000,
			';': 0x0C18180000180000,
			'<': 0x0070180C060C1870,
			'=': 0x0000007E007E0000,
			'>': 0x000E18306030180E,
			'?': 0x001800183060663C,
			'@': 0x003C46067676663C,

			'[':  0x003C0C0C0C0C0C3C,
			'\\': 0x00C06030180C0600,
			']':  0x003C30303030303C,
			'^':  0x0000000000663C18,
			'_':  0x00FF000000000000,

			'{': 0x00381C180E181C38,
			'|': 0x0018181818181818,
			'}': 0x001C38183018381C,
			'~': 0x000000337ECC0000,
		},
	}
}
