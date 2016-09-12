package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/mumax/3/cuda/cu"
	"github.com/mumax/3/timer"
	"sync"
	"unsafe"
)

// CUDA handle for addcubicanisotropy2 kernel
var addcubicanisotropy2_code cu.Function

// Stores the arguments for addcubicanisotropy2 kernel invocation
type addcubicanisotropy2_args_t struct {
	arg_Bx      unsafe.Pointer
	arg_By      unsafe.Pointer
	arg_Bz      unsafe.Pointer
	arg_mx      unsafe.Pointer
	arg_my      unsafe.Pointer
	arg_mz      unsafe.Pointer
	arg_Ms_     unsafe.Pointer
	arg_Ms_mul  float32
	arg_k1_     unsafe.Pointer
	arg_k1_mul  float32
	arg_k2_     unsafe.Pointer
	arg_k2_mul  float32
	arg_k3_     unsafe.Pointer
	arg_k3_mul  float32
	arg_c1x_    unsafe.Pointer
	arg_c1x_mul float32
	arg_c1y_    unsafe.Pointer
	arg_c1y_mul float32
	arg_c1z_    unsafe.Pointer
	arg_c1z_mul float32
	arg_c2x_    unsafe.Pointer
	arg_c2x_mul float32
	arg_c2y_    unsafe.Pointer
	arg_c2y_mul float32
	arg_c2z_    unsafe.Pointer
	arg_c2z_mul float32
	arg_N       int
	argptr      [27]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for addcubicanisotropy2 kernel invocation
var addcubicanisotropy2_args addcubicanisotropy2_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	addcubicanisotropy2_args.argptr[0] = unsafe.Pointer(&addcubicanisotropy2_args.arg_Bx)
	addcubicanisotropy2_args.argptr[1] = unsafe.Pointer(&addcubicanisotropy2_args.arg_By)
	addcubicanisotropy2_args.argptr[2] = unsafe.Pointer(&addcubicanisotropy2_args.arg_Bz)
	addcubicanisotropy2_args.argptr[3] = unsafe.Pointer(&addcubicanisotropy2_args.arg_mx)
	addcubicanisotropy2_args.argptr[4] = unsafe.Pointer(&addcubicanisotropy2_args.arg_my)
	addcubicanisotropy2_args.argptr[5] = unsafe.Pointer(&addcubicanisotropy2_args.arg_mz)
	addcubicanisotropy2_args.argptr[6] = unsafe.Pointer(&addcubicanisotropy2_args.arg_Ms_)
	addcubicanisotropy2_args.argptr[7] = unsafe.Pointer(&addcubicanisotropy2_args.arg_Ms_mul)
	addcubicanisotropy2_args.argptr[8] = unsafe.Pointer(&addcubicanisotropy2_args.arg_k1_)
	addcubicanisotropy2_args.argptr[9] = unsafe.Pointer(&addcubicanisotropy2_args.arg_k1_mul)
	addcubicanisotropy2_args.argptr[10] = unsafe.Pointer(&addcubicanisotropy2_args.arg_k2_)
	addcubicanisotropy2_args.argptr[11] = unsafe.Pointer(&addcubicanisotropy2_args.arg_k2_mul)
	addcubicanisotropy2_args.argptr[12] = unsafe.Pointer(&addcubicanisotropy2_args.arg_k3_)
	addcubicanisotropy2_args.argptr[13] = unsafe.Pointer(&addcubicanisotropy2_args.arg_k3_mul)
	addcubicanisotropy2_args.argptr[14] = unsafe.Pointer(&addcubicanisotropy2_args.arg_c1x_)
	addcubicanisotropy2_args.argptr[15] = unsafe.Pointer(&addcubicanisotropy2_args.arg_c1x_mul)
	addcubicanisotropy2_args.argptr[16] = unsafe.Pointer(&addcubicanisotropy2_args.arg_c1y_)
	addcubicanisotropy2_args.argptr[17] = unsafe.Pointer(&addcubicanisotropy2_args.arg_c1y_mul)
	addcubicanisotropy2_args.argptr[18] = unsafe.Pointer(&addcubicanisotropy2_args.arg_c1z_)
	addcubicanisotropy2_args.argptr[19] = unsafe.Pointer(&addcubicanisotropy2_args.arg_c1z_mul)
	addcubicanisotropy2_args.argptr[20] = unsafe.Pointer(&addcubicanisotropy2_args.arg_c2x_)
	addcubicanisotropy2_args.argptr[21] = unsafe.Pointer(&addcubicanisotropy2_args.arg_c2x_mul)
	addcubicanisotropy2_args.argptr[22] = unsafe.Pointer(&addcubicanisotropy2_args.arg_c2y_)
	addcubicanisotropy2_args.argptr[23] = unsafe.Pointer(&addcubicanisotropy2_args.arg_c2y_mul)
	addcubicanisotropy2_args.argptr[24] = unsafe.Pointer(&addcubicanisotropy2_args.arg_c2z_)
	addcubicanisotropy2_args.argptr[25] = unsafe.Pointer(&addcubicanisotropy2_args.arg_c2z_mul)
	addcubicanisotropy2_args.argptr[26] = unsafe.Pointer(&addcubicanisotropy2_args.arg_N)
}

