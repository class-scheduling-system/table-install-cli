/*
 * --------------------------------------------------------------------------------
 * Copyright (c) 2022-NOW(至今) 锋楪技术团队
 * Author: 锋楪技术团队 (https://www.frontleaves.com)
 *
 * 本文件包含锋楪技术团队项目的源代码，项目的所有源代码均遵循 MIT 开源许可证协议。
 * --------------------------------------------------------------------------------
 * 许可证声明：
 *
 * 版权所有 (c) 2022-2025 锋楪技术团队。保留所有权利。
 *
 * 本软件是“按原样”提供的，没有任何形式的明示或暗示的保证，包括但不限于
 * 对适销性、特定用途的适用性和非侵权性的暗示保证。在任何情况下，
 * 作者或版权持有人均不承担因软件或软件的使用或其他交易而产生的、
 * 由此引起的或以任何方式与此软件有关的任何索赔、损害或其他责任。
 *
 * 使用本软件即表示您了解此声明并同意其条款。
 *
 * 有关 MIT 许可证的更多信息，请查看项目根目录下的 LICENSE 文件或访问：
 * https://opensource.org/licenses/MIT
 * --------------------------------------------------------------------------------
 * 免责声明：
 *
 * 使用本软件的风险由用户自担。作者或版权持有人在法律允许的最大范围内，
 * 对因使用本软件内容而导致的任何直接或间接的损失不承担任何责任。
 * --------------------------------------------------------------------------------
 */

package setup

import "math/rand/v2"

