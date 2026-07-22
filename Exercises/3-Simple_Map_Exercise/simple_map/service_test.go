package simplemap

import (
	"testing"
)

type Class struct {
	name string
	students []string
}

type TestCases struct {
	name string
	input Class
}

func NewTestCases() []TestCases {
	newTestCases := []TestCases{
		{
			name: "Test One",
			input: Class{
				name: "Mathematics 101",
				students: []string{
					"Meghan",
					"Laurent",
					"Noelle",
					"Achilles",
					"Rebecca",
					"Steven",
					"Christina",
				},
			},
		},
		{
			name: "Test Two",
			input: Class{
				name: "Physics 101",
				students: []string{
					"Gabriel",
					"Violet",
					"Sierra",
					"Brandon",
					"Rebecca",
					"Steven",
					"Christina",
				},
			},
		},
		{
			name: "Test Three",
			input: Class{
				name: "Biology 101",
				students: []string{
					"Meghan",
					"Laurent",
					"Noelle",
					"Iman",
					"Garrett",
					"Fionna",
					"Carmichael",
				},
			},
		},
	}

	return newTestCases
}

func TestCreateClass(t *testing.T) {
	Classes := NewStudentClass()

	tests := NewTestCases()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			NewClass := Classes.CreateClass(test.input.name, test.input.students)
			// t.Errorf("New Class Created: %v", NewClass)

			if NewClass.className != test.input.name {
				t.Errorf("\n Got: %v \n Want: %v \n", NewClass.className, test.input.name)
				t.Logf("\n Got: %v \n Want: %v \n", NewClass.className, test.input.name)
			}

			for _, student := range test.input.students {
				if NewClass.attendance[student] {
					t.Errorf("\n Got: %v \n Want: %v \n", true, false)
					t.Logf("\n Got: %v \n Want: %v \n", true, false)
				}
			}
		})
	}
}

func TestAttendance(t *testing.T) {
}