package main

type philosopherPosition int

const (
	north philosopherPosition = iota
	east
	south
	west
)

type chopstickPosition int

const (
	nw chopstickPosition = iota
	ne
	se
	sw
)

type philosopher struct {
	pos      philosopherPosition
	left     *chopstick
	right    *chopstick
	eating   bool
	hasEaten bool
}

type chopstick struct {
	number int
	owner  *philosopher
}

func (p *philosopher) ascii() []string {
	if p.eating {
		return asciiData[p.pos][4]
	}
	if p.left.owner == p && p.right.owner == p {
		return asciiData[p.pos][3]
	}
	if p.right.owner == p {
		return asciiData[p.pos][2]
	}
	if p.left.owner == p {
		return asciiData[p.pos][1]
	}
	return asciiData[p.pos][0]
}

func (p *philosopher) firstStick() *chopstick {
	if p.left.number < p.right.number {
		return p.left
	}
	return p.right
}

func (p *philosopher) secondStick() *chopstick {
	if p.left.number < p.right.number {
		return p.right
	}
	return p.left
}

func (p *philosopher) pickUp(c *chopstick) bool {
	if c.owner == nil {
		c.owner = p
		return true
	}
	return false
}

var asciiData [4][5][]string = [4][5][]string{
	[5][]string{
		[]string{
			" ╭─────╮ ",
			" │     │ ",
			" ╰─────╯ ",
			"   ╭─╮   ",
			"   ╰─╯   ",
		},
		[]string{
			" ╭─────╮ ",
			" │     ││",
			" ╰─────╯│",
			"   ╭─╮   ",
			"   ╰─╯   ",
		},
		[]string{
			" ╭─────╮ ",
			"││     │ ",
			"│╰─────╯ ",
			"   ╭─╮   ",
			"   ╰─╯   ",
		},
		[]string{
			" ╭─────╮ ",
			"││     ││",
			"│╰─────╯│",
			"   ╭─╮   ",
			"   ╰─╯   ",
		},
		[]string{
			" ╭─────╮ ",
			" │     │ ",
			"╲╰─────╯╱",
			" ╲ ╭─╮ ╱ ",
			"   ╰─╯   ",
		},
	},
	[5][]string{
		[]string{
			"          ",
			"    ╭────╮",
			"╭─╮ │    │",
			"╰─╯ │    │",
			"    ╰────╯",
			"          ",
		},
		[]string{
			"          ",
			"    ╭────╮",
			"╭─╮ │    │",
			"╰─╯ │    │",
			"    ╰────╯",
			"  ──────  ",
		},
		[]string{
			"  ──────  ",
			"    ╭────╮",
			"╭─╮ │    │",
			"╰─╯ │    │",
			"    ╰────╯",
			"          ",
		},
		[]string{
			"  ──────  ",
			"    ╭────╮",
			"╭─╮ │    │",
			"╰─╯ │    │",
			"    ╰────╯",
			"  ──────  ",
		},
		[]string{
			"    ╱     ",
			"   ╱╭────╮",
			"╭─╮ │    │",
			"╰─╯ │    │",
			"   ╲╰────╯",
			"    ╲     ",
		},
	},
	[5][]string{
		[]string{
			"   ╭─╮   ",
			"   ╰─╯   ",
			" ╭─────╮ ",
			" │     │ ",
			" ╰─────╯ ",
		},
		[]string{
			"   ╭─╮   ",
			"   ╰─╯   ",
			"│╭─────╮ ",
			"││     │ ",
			" ╰─────╯ ",
		},
		[]string{
			"   ╭─╮   ",
			"   ╰─╯   ",
			" ╭─────╮│",
			" │     ││",
			" ╰─────╯ ",
		},
		[]string{
			"   ╭─╮   ",
			"   ╰─╯   ",
			"│╭─────╮│",
			"││     ││",
			" ╰─────╯ ",
		},
		[]string{
			"   ╭─╮   ",
			" ╱ ╰─╯ ╲ ",
			"╱╭─────╮╲",
			" │     │ ",
			" ╰─────╯ ",
		},
	},
	[5][]string{
		[]string{
			"          ",
			"╭────╮    ",
			"│    │ ╭─╮",
			"│    │ ╰─╯",
			"╰────╯    ",
			"          ",
		},
		[]string{
			"  ──────  ",
			"╭────╮    ",
			"│    │ ╭─╮",
			"│    │ ╰─╯",
			"╰────╯    ",
			"          ",
		},
		[]string{
			"          ",
			"╭────╮    ",
			"│    │ ╭─╮",
			"│    │ ╰─╯",
			"╰────╯    ",
			"  ──────  ",
		},
		[]string{
			"  ──────  ",
			"╭────╮    ",
			"│    │ ╭─╮",
			"│    │ ╰─╯",
			"╰────╯    ",
			"  ──────  ",
		},
		[]string{
			"     ╲    ",
			"╭────╮╲   ",
			"│    │ ╭─╮",
			"│    │ ╰─╯",
			"╰────╯╱   ",
			"     ╱    ",
		},
	},
}
