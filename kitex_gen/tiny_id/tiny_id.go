// Code generated by thriftgo (0.3.4). DO NOT EDIT.

package tiny_id

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"strings"
)

type BaseResp struct {
	Code int64  `thrift:"code,1" frugal:"1,default,i64" json:"code"`
	Msg  string `thrift:"msg,2" frugal:"2,default,string" json:"msg"`
}

func NewBaseResp() *BaseResp {
	return &BaseResp{}
}

func (p *BaseResp) InitDefault() {
	*p = BaseResp{}
}

func (p *BaseResp) GetCode() (v int64) {
	return p.Code
}

func (p *BaseResp) GetMsg() (v string) {
	return p.Msg
}
func (p *BaseResp) SetCode(val int64) {
	p.Code = val
}
func (p *BaseResp) SetMsg(val string) {
	p.Msg = val
}

var fieldIDToName_BaseResp = map[int16]string{
	1: "code",
	2: "msg",
}

func (p *BaseResp) Read(iprot thrift.TProtocol) (err error) {

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
			if fieldTypeId == thrift.I64 {
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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_BaseResp[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *BaseResp) ReadField1(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.Code = v
	}
	return nil
}
func (p *BaseResp) ReadField2(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Msg = v
	}
	return nil
}

func (p *BaseResp) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("BaseResp"); err != nil {
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

func (p *BaseResp) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("code", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.Code); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *BaseResp) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("msg", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Msg); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *BaseResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("BaseResp(%+v)", *p)

}

func (p *BaseResp) DeepEqual(ano *BaseResp) bool {
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

func (p *BaseResp) Field1DeepEqual(src int64) bool {

	if p.Code != src {
		return false
	}
	return true
}
func (p *BaseResp) Field2DeepEqual(src string) bool {

	if strings.Compare(p.Msg, src) != 0 {
		return false
	}
	return true
}

type GetMaxIDRequest struct {
	BizType int64 `thrift:"bizType,1" frugal:"1,default,i64" json:"bizType"`
}

func NewGetMaxIDRequest() *GetMaxIDRequest {
	return &GetMaxIDRequest{}
}

func (p *GetMaxIDRequest) InitDefault() {
	*p = GetMaxIDRequest{}
}

func (p *GetMaxIDRequest) GetBizType() (v int64) {
	return p.BizType
}
func (p *GetMaxIDRequest) SetBizType(val int64) {
	p.BizType = val
}

var fieldIDToName_GetMaxIDRequest = map[int16]string{
	1: "bizType",
}

func (p *GetMaxIDRequest) Read(iprot thrift.TProtocol) (err error) {

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
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField1(iprot); err != nil {
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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GetMaxIDRequest[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *GetMaxIDRequest) ReadField1(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.BizType = v
	}
	return nil
}

func (p *GetMaxIDRequest) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GetMaxIDRequest"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
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

func (p *GetMaxIDRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("bizType", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.BizType); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *GetMaxIDRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetMaxIDRequest(%+v)", *p)

}

func (p *GetMaxIDRequest) DeepEqual(ano *GetMaxIDRequest) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.BizType) {
		return false
	}
	return true
}

func (p *GetMaxIDRequest) Field1DeepEqual(src int64) bool {

	if p.BizType != src {
		return false
	}
	return true
}

type GetMaxIDResponse struct {
	Base  *BaseResp `thrift:"base,1" frugal:"1,default,BaseResp" json:"base"`
	MaxID int64     `thrift:"maxID,2" frugal:"2,default,i64" json:"maxID"`
}

func NewGetMaxIDResponse() *GetMaxIDResponse {
	return &GetMaxIDResponse{}
}

func (p *GetMaxIDResponse) InitDefault() {
	*p = GetMaxIDResponse{}
}

var GetMaxIDResponse_Base_DEFAULT *BaseResp

func (p *GetMaxIDResponse) GetBase() (v *BaseResp) {
	if !p.IsSetBase() {
		return GetMaxIDResponse_Base_DEFAULT
	}
	return p.Base
}

func (p *GetMaxIDResponse) GetMaxID() (v int64) {
	return p.MaxID
}
func (p *GetMaxIDResponse) SetBase(val *BaseResp) {
	p.Base = val
}
func (p *GetMaxIDResponse) SetMaxID(val int64) {
	p.MaxID = val
}

var fieldIDToName_GetMaxIDResponse = map[int16]string{
	1: "base",
	2: "maxID",
}

func (p *GetMaxIDResponse) IsSetBase() bool {
	return p.Base != nil
}

func (p *GetMaxIDResponse) Read(iprot thrift.TProtocol) (err error) {

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
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.I64 {
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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GetMaxIDResponse[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *GetMaxIDResponse) ReadField1(iprot thrift.TProtocol) error {
	p.Base = NewBaseResp()
	if err := p.Base.Read(iprot); err != nil {
		return err
	}
	return nil
}
func (p *GetMaxIDResponse) ReadField2(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.MaxID = v
	}
	return nil
}

func (p *GetMaxIDResponse) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GetMaxIDResponse"); err != nil {
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

func (p *GetMaxIDResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("base", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Base.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *GetMaxIDResponse) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("maxID", thrift.I64, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.MaxID); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *GetMaxIDResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetMaxIDResponse(%+v)", *p)

}

func (p *GetMaxIDResponse) DeepEqual(ano *GetMaxIDResponse) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Base) {
		return false
	}
	if !p.Field2DeepEqual(ano.MaxID) {
		return false
	}
	return true
}

func (p *GetMaxIDResponse) Field1DeepEqual(src *BaseResp) bool {

	if !p.Base.DeepEqual(src) {
		return false
	}
	return true
}
func (p *GetMaxIDResponse) Field2DeepEqual(src int64) bool {

	if p.MaxID != src {
		return false
	}
	return true
}