// Wrapper for addcubicanisotropy2 CUDA kernel, asynchronous.
func k_addcubicanisotropy2_async(Bx unsafe.Pointer, By unsafe.Pointer, Bz unsafe.Pointer, mx unsafe.Pointer, my unsafe.Pointer, mz unsafe.Pointer, Ms_ unsafe.Pointer, Ms_mul float32, k1_ unsafe.Pointer, k1_mul float32, k2_ unsafe.Pointer, k2_mul float32, k3_ unsafe.Pointer, k3_mul float32, c1x_ unsafe.Pointer, c1x_mul float32, c1y_ unsafe.Pointer, c1y_mul float32, c1z_ unsafe.Pointer, c1z_mul float32, c2x_ unsafe.Pointer, c2x_mul float32, c2y_ unsafe.Pointer, c2y_mul float32, c2z_ unsafe.Pointer, c2z_mul float32, N int, cfg *config) {
	if Synchronous { // debug
		Sync()
		timer.Start("addcubicanisotropy2")
	}

	addcubicanisotropy2_args.Lock()
	defer addcubicanisotropy2_args.Unlock()

	if addcubicanisotropy2_code == 0 {
		addcubicanisotropy2_code = fatbinLoad(addcubicanisotropy2_map, "addcubicanisotropy2")
	}

	addcubicanisotropy2_args.arg_Bx = Bx
	addcubicanisotropy2_args.arg_By = By
	addcubicanisotropy2_args.arg_Bz = Bz
	addcubicanisotropy2_args.arg_mx = mx
	addcubicanisotropy2_args.arg_my = my
	addcubicanisotropy2_args.arg_mz = mz
	addcubicanisotropy2_args.arg_Ms_ = Ms_
	addcubicanisotropy2_args.arg_Ms_mul = Ms_mul
	addcubicanisotropy2_args.arg_k1_ = k1_
	addcubicanisotropy2_args.arg_k1_mul = k1_mul
	addcubicanisotropy2_args.arg_k2_ = k2_
	addcubicanisotropy2_args.arg_k2_mul = k2_mul
	addcubicanisotropy2_args.arg_k3_ = k3_
	addcubicanisotropy2_args.arg_k3_mul = k3_mul
	addcubicanisotropy2_args.arg_c1x_ = c1x_
	addcubicanisotropy2_args.arg_c1x_mul = c1x_mul
	addcubicanisotropy2_args.arg_c1y_ = c1y_
	addcubicanisotropy2_args.arg_c1y_mul = c1y_mul
	addcubicanisotropy2_args.arg_c1z_ = c1z_
	addcubicanisotropy2_args.arg_c1z_mul = c1z_mul
	addcubicanisotropy2_args.arg_c2x_ = c2x_
	addcubicanisotropy2_args.arg_c2x_mul = c2x_mul
	addcubicanisotropy2_args.arg_c2y_ = c2y_
	addcubicanisotropy2_args.arg_c2y_mul = c2y_mul
	addcubicanisotropy2_args.arg_c2z_ = c2z_
	addcubicanisotropy2_args.arg_c2z_mul = c2z_mul
	addcubicanisotropy2_args.arg_N = N

	args := addcubicanisotropy2_args.argptr[:]
	cu.LaunchKernel(addcubicanisotropy2_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
		timer.Stop("addcubicanisotropy2")
	}
}

// maps compute capability on PTX code for addcubicanisotropy2 kernel.
var addcubicanisotropy2_map = map[int]string{0: "",
	20: addcubicanisotropy2_ptx_20,
	30: addcubicanisotropy2_ptx_30,
	35: addcubicanisotropy2_ptx_35,
	50: addcubicanisotropy2_ptx_50,
	52: addcubicanisotropy2_ptx_52,
	53: addcubicanisotropy2_ptx_53}

