package base

type ErrNilPointer struct {
	error
	ErrMsg string
}

func (e *ErrNilPointer) Error() string {
	return "ErrNilPointer : " + e.ErrMsg
}

type ErrEmpty struct {
	error
	ErrMsg string
}

func (e *ErrEmpty) Error() string {
	return "ErrEmpty : " + e.ErrMsg
}

type ErrElementNotFound struct {
	error
	ErrMsg string
}

func (e *ErrElementNotFound) Error() string {
	return "ErrElementNotFound : " + e.ErrMsg
}

type ErrFull struct {
	error
	ErrMsg string
}

func (e *ErrFull) Error() string {
	return "ErrFull : " + e.ErrMsg
}

type ErrNoSuchElem struct {
	error
	ErrMsg string
}

func (e *ErrNoSuchElem) Error() string {
	return "ErrNoSuchElem : " + e.ErrMsg
}

type ErrIllegalArgument struct {
	error
	ErrMsg string
}

func (e *ErrIllegalArgument) Error() string {
	return "ErrIllegalArgument : " + e.ErrMsg
}

type ErrIllegalState struct {
	error
	ErrMsg string
}

func (e *ErrIllegalState) Error() string {
	return "ErrIllegalState : " + e.ErrMsg
}

type ErrNotImplemented struct {
	error
	ErrMsg string
}

func (e *ErrNotImplemented) Error() string {
	return "ErrNotImplemented : " + e.ErrMsg
}

type ErrOutOfRange struct {
	error
	ErrMsg string
}

func (e *ErrOutOfRange) Error() string {
	return "ErrOutOfRange : " + e.ErrMsg
}

type ErrCallNotSatisfied struct {
	error
	ErrMsg string
}

func (e *ErrCallNotSatisfied) Error() string {
	return "ErrCallNotSatisfied : " + e.ErrMsg
}
