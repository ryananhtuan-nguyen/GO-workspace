// The function is not returning the correct values. Can you figure out why?

// Example (Input --> Output ):

// 3 --> "Earth"

//=============WRONG ONE
// package kata

// func GetPlanetName(ID int) string {
//   	case ID {
//         when 1:
//             return "Mercury",
//         when 2:
//             return "Venus",
//         when 3:
//             return "Earth",
//         when 4:
//             return "Mars",
//         when 5:
//             return "Jupiter",
//         when 6:
//             return "Saturn",
//         when 7:
//             return "Uranus",
//         when 8:
//     		    return "Neptune",
//     	  else:
//     		    return "Pluto", // ;-)
//     }
// }

// ================CORRECTED
package kata

func GetPlanetName(ID int) string {
	switch ID {
	case 1:
		return "Mercury"
	case 2:
		return "Venus"
	case 3:
		return "Earth"
	case 4:
		return "Mars"
	case 5:
		return "Jupiter"
	case 6:
		return "Saturn"
	case 7:
		return "Uranus"
	case 8:
		return "Neptune"
	default:
		return "Pluto"
	}
}
