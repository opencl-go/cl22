package cl22

// #include "api.h"
// extern cl_context cl22CreateContext(cl_context_properties *properties,
//     cl_uint numDevices, cl_device_id *devices,
//     void *notify, intptr_t userData,
//     cl_int *errcodeReturn);
// extern cl_context cl22CreateContextFromType(cl_context_properties *properties,
//     cl_device_type deviceType,
//     void *notify, intptr_t userData,
//     cl_int *errcodeReturn);
// extern cl_int cl22SetContextDestructorCallback(cl_context context, void *notify, void *userData);
import "C"
import (
	"fmt"
	"unsafe"
)

// A Context is used by the OpenCL runtime for managing objects such as command-queues, memory,
// program and kernel objects, and for executing kernels on one or more devices specified in the context.
type Context uintptr

func (c Context) handle() C.cl_context {
	return *((*C.cl_context)(unsafe.Pointer(&c)))
}

// String provides a readable presentation of the platform identifier.
// It is based on the numerical value of the underlying pointer.
func (c Context) String() string {
	return fmt.Sprintf("0x%X", uintptr(c))
}

const (
	// ContextPlatform specifies the platform to use.
	//
	// Use OnPlatform() for convenience.
	//
	// Property value type: PlatformID
	ContextPlatform uintptr = C.CL_CONTEXT_PLATFORM
	// ContextInteropUserSync specifies whether the user is responsible for synchronization between OpenCL and
	// other APIs. Please refer to the specific sections in the OpenCL Extension Specification that describe sharing
	// with other APIs for restrictions on using this flag.
	//
	// Use WithInteropUserSync() for convenience.
	//
	// Property value type: Bool
	// Since: 1.2
	ContextInteropUserSync uintptr = C.CL_CONTEXT_INTEROP_USER_SYNC
)

// ContextProperty is one entry of properties which are taken into account when creating context objects.
type ContextProperty []uintptr

// OnPlatform is a convenience function to handle the ContextPlatform property.
// Use it in combination with CreateContext() or CreateContextFromType().
func OnPlatform(id PlatformID) ContextProperty {
	return ContextProperty{ContextPlatform, uintptr(id)}
}

// WithInteropUserSync is a convenience function to handle the ContextInteropUserSync property.
// Use it in combination with CreateContext() or CreateContextFromType().
func WithInteropUserSync(value bool) ContextProperty {
	return ContextProperty{ContextInteropUserSync, uintptr(BoolFrom(value))}
}

// CreateContext creates an OpenCL context for the specified devices.
//
// The callback is an optional receiver for any errors that happen during creation or execution of the context.
// It is possible that one registered callback is re-used for multiple context objects. Synchronization and separation
// between the contexts is left to the user.
// It is possible to call ContextErrorCallback.Release() while any context the callback is registered with is still
// active.
//
// See also: https://registry.khronos.org/OpenCL/sdk/2.2/docs/man/html/clCreateContext.html
func CreateContext(deviceIds []DeviceID, callback *ContextErrorCallback, properties ...ContextProperty) (Context, error) {
	var rawPropertyList []uintptr
	for _, property := range properties {
		rawPropertyList = append(rawPropertyList, property...)
	}
	var rawProperties unsafe.Pointer
	if len(properties) > 0 {
		rawPropertyList = append(rawPropertyList, 0)
		rawProperties = unsafe.Pointer(&rawPropertyList[0])
	}
	var rawDeviceIds unsafe.Pointer
	if len(deviceIds) > 0 {
		rawDeviceIds = unsafe.Pointer(&deviceIds[0])
	}
	var callbackFunc unsafe.Pointer
	var callbackKey uintptr
	if callback != nil {
		callbackFunc = cContextErrorCallbackFunc()
		callbackKey = callback.key
	}
	var status C.cl_int
	context := C.cl22CreateContext(
		(*C.cl_context_properties)(rawProperties),
		C.uint(len(deviceIds)),
		(*C.cl_device_id)(rawDeviceIds),
		callbackFunc, (C.intptr_t)(callbackKey),
		&status)
	if status != C.CL_SUCCESS {
		return 0, StatusError(status)
	}
	return Context(*((*uintptr)(unsafe.Pointer(&context)))), nil
}

