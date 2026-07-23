package simplemap

import (
	"testing"
)

type Class struct {
	name     string
	students []string
}

type TestCases struct {
	name  string
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
	Classes := NewStudentClass()

	tests := NewTestCases()

	var mathClass []string
	var physicsClass []string
	var biologyClass []string

	var mathAbsentees []string
	var mathStudents []string
	var physicsAbsentees []string
	var physicsStudents []string
	var biologyAbsentees []string
	var biologyStudents []string

	for _, test := range tests {
		switch test.name {
		case "Test One":
			mathClass = test.input.students
		case "Test Two":
			physicsClass = test.input.students
		default:
			biologyClass = test.input.students
		}
	}

	for i := 0; i < len(mathClass); i++ {
		if mathClass[i] == physicsClass[i] {
			mathAbsentees = append(mathAbsentees, mathClass[i])
			physicsStudents = append(physicsStudents, physicsClass[i])
		} else {
			mathStudents = append(mathStudents, mathClass[i])
			physicsAbsentees = append(physicsAbsentees, physicsClass[i])
		}
		if mathClass[i] == biologyClass[i] {
			biologyAbsentees = append(biologyAbsentees, biologyClass[i])
		} else {
			biologyStudents = append(biologyStudents, biologyClass[i])
		}
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			switch test.input.name {
			case "Mathematics 101":
				mathClass := Classes.CreateClass(test.input.name, test.input.students)
				Classes.TakeAttendance(mathClass, mathStudents)
				testMathAbsentees := Classes.CheckAttendance(mathClass)
				if !EqualIgnortOrderSort(mathAbsentees, testMathAbsentees) {
					t.Errorf("\n Got: %v \n Want: %v \n", testMathAbsentees, mathAbsentees)
				}
			case "Physics 101":
				physicsClass := Classes.CreateClass(test.input.name, test.input.students)
				Classes.TakeAttendance(physicsClass, physicsStudents)
				testPhysicsAbsentees := Classes.CheckAttendance(physicsClass)
				if !EqualIgnortOrderSort(physicsAbsentees, testPhysicsAbsentees) {
					t.Errorf("\n Got: %v \n Want: %v \n", testPhysicsAbsentees, physicsAbsentees)
				}
			default:
				biologyClass := Classes.CreateClass(test.input.name, test.input.students)
				Classes.TakeAttendance(biologyClass, biologyStudents)
				testBiologyAbsentees := Classes.CheckAttendance(biologyClass)
				if !EqualIgnortOrderSort(biologyAbsentees, testBiologyAbsentees) {
					t.Errorf("\n Got: %v \n Want: %v \n", testBiologyAbsentees, biologyAbsentees)
				}
			}
		})
	}

}
