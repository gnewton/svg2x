package pkg

import (
	"errors"
	"log"
	"regexp"
	"strconv"
)

type Point struct {
	X, Y float64
}

type State int

const (
	start State = iota
	inInteger
	inMantissa
)

func (s State) String() string {
	return [...]string{"start", "inInteger", "inMantissa", "newNumber"}[s]
}

var numRE = regexp.MustCompile("[0-9]")

var cb = make([]byte, 1)

// l.37,0,1-2
// c-2.232,1.152-6.913,2.304-12.817,2.304
//c-13.682,0-23.906-8.641-23.906-24.626' +
//     'c0-15.266,10.297-25.49,25.346-25.49c5.977,0,9.865,1.296,11.521,2.16l-1.584,5.112C66.747,9.134,63.363,8.27,59.33,8.27' +
//     'c-11.377,0-18.938,7.272-18.938,20.018c0,11.953,6.841,19.514,18.578,19.514c3.888,0,7.777-0.792,10.297-2.016L70.491,50.826z',
// 'M10,10',

// Assume leading command character is still in place
func ParseNumbers(s string) ([]float64, error) {
	if len(s) == 0 {
		return nil, errors.New("Empty string")
	}

	if len(s) == 1 {
		return nil, errors.New("String too short; len=1")
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	numbers := make([]float64, 0)
	state := start
	thisNum := ""
	log.Println(s)
	for i := 1; i < len(s); i++ {
		c := s[i]
		switch c {
		case '-':
			log.Println("minus", string(c), state)
			if state == start {
				state = inInteger
				thisNum += "-"
			} else {
				if state == inInteger || state == inMantissa {
					// new number
					thisFloat, err := strconv.ParseFloat(thisNum, 64)
					if err != nil {
						log.Println(err)
						return nil, err
					}
					numbers = append(numbers, thisFloat)
					thisNum = "-"
					state = inInteger
				}
			}
		case '.':
			log.Println("period", string(c), state)
			if state == start || state == inInteger {
				state = inMantissa
				thisNum += "."
			} else {
				if state == inMantissa {
					// new number
					thisFloat, err := strconv.ParseFloat(thisNum, 64)
					if err != nil {
						log.Println(err)
						return nil, err
					}
					numbers = append(numbers, thisFloat)
					thisNum = "."
					state = inMantissa
				}
			}

		case ',', ' ', '\t', '\n', '\f', '\r':
			log.Println("comma whitespace", string(c), state)
			if state == start {
				continue
			}
			// new number
			thisFloat, err := strconv.ParseFloat(thisNum, 64)
			if err != nil {
				return nil, err
			}
			numbers = append(numbers, thisFloat)
			thisNum = ""
			state = start

		default:
			log.Println("default", string(c), state)
			cb[0] = c
			if numRE.Match(cb) {
				thisNum += string(c)
				if state == start {
					state = inInteger
				}
			}
		}
	}
	log.Println("thisNum", thisNum, len(thisNum))
	if len(thisNum) > 0 {
		thisFloat, err := strconv.ParseFloat(thisNum, 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, thisFloat)
	}
	return numbers, nil
}
