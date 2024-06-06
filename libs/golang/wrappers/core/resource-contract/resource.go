package wrappersresourcecontract

// Resource defines the interface that all resources should implement
type Resource interface {
    // Init initializes the resource and returns an error if any occurs.
    Init() error
    
    // GetClient returns the underlying client of the resource.
    GetClient() interface{}
}
