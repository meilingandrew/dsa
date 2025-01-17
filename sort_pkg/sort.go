package sort_pkg

const (
	BULKY_VOL       = 1000000 // cm^3
	BULKY_DIMENSION = 150     // cm
)

type PackageCriteria int

const (
	// If it's volume (Width x Height x Length) is greater than or equal to 1,000,000 cm^3
	// or when one of its dimensions is greater or equal to 150cm
	Default PackageCriteria = iota + 1
	Bulky
	// When it's mass >= 20kg
	Heavy
)

func (p PackageCriteria) String() string {
	return [...]string{"bulky", "heavy"}[p-1]
}

type Stack int

const (
	// not `Bulky` or `Heavy`
	Standard Stack = iota + 1
	// Either `Bulky` or `Heavy`
	Special
	// Both
	Rejected
)

func (s Stack) String() string {
	return [...]string{"standard", "special", "rejected"}[s-1]
}

func sort(width, height, length, mass int) string {
	var (
		sizeCategory   PackageCriteria = Default
		weightCategory PackageCriteria = Default
		stack          Stack           = Standard
	)

	if width <= 0 || height <= 0 || length <= 0 || mass <= 0 {
		stack = Rejected
		return stack.String()
	}

	if (width >= BULKY_DIMENSION || height >= BULKY_DIMENSION || length >= BULKY_DIMENSION) || ((width * height * length) >= BULKY_VOL) {
		sizeCategory = Bulky
	}

	if mass >= 20 {
		weightCategory = Heavy
	}

	if sizeCategory == Default && weightCategory == Default {
		stack = Standard
	} else if sizeCategory == Bulky && weightCategory == Heavy {
		stack = Rejected
	} else {
		stack = Special
	}

	return stack.String()
}
