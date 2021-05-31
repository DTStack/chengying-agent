package apibase

import "fmt"

type ApiParameterErrors struct {
	errors map[string]error
}

func NewApiParameterErrors() *ApiParameterErrors {
	return &ApiParameterErrors{
		errors: map[string]error{},
	}
}

func (errs *ApiParameterErrors) Error() string {
	str := ""
	for param, err := range errs.errors {
		str += fmt.Sprintf("parameter(%s): %v, ", param, err)
	}
	return str
}

func (errs *ApiParameterErrors) AppendError(name string, err interface{}, args ...interface{}) {
	if e, ok := err.(error); ok {
		errs.errors[name] = e
	} else if s, ok := err.(string); ok {
		errs.errors[name] = fmt.Errorf(s, args...)
	} else {
		errs.errors[name] = fmt.Errorf("%+v", err)
	}
}

func IsApiParameterErrors(err interface{}) bool {
	_, ok := err.(*ApiParameterErrors)
	return ok
}

func (err *ApiParameterErrors) CheckAndThrowApiParameterErrors() {
	if err != nil && len(err.errors) > 0 {
		panic(err)
	}
}

type DBModelError struct {
	err error
}

func (e *DBModelError) Error() string {
	if e.err != nil {
		return e.err.Error()
	} else {
		return "未知的数据错误: Unknown error from DB model"
	}
}

func ThrowDBModelError(errArgs ...interface{}) {
	var err error
	if len(errArgs) > 0 {
		if e, ok := errArgs[0].(error); ok {
			err = e
			err = fmt.Errorf("获取agent数据失败: %s", e.Error())
		} else if format, ok := errArgs[0].(string); ok {
			if len(errArgs) > 1 {
				err = fmt.Errorf(format, errArgs[1:])
			} else {
				err = fmt.Errorf(format)
			}
		}
	} else {
		err = fmt.Errorf("未知的数据库错误: Unknown DB Error")
	}
	panic(&DBModelError{err})
}

func IsDBModelError(err interface{}) bool {
	_, ok := err.(*DBModelError)
	return ok
}

type RpcHandleError struct {
	err error
}

func (e *RpcHandleError) Error() string {
	if e.err != nil {
		return e.err.Error()
	} else {

		return "Unknown error from RPC handle"
	}
}

func ThrowRpcHandleError(errArgs ...interface{}) {
	var err error
	if len(errArgs) > 0 {
		if e, ok := errArgs[0].(error); ok {
			err = fmt.Errorf("agent执行命令失败: %s", e.Error())
		} else if format, ok := errArgs[0].(string); ok {
			if len(errArgs) > 1 {
				err = fmt.Errorf(format, errArgs[1:])
			} else {
				err = fmt.Errorf(format)
			}
		}
	} else {
		err = fmt.Errorf("Unknown RPC handle Error")
	}
	panic(&RpcHandleError{err})
}

func IsRpcHandleError(err interface{}) bool {
	_, ok := err.(*RpcHandleError)
	return ok
}

type SshHandleError struct {
	err error
}

func (e *SshHandleError) Error() string {
	if e.err != nil {
		return e.err.Error()
	} else {
		return "Unknown error from ssh handler"
	}
}

func ThrowSshHandleError(errArgs ...interface{}) {
	var err error
	if len(errArgs) > 0 {
		if e, ok := errArgs[0].(error); ok {
			err = e
		} else if format, ok := errArgs[0].(string); ok {
			if len(errArgs) > 1 {
				err = fmt.Errorf(format, errArgs[1:])
			} else {
				err = fmt.Errorf(format)
			}
		}
	} else {
		err = fmt.Errorf("Unknown ssh handle Error")
	}
	panic(&SshHandleError{err})
}

func IsSshHandleError(err interface{}) bool {
	_, ok := err.(*SshHandleError)
	return ok
}
