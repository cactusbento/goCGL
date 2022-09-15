package main

func isAlive(c cell) bool {
	if _, ok := Alive[c]; ok {
		return true 
	} else {
		return false 
	}
}

func checkNeighbors(c cell) (number int, deadNeighbors []cell) {
	number = 0
	for ly :=  c.y - 1 ; ly <= c.y + 1 ; ly += 1 {
		for lx := c.x - 1 ; lx <= c.x + 1 ; lx += 1 {
			if ( lx == c.x && ly == c.y ) {
				continue
			}

			if isAlive( cell{x: lx, y: ly} ) {
				number += 1
			} else {
				deadNeighbors = append(deadNeighbors, cell{x: lx, y: ly} )
			}
		}
	}
	return
}

func isInBounds(c cell) bool {
	if c.x < 0 || c.y < 0 || 
		c.x > W || c.y > H {
			return false 
	} else {
		return true 
	}
}

// Main will handle replacing
// Alive with nextGen()
func nextGen() map[cell]bool {
	ng := make(map[cell]bool)

	for c := range Alive {
		var n, d = checkNeighbors(c)

		// Cell will continue into the next generation
		if isAlive(c) && (n == 2 || n == 3) {
			ng[c] = true 
		} 
		
		// Revive a dead cell if eligable
		for _, cd := range d {
			var nd, _ = checkNeighbors(cd)
			if nd == 3 && isInBounds(cd) { ng[cd] = true }
		}

	}

	// No need to kill cells since they
	// never existed in the next generation 

	return ng
}
