package image

import (
	"bytes"
	"encoding/base64"
	"image/color"
	"strings"

	"github.com/jbuchbinder/gg"
	"github.com/jiro4989/textimg/v3/token"
)

func Draw(tokens token.Tokens) (string, error) {
	foreground := color.RGBA{205, 214, 244, 255}
	background := color.RGBA{30, 30, 46, 255}

	dc := gg.NewContext(1200, 630)
	fgCol := foreground
	bgCol := background
	dc.SetColor(bgCol)
	if err := dc.LoadFontFace("fonts/FiraCode-Regular.ttf", 14); err != nil {
		return "", err
	}
	dc.Clear()
	curX, curY := 0.0, 0.0

	for _, t := range tokens {
		switch t.Kind {
		case token.KindColor:
			switch t.ColorType {
			case token.ColorTypeReset:
				fgCol = foreground
				bgCol = background
			case token.ColorTypeResetForeground:
				fgCol = foreground
			case token.ColorTypeResetBackground:
				bgCol = background
			case token.ColorTypeReverse:
				fgCol, bgCol = bgCol, fgCol
			case token.ColorTypeForeground:
				fgCol = color.RGBA(t.Color)
			case token.ColorTypeBackground:
				bgCol = color.RGBA(t.Color)
			}
		case token.KindText:
			w, h := dc.MeasureMultilineString(t.Text, 1.0)
			dc.SetColor(bgCol)
			dc.DrawRectangle(curX, curY, w, h)
			dc.SetColor(fgCol)
			dc.DrawStringAnchored(strings.ReplaceAll(strings.ReplaceAll(t.Text, "\t", "    "), "\n", ""), curX, curY, 0.0, 1.0)
			curX += w
			if strings.Contains(t.Text, "\n") {
				curY += h
				curX = 0
			}
		}
	}
	buffer := new(bytes.Buffer)
	dc.EncodePNG(buffer)
	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(buffer.Bytes()), nil
}
