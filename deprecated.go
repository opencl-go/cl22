package cl22

// #cgo CFLAGS: -DCL_USE_DEPRECATED_OPENCL_1_2_APIS
// #cgo CXXFLAGS: -DCL_USE_DEPRECATED_OPENCL_1_2_APIS
// #cgo CPPFLAGS: -DCL_USE_DEPRECATED_OPENCL_1_2_APIS
// #include "api.h"
import "C"
import "unsafe"

// CreateCommandQueue creates a command-queue on a specific device.
//
// Deprecated: 1.2; Use CreateCommandQueueWithProperties() instead.
// See also: https://registry.khronos.org/OpenCL/sdk/2.2/docs/man/html/clCreateCommandQueue.html
func CreateCommandQueue(context Context, deviceID DeviceID, properties CommandQueuePropertiesFlags) (CommandQueue, error) {
	var status C.cl_int
	commandQueue := C.clCreateCommandQueue(
		context.handle(),
		deviceID.handle(),
		C.cl_command_queue_properties(properties),
		&status)
	if status != C.CL_SUCCESS {
		return 0, StatusError(status)
	}
	return CommandQueue(*((*uintptr)(unsafe.Pointer(&commandQueue)))), nil
}

// CreateSampler creates a sampler object.
//
// Deprecated: 1.2; Use CreateSamplerWithProperties() instead.
// See also: https://registry.khronos.org/OpenCL/sdk/2.2/docs/man/html/clCreateSampler.html
func CreateSampler(context Context, normalizedCoords bool, addressingMode SamplerAddressingMode, filterMode SamplerFilterMode) (Sampler, error) {
	var status C.cl_int
	sampler := C.clCreateSampler(
		context.handle(),
		C.cl_bool(BoolFrom(normalizedCoords)),
		C.cl_addressing_mode(addressingMode),
		C.cl_filter_mode(filterMode),
		&status)
	if status != C.CL_SUCCESS {
		return 0, StatusError(status)
	}
	return Sampler(*((*uintptr)(unsafe.Pointer(&sampler)))), nil
}

// EnqueueTask enqueues a command to execute a kernel, using a single work-item, on a device.
//
// EnqueueTask() is equivalent to calling EnqueueNDRangeKernel() with one WorkDimension that has
// GlobalOffset = 0, GlobalSize = 1, and LocalSize = 1.
//
// Deprecated: 1.2; Use EnqueueNDRangeKernel() instead.
// See also: https://registry.khronos.org/OpenCL/sdk/2.2/docs/man/html/clEnqueueTask.html
func EnqueueTask(commandQueue CommandQueue, kernel Kernel, waitList []Event, event *Event) error {
	var rawWaitList unsafe.Pointer
	if len(waitList) > 0 {
		rawWaitList = unsafe.Pointer(&waitList[0])
	}
	status := C.clEnqueueTask(
		commandQueue.handle(),
		kernel.handle(),
		C.cl_uint(len(waitList)),
		(*C.cl_event)(rawWaitList),
		(*C.cl_event)(unsafe.Pointer(event)))
	if status != C.CL_SUCCESS {
		return StatusError(status)
	}
	return nil
}
