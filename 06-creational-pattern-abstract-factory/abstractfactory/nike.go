package abstractfactory

type Nike struct {
}

func (a *Nike) MakeShoe() IShoe {
	return &nikeShoe{
		shoe: shoe{
			logo: "nike",
			size: 12,
		},
	}
}

func (a *Nike) MakeShort() IShort {
	return &nikeShort{
		short: short{
			logo: "nike",
			size: 12,
		},
	}
}
