// Code generated by thriftgo (0.3.14). DO NOT EDIT.

package typo

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

type GitTypoReq struct {
	GitLink string `thrift:"GitLink,1" json:"gitLink" query:"gitLink" form:"GitLink" `
	Config  string `thrift:"Config,2" json:"Config" query:"config"`
}

func NewGitTypoReq() *GitTypoReq {
	return &GitTypoReq{}
}

func (p *GitTypoReq) InitDefault() {
}

func (p *GitTypoReq) GetGitLink() (v string) {
	return p.GitLink
}

func (p *GitTypoReq) GetConfig() (v string) {
	return p.Config
}

var fieldIDToName_GitTypoReq = map[int16]string{
	1: "GitLink",
	2: "Config",
}

func (p *GitTypoReq) Read(iprot thrift.TProtocol) (err error) {

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
			if fieldTypeId == thrift.STRING {
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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GitTypoReq[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *GitTypoReq) ReadField1(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.GitLink = _field
	return nil
}
func (p *GitTypoReq) ReadField2(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Config = _field
	return nil
}

func (p *GitTypoReq) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GitTypoReq"); err != nil {
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

func (p *GitTypoReq) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("GitLink", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.GitLink); err != nil {
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

func (p *GitTypoReq) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Config", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Config); err != nil {
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

func (p *GitTypoReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GitTypoReq(%+v)", *p)

}

//	struct GitTypoReqResp {
//	    1: string RespBody;
//	}
type GitTypoService interface {
	GitTypoMethod(ctx context.Context, request *GitTypoReq) (r string, err error)
}

type GitTypoServiceClient struct {
	c thrift.TClient
}

func NewGitTypoServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *GitTypoServiceClient {
	return &GitTypoServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewGitTypoServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *GitTypoServiceClient {
	return &GitTypoServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewGitTypoServiceClient(c thrift.TClient) *GitTypoServiceClient {
	return &GitTypoServiceClient{
		c: c,
	}
}

func (p *GitTypoServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *GitTypoServiceClient) GitTypoMethod(ctx context.Context, request *GitTypoReq) (r string, err error) {
	var _args GitTypoServiceGitTypoMethodArgs
	_args.Request = request
	var _result GitTypoServiceGitTypoMethodResult
	if err = p.Client_().Call(ctx, "GitTypoMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type GitTypoServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      GitTypoService
}

func (p *GitTypoServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *GitTypoServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *GitTypoServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewGitTypoServiceProcessor(handler GitTypoService) *GitTypoServiceProcessor {
	self := &GitTypoServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("GitTypoMethod", &gitTypoServiceProcessorGitTypoMethod{handler: handler})
	return self
}
func (p *GitTypoServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
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

type gitTypoServiceProcessorGitTypoMethod struct {
	handler GitTypoService
}

func (p *gitTypoServiceProcessorGitTypoMethod) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := GitTypoServiceGitTypoMethodArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("GitTypoMethod", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := GitTypoServiceGitTypoMethodResult{}
	var retval string
	if retval, err2 = p.handler.GitTypoMethod(ctx, args.Request); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing GitTypoMethod: "+err2.Error())
		oprot.WriteMessageBegin("GitTypoMethod", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("GitTypoMethod", thrift.REPLY, seqId); err2 != nil {
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

type GitTypoServiceGitTypoMethodArgs struct {
	Request *GitTypoReq `thrift:"request,1"`
}

func NewGitTypoServiceGitTypoMethodArgs() *GitTypoServiceGitTypoMethodArgs {
	return &GitTypoServiceGitTypoMethodArgs{}
}

func (p *GitTypoServiceGitTypoMethodArgs) InitDefault() {
}

var GitTypoServiceGitTypoMethodArgs_Request_DEFAULT *GitTypoReq

func (p *GitTypoServiceGitTypoMethodArgs) GetRequest() (v *GitTypoReq) {
	if !p.IsSetRequest() {
		return GitTypoServiceGitTypoMethodArgs_Request_DEFAULT
	}
	return p.Request
}

var fieldIDToName_GitTypoServiceGitTypoMethodArgs = map[int16]string{
	1: "request",
}

func (p *GitTypoServiceGitTypoMethodArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *GitTypoServiceGitTypoMethodArgs) Read(iprot thrift.TProtocol) (err error) {

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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GitTypoServiceGitTypoMethodArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *GitTypoServiceGitTypoMethodArgs) ReadField1(iprot thrift.TProtocol) error {
	_field := NewGitTypoReq()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Request = _field
	return nil
}

func (p *GitTypoServiceGitTypoMethodArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GitTypoMethod_args"); err != nil {
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

func (p *GitTypoServiceGitTypoMethodArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("request", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Request.Write(oprot); err != nil {
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

func (p *GitTypoServiceGitTypoMethodArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GitTypoServiceGitTypoMethodArgs(%+v)", *p)

}

type GitTypoServiceGitTypoMethodResult struct {
	Success *string `thrift:"success,0,optional"`
}

func NewGitTypoServiceGitTypoMethodResult() *GitTypoServiceGitTypoMethodResult {
	return &GitTypoServiceGitTypoMethodResult{}
}

func (p *GitTypoServiceGitTypoMethodResult) InitDefault() {
}

var GitTypoServiceGitTypoMethodResult_Success_DEFAULT string

func (p *GitTypoServiceGitTypoMethodResult) GetSuccess() (v string) {
	if !p.IsSetSuccess() {
		return GitTypoServiceGitTypoMethodResult_Success_DEFAULT
	}
	return *p.Success
}

var fieldIDToName_GitTypoServiceGitTypoMethodResult = map[int16]string{
	0: "success",
}

func (p *GitTypoServiceGitTypoMethodResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GitTypoServiceGitTypoMethodResult) Read(iprot thrift.TProtocol) (err error) {

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
			if fieldTypeId == thrift.STRING {
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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GitTypoServiceGitTypoMethodResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *GitTypoServiceGitTypoMethodResult) ReadField0(iprot thrift.TProtocol) error {

	var _field *string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = &v
	}
	p.Success = _field
	return nil
}

func (p *GitTypoServiceGitTypoMethodResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GitTypoMethod_result"); err != nil {
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

func (p *GitTypoServiceGitTypoMethodResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.Success); err != nil {
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

func (p *GitTypoServiceGitTypoMethodResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GitTypoServiceGitTypoMethodResult(%+v)", *p)

}
