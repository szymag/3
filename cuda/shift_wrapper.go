package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var shift_code cu.Function

type shift_args struct {
	arg_dst unsafe.Pointer
	arg_src unsafe.Pointer
	arg_N0  int
	arg_N1  int
	arg_N2  int
	arg_sh0 int
	arg_sh1 int
	arg_sh2 int
	argptr  [8]unsafe.Pointer
}

// Wrapper for shift CUDA kernel, asynchronous.
func k_shift_async(dst unsafe.Pointer, src unsafe.Pointer, N0 int, N1 int, N2 int, sh0 int, sh1 int, sh2 int, cfg *config, str cu.Stream) {
	if shift_code == 0 {
		shift_code = fatbinLoad(shift_map, "shift")
	}

	var a shift_args

	a.arg_dst = dst
	a.argptr[0] = unsafe.Pointer(&a.arg_dst)
	a.arg_src = src
	a.argptr[1] = unsafe.Pointer(&a.arg_src)
	a.arg_N0 = N0
	a.argptr[2] = unsafe.Pointer(&a.arg_N0)
	a.arg_N1 = N1
	a.argptr[3] = unsafe.Pointer(&a.arg_N1)
	a.arg_N2 = N2
	a.argptr[4] = unsafe.Pointer(&a.arg_N2)
	a.arg_sh0 = sh0
	a.argptr[5] = unsafe.Pointer(&a.arg_sh0)
	a.arg_sh1 = sh1
	a.argptr[6] = unsafe.Pointer(&a.arg_sh1)
	a.arg_sh2 = sh2
	a.argptr[7] = unsafe.Pointer(&a.arg_sh2)

	args := a.argptr[:]
	cu.LaunchKernel(shift_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, str, args)
}

// Wrapper for shift CUDA kernel, synchronized.
func k_shift(dst unsafe.Pointer, src unsafe.Pointer, N0 int, N1 int, N2 int, sh0 int, sh1 int, sh2 int, cfg *config) {
	str := stream()
	k_shift_async(dst, src, N0, N1, N2, sh0, sh1, sh2, cfg, str)
	syncAndRecycle(str)
}

var shift_map = map[int]string{0: "",
	20: shift_ptx_20,
	30: shift_ptx_30,
	35: shift_ptx_35}

