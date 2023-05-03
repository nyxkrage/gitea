package image

import (
	"bytes"
	"image/color"
	"strings"

	"github.com/jbuchbinder/gg"
	"github.com/jiro4989/textimg/v3/token"
)

func Draw(tokens token.Tokens) ([]byte, error) {
	foreground := color.RGBA{205, 214, 244, 255}
	background := color.RGBA{30, 30, 46, 255}

	dc := gg.NewContext(1200, 630)
	fgCol := foreground
	bgCol := background
	dc.SetColor(bgCol)
	if err := dc.LoadFontFace("fonts/FiraCode-Regular.ttf", 14); err != nil {
		return nil, err
	}
	dc.Clear()
	curX, curY := 50.0, 50.0

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
			dc.Push()
			dc.SetColor(bgCol)
			dc.DrawRectangle(curX, curY, w, h)
			dc.Fill()
			dc.Pop()
			dc.SetColor(fgCol)
			dc.DrawStringAnchored(strings.ReplaceAll(strings.ReplaceAll(t.Text, "\t", "    "), "\n", ""), curX, curY, 0.0, 1.0)
			curX += w
			if strings.Contains(t.Text, "\n") {
				curY += h
				curX = 50
			}
		}
	}
	dc.Push()
	dc.SetColor(background)
	dc.DrawRectangle(1150, 0, 50, 630)
	dc.DrawRectangle(0, 580, 1200, 50)
	dc.Fill()
	dc.Pop()
	buffer := new(bytes.Buffer)
	dc.EncodePNG(buffer)
	return buffer.Bytes(), nil
}
