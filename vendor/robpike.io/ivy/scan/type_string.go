// Code generated by "stringer -type Type"; DO NOT EDIT.

package scan

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[EOF-0]
	_ = x[Error-1]
	_ = x[Newline-2]
	_ = x[Assign-3]
	_ = x[Char-4]
	_ = x[Identifier-5]
	_ = x[LeftBrack-6]
	_ = x[LeftParen-7]
	_ = x[Number-8]
	_ = x[Operator-9]
	_ = x[Op-10]
	_ = x[Rational-11]
	_ = x[Complex-12]
	_ = x[RightBrack-13]
	_ = x[RightParen-14]
	_ = x[Semicolon-15]
	_ = x[String-16]
	_ = x[Colon-17]
}

const _Type_name = "EOFErrorNewlineAssignCharIdentifierLeftBrackLeftParenNumberOperatorOpRationalComplexRightBrackRightParenSemicolonStringColon"

var _Type_index = [...]uint8{0, 3, 8, 15, 21, 25, 35, 44, 53, 59, 67, 69, 77, 84, 94, 104, 113, 119, 124}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}