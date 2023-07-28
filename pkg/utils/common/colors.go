package rakstar

const (
	ErrorColor   = 0xeb3434 << 8
	SuccessColor = 0x93eb34 << 8
	ServerColor  = 0xdb34eb << 8
)

const (
	ErrorColorStr       = "{eb3434}"
	SuccessColorStr     = "{93eb34}"
	WarnColorStr        = "{ffac12}"
	WhiteColorStr       = "{ffffff}"
	ServerColorStr      = "{db34eb}"
	ChatLocalColorStr   = "{92caf0}"
	InteractionColorStr = "{ff5ef2}"
)

func RGA(red, green, blue int) int {
	red = clamp(red, 0, 255)
	green = clamp(green, 0, 255)
	blue = clamp(blue, 0, 255)

	color := (int(red) << 16) | (int(green) << 8) | int(blue)
	return color
}

func RGBA(red, green, blue, alpha int) int {
	red = clamp(red, 0, 255)
	green = clamp(green, 0, 255)
	blue = clamp(blue, 0, 255)
	alpha = clamp(alpha, 0, 255)

	color := (int(red) << 24) | (int(green) << 16) | (int(blue) << 8) | int(alpha)
	return color
}

func clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
