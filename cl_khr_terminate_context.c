#include "api.h"

cl_int cl22ExtTerminateContextKHR(void *fn, cl_context context)
{
    return ((clTerminateContextKHR_fn)(fn))(context);
}
