package cl22

// #include "api.h"
import "C"
import (
	"unsafe"
)

// CreateBuffer creates a buffer object.
//
// See also: https://registry.khronos.org/OpenCL/sdk/2.2/docs/man/html/clCreateBuffer.html
func CreateBuffer(context Context, flags MemFlags, size int, hostPtr unsafe.Pointer) (MemObject, error) {
	var status C.cl_int
	mem := C.clCreateBuffer(
		context.handle(),
		C.cl_mem_flags(flags),
		C.size_t(size),
		hostPtr,
		&status)
	if status != C.CL_SUCCESS {
		return 0, StatusError(status)
	}
	return MemObject(*((*uintptr)(unsafe.Pointer(&mem)))), nil
}

// BufferCreateType determines the kind of sub-buffer object.
type BufferCreateType C.cl_buffer_create_type

const (
	// BufferCreateTypeRegion describes a buffer object that represents a specific region in buffer.
	//
	// Creation data type: BufferRegion
	// Since: 1.1
	BufferCreateTypeRegion BufferCreateType = C.CL_BUFFER_CREATE_TYPE_REGION
)

// BufferRegionByteSize is the size, in bytes, of the BufferRegion structure.
const BufferRegionByteSize = unsafe.Sizeof(C.cl_buffer_region{})

// BufferRegion describes a subset of a buffer.
//
// Since: 1.1
type BufferRegion struct {
	Origin uintptr
	Size   uintptr
}

// CreateSubBuffer creates a new buffer object (referred to as a sub-buffer object) from an existing buffer object.
//
// The createInfo parameter is dependent on the specified createType.
//
// Since: 1.1
// See also: https://registry.khronos.org/OpenCL/sdk/2.2/docs/man/html/clCreateSubBuffer.html
func CreateSubBuffer(buffer MemObject, flags MemFlags, createType BufferCreateType, createInfo unsafe.Pointer) (MemObject, error) {
	var status C.cl_int
	mem := C.clCreateSubBuffer(
		buffer.handle(),
		C.cl_mem_flags(flags),
		C.cl_buffer_create_type(createType),
		createInfo,
		&status)
	if status != C.CL_SUCCESS {
		return 0, StatusError(status)
	}
	return MemObject(*((*uintptr)(unsafe.Pointer(&mem)))), nil
}

// EnqueueMapBuffer enqueues a command to map a region of a buffer object into the host address space and
// returns a pointer to this mapped region.
//
// See also: https://registry.khronos.org/OpenCL/sdk/2.2/docs/man/html/clEnqueueMapBuffer.html
func EnqueueMapBuffer(commandQueue CommandQueue,
	buffer MemObject, blocking bool, flags MapFlags, offset, size uintptr,
	waitList []Event, event *Event) (unsafe.Pointer, error) {
	var rawWaitList unsafe.Pointer
	if len(waitList) > 0 {
		rawWaitList = unsafe.Pointer(&waitList[0])
	}
	var status C.cl_int
	ptr := C.clEnqueueMapBuffer(
		commandQueue.handle(),
		buffer.handle(),
		C.cl_bool(BoolFrom(blocking)),
		C.cl_map_flags(flags),
		C.size_t(offset),
		C.size_t(size),
		C.cl_uint(len(waitList)),
		(*C.cl_event)(rawWaitList),
		(*C.cl_event)(unsafe.Pointer(event)),
		&status)
	if status != C.CL_SUCCESS {
		return nil, StatusError(status)
	}
	return ptr, nil
}

// EnqueueReadBuffer enqueues a command to read from a buffer object to host memory.
//
// See also: https://registry.khronos.org/OpenCL/sdk/2.2/docs/man/html/clEnqueueReadBuffer.html
func EnqueueReadBuffer(commandQueue CommandQueue, mem MemObject, blockingRead bool, offset, size uintptr, data unsafe.Pointer,
	waitList []Event, event *Event) error {
	var rawWaitList unsafe.Pointer
	if len(waitList) > 0 {
		rawWaitList = unsafe.Pointer(&waitList[0])
	}
	status := C.clEnqueueReadBuffer(
		commandQueue.handle(),
		mem.handle(),
		C.cl_bool(BoolFrom(blockingRead)),
		C.size_t(offset),
		C.size_t(size),
		data,
		C.cl_uint(len(waitList)),
		(*C.cl_event)(rawWaitList),
		(*C.cl_event)(unsafe.Pointer(event)))
	if status != C.CL_SUCCESS {
		return StatusError(status)
	}
	return nil
}

