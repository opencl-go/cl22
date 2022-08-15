#include "api.h"

extern void cl22GoSvmFreeCallback(cl_command_queue commandQueue, cl_uint svmPointerCount, void *svmPointers, uintptr_t *userData);

static CL_CALLBACK void cl22CSvmFreeCallback(
    cl_command_queue commandQueue,
    cl_uint svmPointerCount, void *svmPointers[],
    void *userData)
{
    cl22GoSvmFreeCallback(commandQueue, svmPointerCount, svmPointers, (uintptr_t *)(userData));
}

cl_int cl22EnqueueSVMFree(cl_command_queue commandQueue,
    cl_uint svmPointerCount, void *svmPointers,
    uintptr_t *userData,
    cl_uint waitListCount, cl_event const *waitList,
    cl_event *event)
{
    return clEnqueueSVMFree(
        commandQueue,
        svmPointerCount, (void **)(svmPointers),
        (userData != NULL) ? cl22CSvmFreeCallback : NULL, userData,
        waitListCount, waitList,
        event);
}

cl_int cl22EnqueueSVMMigrateMem(
    cl_command_queue commandQueue,
    cl_uint svmPointerCount, void *svmPointers,
    size_t *sizes, cl_mem_migration_flags flags,
    cl_uint waitListCount, cl_event const *waitList,
    cl_event *event)
{
    return clEnqueueSVMMigrateMem(
        commandQueue,
        svmPointerCount, (void const **)(svmPointers),
        sizes, flags,
        waitListCount, waitList,
        event);
}