// CreateContextFromType creates an OpenCL context for devices that match the given device type.
// The context does not reference any sub-devices that may have been created from these devices.
//
// The callback is an optional receiver for any errors that happen during creation or execution of the context.
// It is possible that one registered callback is re-used for multiple context objects. Synchronization and separation
// between the contexts is left to the user.
// It is possible to call ContextErrorCallback.Release() while any context the callback is registered with is still
// active.
//
// See also: https://registry.khronos.org/OpenCL/sdk/2.2/docs/man/html/clCreateContextFromType.html
func CreateContextFromType(deviceType DeviceTypeFlags, callback *ContextErrorCallback, properties ...ContextProperty) (Context, error) {
	var rawPropertyList []uintptr
	for _, property := range properties {
		rawPropertyList = append(rawPropertyList, property...)
	}
	var rawProperties unsafe.Pointer
	if len(properties) > 0 {
		rawPropertyList = append(rawPropertyList, 0)
		rawProperties = unsafe.Pointer(&rawPropertyList[0])
	}
	var callbackFunc unsafe.Pointer
	var callbackKey uintptr
	if callback != nil {
		callbackFunc = cContextErrorCallbackFunc()
		callbackKey = callback.key
	}
	var status C.cl_int
	context := C.cl22CreateContextFromType(
		(*C.cl_context_properties)(rawProperties),
		C.cl_device_type(deviceType),
		callbackFunc, (C.intptr_t)(callbackKey),
		&status)
	if status != C.CL_SUCCESS {
		return 0, StatusError(status)
	}
	return Context(*((*uintptr)(unsafe.Pointer(&context)))), nil
}

// RetainContext increments the context reference count.
//
// CreateContext() and CreateContextFromType() perform an implicit retain. This is very helpful for 3rd party
// libraries, which typically get a context passed to them by the application. However, it is possible that the
// application may delete the context without informing the library. Allowing functions to attach to (i.e. retain)
// and release a context solves the problem of a context being used by a library no longer being valid.
//
// See also: https://registry.khronos.org/OpenCL/sdk/2.2/docs/man/html/clRetainContext.html
func RetainContext(context Context) error {
	status := C.clRetainContext(context.handle())
	if status != C.CL_SUCCESS {
		return StatusError(status)
	}
	return nil
}

// ReleaseContext decrements the context reference count.
//
// After the reference count becomes zero and all the objects attached to context (such as memory objects,
// command-queues) are released, the context is deleted.
//
// See also: https://registry.khronos.org/OpenCL/sdk/2.2/docs/man/html/clReleaseContext.html
func ReleaseContext(context Context) error {
	status := C.clReleaseContext(context.handle())
	if status != C.CL_SUCCESS {
		return StatusError(status)
	}
	return nil
}

// ContextInfoName identifies properties of a context, which can be queried with ContextInfo().
type ContextInfoName C.cl_context_info

const (
	// ContextReferenceCount returns the context reference count.
	//
	// Note: The reference count returned should be considered immediately stale. It is unsuitable for
	// general use in applications. This feature is provided for identifying memory leaks.
	//
	// Returned type: Uint
	ContextReferenceCount ContextInfoName = C.CL_CONTEXT_REFERENCE_COUNT
	// ContextDevices returns the list of devices and sub-devices in context.
	//
	// Returned type: []DeviceID
	ContextDevices ContextInfoName = C.CL_CONTEXT_DEVICES
	// ContextNumDevices returns the number of devices in context.
	//
	// Returned type: Uint
	// Since: 1.1
	ContextNumDevices ContextInfoName = C.CL_CONTEXT_NUM_DEVICES
	// ContextProperties returns the properties argument specified in CreateContext() or CreateContextFromType().
	//
	// The returned list is the concatenated list of all the properties provided at creation.
	// This list also includes the terminating zero value.
	//
	// Returned type: []uintptr
	ContextProperties ContextInfoName = C.CL_CONTEXT_PROPERTIES
)

// ContextInfo queries information about a context.
//
// The provided size need to specify the size of the available space pointed to the provided value in bytes.
//
// The returned number is the required size, in bytes, for the queried information.
// Call the function with a zero size and nil value to request the required size. This helps in determining
// the necessary space for dynamic information, such as arrays.
//
// Raw strings are with a terminating NUL character. For convenience, use ContextInfoString().
//
// See also: https://registry.khronos.org/OpenCL/sdk/2.2/docs/man/html/clGetContextInfo.html
func ContextInfo(context Context, paramName ContextInfoName, paramSize uint, paramValue unsafe.Pointer) (uint, error) {
	sizeReturn := C.size_t(0)
	status := C.clGetContextInfo(
		context.handle(),
		C.cl_context_info(paramName),
		C.size_t(paramSize),
		paramValue,
		&sizeReturn)
	if status != C.CL_SUCCESS {
		return 0, StatusError(status)
	}
	return uint(sizeReturn), nil
}

// ContextInfoString is a convenience method for ContextInfo() to query information values that are string-based.
//
// This function does not verify the queried information is indeed of type string. It assumes the information is
// a NUL terminated raw string and will extract the bytes as characters before that.
func ContextInfoString(context Context, paramName ContextInfoName) (string, error) {
	return queryString(func(paramSize uint, paramValue unsafe.Pointer) (uint, error) {
		return ContextInfo(context, paramName, paramSize, paramValue)
	})
}