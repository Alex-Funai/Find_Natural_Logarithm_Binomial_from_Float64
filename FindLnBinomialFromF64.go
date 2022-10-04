package main

import (
	"fmt"
	"math"
)

// ////////////// ENTER TARGET VALUE HERE ////////////////
const target float64 = 9.170558458320164 // <--------
// const target float64 = 4.829441541679835928251696364374 // <--------

func main() {

	//////////////// CREATE ARRAY RANGES HERE ////////////////
	// 	Recommendations: ln = 10  |  rn = 1000-50000
	var ln = createNaturalLogarithmArray(10)
	var rn = createDecimalRemainderArray(1125)

	//////////////// TEST ARRAY RANGES HERE //////////////////
	// 	(Optional: for visualizing ranges you'll compare)
	printArrayF64(ln)
	//printArrayF64(rn)

	//////////////// EXECUTE TO FIND RESULTS /////////////////
	//	findTarget (target, logarithmArray, remainderArray )
	//	@Param	target float64
	//			-- the decimal value to find a nomial logarithmic expression for
	//	@Param	logarithmArray[] float64
	//			-- the logarithmic array holding ln(0):ln(10)
	//	@Param	remainderArray[] float64
	//			-- the array of decimal values to find remainder values to ln()
	findTarget(target, ln, rn)

	//////////////// DEPRECATIONS ///////////////////
	//	(Deprecated: conversion occurs in findTarget
	//  decimalToFraction(rn[1125])
}

/**
 *	createDecimalRemainderArray()
 *	@Param	size int
 *			-- int specifying size/range of decimals to find remainder of results
 *	@ImpSpec	recommend: size int = 1000:50000
 *	@Return	[]float64
 *			-- array of float64 values %.2f
 */
func createDecimalRemainderArray(size int) []float64 {
	slice := make([]float64, size+1)
	for i := 0; i <= size; i++ {
		slice[i] = float64(i) * .01
	}
	return slice
}

/**
 *	createNaturalLogarithmArray()
 *	@Param size int
 *			-- int specifying size/range of natural logarithmic values.
 *	@ImpSpec	recommend: size int = 10
 *	@Return []float64
 *			-- array of float64 values representing natural logarithmic values.
 */
func createNaturalLogarithmArray(size int) []float64 {
	slice := make([]float64, size)
	for i := 0; i <= size-1; i++ {
		slice[i] = math.Log(float64(i))
	}
	return slice
}

/**
 *	findTarget()
 *	@Param target float64
 *		-- input float64 to searech for a ln binomial.
 *	@Param logarithmArray[] float64
 *		-- input logarithm array to compute against decimal for a ln binomial.
 *	@Param decimalArray[] float64
 *		== input decimalArray to compute against logarithms for a ln binomial.
 */
func findTarget(target float64, logarithmArray, decimalArray []float64) {
	fmt.Println("::::: COMPARING ARRAYS FOR TARGET :::::")
	for i, v := range decimalArray {
		fmt.Println("--- CURRENT ---")
		fmt.Println("DECIMAL |", "Index:", i, "|", "Value:", v)
		for j, k := range logarithmArray {
			fmt.Println("Lns", "|", "ln(", j, ")", "|", "Value:", k,
				"\nk+v:", k+v, "\nk-v:", k-v, "\nv-k:", v-k)
			//fmt.Println(j, k)
			//fmt.Println(v * k)
			if -v+k == target {
				fmt.Println(
					"\n::::: Target Found! :::::",
					"\nTarget:", target,
					"\nEquation:", "-", v, "+", k,
					"\nSimplified:", "ln(", j, ")", "-", i, "/", "100",
					//"\nln(<value>):", "ln(", j, ")",
					//"\nFraction:", i, "/", "100"
				)
				return
			} else if v-k == target {
				fmt.Println(
					"\n::::: Target Found! :::::",
					"\nTarget:", target,
					"\nEquation:", v, "-", k,
					"\nSimplified:", i, "/", "100", "-", "ln(", j, ")",
					//"\nln(<value>):", "ln(", j, ")",
					//"\nFraction:", i, "/", "100"
				)
				return
			} else if v+k == target {
				fmt.Println(
					"\n::::: Target Found! :::::",
					"\nTarget:", target,
					"\nEquation:", v, "+", k,
					"\nSimplified:", i, "/", "100", "+", "ln(", j, ")",
					//"\nln(<value>):", "ln(", j, ")",
					//"\nFraction:", i, "/", "100"
				)
				return
			}
		}
	}
}

/**
 *	printArrayF64()
 *	@Param array[] float64
 *		-- an array of float64 values to print.
 */
func printArrayF64(array []float64) {
	println(array)
	for i, v := range array {
		fmt.Println("Index:", i, "| Value:", v)
	}
}

/** @DEPRECATED
 *	f64to32()
 *	@Param f64 float64
 *		-- a float64 to be converted into float32.
 *	@Return float32
		-- a float32 value that was converted
*/ /*
func f64tof32(f64 float64) float32 {
	var f32 = float32(f64)
	return f32
}
*/
/**	@DEPRECATED
 *	decimalToFraction()
 *	@Param decimal float64
 *		-- a decimal value to be printed as a fraction.
 */ /*
func decimalToFraction(decimal float64) {
	var numerator, denominator int
	numerator = int(decimal * 100)
	denominator = 100
	println("\n\t::FRACTION::\nNumerator:", numerator, "|", "Denominator:", denominator,
		"|", "Fraction:", numerator, "/", denominator)
}
*/
