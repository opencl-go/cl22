#include "api.h"

extern void cl22GoEventCallback(cl_event, cl_int, void*);

static CL_CALLBACK void cl22CEventCallback(cl_event event, cl_int commandStatus, void *userData)
{
    cl22GoEventCallback(event, commandStatus, (uintptr_t *)(userData));
}

cl_int cl22SetEventCallback(cl_event event, cl_int callbackType, uintptr_t *userData)
{
    return clSetEventCallback(event, callbackType, cl22CEventCallback, userData);
}
