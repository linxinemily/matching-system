package domain

import (
	"matching-system/domain/enum"
	"strconv"
	"strings"
)

type Individual struct {
	id     int
	gender enum.Gender
	age    int
	intro  string
	habits string
	coord  string
}

func NewIndividual(id int, gender enum.Gender, age int, intro string, habits string, coord string) *Individual {
	return &Individual{
		id,
		gender,
		age,
		intro,
		habits,
		coord,
	}
}

func (i *Individual) GetXY() (x, y int) {
	s := strings.Split(i.coord, ",")

	_x := strings.TrimPrefix(s[0], "(")
	_y := strings.TrimSuffix(s[1], ")")

	intX, _ := strconv.Atoi(_x)
	intY, _ := strconv.Atoi(_y)

	return intX, intY
}

func (i *Individual) GetHabits() []string {
	return strings.Split(i.habits, ", ")
}

func (i *Individual) GetId() int {
	return i.id
}
