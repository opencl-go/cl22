#include "api.h"

extern void cl22GoMemObjectDestructorCallback(cl_mem, uintptr_t *);

static CL_CALLBACK void cl22CMemObjectDestructorCallback(cl_mem mem, void *userData)
{
    cl22GoMemObjectDestructorCallback(mem, (uintptr_t *)(userData));
}

cl_int cl22SetMemObjectDestructorCallback(cl_mem mem, uintptr_t *userData)
{
    return clSetMemObjectDestructorCallback(mem, cl22CMemObjectDestructorCallback, userData);
}
