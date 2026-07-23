package simplemap

import (
	"fmt"
	"slices"
)

type StudentClass struct {
	className string
	attendance map[string]bool
}

func NewStudentClass() StudentClass {
	return StudentClass{}
}

func (sc *StudentClass) CreateClass(className string, students []string) StudentClass {
	StudentsInClass := map[string]bool{}
	for _, student := range students {
		StudentsInClass[student] = false
	}

	fmt.Println(StudentsInClass)

	NewClass := StudentClass{
		className: className,
		attendance: StudentsInClass,
	}
	return NewClass
}

func(sc *StudentClass) TakeAttendance(class StudentClass, students []string) {
	for _, student := range students {
		class.attendance[student] = true
	}
}

func(sc *StudentClass) CheckAttendance(class StudentClass) []string {
	var absentees []string

	for student, isPresent := range class.attendance {
		if !isPresent {
			absentees = append(absentees, student)
		}
	}

	return absentees
}

func EqualIgnortOrderSort(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	// Create copies to avoid mutating the original data
	aCopy := slices.Clone(a)
	bCopy := slices.Clone(b)

	// Sort both collections lexicographically
	slices.Sort(aCopy)
	slices.Sort(bCopy)

	// Perform an index-to-index comparison
	return slices.Equal(aCopy, bCopy)
}