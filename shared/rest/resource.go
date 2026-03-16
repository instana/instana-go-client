package rest

// InstanaDataObject is a marker interface for any data object provided by any resource of the Instana REST API
type InstanaDataObject interface {
	GetIDForResourcePath() string
}

// RestResource interface definition of a instana REST resource.
type RestResource[T InstanaDataObject] interface {
	GetAll() (*[]T, error)
	GetOne(id string) (T, error)
	Create(data T) (T, error)
	Update(data T) (T, error)
	Delete(data T) error
	DeleteByID(id string) error
}

// ReadOnlyRestResource interface definition for a read only REST resource. The resource at instana might
// implement more methods but the implementation of the provider is limited to read only.
type ReadOnlyRestResource[T InstanaDataObject] interface {
	GetAll() (*[]T, error)
	GetByQuery(queryParams map[string]string) (*[]T, error)
	GetOne(id string) (T, error)
}

// JSONUnmarshaller interface definition for unmarshalling that unmarshalls JSON to go data structures
type JSONUnmarshaller[T any] interface {
	// Unmarshal converts the provided json bytes into the go data structure as provided in the target
	Unmarshal(data []byte) (T, error)
	// UnmarshalArray converts the provided json bytes into the go data structure as provided in the target
	UnmarshalArray(data []byte) (*[]T, error)
}

// DataFilterFunc function definition for filtering data received from Instana API
type DataFilterFunc func(o InstanaDataObject) bool

// RestClient interface to access REST resources of the Instana API
type RestClient interface {
	Get(resourcePath string) ([]byte, error)
	GetOne(id string, resourcePath string) ([]byte, error)
	Post(data InstanaDataObject, resourcePath string) ([]byte, error)
	PostWithID(data InstanaDataObject, resourcePath string) ([]byte, error)
	Put(data InstanaDataObject, resourcePath string) ([]byte, error)
	Delete(resourceID string, resourceBasePath string) error
	GetByQuery(resourcePath string, queryParams map[string]string) ([]byte, error)
	PostByQuery(resourcePath string, queryParams map[string]string) ([]byte, error)
	PutByQuery(resourcePath string, is string, queryParams map[string]string) ([]byte, error)
}

// Made with Bob
