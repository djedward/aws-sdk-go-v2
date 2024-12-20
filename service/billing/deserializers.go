// Code generated by smithy-go-codegen DO NOT EDIT.

package billing

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/billing/types"
	smithy "github.com/aws/smithy-go"
	smithyio "github.com/aws/smithy-go/io"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithytime "github.com/aws/smithy-go/time"
	"github.com/aws/smithy-go/tracing"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
	"strings"
	"time"
)

func deserializeS3Expires(v string) (*time.Time, error) {
	t, err := smithytime.ParseHTTPDate(v)
	if err != nil {
		return nil, nil
	}
	return &t, nil
}

type awsAwsjson10_deserializeOpListBillingViews struct {
}

func (*awsAwsjson10_deserializeOpListBillingViews) ID() string {
	return "OperationDeserializer"
}

func (m *awsAwsjson10_deserializeOpListBillingViews) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
	defer endTimer()
	defer span.End()
	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsAwsjson10_deserializeOpErrorListBillingViews(response, &metadata)
	}
	output := &ListBillingViewsOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	err = awsAwsjson10_deserializeOpDocumentListBillingViewsOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	return out, metadata, err
}

func awsAwsjson10_deserializeOpErrorListBillingViews(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	headerCode := response.Header.Get("X-Amzn-ErrorType")

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	bodyInfo, err := getProtocolErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if typ, ok := resolveProtocolErrorType(headerCode, bodyInfo); ok {
		errorCode = restjson.SanitizeErrorCode(typ)
	}
	if len(bodyInfo.Message) != 0 {
		errorMessage = bodyInfo.Message
	}
	switch {
	case strings.EqualFold("AccessDeniedException", errorCode):
		return awsAwsjson10_deserializeErrorAccessDeniedException(response, errorBody)

	case strings.EqualFold("InternalServerException", errorCode):
		return awsAwsjson10_deserializeErrorInternalServerException(response, errorBody)

	case strings.EqualFold("ThrottlingException", errorCode):
		return awsAwsjson10_deserializeErrorThrottlingException(response, errorBody)

	case strings.EqualFold("ValidationException", errorCode):
		return awsAwsjson10_deserializeErrorValidationException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsAwsjson10_deserializeErrorAccessDeniedException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	output := &types.AccessDeniedException{}
	err := awsAwsjson10_deserializeDocumentAccessDeniedException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	return output
}

func awsAwsjson10_deserializeErrorInternalServerException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	output := &types.InternalServerException{}
	err := awsAwsjson10_deserializeDocumentInternalServerException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	return output
}

func awsAwsjson10_deserializeErrorThrottlingException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	output := &types.ThrottlingException{}
	err := awsAwsjson10_deserializeDocumentThrottlingException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	return output
}

func awsAwsjson10_deserializeErrorValidationException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	output := &types.ValidationException{}
	err := awsAwsjson10_deserializeDocumentValidationException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	return output
}

func awsAwsjson10_deserializeDocumentAccessDeniedException(v **types.AccessDeniedException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.AccessDeniedException
	if *v == nil {
		sv = &types.AccessDeniedException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "message", "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ErrorMessage to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentBillingViewList(v *[]types.BillingViewListElement, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var cv []types.BillingViewListElement
	if *v == nil {
		cv = []types.BillingViewListElement{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col types.BillingViewListElement
		destAddr := &col
		if err := awsAwsjson10_deserializeDocumentBillingViewListElement(&destAddr, value); err != nil {
			return err
		}
		col = *destAddr
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsAwsjson10_deserializeDocumentBillingViewListElement(v **types.BillingViewListElement, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.BillingViewListElement
	if *v == nil {
		sv = &types.BillingViewListElement{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "arn":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected BillingViewArn to be of type string, got %T instead", value)
				}
				sv.Arn = ptr.String(jtv)
			}

		case "billingViewType":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected BillingViewType to be of type string, got %T instead", value)
				}
				sv.BillingViewType = types.BillingViewType(jtv)
			}

		case "name":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected BillingViewName to be of type string, got %T instead", value)
				}
				sv.Name = ptr.String(jtv)
			}

		case "ownerAccountId":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected AccountId to be of type string, got %T instead", value)
				}
				sv.OwnerAccountId = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentInternalServerException(v **types.InternalServerException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.InternalServerException
	if *v == nil {
		sv = &types.InternalServerException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "message", "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ErrorMessage to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentThrottlingException(v **types.ThrottlingException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.ThrottlingException
	if *v == nil {
		sv = &types.ThrottlingException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "message", "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ErrorMessage to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentValidationException(v **types.ValidationException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.ValidationException
	if *v == nil {
		sv = &types.ValidationException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "fieldList":
			if err := awsAwsjson10_deserializeDocumentValidationExceptionFieldList(&sv.FieldList, value); err != nil {
				return err
			}

		case "message", "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ErrorMessage to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		case "reason":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ValidationExceptionReason to be of type string, got %T instead", value)
				}
				sv.Reason = types.ValidationExceptionReason(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentValidationExceptionField(v **types.ValidationExceptionField, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.ValidationExceptionField
	if *v == nil {
		sv = &types.ValidationExceptionField{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ErrorMessage to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		case "name":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected FieldName to be of type string, got %T instead", value)
				}
				sv.Name = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentValidationExceptionFieldList(v *[]types.ValidationExceptionField, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var cv []types.ValidationExceptionField
	if *v == nil {
		cv = []types.ValidationExceptionField{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col types.ValidationExceptionField
		destAddr := &col
		if err := awsAwsjson10_deserializeDocumentValidationExceptionField(&destAddr, value); err != nil {
			return err
		}
		col = *destAddr
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsAwsjson10_deserializeOpDocumentListBillingViewsOutput(v **ListBillingViewsOutput, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *ListBillingViewsOutput
	if *v == nil {
		sv = &ListBillingViewsOutput{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "billingViews":
			if err := awsAwsjson10_deserializeDocumentBillingViewList(&sv.BillingViews, value); err != nil {
				return err
			}

		case "nextToken":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected PageToken to be of type string, got %T instead", value)
				}
				sv.NextToken = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

type protocolErrorInfo struct {
	Type    string `json:"__type"`
	Message string
	Code    any // nonstandard for awsjson but some services do present the type here
}

func getProtocolErrorInfo(decoder *json.Decoder) (protocolErrorInfo, error) {
	var errInfo protocolErrorInfo
	if err := decoder.Decode(&errInfo); err != nil {
		if err == io.EOF {
			return errInfo, nil
		}
		return errInfo, err
	}

	return errInfo, nil
}

func resolveProtocolErrorType(headerType string, bodyInfo protocolErrorInfo) (string, bool) {
	if len(headerType) != 0 {
		return headerType, true
	} else if len(bodyInfo.Type) != 0 {
		return bodyInfo.Type, true
	} else if code, ok := bodyInfo.Code.(string); ok && len(code) != 0 {
		return code, true
	}
	return "", false
}
