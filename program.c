#include "api.h"

extern void cl22GoProgramBuildCallback(cl_program, uintptr_t *);

static CL_CALLBACK void cl22CProgramBuildCallback(cl_program program, void *userData)
{
    cl22GoProgramBuildCallback(program, (uintptr_t *)(userData));
}

cl_int cl22BuildProgram(cl_program program,
    cl_uint numDevices, cl_device_id *devices,
    char *options, uintptr_t *userData)
{
    return clBuildProgram(program, numDevices, devices, options,
        (userData != NULL) ? cl22CProgramBuildCallback : NULL, userData);
}

extern void cl22GoProgramCompileCallback(cl_program, uintptr_t *);

static CL_CALLBACK void cl22CProgramCompileCallback(cl_program program, void *userData)
{
    cl22GoProgramCompileCallback(program, (uintptr_t *)(userData));
}

cl_int cl22CompileProgram(cl_program program,
    cl_uint numDevices, cl_device_id *devices,
    char *options,
    cl_uint numInputHeaders, cl_program *headers, char const **includeNames,
    uintptr_t *userData)
{
    return clCompileProgram(program, numDevices, devices, options,
        numInputHeaders, headers, includeNames,
        (userData != NULL) ? cl22CProgramCompileCallback : NULL, userData);
}

extern void cl22GoProgramLinkCallback(cl_program, uintptr_t *);

static CL_CALLBACK void cl22CProgramLinkCallback(cl_program program, void *userData)
{
    cl22GoProgramLinkCallback(program, (uintptr_t *)(userData));
}

cl_program cl22LinkProgram(cl_context context,
    cl_uint numDevices, cl_device_id *devices,
    char *options,
    cl_uint numInputPrograms, cl_program *programs,
    uintptr_t *userData,
    cl_int *errReturn)
{
    return clLinkProgram(context, numDevices, devices, options,
        numInputPrograms, programs,
        (userData != NULL) ? cl22CProgramLinkCallback : NULL, userData,
        errReturn);
}

extern void cl22GoProgramReleaseCallback(cl_program, uintptr_t *);

static CL_CALLBACK void cl22CProgramReleaseCallback(cl_program program, void *userData)
{
    cl22GoProgramReleaseCallback(program, (uintptr_t *)(userData));
}

cl_int cl22SetProgramReleaseCallback(cl_program program, uintptr_t *userData)
{
    return clSetProgramReleaseCallback(program, cl22CProgramReleaseCallback, userData);
}
