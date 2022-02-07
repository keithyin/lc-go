package algos

import (
	"fmt"
	"strconv"
)

type Stack struct {
	values []interface{}
}

func CreateStack() Stack {
	return Stack{values: make([]interface{}, 0)}
}

func (this *Stack) Push(value interface{}) {
	this.values = append(this.values, value)
}

func (this *Stack) Top() interface{} {
	if this.isEmpty() {
		return nil
	}
	return this.values[len(this.values)-1]
}

func (this *Stack) Pop() interface{} {
	if this.isEmpty() {
		return nil
	}

	val := this.values[len(this.values)-1]
	this.values = this.values[:len(this.values)-1]
	return val
}

func (this *Stack) Size() int {
	return len(this.values)
}

func (this *Stack) isEmpty() bool {
	return len(this.values) == 0
}

func DecodeString(s string) string {
	result := decodeString(s)
	return result
}

func decodeString(s string) string {
	s = "1[" + s + "]"
	numStack := CreateStack()
	valueStack := CreateStack()
	begin := 0
	for begin < len(s) {
		nextNonIntPos, scanedInt := scanInt(s, begin)
		nextNonStrPos, scanedStr := scanChars(s, begin)
		if nextNonIntPos != begin {
			numStack.Push(parseInt(scanedInt))
			valueStack.Push(s[nextNonIntPos : nextNonIntPos+1])
			begin = nextNonIntPos + 1
			continue
		}
		if nextNonStrPos != begin {
			valueStack.Push(scanedStr)
			begin = nextNonStrPos
			continue
		}

		if s[nextNonIntPos:nextNonIntPos+1] == "[" || s[nextNonStrPos:nextNonStrPos+1] == "[" {
			valueStack.Push("[")
			begin = nextNonIntPos + 1
			continue
		}

		if s[nextNonStrPos] == ']' {
			tmpStack := CreateStack()
			for !valueStack.isEmpty() {
				top := valueStack.Pop().(string)
				if top != "[" {
					tmpStack.Push(top)
				} else {
					tmpRes := ""
					for !tmpStack.isEmpty() {
						tmpRes += tmpStack.Pop().(string)
					}
					times := numStack.Pop().(int)
					for i := 0; i < times; i++ {
						valueStack.Push(tmpRes)
					}
					break
				}

			}
			begin = nextNonStrPos + 1
		}
	}
	return valueStack.Pop().(string)

}

func scanInt(s string, begin int) (nextNonIntPos int, intStr string) {
	nextNonIntPos = begin

	for isInt(s[nextNonIntPos]) && nextNonIntPos < len(s) {
		nextNonIntPos++
	}

	return nextNonIntPos, s[begin:nextNonIntPos]
}

func parseInt(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("parse str[%s] to int error", str))
	}
	return result
}

func scanChars(s string, begin int) (nextNonCharPos int, scanedStr string) {
	nextNonCharPos = begin
	for nextNonCharPos < len(s) && isAlphabet(s[nextNonCharPos]) {
		nextNonCharPos++
	}

	return nextNonCharPos, s[begin:nextNonCharPos]
}

func isInt(b byte) bool {
	return b >= '0' && b <= '9'
}

func isAlphabet(b byte) bool {
	return b >= 'a' && b <= 'z'
}
