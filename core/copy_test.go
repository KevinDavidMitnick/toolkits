package core

import (
	"testing"
)

type Person struct {
	Name string
	Age  int
}

type Student struct {
	Person
	NO int
}

func TestDeepCopy(t *testing.T) {
	var a, b int
	b = 1
	DeepCopy(&a, b)
	if a != b {
		t.Errorf("TestDeepCopy:\n Expect => %v\n Got    => %v\n", b, a)
	}

	var str1, str2 string
	str2 = "hello world."
	DeepCopy(&str1, str2)
	if str1 != str2 {
		t.Errorf("TestDeepCopy:\n Expect => %v\n Got    => %v\n", b, a)
	}

	var map1, map2 map[string]map[string]int
	inner := make(map[string]int)
	inner["a"] = 1
	inner["b"] = 2
	map2 = map[string]map[string]int{"c": inner}
	DeepCopy(&map1, map2)
	if map1["c"]["a"] != map2["c"]["a"] {
		t.Errorf("TestDeepCopy:\n Expect => %v\n Got    => %v\n",
			map2["c"]["a"], map1["c"]["a"])
	}

	if map1["c"]["b"] != map2["c"]["b"] {
		t.Errorf("TestDeepCopy:\n Expect => %v\n Got    => %v\n",
			map2["c"]["b"], map1["c"]["b"])
	}

	var slice1, slice2 []string
	slice2 = []string{"a", "b", "c"}
	DeepCopy(&slice1, slice2)
	if slice1[2] != slice2[2] {
		t.Errorf("TestDeepCopy:\n Expect => %v\n Got    => %v\n", slice2[2], slice1[2])
	}

	var stu1, stu2 Student
	stu2.Age = 27
	stu2.NO = 1
	stu2.Name = "管庆敏"

	DeepCopy(&stu1, stu2)
	if stu1.Name != "管庆敏" {
		t.Errorf("TestDeepCopy:\n Expect => %v\n Got    => %v\n", stu2.Name, stu1.Name)
	}
}