// EnqueueReadBufferRect enqueues a command to read from a 2D or 3D rectangular region of a buffer object to
// host memory.
//
// Since: 1.1
// See also: https://registry.khronos.org/OpenCL/sdk/2.2/docs/man/html/clEnqueueReadBufferRect.html
func EnqueueReadBufferRect(commandQueue CommandQueue, mem MemObject, blockingRead bool, bufferOrigin, hostOrigin, region [3]uintptr,
	bufferRowPitch, bufferSlicePitch, hostRowPitch, hostSlicePitch uintptr, data unsafe.Pointer, waitList []Event, event *Event) error {
	var rawWaitList unsafe.Pointer
	if len(waitList) > 0 {
		rawWaitList = unsafe.Pointer(&waitList[0])
	}
	status := C.clEnqueueReadBufferRect(
		commandQueue.handle(),
		mem.handle(),
		C.cl_bool(BoolFrom(blockingRead)),
		(*C.size_t)(unsafe.Pointer(&bufferOrigin[0])),
		(*C.size_t)(unsafe.Pointer(&hostOrigin[0])),
		(*C.size_t)(unsafe.Pointer(&region[0])),
		C.size_t(bufferRowPitch),
		C.size_t(bufferSlicePitch),
		C.size_t(hostRowPitch),
		C.size_t(hostSlicePitch),
		data,
		C.cl_uint(len(waitList)),
		(*C.cl_event)(rawWaitList),
		(*C.cl_event)(unsafe.Pointer(event)))
	if status != C.CL_SUCCESS {
		return StatusError(status)
	}
	return nil
}

// EnqueueWriteBuffer enqueues a command to write to a buffer object from host memory.
//
// See also: https://registry.khronos.org/OpenCL/sdk/2.2/docs/man/html/clEnqueueWriteBuffer.html
func EnqueueWriteBuffer(commandQueue CommandQueue, mem MemObject, blockingRead bool, offset, size uintptr, data unsafe.Pointer,
	waitList []Event, event *Event) error {
	var rawWaitList unsafe.Pointer
	if len(waitList) > 0 {
		rawWaitList = unsafe.Pointer(&waitList[0])
	}
	status := C.clEnqueueWriteBuffer(
		commandQueue.handle(),
		mem.handle(),
		C.cl_bool(BoolFrom(blockingRead)),
		C.size_t(offset),
		C.size_t(size),
		data,
		C.cl_uint(len(waitList)),
		(*C.cl_event)(rawWaitList),
		(*C.cl_event)(unsafe.Pointer(event)))
	if status != C.CL_SUCCESS {
		return StatusError(status)
	}
	return nil
}

// EnqueueWriteBufferRect enqueues a command to write to a 2D or 3D rectangular region of a buffer object from
// host memory.
//
// Since: 1.1
// See also: https://registry.khronos.org/OpenCL/sdk/2.2/docs/man/html/clEnqueueWriteBufferRect.html
func EnqueueWriteBufferRect(commandQueue CommandQueue, mem MemObject, blockingRead bool, bufferOrigin, hostOrigin, region [3]uintptr,
	bufferRowPitch, bufferSlicePitch, hostRowPitch, hostSlicePitch uintptr, data unsafe.Pointer, waitList []Event, event *Event) error {
	var rawWaitList unsafe.Pointer
	if len(waitList) > 0 {
		rawWaitList = unsafe.Pointer(&waitList[0])
	}
	status := C.clEnqueueWriteBufferRect(
		commandQueue.handle(),
		mem.handle(),
		C.cl_bool(BoolFrom(blockingRead)),
		(*C.size_t)(unsafe.Pointer(&bufferOrigin[0])),
		(*C.size_t)(unsafe.Pointer(&hostOrigin[0])),
		(*C.size_t)(unsafe.Pointer(&region[0])),
		C.size_t(bufferRowPitch),
		C.size_t(bufferSlicePitch),
		C.size_t(hostRowPitch),
		C.size_t(hostSlicePitch),
		data,
		C.cl_uint(len(waitList)),
		(*C.cl_event)(rawWaitList),
		(*C.cl_event)(unsafe.Pointer(event)))
	if status != C.CL_SUCCESS {
		return StatusError(status)
	}
	return nil
}

