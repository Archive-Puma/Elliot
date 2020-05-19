package cli

func calculatePosition(width int, height int, view View) (int, int, int, int) {
	x := view.x
	if x < -1 {
		x = width + x
	}
	w := view.w
	if w <= 0 {
		w = width + w
	}
	y := view.y
	if y < -1 {
		y = height + y
	}
	h := view.h
	if h <= 0 {
		h = height + h
	}

	return x, y, w, h
}
