// Code generated by thriftgo (0.3.13). DO NOT EDIT.

package base

import (
	"fmt"
	"strings"

	"github.com/apache/thrift/lib/go/thrift"
)

type Base struct {
	Code *int32  `thrift:"Code,1,optional" frugal:"1,optional,i32" json:"code,omitempty"`
	Msg  *string `thrift:"Msg,2,optional" frugal:"2,optional,string" json:"msg,omitempty"`
}

func NewBase() *Base {
	return &Base{}
}

func (p *Base) InitDefault() {
}

var Base_Code_DEFAULT int32

func (p *Base) GetCode() (v int32) {
	if !p.IsSetCode() {
		return Base_Code_DEFAULT
	}
	return *p.Code
}

var Base_Msg_DEFAULT string

func (p *Base) GetMsg() (v string) {
	if !p.IsSetMsg() {
		return Base_Msg_DEFAULT
	}
	return *p.Msg
}
func (p *Base) SetCode(val *int32) {
	p.Code = val
}
func (p *Base) SetMsg(val *string) {
	p.Msg = val
}

var fieldIDToName_Base = map[int16]string{
	1: "Code",
	2: "Msg",
}

func (p *Base) IsSetCode() bool {
	return p.Code != nil
}

func (p *Base) IsSetMsg() bool {
	return p.Msg != nil
}

func (p *Base) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_Base[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *Base) ReadField1(iprot thrift.TProtocol) error {

	var _field *int32
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		_field = &v
	}
	p.Code = _field
	return nil
}
func (p *Base) ReadField2(iprot thrift.TProtocol) error {

	var _field *string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = &v
	}
	p.Msg = _field
	return nil
}

func (p *Base) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("Base"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *Base) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetCode() {
		if err = oprot.WriteFieldBegin("Code", thrift.I32, 1); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteI32(*p.Code); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *Base) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetMsg() {
		if err = oprot.WriteFieldBegin("Msg", thrift.STRING, 2); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.Msg); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *Base) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Base(%+v)", *p)

}

func (p *Base) DeepEqual(ano *Base) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Code) {
		return false
	}
	if !p.Field2DeepEqual(ano.Msg) {
		return false
	}
	return true
}

func (p *Base) Field1DeepEqual(src *int32) bool {

	if p.Code == src {
		return true
	} else if p.Code == nil || src == nil {
		return false
	}
	if *p.Code != *src {
		return false
	}
	return true
}
func (p *Base) Field2DeepEqual(src *string) bool {

	if p.Msg == src {
		return true
	} else if p.Msg == nil || src == nil {
		return false
	}
	if strings.Compare(*p.Msg, *src) != 0 {
		return false
	}
	return true
}