// addcubicanisotropy2 PTX code for various compute capabilities.
const (
	addcubicanisotropy2_ptx_20 = `
.version 4.3
.target sm_20
.address_size 64

	// .globl	addcubicanisotropy2

.visible .entry addcubicanisotropy2(
	.param .u64 addcubicanisotropy2_param_0,
	.param .u64 addcubicanisotropy2_param_1,
	.param .u64 addcubicanisotropy2_param_2,
	.param .u64 addcubicanisotropy2_param_3,
	.param .u64 addcubicanisotropy2_param_4,
	.param .u64 addcubicanisotropy2_param_5,
	.param .u64 addcubicanisotropy2_param_6,
	.param .f32 addcubicanisotropy2_param_7,
	.param .u64 addcubicanisotropy2_param_8,
	.param .f32 addcubicanisotropy2_param_9,
	.param .u64 addcubicanisotropy2_param_10,
	.param .f32 addcubicanisotropy2_param_11,
	.param .u64 addcubicanisotropy2_param_12,
	.param .f32 addcubicanisotropy2_param_13,
	.param .u64 addcubicanisotropy2_param_14,
	.param .f32 addcubicanisotropy2_param_15,
	.param .u64 addcubicanisotropy2_param_16,
	.param .f32 addcubicanisotropy2_param_17,
	.param .u64 addcubicanisotropy2_param_18,
	.param .f32 addcubicanisotropy2_param_19,
	.param .u64 addcubicanisotropy2_param_20,
	.param .f32 addcubicanisotropy2_param_21,
	.param .u64 addcubicanisotropy2_param_22,
	.param .f32 addcubicanisotropy2_param_23,
	.param .u64 addcubicanisotropy2_param_24,
	.param .f32 addcubicanisotropy2_param_25,
	.param .u32 addcubicanisotropy2_param_26
)
{



	ret;
}


`
	addcubicanisotropy2_ptx_30 = `
.version 4.3
.target sm_30
.address_size 64

	// .globl	addcubicanisotropy2

.visible .entry addcubicanisotropy2(
	.param .u64 addcubicanisotropy2_param_0,
	.param .u64 addcubicanisotropy2_param_1,
	.param .u64 addcubicanisotropy2_param_2,
	.param .u64 addcubicanisotropy2_param_3,
	.param .u64 addcubicanisotropy2_param_4,
	.param .u64 addcubicanisotropy2_param_5,
	.param .u64 addcubicanisotropy2_param_6,
	.param .f32 addcubicanisotropy2_param_7,
	.param .u64 addcubicanisotropy2_param_8,
	.param .f32 addcubicanisotropy2_param_9,
	.param .u64 addcubicanisotropy2_param_10,
	.param .f32 addcubicanisotropy2_param_11,
	.param .u64 addcubicanisotropy2_param_12,
	.param .f32 addcubicanisotropy2_param_13,
	.param .u64 addcubicanisotropy2_param_14,
	.param .f32 addcubicanisotropy2_param_15,
	.param .u64 addcubicanisotropy2_param_16,
	.param .f32 addcubicanisotropy2_param_17,
	.param .u64 addcubicanisotropy2_param_18,
	.param .f32 addcubicanisotropy2_param_19,
	.param .u64 addcubicanisotropy2_param_20,
	.param .f32 addcubicanisotropy2_param_21,
	.param .u64 addcubicanisotropy2_param_22,
	.param .f32 addcubicanisotropy2_param_23,
	.param .u64 addcubicanisotropy2_param_24,
	.param .f32 addcubicanisotropy2_param_25,
	.param .u32 addcubicanisotropy2_param_26
)
{



	ret;
}


`
	addcubicanisotropy2_ptx_35 = `
.version 4.3
.target sm_35
.address_size 64

	// .weak	cudaMalloc

.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaFuncGetAttributes
.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaDeviceGetAttribute
.weak .func  (.param .b32 func_retval0) cudaDeviceGetAttribute(
	.param .b64 cudaDeviceGetAttribute_param_0,
	.param .b32 cudaDeviceGetAttribute_param_1,
	.param .b32 cudaDeviceGetAttribute_param_2
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaGetDevice
.weak .func  (.param .b32 func_retval0) cudaGetDevice(
	.param .b64 cudaGetDevice_param_0
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaOccupancyMaxActiveBlocksPerMultiprocessor
.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessor(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_3
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags
.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_3,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_4
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .globl	addcubicanisotropy2
.visible .entry addcubicanisotropy2(
	.param .u64 addcubicanisotropy2_param_0,
	.param .u64 addcubicanisotropy2_param_1,
	.param .u64 addcubicanisotropy2_param_2,
	.param .u64 addcubicanisotropy2_param_3,
	.param .u64 addcubicanisotropy2_param_4,
	.param .u64 addcubicanisotropy2_param_5,
	.param .u64 addcubicanisotropy2_param_6,
	.param .f32 addcubicanisotropy2_param_7,
	.param .u64 addcubicanisotropy2_param_8,
	.param .f32 addcubicanisotropy2_param_9,
	.param .u64 addcubicanisotropy2_param_10,
	.param .f32 addcubicanisotropy2_param_11,
	.param .u64 addcubicanisotropy2_param_12,
	.param .f32 addcubicanisotropy2_param_13,
	.param .u64 addcubicanisotropy2_param_14,
	.param .f32 addcubicanisotropy2_param_15,
	.param .u64 addcubicanisotropy2_param_16,
	.param .f32 addcubicanisotropy2_param_17,
	.param .u64 addcubicanisotropy2_param_18,
	.param .f32 addcubicanisotropy2_param_19,
	.param .u64 addcubicanisotropy2_param_20,
	.param .f32 addcubicanisotropy2_param_21,
	.param .u64 addcubicanisotropy2_param_22,
	.param .f32 addcubicanisotropy2_param_23,
	.param .u64 addcubicanisotropy2_param_24,
	.param .f32 addcubicanisotropy2_param_25,
	.param .u32 addcubicanisotropy2_param_26
)
{



	ret;
}


`
	addcubicanisotropy2_ptx_50 = `
.version 4.3
.target sm_50
.address_size 64

	// .weak	cudaMalloc

.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaFuncGetAttributes
.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaDeviceGetAttribute
.weak .func  (.param .b32 func_retval0) cudaDeviceGetAttribute(
	.param .b64 cudaDeviceGetAttribute_param_0,
	.param .b32 cudaDeviceGetAttribute_param_1,
	.param .b32 cudaDeviceGetAttribute_param_2
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaGetDevice
.weak .func  (.param .b32 func_retval0) cudaGetDevice(
	.param .b64 cudaGetDevice_param_0
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaOccupancyMaxActiveBlocksPerMultiprocessor
.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessor(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_3
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags
.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_3,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_4
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .globl	addcubicanisotropy2
.visible .entry addcubicanisotropy2(
	.param .u64 addcubicanisotropy2_param_0,
	.param .u64 addcubicanisotropy2_param_1,
	.param .u64 addcubicanisotropy2_param_2,
	.param .u64 addcubicanisotropy2_param_3,
	.param .u64 addcubicanisotropy2_param_4,
	.param .u64 addcubicanisotropy2_param_5,
	.param .u64 addcubicanisotropy2_param_6,
	.param .f32 addcubicanisotropy2_param_7,
	.param .u64 addcubicanisotropy2_param_8,
	.param .f32 addcubicanisotropy2_param_9,
	.param .u64 addcubicanisotropy2_param_10,
	.param .f32 addcubicanisotropy2_param_11,
	.param .u64 addcubicanisotropy2_param_12,
	.param .f32 addcubicanisotropy2_param_13,
	.param .u64 addcubicanisotropy2_param_14,
	.param .f32 addcubicanisotropy2_param_15,
	.param .u64 addcubicanisotropy2_param_16,
	.param .f32 addcubicanisotropy2_param_17,
	.param .u64 addcubicanisotropy2_param_18,
	.param .f32 addcubicanisotropy2_param_19,
	.param .u64 addcubicanisotropy2_param_20,
	.param .f32 addcubicanisotropy2_param_21,
	.param .u64 addcubicanisotropy2_param_22,
	.param .f32 addcubicanisotropy2_param_23,
	.param .u64 addcubicanisotropy2_param_24,
	.param .f32 addcubicanisotropy2_param_25,
	.param .u32 addcubicanisotropy2_param_26
)
{



	ret;
}


`
	addcubicanisotropy2_ptx_52 = `
.version 4.3
.target sm_52
.address_size 64

	// .weak	cudaMalloc

.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaFuncGetAttributes
.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaDeviceGetAttribute
.weak .func  (.param .b32 func_retval0) cudaDeviceGetAttribute(
	.param .b64 cudaDeviceGetAttribute_param_0,
	.param .b32 cudaDeviceGetAttribute_param_1,
	.param .b32 cudaDeviceGetAttribute_param_2
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaGetDevice
.weak .func  (.param .b32 func_retval0) cudaGetDevice(
	.param .b64 cudaGetDevice_param_0
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaOccupancyMaxActiveBlocksPerMultiprocessor
.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessor(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_3
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags
.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_3,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_4
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .globl	addcubicanisotropy2
.visible .entry addcubicanisotropy2(
	.param .u64 addcubicanisotropy2_param_0,
	.param .u64 addcubicanisotropy2_param_1,
	.param .u64 addcubicanisotropy2_param_2,
	.param .u64 addcubicanisotropy2_param_3,
	.param .u64 addcubicanisotropy2_param_4,
	.param .u64 addcubicanisotropy2_param_5,
	.param .u64 addcubicanisotropy2_param_6,
	.param .f32 addcubicanisotropy2_param_7,
	.param .u64 addcubicanisotropy2_param_8,
	.param .f32 addcubicanisotropy2_param_9,
	.param .u64 addcubicanisotropy2_param_10,
	.param .f32 addcubicanisotropy2_param_11,
	.param .u64 addcubicanisotropy2_param_12,
	.param .f32 addcubicanisotropy2_param_13,
	.param .u64 addcubicanisotropy2_param_14,
	.param .f32 addcubicanisotropy2_param_15,
	.param .u64 addcubicanisotropy2_param_16,
	.param .f32 addcubicanisotropy2_param_17,
	.param .u64 addcubicanisotropy2_param_18,
	.param .f32 addcubicanisotropy2_param_19,
	.param .u64 addcubicanisotropy2_param_20,
	.param .f32 addcubicanisotropy2_param_21,
	.param .u64 addcubicanisotropy2_param_22,
	.param .f32 addcubicanisotropy2_param_23,
	.param .u64 addcubicanisotropy2_param_24,
	.param .f32 addcubicanisotropy2_param_25,
	.param .u32 addcubicanisotropy2_param_26
)
{



	ret;
}


`
	addcubicanisotropy2_ptx_53 = `
.version 4.3
.target sm_53
.address_size 64

	// .weak	cudaMalloc

.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaFuncGetAttributes
.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaDeviceGetAttribute
.weak .func  (.param .b32 func_retval0) cudaDeviceGetAttribute(
	.param .b64 cudaDeviceGetAttribute_param_0,
	.param .b32 cudaDeviceGetAttribute_param_1,
	.param .b32 cudaDeviceGetAttribute_param_2
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaGetDevice
.weak .func  (.param .b32 func_retval0) cudaGetDevice(
	.param .b64 cudaGetDevice_param_0
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaOccupancyMaxActiveBlocksPerMultiprocessor
.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessor(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_3
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags
.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_3,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_4
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .globl	addcubicanisotropy2
.visible .entry addcubicanisotropy2(
	.param .u64 addcubicanisotropy2_param_0,
	.param .u64 addcubicanisotropy2_param_1,
	.param .u64 addcubicanisotropy2_param_2,
	.param .u64 addcubicanisotropy2_param_3,
	.param .u64 addcubicanisotropy2_param_4,
	.param .u64 addcubicanisotropy2_param_5,
	.param .u64 addcubicanisotropy2_param_6,
	.param .f32 addcubicanisotropy2_param_7,
	.param .u64 addcubicanisotropy2_param_8,
	.param .f32 addcubicanisotropy2_param_9,
	.param .u64 addcubicanisotropy2_param_10,
	.param .f32 addcubicanisotropy2_param_11,
	.param .u64 addcubicanisotropy2_param_12,
	.param .f32 addcubicanisotropy2_param_13,
	.param .u64 addcubicanisotropy2_param_14,
	.param .f32 addcubicanisotropy2_param_15,
	.param .u64 addcubicanisotropy2_param_16,
	.param .f32 addcubicanisotropy2_param_17,
	.param .u64 addcubicanisotropy2_param_18,
	.param .f32 addcubicanisotropy2_param_19,
	.param .u64 addcubicanisotropy2_param_20,
	.param .f32 addcubicanisotropy2_param_21,
	.param .u64 addcubicanisotropy2_param_22,
	.param .f32 addcubicanisotropy2_param_23,
	.param .u64 addcubicanisotropy2_param_24,
	.param .f32 addcubicanisotropy2_param_25,
	.param .u32 addcubicanisotropy2_param_26
)
{



	ret;
}


`
)