// EnqueueFillBuffer enqueues a command to fill a buffer object with a pattern of a given pattern size.
//
// Since: 1.2
// See also: https://registry.khronos.org/OpenCL/sdk/2.2/docs/man/html/clEnqueueFillBuffer.html
func EnqueueFillBuffer(commandQueue CommandQueue, mem MemObject, pattern unsafe.Pointer, patternSize, offset, size uintptr,
	waitList []Event, event *Event) error {
	var rawWaitList unsafe.Pointer
	if len(waitList) > 0 {
		rawWaitList = unsafe.Pointer(&waitList[0])
	}
	status := C.clEnqueueFillBuffer(
		commandQueue.handle(),
		mem.handle(),
		pattern,
		C.size_t(patternSize),
		C.size_t(offset),
		C.size_t(size),
		C.cl_uint(len(waitList)),
		(*C.cl_event)(rawWaitList),
		(*C.cl_event)(unsafe.Pointer(event)))
	if status != C.CL_SUCCESS {
		return StatusError(status)
	}
	return nil
}

// EnqueueCopyBuffer enqueues a command to copy from one buffer object to another.
//
// See also: https://registry.khronos.org/OpenCL/sdk/2.2/docs/man/html/clEnqueueCopyBuffer.html
func EnqueueCopyBuffer(commandQueue CommandQueue, src, dst MemObject, srcOffset, dstOffset, size uintptr, waitList []Event, event *Event) error {
	var rawWaitList unsafe.Pointer
	if len(waitList) > 0 {
		rawWaitList = unsafe.Pointer(&waitList[0])
	}
	status := C.clEnqueueCopyBuffer(
		commandQueue.handle(),
		src.handle(),
		dst.handle(),
		C.size_t(srcOffset),
		C.size_t(dstOffset),
		C.size_t(size),
		C.cl_uint(len(waitList)),
		(*C.cl_event)(rawWaitList),
		(*C.cl_event)(unsafe.Pointer(event)))
	if status != C.CL_SUCCESS {
		return StatusError(status)
	}
	return nil
}

// EnqueueCopyBufferRect enqueues a command to copy a 2D or 3D rectangular region from a buffer object to another
// buffer object.
//
// See also: https://registry.khronos.org/OpenCL/sdk/2.2/docs/man/html/clEnqueueCopyBufferRect.html
func EnqueueCopyBufferRect(commandQueue CommandQueue, src, dst MemObject, srcOrigin, dstOrigin, region [3]uintptr,
	srcRowPitch, srcSlicePitch, dstRowPitch, dstSlicePitch uintptr,
	waitList []Event, event *Event) error {
	var rawWaitList unsafe.Pointer
	if len(waitList) > 0 {
		rawWaitList = unsafe.Pointer(&waitList[0])
	}
	status := C.clEnqueueCopyBufferRect(
		commandQueue.handle(),
		src.handle(),
		dst.handle(),
		(*C.size_t)(unsafe.Pointer(&srcOrigin[0])),
		(*C.size_t)(unsafe.Pointer(&dstOrigin[0])),
		(*C.size_t)(unsafe.Pointer(&region[0])),
		C.size_t(srcRowPitch),
		C.size_t(srcSlicePitch),
		C.size_t(dstRowPitch),
		C.size_t(dstSlicePitch),
		C.cl_uint(len(waitList)),
		(*C.cl_event)(rawWaitList),
		(*C.cl_event)(unsafe.Pointer(event)))
	if status != C.CL_SUCCESS {
		return StatusError(status)
	}
	return nil
}