type TinyIDService interface {
	GetMaxID(ctx context.Context, req *GetMaxIDRequest) (r *GetMaxIDResponse, err error)
}

type TinyIDServiceClient struct {
	c thrift.TClient
}

func NewTinyIDServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *TinyIDServiceClient {
	return &TinyIDServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewTinyIDServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *TinyIDServiceClient {
	return &TinyIDServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewTinyIDServiceClient(c thrift.TClient) *TinyIDServiceClient {
	return &TinyIDServiceClient{
		c: c,
	}
}

func (p *TinyIDServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *TinyIDServiceClient) GetMaxID(ctx context.Context, req *GetMaxIDRequest) (r *GetMaxIDResponse, err error) {
	var _args TinyIDServiceGetMaxIDArgs
	_args.Req = req
	var _result TinyIDServiceGetMaxIDResult
	if err = p.Client_().Call(ctx, "GetMaxID", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type TinyIDServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      TinyIDService
}

func (p *TinyIDServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *TinyIDServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *TinyIDServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewTinyIDServiceProcessor(handler TinyIDService) *TinyIDServiceProcessor {
	self := &TinyIDServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("GetMaxID", &tinyIDServiceProcessorGetMaxID{handler: handler})
	return self
}
func (p *TinyIDServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type tinyIDServiceProcessorGetMaxID struct {
	handler TinyIDService
}

func (p *tinyIDServiceProcessorGetMaxID) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := TinyIDServiceGetMaxIDArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("GetMaxID", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := TinyIDServiceGetMaxIDResult{}
	var retval *GetMaxIDResponse
	if retval, err2 = p.handler.GetMaxID(ctx, args.Req); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing GetMaxID: "+err2.Error())
		oprot.WriteMessageBegin("GetMaxID", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("GetMaxID", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type TinyIDServiceGetMaxIDArgs struct {
	Req *GetMaxIDRequest `thrift:"req,1" frugal:"1,default,GetMaxIDRequest" json:"req"`
}

func NewTinyIDServiceGetMaxIDArgs() *TinyIDServiceGetMaxIDArgs {
	return &TinyIDServiceGetMaxIDArgs{}
}

func (p *TinyIDServiceGetMaxIDArgs) InitDefault() {
	*p = TinyIDServiceGetMaxIDArgs{}
}

var TinyIDServiceGetMaxIDArgs_Req_DEFAULT *GetMaxIDRequest

func (p *TinyIDServiceGetMaxIDArgs) GetReq() (v *GetMaxIDRequest) {
	if !p.IsSetReq() {
		return TinyIDServiceGetMaxIDArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *TinyIDServiceGetMaxIDArgs) SetReq(val *GetMaxIDRequest) {
	p.Req = val
}

var fieldIDToName_TinyIDServiceGetMaxIDArgs = map[int16]string{
	1: "req",
}

func (p *TinyIDServiceGetMaxIDArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *TinyIDServiceGetMaxIDArgs) Read(iprot thrift.TProtocol) (err error) {

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
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_TinyIDServiceGetMaxIDArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *TinyIDServiceGetMaxIDArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Req = NewGetMaxIDRequest()
	if err := p.Req.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *TinyIDServiceGetMaxIDArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GetMaxID_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
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

func (p *TinyIDServiceGetMaxIDArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Req.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *TinyIDServiceGetMaxIDArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TinyIDServiceGetMaxIDArgs(%+v)", *p)

}

func (p *TinyIDServiceGetMaxIDArgs) DeepEqual(ano *TinyIDServiceGetMaxIDArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Req) {
		return false
	}
	return true
}

func (p *TinyIDServiceGetMaxIDArgs) Field1DeepEqual(src *GetMaxIDRequest) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

type TinyIDServiceGetMaxIDResult struct {
	Success *GetMaxIDResponse `thrift:"success,0,optional" frugal:"0,optional,GetMaxIDResponse" json:"success,omitempty"`
}

func NewTinyIDServiceGetMaxIDResult() *TinyIDServiceGetMaxIDResult {
	return &TinyIDServiceGetMaxIDResult{}
}

func (p *TinyIDServiceGetMaxIDResult) InitDefault() {
	*p = TinyIDServiceGetMaxIDResult{}
}

var TinyIDServiceGetMaxIDResult_Success_DEFAULT *GetMaxIDResponse

func (p *TinyIDServiceGetMaxIDResult) GetSuccess() (v *GetMaxIDResponse) {
	if !p.IsSetSuccess() {
		return TinyIDServiceGetMaxIDResult_Success_DEFAULT
	}
	return p.Success
}
func (p *TinyIDServiceGetMaxIDResult) SetSuccess(x interface{}) {
	p.Success = x.(*GetMaxIDResponse)
}

var fieldIDToName_TinyIDServiceGetMaxIDResult = map[int16]string{
	0: "success",
}

func (p *TinyIDServiceGetMaxIDResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *TinyIDServiceGetMaxIDResult) Read(iprot thrift.TProtocol) (err error) {

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
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_TinyIDServiceGetMaxIDResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *TinyIDServiceGetMaxIDResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = NewGetMaxIDResponse()
	if err := p.Success.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *TinyIDServiceGetMaxIDResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GetMaxID_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
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

func (p *TinyIDServiceGetMaxIDResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *TinyIDServiceGetMaxIDResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TinyIDServiceGetMaxIDResult(%+v)", *p)

}

func (p *TinyIDServiceGetMaxIDResult) DeepEqual(ano *TinyIDServiceGetMaxIDResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *TinyIDServiceGetMaxIDResult) Field0DeepEqual(src *GetMaxIDResponse) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}
