#include "api.h"

extern void cl22GoContextErrorCallback(char *, uint8_t *, size_t, uintptr_t *);

static CL_CALLBACK void cl22CContextErrorCallback(char const *errorInfo,
    void const *privateInfoPtr, size_t privateInfoLen,
    void *userData)
{
    cl22GoContextErrorCallback((char *)(errorInfo), (uint8_t *)(privateInfoPtr), privateInfoLen, (uintptr_t *)(userData));
}

cl_context cl22CreateContext(cl_context_properties *properties,
    cl_uint numDevices, cl_device_id *devices,
    uintptr_t *userData,
    cl_int *errcodeReturn)
{
    return clCreateContext(properties, numDevices, devices,
        (userData != NULL) ? cl22CContextErrorCallback : NULL, userData,
        errcodeReturn);
}

cl_context cl22CreateContextFromType(cl_context_properties *properties,
    cl_device_type deviceType,
    uintptr_t *userData,
    cl_int *errcodeReturn)
{
    return clCreateContextFromType(properties, deviceType,
        (userData != NULL) ? cl22CContextErrorCallback : NULL, userData,
        errcodeReturn);
}
