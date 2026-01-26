package ui

const (
	leftColumnRatio  = 40
	RightColumnRatio = 60

	Widget1HeightRatio = 10
	Widget2HeightRatio = 45
	Widget3HeightRatio = 45

	Widget4HeightRatio = 70
	Widget5HeightRatio = 30
)

func NewDefaultLayout(maxX, maxY int) []*Widget {
	leftWidth := maxX * leftColumnRatio / 100
	rightStart := leftWidth

	Widget1Height := maxY * Widget1HeightRatio / 100
	Widget2Height := maxY * Widget2HeightRatio / 100

	Widget4Height := maxY * Widget4HeightRatio / 100

	return []*Widget{
		{
			Name: "widget1",
			Body: "widget1",
			X0:   0,
			Y0:   0,
			X1:   leftWidth - 1,
			Y1:   Widget1Height - 1,
		},
		{
			Name: "widget2",
			Body: "widget2",
			X0:   0,
			Y0:   Widget1Height,
			X1:   leftWidth - 1,
			Y1:   Widget1Height + Widget2Height - 1,
		},
		{
			Name: "widget3",
			Body: "widget3",
			X0:   0,
			Y0:   Widget1Height + Widget2Height,
			X1:   leftWidth - 1,
			Y1:   maxY - 1,
		},
		{
			Name: "widget4",
			Body: "widget4",
			X0:   rightStart,
			Y0:   0,
			X1:   maxX - 1,
			Y1:   Widget4Height - 1,
		},
		{
			Name: "widget5",
			Body: "widget5",
			X0:   rightStart,
			Y0:   Widget4Height,
			X1:   maxX - 1,
			Y1:   maxY - 1,
		},
	}
}