const (
	shift_ptx_20 = `
.version 3.1
.target sm_20
.address_size 64


.visible .entry shift(
	.param .u64 shift_param_0,
	.param .u64 shift_param_1,
	.param .u32 shift_param_2,
	.param .u32 shift_param_3,
	.param .u32 shift_param_4,
	.param .u32 shift_param_5,
	.param .u32 shift_param_6,
	.param .u32 shift_param_7
)
{
	.reg .pred 	%p<7>;
	.reg .s32 	%r<46>;
	.reg .f32 	%f<2>;
	.reg .s64 	%rd<9>;


	ld.param.u64 	%rd3, [shift_param_0];
	ld.param.u64 	%rd4, [shift_param_1];
	ld.param.u32 	%r21, [shift_param_2];
	ld.param.u32 	%r22, [shift_param_3];
	ld.param.u32 	%r23, [shift_param_4];
	ld.param.u32 	%r24, [shift_param_5];
	ld.param.u32 	%r25, [shift_param_6];
	ld.param.u32 	%r26, [shift_param_7];
	cvta.to.global.u64 	%rd1, %rd3;
	cvta.to.global.u64 	%rd2, %rd4;
	.loc 2 10 1
	mov.u32 	%r1, %ntid.y;
	mov.u32 	%r2, %ctaid.y;
	mov.u32 	%r3, %tid.y;
	mad.lo.s32 	%r4, %r1, %r2, %r3;
	.loc 2 11 1
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %ctaid.x;
	mov.u32 	%r7, %tid.x;
	mad.lo.s32 	%r8, %r5, %r6, %r7;
	.loc 2 13 1
	setp.lt.s32 	%p1, %r8, %r23;
	setp.lt.s32 	%p2, %r4, %r22;
	and.pred  	%p3, %p2, %p1;
	.loc 2 18 1
	setp.gt.s32 	%p4, %r21, 0;
	.loc 2 13 1
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_3;
	bra.uni 	BB0_1;

BB0_1:
	.loc 2 19 1
	add.s32 	%r9, %r21, -1;
	sub.s32 	%r28, %r4, %r25;
	mov.u32 	%r27, 0;
	.loc 3 238 5
	max.s32 	%r29, %r28, %r27;
	.loc 2 19 1
	add.s32 	%r30, %r22, -1;
	.loc 3 210 5
	min.s32 	%r10, %r29, %r30;
	.loc 2 19 1
	sub.s32 	%r31, %r8, %r26;
	.loc 3 238 5
	max.s32 	%r32, %r31, %r27;
	.loc 2 19 1
	add.s32 	%r33, %r23, -1;
	.loc 3 210 5
	min.s32 	%r11, %r32, %r33;
	.loc 2 18 1
	neg.s32 	%r44, %r24;
	mad.lo.s32 	%r43, %r23, %r4, %r8;
	mul.lo.s32 	%r14, %r23, %r22;
	mov.u32 	%r45, %r27;

BB0_2:
	.loc 3 238 5
	mov.u32 	%r17, %r45;
	max.s32 	%r37, %r44, %r27;
	.loc 3 210 5
	min.s32 	%r38, %r37, %r9;
	.loc 2 19 1
	mad.lo.s32 	%r39, %r38, %r22, %r10;
	mad.lo.s32 	%r40, %r39, %r23, %r11;
	mul.wide.s32 	%rd5, %r40, 4;
	add.s64 	%rd6, %rd2, %rd5;
	mul.wide.s32 	%rd7, %r43, 4;
	add.s64 	%rd8, %rd1, %rd7;
	ld.global.f32 	%f1, [%rd6];
	st.global.f32 	[%rd8], %f1;
	.loc 2 18 1
	add.s32 	%r44, %r44, 1;
	add.s32 	%r43, %r43, %r14;
	.loc 2 18 18
	add.s32 	%r20, %r17, 1;
	.loc 2 18 1
	setp.lt.s32 	%p6, %r20, %r21;
	mov.u32 	%r45, %r20;
	@%p6 bra 	BB0_2;

BB0_3:
	.loc 2 21 2
	ret;
}


`
	shift_ptx_30 = `
.version 3.1
.target sm_30
.address_size 64


.visible .entry shift(
	.param .u64 shift_param_0,
	.param .u64 shift_param_1,
	.param .u32 shift_param_2,
	.param .u32 shift_param_3,
	.param .u32 shift_param_4,
	.param .u32 shift_param_5,
	.param .u32 shift_param_6,
	.param .u32 shift_param_7
)
{
	.reg .pred 	%p<7>;
	.reg .s32 	%r<46>;
	.reg .f32 	%f<2>;
	.reg .s64 	%rd<9>;


	ld.param.u64 	%rd3, [shift_param_0];
	ld.param.u64 	%rd4, [shift_param_1];
	ld.param.u32 	%r21, [shift_param_2];
	ld.param.u32 	%r22, [shift_param_3];
	ld.param.u32 	%r23, [shift_param_4];
	ld.param.u32 	%r24, [shift_param_5];
	ld.param.u32 	%r25, [shift_param_6];
	ld.param.u32 	%r26, [shift_param_7];
	cvta.to.global.u64 	%rd1, %rd3;
	cvta.to.global.u64 	%rd2, %rd4;
	.loc 2 10 1
	mov.u32 	%r1, %ntid.y;
	mov.u32 	%r2, %ctaid.y;
	mov.u32 	%r3, %tid.y;
	mad.lo.s32 	%r4, %r1, %r2, %r3;
	.loc 2 11 1
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %ctaid.x;
	mov.u32 	%r7, %tid.x;
	mad.lo.s32 	%r8, %r5, %r6, %r7;
	.loc 2 13 1
	setp.lt.s32 	%p1, %r8, %r23;
	setp.lt.s32 	%p2, %r4, %r22;
	and.pred  	%p3, %p2, %p1;
	.loc 2 18 1
	setp.gt.s32 	%p4, %r21, 0;
	.loc 2 13 1
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_3;
	bra.uni 	BB0_1;

BB0_1:
	.loc 2 19 1
	add.s32 	%r9, %r21, -1;
	sub.s32 	%r28, %r4, %r25;
	mov.u32 	%r27, 0;
	.loc 3 238 5
	max.s32 	%r29, %r28, %r27;
	.loc 2 19 1
	add.s32 	%r30, %r22, -1;
	.loc 3 210 5
	min.s32 	%r10, %r29, %r30;
	.loc 2 19 1
	sub.s32 	%r31, %r8, %r26;
	.loc 3 238 5
	max.s32 	%r32, %r31, %r27;
	.loc 2 19 1
	add.s32 	%r33, %r23, -1;
	.loc 3 210 5
	min.s32 	%r11, %r32, %r33;
	.loc 2 18 1
	neg.s32 	%r44, %r24;
	mad.lo.s32 	%r43, %r23, %r4, %r8;
	mul.lo.s32 	%r14, %r23, %r22;
	mov.u32 	%r45, %r27;

BB0_2:
	.loc 3 238 5
	mov.u32 	%r17, %r45;
	max.s32 	%r37, %r44, %r27;
	.loc 3 210 5
	min.s32 	%r38, %r37, %r9;
	.loc 2 19 1
	mad.lo.s32 	%r39, %r38, %r22, %r10;
	mad.lo.s32 	%r40, %r39, %r23, %r11;
	mul.wide.s32 	%rd5, %r40, 4;
	add.s64 	%rd6, %rd2, %rd5;
	mul.wide.s32 	%rd7, %r43, 4;
	add.s64 	%rd8, %rd1, %rd7;
	ld.global.f32 	%f1, [%rd6];
	st.global.f32 	[%rd8], %f1;
	.loc 2 18 1
	add.s32 	%r44, %r44, 1;
	add.s32 	%r43, %r43, %r14;
	.loc 2 18 18
	add.s32 	%r20, %r17, 1;
	.loc 2 18 1
	setp.lt.s32 	%p6, %r20, %r21;
	mov.u32 	%r45, %r20;
	@%p6 bra 	BB0_2;

BB0_3:
	.loc 2 21 2
	ret;
}


`
	shift_ptx_35 = `
.version 3.1
.target sm_35
.address_size 64


.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 66 3
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 71 3
	ret;
}

.visible .entry shift(
	.param .u64 shift_param_0,
	.param .u64 shift_param_1,
	.param .u32 shift_param_2,
	.param .u32 shift_param_3,
	.param .u32 shift_param_4,
	.param .u32 shift_param_5,
	.param .u32 shift_param_6,
	.param .u32 shift_param_7
)
{
	.reg .pred 	%p<7>;
	.reg .s32 	%r<45>;
	.reg .f32 	%f<2>;
	.reg .s64 	%rd<9>;


	ld.param.u64 	%rd3, [shift_param_0];
	ld.param.u64 	%rd4, [shift_param_1];
	ld.param.u32 	%r21, [shift_param_2];
	ld.param.u32 	%r22, [shift_param_3];
	ld.param.u32 	%r23, [shift_param_4];
	ld.param.u32 	%r24, [shift_param_5];
	ld.param.u32 	%r25, [shift_param_6];
	ld.param.u32 	%r26, [shift_param_7];
	cvta.to.global.u64 	%rd1, %rd3;
	cvta.to.global.u64 	%rd2, %rd4;
	.loc 3 10 1
	mov.u32 	%r1, %ntid.y;
	mov.u32 	%r2, %ctaid.y;
	mov.u32 	%r3, %tid.y;
	mad.lo.s32 	%r4, %r1, %r2, %r3;
	.loc 3 11 1
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %ctaid.x;
	mov.u32 	%r7, %tid.x;
	mad.lo.s32 	%r8, %r5, %r6, %r7;
	.loc 3 13 1
	setp.lt.s32 	%p1, %r8, %r23;
	setp.lt.s32 	%p2, %r4, %r22;
	and.pred  	%p3, %p2, %p1;
	.loc 3 18 1
	setp.gt.s32 	%p4, %r21, 0;
	.loc 3 13 1
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB2_3;
	bra.uni 	BB2_1;

BB2_1:
	.loc 3 19 1
	add.s32 	%r9, %r21, -1;
	sub.s32 	%r28, %r4, %r25;
	mov.u32 	%r27, 0;
	.loc 4 238 5
	max.s32 	%r29, %r28, %r27;
	.loc 3 19 1
	add.s32 	%r30, %r22, -1;
	.loc 4 210 5
	min.s32 	%r10, %r29, %r30;
	.loc 3 19 1
	sub.s32 	%r31, %r8, %r26;
	.loc 4 238 5
	max.s32 	%r32, %r31, %r27;
	.loc 3 19 1
	add.s32 	%r33, %r23, -1;
	.loc 4 210 5
	min.s32 	%r11, %r32, %r33;
	.loc 3 18 1
	neg.s32 	%r43, %r24;
	mad.lo.s32 	%r42, %r23, %r4, %r8;
	mul.lo.s32 	%r14, %r23, %r22;
	mov.u32 	%r44, %r27;

BB2_2:
	.loc 4 238 5
	mov.u32 	%r17, %r44;
	max.s32 	%r37, %r43, %r27;
	.loc 4 210 5
	min.s32 	%r38, %r37, %r9;
	.loc 3 19 1
	mad.lo.s32 	%r39, %r38, %r22, %r10;
	mad.lo.s32 	%r40, %r39, %r23, %r11;
	mul.wide.s32 	%rd5, %r40, 4;
	add.s64 	%rd6, %rd2, %rd5;
	ld.global.nc.f32 	%f1, [%rd6];
	mul.wide.s32 	%rd7, %r42, 4;
	add.s64 	%rd8, %rd1, %rd7;
	st.global.f32 	[%rd8], %f1;
	.loc 3 18 1
	add.s32 	%r43, %r43, 1;
	add.s32 	%r42, %r42, %r14;
	.loc 3 18 18
	add.s32 	%r20, %r17, 1;
	.loc 3 18 1
	setp.lt.s32 	%p6, %r20, %r21;
	mov.u32 	%r44, %r20;
	@%p6 bra 	BB2_2;

BB2_3:
	.loc 3 21 2
	ret;
}


`
)
