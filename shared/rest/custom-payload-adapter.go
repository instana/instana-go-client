package rest

import "github.com/instana/instana-go-client/shared/types"

// CustomPayloadFieldsAware interface for objects that have custom payload fields
type CustomPayloadFieldsAware interface {
	GetCustomerPayloadFields() []types.CustomPayloadField[any]
	SetCustomerPayloadFields([]types.CustomPayloadField[any])
}

// customPayloadFieldsAwareInstanaDataObject combines CustomPayloadFieldsAware and InstanaDataObject
type customPayloadFieldsAwareInstanaDataObject interface {
	CustomPayloadFieldsAware
	InstanaDataObject
}

// NewCustomPayloadFieldsUnmarshallerAdapter creates a new Unmarshaller instance which can be added as an adapter to the default unmarshallers to map custom payload fields
func NewCustomPayloadFieldsUnmarshallerAdapter[T customPayloadFieldsAwareInstanaDataObject](unmarshaller JSONUnmarshaller[T]) JSONUnmarshaller[T] {
	return &customPayloadFieldsUnmarshallerAdapter[T]{unmarshaller: unmarshaller}
}

type customPayloadFieldsUnmarshallerAdapter[T customPayloadFieldsAwareInstanaDataObject] struct {
	unmarshaller JSONUnmarshaller[T]
}

// UnmarshalArray Unmarshaller interface implementation
func (a *customPayloadFieldsUnmarshallerAdapter[T]) UnmarshalArray(data []byte) (*[]T, error) {
	temp, err := a.unmarshaller.UnmarshalArray(data)
	if err != nil {
		return temp, err
	}
	if temp != nil {
		for _, v := range *temp {
			a.mapCustomPayloadFields(v)
		}
	}
	return temp, nil
}

// Unmarshal Unmarshaller interface implementation
func (a *customPayloadFieldsUnmarshallerAdapter[T]) Unmarshal(data []byte) (T, error) {
	temp, err := a.unmarshaller.Unmarshal(data)
	if err != nil {
		return temp, err
	}
	a.mapCustomPayloadFields(temp)
	return temp, nil
}

func (a *customPayloadFieldsUnmarshallerAdapter[T]) mapCustomPayloadFields(temp T) {
	customFields := temp.GetCustomerPayloadFields()
	for i, v := range customFields {
		customFields[i] = a.mapCustomPayloadField(v)
	}
	temp.SetCustomerPayloadFields(customFields)
}

func (a *customPayloadFieldsUnmarshallerAdapter[T]) mapCustomPayloadField(field types.CustomPayloadField[any]) types.CustomPayloadField[any] {
	if field.Type == types.DynamicCustomPayloadType {
		data := field.Value.(map[string]interface{})
		var keyPtr *string
		if val, ok := data["key"]; ok && val != nil {
			key := val.(string)
			keyPtr = &key
		}
		field.Value = types.DynamicCustomPayloadFieldValue{
			TagName: data["tagName"].(string),
			Key:     keyPtr,
		}
	} else {
		value := field.Value.(string)
		field.Value = types.StaticStringCustomPayloadFieldValue(value)
	}
	return field
}