func (is *InitStruct) OperateSetupClassroom() {

	is.operate.InitClassroom(
		"HXGCZX201", // number
		"HXGC2#201-化工分析实验室（一）", // name
		"铁门关校区",                // campus
		"化学工程实训中心",             // building
		2,                      // floor
		"多媒体教室",                // classroomType
		nil,                    // classroomTag
		50,                     // capacity
		nil,                    // examinationRoomCapacity
		false,                  // examinationRoom
		false,                  // isMultimedia
		true,                   // isAirConditioned
		false,                  // status
		nil,                    // description
		nil,                    // managementDepartment
		rand.Float64()*100,     // area
		nil,                    // tablesChairsType
	)

	is.operate.InitClassroom(
		"HXGCZX202",        // number
		"HXGC2#202",        // name
		"铁门关校区",            // campus
		"化学工程实训中心",         // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"HXGCZX203", // number
		"HXGC2#203-化学分析实验室（二）", // name
		"铁门关校区",                // campus
		"化学工程实训中心",             // building
		2,                      // floor
		"多媒体教室",                // classroomType
		nil,                    // classroomTag
		50,                     // capacity
		nil,                    // examinationRoomCapacity
		false,                  // examinationRoom
		false,                  // isMultimedia
		true,                   // isAirConditioned
		false,                  // status
		nil,                    // description
		nil,                    // managementDepartment
		rand.Float64()*100,     // area
		nil,                    // tablesChairsType
	)

	is.operate.InitClassroom(
		"HXGCZX204",        // number
		"HXGC2#204",        // name
		"铁门关校区",            // campus
		"化学工程实训中心",         // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"HXGCZX205",        // number
		"HXGC2#205",        // name
		"铁门关校区",            // campus
		"化学工程实训中心",         // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"HXGCZX207",        // number
		"HXGC2#207",        // name
		"铁门关校区",            // campus
		"化学工程实训中心",         // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"HXGCZX209",        // number
		"HXGC2#209",        // name
		"铁门关校区",            // campus
		"化学工程实训中心",         // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"HXGCZX211",        // number
		"HXGC2#211",        // name
		"铁门关校区",            // campus
		"化学工程实训中心",         // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JDGCZX105",        // number
		"JDGC1#105-智能焊接加工", // name
		"铁门关校区",            // campus
		"机电工程实训中心",         // building
		1,                  // floor
		"实训车间",             // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JDGCZX106",        // number
		"JDGC1#106-数控仿真机房", // name
		"铁门关校区",            // campus
		"机电工程实训中心",         // building
		1,                  // floor
		"机房",               // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"XXJSZX304",        // number
		"XXJS3#304电子商务机房",  // name
		"铁门关校区",            // campus
		"信息技术中心",           // building
		3,                  // floor
		"机房",               // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		true,               // examinationRoom
		true,               // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"XXJSZX303",        // number
		"XXJS3#303机房",      // name
		"铁门关校区",            // campus
		"信息技术中心",           // building
		3,                  // floor
		"机房",               // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"XXJSZX301",        // number
		"XXJS3#301多媒体教室",   // name
		"铁门关校区",            // campus
		"信息技术中心",           // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"XXJSZX302",        // number
		"XXJS3#302录播教室",    // name
		"铁门关校区",            // campus
		"信息技术中心",           // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JDGCZX104",        // number
		"JDGC1#104",        // name
		"铁门关校区",            // campus
		"机电工程实训中心",         // building
		1,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL4-406",         // number
		"4#406",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL3-302",         // number
		"3#302",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"XDNYZX302",        // number
		"XDNY3#302-公共机房",   // name
		"铁门关校区",            // campus
		"现代农业实训中心",         // building
		3,                  // floor
		"机房",               // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"XDNYZX201",        // number
		"XDNY2#201",        // name
		"铁门关校区",            // campus
		"现代农业实训中心",         // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		35,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"CJ1-hangongfang2", // number
		"CJ1-焊工房2",         // name
		"库尔勒校区",            // campus
		"车间",               // building
		1,                  // floor
		"实训车间",             // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"CJ1-hangongfang1", // number
		"CJ1-焊工房1",         // name
		"库尔勒校区",            // campus
		"车间",               // building
		1,                  // floor
		"实训车间",             // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL5-511",         // number
		"5-511",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JDGCZX207",        // number
		"JDGC2#207",        // name
		"铁门关校区",            // campus
		"机电工程实训中心",         // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"QCGCZX3-301",          // number
		"QCGCZX3-301-新能源电机实训室", // name
		"铁门关校区",                // campus
		"汽车工程实训中心",             // building
		3,                      // floor
		"汽修实训室",                // classroomType
		nil,                    // classroomTag
		50,                     // capacity
		nil,                    // examinationRoomCapacity
		false,                  // examinationRoom
		false,                  // isMultimedia
		true,                   // isAirConditioned
		false,                  // status
		nil,                    // description
		nil,                    // managementDepartment
		rand.Float64()*100,     // area
		nil,                    // tablesChairsType
	)

	is.operate.InitClassroom(
		"CJ1-qiangongfang", // number
		"CJ1-钳工房",          // name
		"库尔勒校区",            // campus
		"车间",               // building
		1,                  // floor
		"实训车间",             // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL519",           // number
		"JXL5#519",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL420",           // number
		"JXL4#420",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"XDNYZX306",        // number
		"XDNY3#306",        // name
		"铁门关校区",            // campus
		"现代农业实训中心",         // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		35,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"SXL3-305",         // number
		"3#305",            // name
		"库尔勒校区",            // campus
		"实训楼",              // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"XDNYZX209",        // number
		"XDNY2#209",        // name
		"铁门关校区",            // campus
		"现代农业实训中心",         // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		35,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL3-309",         // number
		"3#309",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JNL4#402-蒙台梭立实训室", // number
		"4#402-蒙台梭立实训室",    // name
		"库尔勒校区",            // campus
		"技能训练中心",           // building
		4,                  // floor
		"蒙台梭立实训室",          // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL201",           // number
		"JXL2#201",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"HXGCZX307",        // number
		"HXGC3#307",        // name
		"铁门关校区",            // campus
		"化学工程实训中心",         // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"QCGCZX2-205", // number
		"QCGCZX2-205-汽车钣金综合实训室", // name
		"铁门关校区",                 // campus
		"汽车工程实训中心",              // building
		2,                       // floor
		"汽修实训室",                 // classroomType
		nil,                     // classroomTag
		80,                      // capacity
		nil,                     // examinationRoomCapacity
		false,                   // examinationRoom
		false,                   // isMultimedia
		true,                    // isAirConditioned
		false,                   // status
		nil,                     // description
		nil,                     // managementDepartment
		rand.Float64()*100,      // area
		nil,                     // tablesChairsType
	)

	is.operate.InitClassroom(
		"JDGCZX305",        // number
		"JDGC3#305",        // name
		"铁门关校区",            // campus
		"机电工程实训中心",         // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL404",           // number
		"JXL4#404",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JYYSZX203", // number
		"JYYS2#203-幼儿园虚拟仿真实训室", // name
		"铁门关校区",                // campus
		"教育艺术中心",               // building
		2,                      // floor
		"虚拟仿真实训室",              // classroomType
		nil,                    // classroomTag
		80,                     // capacity
		nil,                    // examinationRoomCapacity
		true,                   // examinationRoom
		true,                   // isMultimedia
		true,                   // isAirConditioned
		false,                  // status
		nil,                    // description
		nil,                    // managementDepartment
		rand.Float64()*100,     // area
		nil,                    // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL413",           // number
		"JXL4#413-古筝实训室",   // name
		"铁门关校区",            // campus
		"教学楼",              // building
		4,                  // floor
		"琴房",               // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"HXGCZX306",        // number
		"HXGC3#306",        // name
		"铁门关校区",            // campus
		"化学工程实训中心",         // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JDGCZX312",        // number
		"JDGC3#312-PLC实训室", // name
		"铁门关校区",            // campus
		"机电工程实训中心",         // building
		3,                  // floor
		"电工实训室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JDGCZX311", // number
		"JDGC3#311-电气控制实训室（二）", // name
		"铁门关校区",                // campus
		"机电工程实训中心",             // building
		3,                      // floor
		"电工实训室",                // classroomType
		nil,                    // classroomTag
		50,                     // capacity
		nil,                    // examinationRoomCapacity
		false,                  // examinationRoom
		false,                  // isMultimedia
		true,                   // isAirConditioned
		false,                  // status
		nil,                    // description
		nil,                    // managementDepartment
		rand.Float64()*100,     // area
		nil,                    // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL5#511",         // number
		"5#511",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JNL3#309-3号琴房（电钢）", // number
		"3#309-3号琴房（电钢）",    // name
		"库尔勒校区",             // campus
		"技能训练中心",            // building
		3,                   // floor
		"琴房",                // classroomType
		nil,                 // classroomTag
		80,                  // capacity
		nil,                 // examinationRoomCapacity
		false,               // examinationRoom
		false,               // isMultimedia
		true,                // isAirConditioned
		false,               // status
		nil,                 // description
		nil,                 // managementDepartment
		rand.Float64()*100,  // area
		nil,                 // tablesChairsType
	)

	is.operate.InitClassroom(
		"XDNYZX101",        // number
		"XDNY1#101",        // name
		"铁门关校区",            // campus
		"现代农业实训中心",         // building
		1,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		35,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL205",           // number
		"JXL2#205",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"XDNYZX206",        // number
		"XDNY2#206",        // name
		"铁门关校区",            // campus
		"现代农业实训中心",         // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		35,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL210",            // number
		"JXL2#210-美术实训室（一）", // name
		"铁门关校区",             // campus
		"教学楼",               // building
		2,                   // floor
		"美术教室",              // classroomType
		nil,                 // classroomTag
		50,                  // capacity
		nil,                 // examinationRoomCapacity
		false,               // examinationRoom
		false,               // isMultimedia
		true,                // isAirConditioned
		false,               // status
		nil,                 // description
		nil,                 // managementDepartment
		rand.Float64()*100,  // area
		nil,                 // tablesChairsType
	)

	is.operate.InitClassroom(
		"JDGCZX209",         // number
		"JDGC2#209-现代电气实训室", // name
		"铁门关校区",             // campus
		"机电工程实训中心",          // building
		2,                   // floor
		"现代电气控制实训室",         // classroomType
		nil,                 // classroomTag
		50,                  // capacity
		nil,                 // examinationRoomCapacity
		false,               // examinationRoom
		false,               // isMultimedia
		true,                // isAirConditioned
		false,               // status
		nil,                 // description
		nil,                 // managementDepartment
		rand.Float64()*100,  // area
		nil,                 // tablesChairsType
	)

	is.operate.InitClassroom(
		"XDNYZX203",        // number
		"XDNY2#203",        // name
		"铁门关校区",            // campus
		"现代农业实训中心",         // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		35,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL3-301",         // number
		"3#301",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL5#507",         // number
		"5#507",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"CJ1-huagongchejian", // number
		"CJ1-化工车间",           // name
		"库尔勒校区",              // campus
		"车间",                 // building
		1,                    // floor
		"实训车间",               // classroomType
		nil,                  // classroomTag
		80,                   // capacity
		nil,                  // examinationRoomCapacity
		false,                // examinationRoom
		false,                // isMultimedia
		true,                 // isAirConditioned
		false,                // status
		nil,                  // description
		nil,                  // managementDepartment
		rand.Float64()*100,   // area
		nil,                  // tablesChairsType
	)

	is.operate.InitClassroom(
		"XDNYZX103",        // number
		"XDNY1#103",        // name
		"铁门关校区",            // campus
		"现代农业实训中心",         // building
		1,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		35,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL5-509",         // number
		"5-509",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JNL3#306-1号舞蹈室",   // number
		"3#306-1号舞蹈室",      // name
		"库尔勒校区",            // campus
		"技能训练中心",           // building
		3,                  // floor
		"舞蹈教室",             // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"HXGCZX210",        // number
		"HXGC2#210",        // name
		"铁门关校区",            // campus
		"化学工程实训中心",         // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL403",           // number
		"JXL4#403",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL411",           // number
		"JXL4#411",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL417",           // number
		"JXL4#417",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"SXL4#407",         // number
		"4#407",            // name
		"库尔勒校区",            // campus
		"实训楼",              // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL321",           // number
		"JXL3#321",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"HXGCZX206",        // number
		"HXGC2#206",        // name
		"铁门关校区",            // campus
		"化学工程实训中心",         // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL503",           // number
		"JXL5#503",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL3-315",         // number
		"3#315",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JDGCZX107",        // number
		"JDGC1#107-数控实训室",  // name
		"铁门关校区",            // campus
		"机电工程实训中心",         // building
		1,                  // floor
		"实训车间",             // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL405",           // number
		"JXL4#405",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"HXGCZX208",        // number
		"HXGC2#208",        // name
		"铁门关校区",            // campus
		"化学工程实训中心",         // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL506",           // number
		"JXL5#506",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL504",           // number
		"JXL5#504",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JDGCZX110",        // number
		"1#110-焊接实训中心",     // name
		"铁门关校区",            // campus
		"机电工程实训中心",         // building
		1,                  // floor
		"实训车间",             // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		false,              // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"HXGCZX303",        // number
		"HXGC3#303",        // name
		"铁门关校区",            // campus
		"化学工程实训中心",         // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL209",           // number
		"JXL2#209",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		2,                  // floor
		"机房",               // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"HXGCZX302",        // number
		"HXGC3#302",        // name
		"铁门关校区",            // campus
		"化学工程实训中心",         // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"XDNYZX104",        // number
		"XDNY1#104",        // name
		"铁门关校区",            // campus
		"现代农业实训中心",         // building
		1,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		35,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL2-211",         // number
		"2#211",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JNL3#304-1号琴房",    // number
		"3#304-1号琴房",       // name
		"库尔勒校区",            // campus
		"技能训练中心",           // building
		3,                  // floor
		"琴房",               // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"XDNYZX301",        // number
		"XDNY3#301-公共机房",   // name
		"铁门关校区",            // campus
		"现代农业实训中心",         // building
		3,                  // floor
		"机房",               // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"XDNYZX210",        // number
		"XDNY2#210",        // name
		"铁门关校区",            // campus
		"现代农业实训中心",         // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		35,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JNL4#401-3号舞蹈室",   // number
		"4#401-3号舞蹈室",      // name
		"库尔勒校区",            // campus
		"技能训练中心",           // building
		4,                  // floor
		"舞蹈教室",             // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"HXGCZX108",           // number
		"HXGC1#108-化工实训车间(二)", // name
		"铁门关校区",               // campus
		"化学工程实训中心",            // building
		1,                     // floor
		"实训车间",                // classroomType
		nil,                   // classroomTag
		120,                   // capacity
		nil,                   // examinationRoomCapacity
		false,                 // examinationRoom
		false,                 // isMultimedia
		true,                  // isAirConditioned
		false,                 // status
		nil,                   // description
		nil,                   // managementDepartment
		rand.Float64()*100,    // area
		nil,                   // tablesChairsType
	)

	is.operate.InitClassroom(
		"JYYSZX314",           // number
		"JYYS3#314-电钢琴实训室（五）", // name
		"铁门关校区",               // campus
		"教育艺术中心",              // building
		3,                     // floor
		"琴房",                  // classroomType
		nil,                   // classroomTag
		50,                    // capacity
		nil,                   // examinationRoomCapacity
		false,                 // examinationRoom
		false,                 // isMultimedia
		true,                  // isAirConditioned
		false,                 // status
		nil,                   // description
		nil,                   // managementDepartment
		rand.Float64()*100,    // area
		nil,                   // tablesChairsType
	)

	is.operate.InitClassroom(
		"SXL4#410",         // number
		"4#410",            // name
		"库尔勒校区",            // campus
		"实训楼",              // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"QCGCZX3-304",          // number
		"QCGCZX3-304-低压电工室(2)", // name
		"铁门关校区",                // campus
		"汽车工程实训中心",             // building
		3,                      // floor
		"汽修实训室",                // classroomType
		nil,                    // classroomTag
		50,                     // capacity
		nil,                    // examinationRoomCapacity
		false,                  // examinationRoom
		false,                  // isMultimedia
		true,                   // isAirConditioned
		false,                  // status
		nil,                    // description
		nil,                    // managementDepartment
		rand.Float64()*100,     // area
		nil,                    // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL4-401",         // number
		"4#401",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"XDNYZX204",        // number
		"XDNY2#204",        // name
		"铁门关校区",            // campus
		"现代农业实训中心",         // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		35,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL3-314",         // number
		"3#314",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL521",           // number
		"JXL5#521",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JNL1#-数控加工中心",     // number
		"1#数控加工中心",         // name
		"库尔勒校区",            // campus
		"技能训练中心",           // building
		1,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL223",           // number
		"JXL2#223",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL513",           // number
		"JXL5#513",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"QCGCZX3-309", // number
		"QCGCZX3-309-智能网联综合实训室", // name
		"铁门关校区",                 // campus
		"汽车工程实训中心",              // building
		3,                       // floor
		"汽修实训室",                 // classroomType
		nil,                     // classroomTag
		100,                     // capacity
		nil,                     // examinationRoomCapacity
		false,                   // examinationRoom
		false,                   // isMultimedia
		true,                    // isAirConditioned
		false,                   // status
		nil,                     // description
		nil,                     // managementDepartment
		rand.Float64()*100,      // area
		nil,                     // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL409",           // number
		"JXL4#409",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL515",           // number
		"JXL5#515",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL4-403",         // number
		"4#403",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL517",           // number
		"JXL5#517",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"SXL3#304",         // number
		"3#304",            // name
		"库尔勒校区",            // campus
		"实训楼",              // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"HXGCZX106",            // number
		"HXGC1#106-化工仿真实训室(二)", // name
		"铁门关校区",                // campus
		"化学工程实训中心",             // building
		1,                      // floor
		"化工仿真实训室",              // classroomType
		nil,                    // classroomTag
		40,                     // capacity
		nil,                    // examinationRoomCapacity
		false,                  // examinationRoom
		false,                  // isMultimedia
		true,                   // isAirConditioned
		false,                  // status
		nil,                    // description
		nil,                    // managementDepartment
		rand.Float64()*100,     // area
		nil,                    // tablesChairsType
	)

	is.operate.InitClassroom(
		"JYYSZX302",           // number
		"JYYS3#302-电钢琴实训室（二）", // name
		"铁门关校区",               // campus
		"教育艺术中心",              // building
		3,                     // floor
		"琴房",                  // classroomType
		nil,                   // classroomTag
		50,                    // capacity
		nil,                   // examinationRoomCapacity
		false,                 // examinationRoom
		false,                 // isMultimedia
		true,                  // isAirConditioned
		false,                 // status
		nil,                   // description
		nil,                   // managementDepartment
		rand.Float64()*100,    // area
		nil,                   // tablesChairsType
	)

	is.operate.InitClassroom(
		"JYYSZX305",        // number
		"JYYS3#305-乐理实训室",  // name
		"铁门关校区",            // campus
		"教育艺术中心",           // building
		3,                  // floor
		"声乐教室",             // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"SXL4#409",         // number
		"4#409",            // name
		"库尔勒校区",            // campus
		"实训楼",              // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JNL3#305-1号声乐室",   // number
		"3#305-1号声乐室",      // name
		"库尔勒校区",            // campus
		"技能训练中心",           // building
		3,                  // floor
		"声乐教室",             // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JDGCZX204",           // number
		"JDGC2#204-自动化生产线实训室", // name
		"铁门关校区",               // campus
		"机电工程实训中心",            // building
		2,                     // floor
		"自动化生产线实训室",           // classroomType
		nil,                   // classroomTag
		50,                    // capacity
		nil,                   // examinationRoomCapacity
		false,                 // examinationRoom
		false,                 // isMultimedia
		true,                  // isAirConditioned
		false,                 // status
		nil,                   // description
		nil,                   // managementDepartment
		rand.Float64()*100,    // area
		nil,                   // tablesChairsType
	)

	is.operate.InitClassroom(
		"SXL1#104",         // number
		"1#104",            // name
		"库尔勒校区",            // campus
		"实训楼",              // building
		1,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"CJ1-qixiuchejian", // number
		"CJ1-汽修车间",         // name
		"库尔勒校区",            // campus
		"车间",               // building
		1,                  // floor
		"实训车间",             // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JDGCZX308",          // number
		"JDGC3#308-电气实训室（二）", // name
		"铁门关校区",              // campus
		"机电工程实训中心",           // building
		3,                    // floor
		"电工实训室",              // classroomType
		nil,                  // classroomTag
		50,                   // capacity
		nil,                  // examinationRoomCapacity
		false,                // examinationRoom
		false,                // isMultimedia
		true,                 // isAirConditioned
		false,                // status
		nil,                  // description
		nil,                  // managementDepartment
		rand.Float64()*100,   // area
		nil,                  // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL5#508",         // number
		"5#508",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JNL2#208-液压气动室",   // number
		"2#208-液压气动",       // name
		"库尔勒校区",            // campus
		"技能训练中心",           // building
		2,                  // floor
		"液压气动实训室",          // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL2-xiaojifang",  // number
		"2#小机房",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		2,                  // floor
		"机房",               // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL5-508",         // number
		"5-508",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL415",            // number
		"JXL4#415-乐理实训室（一）", // name
		"铁门关校区",             // campus
		"教学楼",               // building
		4,                   // floor
		"声乐教室",              // classroomType
		nil,                 // classroomTag
		50,                  // capacity
		nil,                 // examinationRoomCapacity
		false,               // examinationRoom
		false,               // isMultimedia
		true,                // isAirConditioned
		false,               // status
		nil,                 // description
		nil,                 // managementDepartment
		rand.Float64()*100,  // area
		nil,                 // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL2-202",         // number
		"2#202",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JNL3#302-美术室",     // number
		"3#302-美术室",        // name
		"库尔勒校区",            // campus
		"技能训练中心",           // building
		3,                  // floor
		"美术教室",             // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JYYSZX104",          // number
		"JYYS1#104-舞蹈实训室（四）", // name
		"铁门关校区",              // campus
		"教育艺术中心",             // building
		1,                    // floor
		"舞蹈教室",               // classroomType
		nil,                  // classroomTag
		50,                   // capacity
		nil,                  // examinationRoomCapacity
		false,                // examinationRoom
		false,                // isMultimedia
		true,                 // isAirConditioned
		false,                // status
		nil,                  // description
		nil,                  // managementDepartment
		rand.Float64()*100,   // area
		nil,                  // tablesChairsType
	)

	is.operate.InitClassroom(
		"HXGCZX309",        // number
		"HXGC3#309",        // name
		"铁门关校区",            // campus
		"化学工程实训中心",         // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"QCGCZX3-303",          // number
		"QCGCZX3-303-低压电工室(1)", // name
		"铁门关校区",                // campus
		"汽车工程实训中心",             // building
		3,                      // floor
		"汽修实训室",                // classroomType
		nil,                    // classroomTag
		50,                     // capacity
		nil,                    // examinationRoomCapacity
		false,                  // examinationRoom
		false,                  // isMultimedia
		true,                   // isAirConditioned
		false,                  // status
		nil,                    // description
		nil,                    // managementDepartment
		rand.Float64()*100,     // area
		nil,                    // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL5-512",         // number
		"5-512",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL5-503",         // number
		"5#503",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JNL3#307-2号舞蹈室",   // number
		"3#307-2号舞蹈室",      // name
		"库尔勒校区",            // campus
		"技能训练中心",           // building
		3,                  // floor
		"舞蹈教室",             // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL1-xiaojifang",  // number
		"1#小机房",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		1,                  // floor
		"机房",               // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL4-410",         // number
		"4#410",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"XDNYZX211",        // number
		"XDNY2#211",        // name
		"铁门关校区",            // campus
		"现代农业实训中心",         // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		35,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL4-404",         // number
		"4#404",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL3-dajifang",    // number
		"3#大机房",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		3,                  // floor
		"机房",               // classroomType
		nil,                // classroomTag
		60,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL216",           // number
		"JXL2#216",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		2,                  // floor
		"机房",               // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		true,               // examinationRoom
		true,               // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL5#513",         // number
		"5#513",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"HXGCZX308",        // number
		"HXGC3#308",        // name
		"铁门关校区",            // campus
		"化学工程实训中心",         // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL1-dajifang",    // number
		"1#大机房",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		1,                  // floor
		"机房",               // classroomType
		nil,                // classroomTag
		60,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JDGCZX203",         // number
		"JDGC2#203-逆向建模实训室", // name
		"铁门关校区",             // campus
		"机电工程实训中心",          // building
		2,                   // floor
		"电工实训室",             // classroomType
		nil,                 // classroomTag
		50,                  // capacity
		nil,                 // examinationRoomCapacity
		false,               // examinationRoom
		false,               // isMultimedia
		true,                // isAirConditioned
		false,               // status
		nil,                 // description
		nil,                 // managementDepartment
		rand.Float64()*100,  // area
		nil,                 // tablesChairsType
	)

	is.operate.InitClassroom(
		"JNL4-演播厅",         // number
		"4#演播厅",            // name
		"库尔勒校区",            // campus
		"技能训练中心",           // building
		4,                  // floor
		"公共教室",             // classroomType
		nil,                // classroomTag
		280,                // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL401",           // number
		"JXL4#401",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JDGCZX108",        // number
		"1#108-机械加工中心",     // name
		"铁门关校区",            // campus
		"机电工程实训中心",         // building
		1,                  // floor
		"电工实训室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		false,              // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"SXL2#208",         // number
		"2#208",            // name
		"库尔勒校区",            // campus
		"实训楼",              // building
		2,                  // floor
		"移动一体机",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JNL1#-报告厅",        // number
		"1#报告厅",            // name
		"库尔勒校区",            // campus
		"技能训练中心",           // building
		1,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JNL2#206-现代电气控制室", // number
		"2#206现代电气控制",      // name
		"库尔勒校区",            // campus
		"技能训练中心",           // building
		2,                  // floor
		"现代电气控制实训室",        // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL4-413",         // number
		"4#413",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL2-207",         // number
		"2#207",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL3-307",         // number
		"3#307",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JYYSZX205",        // number
		"JYYS2#205-玩具实训室",  // name
		"铁门关校区",            // campus
		"教育艺术中心",           // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"XDNYZX208",        // number
		"XDNY2#208",        // name
		"铁门关校区",            // campus
		"现代农业实训中心",         // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		35,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JDGCZX313",        // number
		"JDGC3#313-传感器实训室", // name
		"铁门关校区",            // campus
		"机电工程实训中心",         // building
		3,                  // floor
		"电工实训室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL311",           // number
		"JXL3#311",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"SXL3#308",         // number
		"3#308",            // name
		"库尔勒校区",            // campus
		"实训楼",              // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL2-205",         // number
		"2#205",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"CJ1-qczhcj",       // number
		"CJ1-汽车综合车间",       // name
		"库尔勒校区",            // campus
		"车间",               // building
		1,                  // floor
		"实训车间",             // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"XDNYZX304",        // number
		"XDNY3#304",        // name
		"铁门关校区",            // campus
		"现代农业实训中心",         // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"SXL3#303",         // number
		"3#303",            // name
		"库尔勒校区",            // campus
		"实训楼",              // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL507",           // number
		"JXL5#507",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL3-304",         // number
		"3#304",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"XDNYZX105",        // number
		"XDNY1#105",        // name
		"铁门关校区",            // campus
		"现代农业实训中心",         // building
		1,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		35,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"SXL1#110",         // number
		"1#110",            // name
		"库尔勒校区",            // campus
		"实训楼",              // building
		1,                  // floor
		"移动一体机",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"QCGCZX3-302", // number
		"QCGCZX3-302-新能源动力电池实训室", // name
		"铁门关校区",            // campus
		"汽车工程实训中心",         // building
		3,                  // floor
		"汽修实训室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"SXL2#206",         // number
		"2#206",            // name
		"库尔勒校区",            // campus
		"实训楼",              // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL304",           // number
		"JXL3#304",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL1-xinlizhongxin", // number
		"1#心理活动中心",           // name
		"库尔勒校区",              // campus
		"综合教学楼",              // building
		1,                    // floor
		"活动室",                // classroomType
		nil,                  // classroomTag
		80,                   // capacity
		nil,                  // examinationRoomCapacity
		false,                // examinationRoom
		false,                // isMultimedia
		true,                 // isAirConditioned
		false,                // status
		nil,                  // description
		nil,                  // managementDepartment
		rand.Float64()*100,   // area
		nil,                  // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL322",           // number
		"JXL3#322",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JYYSZX313",           // number
		"JYYS3#313-电钢琴实训室（四）", // name
		"铁门关校区",               // campus
		"教育艺术中心",              // building
		3,                     // floor
		"琴房",                  // classroomType
		nil,                   // classroomTag
		50,                    // capacity
		nil,                   // examinationRoomCapacity
		false,                 // examinationRoom
		false,                 // isMultimedia
		true,                  // isAirConditioned
		false,                 // status
		nil,                   // description
		nil,                   // managementDepartment
		rand.Float64()*100,    // area
		nil,                   // tablesChairsType
	)

	is.operate.InitClassroom(
		"XXJSZX306",        // number
		"XXJS3#306教室",      // name
		"铁门关校区",            // campus
		"信息技术中心",           // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL204",           // number
		"JXL2#204",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"QCGCZX2-202", // number
		"QCGCZX2-202-汽车电器设备实训室", // name
		"铁门关校区",                 // campus
		"汽车工程实训中心",              // building
		2,                       // floor
		"汽修实训室",                 // classroomType
		nil,                     // classroomTag
		80,                      // capacity
		nil,                     // examinationRoomCapacity
		false,                   // examinationRoom
		false,                   // isMultimedia
		true,                    // isAirConditioned
		false,                   // status
		nil,                     // description
		nil,                     // managementDepartment
		rand.Float64()*100,      // area
		nil,                     // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL2-dajifang",    // number
		"2#大机房",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		2,                  // floor
		"机房",               // classroomType
		nil,                // classroomTag
		60,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JNL3#308-2号琴房（电钢）", // number
		"3#308-2号琴房（电钢）",    // name
		"库尔勒校区",             // campus
		"技能训练中心",            // building
		3,                   // floor
		"琴房",                // classroomType
		nil,                 // classroomTag
		80,                  // capacity
		nil,                 // examinationRoomCapacity
		false,               // examinationRoom
		false,               // isMultimedia
		true,                // isAirConditioned
		false,               // status
		nil,                 // description
		nil,                 // managementDepartment
		rand.Float64()*100,  // area
		nil,                 // tablesChairsType
	)

	is.operate.InitClassroom(
		"JYYSZX102",          // number
		"JYYS1#102-舞蹈实训室（二）", // name
		"铁门关校区",              // campus
		"教育艺术中心",             // building
		1,                    // floor
		"舞蹈教室",               // classroomType
		nil,                  // classroomTag
		50,                   // capacity
		nil,                  // examinationRoomCapacity
		false,                // examinationRoom
		false,                // isMultimedia
		true,                 // isAirConditioned
		false,                // status
		nil,                  // description
		nil,                  // managementDepartment
		rand.Float64()*100,   // area
		nil,                  // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL511",           // number
		"JXL5#511",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL5#510",         // number
		"5#510",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL5#509",         // number
		"5#509",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"SXL4#403",         // number
		"4#403",            // name
		"库尔勒校区",            // campus
		"实训楼",              // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL410",           // number
		"JXL4#410",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL4402",          // number
		"4#402",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"SXL2#204",         // number
		"2#204",            // name
		"库尔勒校区",            // campus
		"实训楼",              // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"QCGCZX2-204", // number
		"QCGCZX2-204-汽车发动机综合实训室", // name
		"铁门关校区",            // campus
		"汽车工程实训中心",         // building
		2,                  // floor
		"汽修实训室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"HXGCZX305",        // number
		"HXGC3#305",        // name
		"铁门关校区",            // campus
		"化学工程实训中心",         // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"XXJSZX305",        // number
		"XXJS3#305机房",      // name
		"铁门关校区",            // campus
		"信息技术中心",           // building
		3,                  // floor
		"机房",               // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL3-312",         // number
		"3#312",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL4-409",         // number
		"4#409",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"XDNYZX202",        // number
		"XDNY2#202",        // name
		"铁门关校区",            // campus
		"现代农业实训中心",         // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		35,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL2-209",         // number
		"2#209",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JNL3#303-2号声乐室",   // number
		"3#303-2号声乐室",      // name
		"库尔勒校区",            // campus
		"技能训练中心",           // building
		3,                  // floor
		"声乐教室",             // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JNL3#311-4号琴房（电钢）", // number
		"3#311-4号琴房（电钢）",    // name
		"库尔勒校区",             // campus
		"技能训练中心",            // building
		3,                   // floor
		"琴房",                // classroomType
		nil,                 // classroomTag
		80,                  // capacity
		nil,                 // examinationRoomCapacity
		false,               // examinationRoom
		false,               // isMultimedia
		true,                // isAirConditioned
		false,               // status
		nil,                 // description
		nil,                 // managementDepartment
		rand.Float64()*100,  // area
		nil,                 // tablesChairsType
	)

	is.operate.InitClassroom(
		"SXL3#307",         // number
		"3#307",            // name
		"库尔勒校区",            // campus
		"实训楼",              // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL406",           // number
		"JXL4#406",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL4-411",         // number
		"4#411",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JYYSZX303",           // number
		"JYYS3#303-电钢琴实训室（三）", // name
		"铁门关校区",               // campus
		"教育艺术中心",              // building
		3,                     // floor
		"琴房",                  // classroomType
		nil,                   // classroomTag
		50,                    // capacity
		nil,                   // examinationRoomCapacity
		false,                 // examinationRoom
		false,                 // isMultimedia
		true,                  // isAirConditioned
		false,                 // status
		nil,                   // description
		nil,                   // managementDepartment
		rand.Float64()*100,    // area
		nil,                   // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL5-513",         // number
		"5-513",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL211",           // number
		"JXL2#211",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		2,                  // floor
		"公共教室",             // classroomType
		nil,                // classroomTag
		60,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"CJ1-qcjdwxcj",     // number
		"CJ1-汽车机电维修车间",     // name
		"库尔勒校区",            // campus
		"车间",               // building
		1,                  // floor
		"实训车间",             // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL516",           // number
		"JXL5#516",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"XDNYZX207",        // number
		"XDNY2#207",        // name
		"铁门关校区",            // campus
		"现代农业实训中心",         // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		35,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL313",           // number
		"JXL3#313",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"QCGCZX2-203",      // number
		"QCGCZX2-203",      // name
		"铁门关校区",            // campus
		"汽车工程实训中心",         // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL3-310jifang",   // number
		"3#310机房",          // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		3,                  // floor
		"机房",               // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"SXL2-203",         // number
		"2#203",            // name
		"库尔勒校区",            // campus
		"实训楼",              // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL320",           // number
		"JXL3#320",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL509",           // number
		"JXL5#509",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"SXL1#103",         // number
		"1#103",            // name
		"库尔勒校区",            // campus
		"实训楼",              // building
		1,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL306",           // number
		"JXL3#306",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL3-313",         // number
		"3#313",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"QCGCZX2-201",         // number
		"QCGCZX2-201-汽车底盘实训室", // name
		"铁门关校区",               // campus
		"汽车工程实训中心",            // building
		2,                     // floor
		"汽修实训室",               // classroomType
		nil,                   // classroomTag
		80,                    // capacity
		nil,                   // examinationRoomCapacity
		false,                 // examinationRoom
		false,                 // isMultimedia
		true,                  // isAirConditioned
		false,                 // status
		nil,                   // description
		nil,                   // managementDepartment
		rand.Float64()*100,    // area
		nil,                   // tablesChairsType
	)

	is.operate.InitClassroom(
		"JYYSZX201",          // number
		"JYYS2#201-声乐实训室（一）", // name
		"铁门关校区",              // campus
		"教育艺术中心",             // building
		2,                    // floor
		"声乐教室",               // classroomType
		nil,                  // classroomTag
		50,                   // capacity
		nil,                  // examinationRoomCapacity
		false,                // examinationRoom
		false,                // isMultimedia
		true,                 // isAirConditioned
		false,                // status
		nil,                  // description
		nil,                  // managementDepartment
		rand.Float64()*100,   // area
		nil,                  // tablesChairsType
	)

	is.operate.InitClassroom(
		"JNL2#205-网络实训室",   // number
		"2#205-网络实训室",      // name
		"库尔勒校区",            // campus
		"技能训练中心",           // building
		2,                  // floor
		"网络实训室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL408",           // number
		"JXL4#408",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL5-505",         // number
		"5#505",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JDGCZX304",        // number
		"JDGC3#304",        // name
		"铁门关校区",            // campus
		"机电工程实训中心",         // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"XDNYZX102",        // number
		"XDNY1#102",        // name
		"铁门关校区",            // campus
		"现代农业实训中心",         // building
		1,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		35,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL407",           // number
		"JXL4#407",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL421",           // number
		"JXL4#421",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JDGCZX306",        // number
		"JDGC3#306",        // name
		"铁门关校区",            // campus
		"机电工程实训中心",         // building
		3,                  // floor
		"电工实训室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JDGCZX303",         // number
		"JDGC3#303-智能仿真实训室", // name
		"铁门关校区",             // campus
		"机电工程实训中心",          // building
		3,                   // floor
		"虚拟仿真实训室",           // classroomType
		nil,                 // classroomTag
		50,                  // capacity
		nil,                 // examinationRoomCapacity
		false,               // examinationRoom
		false,               // isMultimedia
		true,                // isAirConditioned
		false,               // status
		nil,                 // description
		nil,                 // managementDepartment
		rand.Float64()*100,  // area
		nil,                 // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL308",           // number
		"JXL3#308",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JDGCZX307",          // number
		"JDGC3#307-电气实训室（一）", // name
		"铁门关校区",              // campus
		"机电工程实训中心",           // building
		3,                    // floor
		"电工实训室",              // classroomType
		nil,                  // classroomTag
		50,                   // capacity
		nil,                  // examinationRoomCapacity
		false,                // examinationRoom
		false,                // isMultimedia
		true,                 // isAirConditioned
		false,                // status
		nil,                  // description
		nil,                  // managementDepartment
		rand.Float64()*100,   // area
		nil,                  // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL419",           // number
		"JXL4#419",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL315",           // number
		"JXL3#315",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL508",           // number
		"JXL5#508",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"SXL3#306",         // number
		"3#306",            // name
		"库尔勒校区",            // campus
		"实训楼",              // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL5-504",         // number
		"5#504",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"SXL3#301",         // number
		"3#301",            // name
		"库尔勒校区",            // campus
		"实训楼",              // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"QCGCZX3-308",      // number
		"QCGCZX3-308",      // name
		"铁门关校区",            // campus
		"汽车工程实训中心",         // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		100,                // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"SXL4#405",         // number
		"4#405",            // name
		"库尔勒校区",            // campus
		"实训楼",              // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL416",            // number
		"JXL4#416-乐理实训室（二）", // name
		"铁门关校区",             // campus
		"教学楼",               // building
		4,                   // floor
		"声乐教室",              // classroomType
		nil,                 // classroomTag
		50,                  // capacity
		nil,                 // examinationRoomCapacity
		false,               // examinationRoom
		false,               // isMultimedia
		true,                // isAirConditioned
		false,               // status
		nil,                 // description
		nil,                 // managementDepartment
		rand.Float64()*100,  // area
		nil,                 // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL3-305",         // number
		"3#305",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL2-lubojiaoshi", // number
		"2#录播教室",           // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JDGCZX310", // number
		"JDGC3#310-电气控制实训室（一）", // name
		"铁门关校区",                // campus
		"机电工程实训中心",             // building
		3,                      // floor
		"电工实训室",                // classroomType
		nil,                    // classroomTag
		50,                     // capacity
		nil,                    // examinationRoomCapacity
		false,                  // examinationRoom
		false,                  // isMultimedia
		true,                   // isAirConditioned
		false,                  // status
		nil,                    // description
		nil,                    // managementDepartment
		rand.Float64()*100,     // area
		nil,                    // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL2-210",         // number
		"2#210",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL224",           // number
		"JXL2#224",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		2,                  // floor
		"公共教室",             // classroomType
		nil,                // classroomTag
		60,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL505",           // number
		"JXL5#505",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL310",           // number
		"JXL3#310",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL5-506",         // number
		"5#506",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL309",           // number
		"JXL3#309",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL4-405",         // number
		"4#405",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"SXL1#107",         // number
		"1#107",            // name
		"库尔勒校区",            // campus
		"实训楼",              // building
		1,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JDGCZX210",         // number
		"JDGC2#210-钳工创新实训室", // name
		"铁门关校区",             // campus
		"机电工程实训中心",          // building
		2,                   // floor
		"实训车间",              // classroomType
		nil,                 // classroomTag
		50,                  // capacity
		nil,                 // examinationRoomCapacity
		false,               // examinationRoom
		false,               // isMultimedia
		true,                // isAirConditioned
		false,               // status
		nil,                 // description
		nil,                 // managementDepartment
		rand.Float64()*100,  // area
		nil,                 // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL3-306",         // number
		"3#306",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"QCGCZX3-306",            // number
		"QCGCZX3-306-虚拟仿真实训室(2)", // name
		"铁门关校区",                  // campus
		"汽车工程实训中心",               // building
		3,                        // floor
		"汽修实训室",                  // classroomType
		nil,                      // classroomTag
		50,                       // capacity
		nil,                      // examinationRoomCapacity
		false,                    // examinationRoom
		false,                    // isMultimedia
		true,                     // isAirConditioned
		false,                    // status
		nil,                      // description
		nil,                      // managementDepartment
		rand.Float64()*100,       // area
		nil,                      // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL5-507",         // number
		"5-507",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL2-212",         // number
		"2#212",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JNL2#201-图书室",     // number
		"2#201-图书室（2）",     // name
		"库尔勒校区",            // campus
		"技能训练中心",           // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL2-201",         // number
		"2#201",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL520",           // number
		"JXL5#520",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"SXL4#408",         // number
		"4#408",            // name
		"库尔勒校区",            // campus
		"实训楼",              // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL121",           // number
		"JXL1#121-阶梯教室",    // name
		"铁门关校区",            // campus
		"教学楼",              // building
		1,                  // floor
		"公共教室",             // classroomType
		nil,                // classroomTag
		215,                // capacity
		nil,                // examinationRoomCapacity
		true,               // examinationRoom
		true,               // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL307",           // number
		"JXL3#307",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL510",           // number
		"JXL5#510",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"QCGCZX3-307",      // number
		"QCGCZX3-307",      // name
		"铁门关校区",            // campus
		"汽车工程实训中心",         // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JNL3#310-婴幼儿卫生保健室", // number
		"3#310-婴幼儿卫生保健",     // name
		"库尔勒校区",             // campus
		"技能训练中心",            // building
		3,                   // floor
		"幼儿保健室",             // classroomType
		nil,                 // classroomTag
		80,                  // capacity
		nil,                 // examinationRoomCapacity
		false,               // examinationRoom
		false,               // isMultimedia
		true,                // isAirConditioned
		false,               // status
		nil,                 // description
		nil,                 // managementDepartment
		rand.Float64()*100,  // area
		nil,                 // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL501",           // number
		"JXL5#501",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL5-501",         // number
		"5#501",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JYYSZX207",         // number
		"JYYS2#207-幼儿游戏实训室", // name
		"铁门关校区",             // campus
		"教育艺术中心",            // building
		2,                   // floor
		"多媒体教室",             // classroomType
		nil,                 // classroomTag
		80,                  // capacity
		nil,                 // examinationRoomCapacity
		false,               // examinationRoom
		false,               // isMultimedia
		true,                // isAirConditioned
		false,               // status
		nil,                 // description
		nil,                 // managementDepartment
		rand.Float64()*100,  // area
		nil,                 // tablesChairsType
	)

	is.operate.InitClassroom(
		"JYYSZX202",          // number
		"JYYS2#202-声乐实训室（二）", // name
		"铁门关校区",              // campus
		"教育艺术中心",             // building
		2,                    // floor
		"声乐教室",               // classroomType
		nil,                  // classroomTag
		50,                   // capacity
		nil,                  // examinationRoomCapacity
		false,                // examinationRoom
		false,                // isMultimedia
		true,                 // isAirConditioned
		false,                // status
		nil,                  // description
		nil,                  // managementDepartment
		rand.Float64()*100,   // area
		nil,                  // tablesChairsType
	)

	is.operate.InitClassroom(
		"JNL2#202-图书室",     // number
		"2#202-图书室（1）",     // name
		"库尔勒校区",            // campus
		"技能训练中心",           // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"HXGCZX301",        // number
		"HXGC3#301",        // name
		"铁门关校区",            // campus
		"化学工程实训中心",         // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL5#512",         // number
		"5#512",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"HXGCZX311",        // number
		"HXGC3#311",        // name
		"铁门关校区",            // campus
		"化学工程实训中心",         // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL2-203",         // number
		"2#203",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"QCGCZX1-101",          // number
		"QCGCZX1-101-汽车综合维修车间", // name
		"铁门关校区",                // campus
		"汽车工程实训中心",             // building
		1,                      // floor
		"实训车间",                 // classroomType
		nil,                    // classroomTag
		150,                    // capacity
		nil,                    // examinationRoomCapacity
		false,                  // examinationRoom
		false,                  // isMultimedia
		true,                   // isAirConditioned
		false,                  // status
		nil,                    // description
		nil,                    // managementDepartment
		rand.Float64()*100,     // area
		nil,                    // tablesChairsType
	)

	is.operate.InitClassroom(
		"JYYSZX101",          // number
		"JYYS1#101-舞蹈实训室（一）", // name
		"铁门关校区",              // campus
		"教育艺术中心",             // building
		1,                    // floor
		"舞蹈教室",               // classroomType
		nil,                  // classroomTag
		50,                   // capacity
		nil,                  // examinationRoomCapacity
		false,                // examinationRoom
		false,                // isMultimedia
		true,                 // isAirConditioned
		false,                // status
		nil,                  // description
		nil,                  // managementDepartment
		rand.Float64()*100,   // area
		nil,                  // tablesChairsType
	)

	is.operate.InitClassroom(
		"JDGCZX205",        // number
		"JDGC2#205",        // name
		"铁门关校区",            // campus
		"机电工程实训中心",         // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL215",           // number
		"JXL2#215",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		2,                  // floor
		"公共教室",             // classroomType
		nil,                // classroomTag
		60,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"HXGCZX105",            // number
		"HXGC1#105-化工仿真实训室(一)", // name
		"铁门关校区",                // campus
		"化学工程实训中心",             // building
		1,                      // floor
		"化工仿真实训室",              // classroomType
		nil,                    // classroomTag
		40,                     // capacity
		nil,                    // examinationRoomCapacity
		false,                  // examinationRoom
		false,                  // isMultimedia
		true,                   // isAirConditioned
		false,                  // status
		nil,                    // description
		nil,                    // managementDepartment
		rand.Float64()*100,     // area
		nil,                    // tablesChairsType
	)

	is.operate.InitClassroom(
		"SXL2#205",         // number
		"2#205",            // name
		"库尔勒校区",            // campus
		"实训楼",              // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL5-502",         // number
		"5#502",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"SXL2#207",         // number
		"2#207",            // name
		"库尔勒校区",            // campus
		"实训楼",              // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"HXGCZX304",        // number
		"HXGC3#304",        // name
		"铁门关校区",            // campus
		"化学工程实训中心",         // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL213",            // number
		"JXL2#213-美术实训室（二）", // name
		"铁门关校区",             // campus
		"教学楼",               // building
		2,                   // floor
		"美术教室",              // classroomType
		nil,                 // classroomTag
		50,                  // capacity
		nil,                 // examinationRoomCapacity
		false,               // examinationRoom
		false,               // isMultimedia
		true,                // isAirConditioned
		false,               // status
		nil,                 // description
		nil,                 // managementDepartment
		rand.Float64()*100,  // area
		nil,                 // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL3-311",         // number
		"3#311",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JYYSZX301",           // number
		"JYYS3#301-电钢琴实训室（一）", // name
		"铁门关校区",               // campus
		"教育艺术中心",              // building
		3,                     // floor
		"琴房",                  // classroomType
		nil,                   // classroomTag
		50,                    // capacity
		nil,                   // examinationRoomCapacity
		false,                 // examinationRoom
		false,                 // isMultimedia
		true,                  // isAirConditioned
		false,                 // status
		nil,                   // description
		nil,                   // managementDepartment
		rand.Float64()*100,    // area
		nil,                   // tablesChairsType
	)

	is.operate.InitClassroom(
		"XDNYZX205",        // number
		"XDNY2#205",        // name
		"铁门关校区",            // campus
		"现代农业实训中心",         // building
		2,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		35,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL305",           // number
		"JXL3#305",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JNL2#204-自动化生产线室", // number
		"2#204-自动化生产线",     // name
		"库尔勒校区",            // campus
		"技能训练中心",           // building
		2,                  // floor
		"自动化生产线实训室",        // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL4-408",         // number
		"4#408",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL3-303",         // number
		"3#303",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"XDNYZX305",        // number
		"XDNY3#305",        // name
		"铁门关校区",            // campus
		"现代农业实训中心",         // building
		3,                  // floor
		"机房",               // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"CJ1-kqf",          // number
		"CJ1-烤漆房",          // name
		"库尔勒校区",            // campus
		"车间",               // building
		1,                  // floor
		"实训车间",             // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL1-jietijiaoshi", // number
		"1#阶梯教室",            // name
		"库尔勒校区",             // campus
		"综合教学楼",             // building
		1,                   // floor
		"多媒体教室",             // classroomType
		nil,                 // classroomTag
		200,                 // capacity
		nil,                 // examinationRoomCapacity
		false,               // examinationRoom
		false,               // isMultimedia
		true,                // isAirConditioned
		false,               // status
		nil,                 // description
		nil,                 // managementDepartment
		rand.Float64()*100,  // area
		nil,                 // tablesChairsType
	)

	is.operate.InitClassroom(
		"JDGCZX221",        // number
		"JDGC2#206-液压实训室",  // name
		"铁门关校区",            // campus
		"机电工程实训中心",         // building
		2,                  // floor
		"液压气动实训室",          // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"QCGCZX3-305",            // number
		"QCGCZX3-305-虚拟仿真实训室(1)", // name
		"铁门关校区",                  // campus
		"汽车工程实训中心",               // building
		3,                        // floor
		"汽修实训室",                  // classroomType
		nil,                      // classroomTag
		50,                       // capacity
		nil,                      // examinationRoomCapacity
		false,                    // examinationRoom
		false,                    // isMultimedia
		true,                     // isAirConditioned
		false,                    // status
		nil,                      // description
		nil,                      // managementDepartment
		rand.Float64()*100,       // area
		nil,                      // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL4-407",         // number
		"4#407",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JYYSZX208",           // number
		"JYYS2#208-早期教育智慧实训室", // name
		"铁门关校区",               // campus
		"教育艺术中心",              // building
		2,                     // floor
		"多媒体教室",               // classroomType
		nil,                   // classroomTag
		100,                   // capacity
		nil,                   // examinationRoomCapacity
		false,                 // examinationRoom
		false,                 // isMultimedia
		true,                  // isAirConditioned
		false,                 // status
		nil,                   // description
		nil,                   // managementDepartment
		rand.Float64()*100,    // area
		nil,                   // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL318",           // number
		"JXL3#318",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JNL2#207-机械装调室",   // number
		"2#207-机械装调",       // name
		"库尔勒校区",            // campus
		"技能训练中心",           // building
		2,                  // floor
		"机械装调实训室",          // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JNL2#203-综合布线室",   // number
		"2#203-综合布线室",      // name
		"库尔勒校区",            // campus
		"技能训练中心",           // building
		2,                  // floor
		"网络综合布线室",          // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL5-510",         // number
		"5-510",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		5,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL203",           // number
		"JXL2#203",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		2,                  // floor
		"公共教室",             // classroomType
		nil,                // classroomTag
		60,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JYYSZX103",          // number
		"JYYS1#103-舞蹈实训室（三）", // name
		"铁门关校区",              // campus
		"教育艺术中心",             // building
		1,                    // floor
		"舞蹈教室",               // classroomType
		nil,                  // classroomTag
		50,                   // capacity
		nil,                  // examinationRoomCapacity
		false,                // examinationRoom
		false,                // isMultimedia
		true,                 // isAirConditioned
		false,                // status
		nil,                  // description
		nil,                  // managementDepartment
		rand.Float64()*100,   // area
		nil,                  // tablesChairsType
	)

	is.operate.InitClassroom(
		"ZHL4-412",         // number
		"4#412",            // name
		"库尔勒校区",            // campus
		"综合教学楼",            // building
		4,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL218",           // number
		"JXL2#218",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		2,                  // floor
		"机房",               // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		true,               // examinationRoom
		true,               // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JNL3#301-3号声乐室",   // number
		"3#301-3号声乐室",      // name
		"库尔勒校区",            // campus
		"技能训练中心",           // building
		3,                  // floor
		"声乐教室",             // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"CJ1-bjcj",         // number
		"CJ1-钣金车间",         // name
		"库尔勒校区",            // campus
		"车间",               // building
		1,                  // floor
		"实训车间",             // classroomType
		nil,                // classroomTag
		80,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"JXL317",           // number
		"JXL3#317",         // name
		"铁门关校区",            // campus
		"教学楼",              // building
		3,                  // floor
		"多媒体教室",            // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"XDNYZX106",        // number
		"XDNY1#106-农业实训车间", // name
		"铁门关校区",            // campus
		"现代农业实训中心",         // building
		1,                  // floor
		"实训车间",             // classroomType
		nil,                // classroomTag
		120,                // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)

	is.operate.InitClassroom(
		"HXGCZX107",           // number
		"HXGC1#107-化工实训车间(一)", // name
		"铁门关校区",               // campus
		"化学工程实训中心",            // building
		1,                     // floor
		"实训车间",                // classroomType
		nil,                   // classroomTag
		120,                   // capacity
		nil,                   // examinationRoomCapacity
		false,                 // examinationRoom
		false,                 // isMultimedia
		true,                  // isAirConditioned
		false,                 // status
		nil,                   // description
		nil,                   // managementDepartment
		rand.Float64()*100,    // area
		nil,                   // tablesChairsType
	)

	is.operate.InitClassroom(
		"XDNYZX303",        // number
		"XDNY3#303-会计实训室",  // name
		"铁门关校区",            // campus
		"现代农业实训中心",         // building
		3,                  // floor
		"机房",               // classroomType
		nil,                // classroomTag
		50,                 // capacity
		nil,                // examinationRoomCapacity
		false,              // examinationRoom
		false,              // isMultimedia
		true,               // isAirConditioned
		false,              // status
		nil,                // description
		nil,                // managementDepartment
		rand.Float64()*100, // area
		nil,                // tablesChairsType
	)
}
