#include "api.h"

extern void cl22GoKernelNativeCallback(void *);

static CL_CALLBACK void cl22CKernelNativeCallback(void *args)
{
    cl22GoKernelNativeCallback(args);
}

cl_int cl22EnqueueNativeKernel(cl_command_queue commandQueue,
    void *args, size_t argsSize,
    cl_uint numMemObjects, cl_mem *memList, void const *argsMemLoc,
    cl_uint waitListCount, cl_event const *waitList,
    cl_event *event)
{
    return clEnqueueNativeKernel(
        commandQueue,
        cl22CKernelNativeCallback, args, argsSize,
        numMemObjects, memList, (void const **)(argsMemLoc),
        waitListCount, waitList,
        event);
}
