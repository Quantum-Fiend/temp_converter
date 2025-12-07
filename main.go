package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)
	
	fmt.Println("🔥 Go Wasm Temperature Converter Initialized")
	
	// Wait for DOM to be ready
	js.Global().Get("document").Call("addEventListener", "DOMContentLoaded", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		setupUI()
		return nil
	}))
	
	<-c
}

func setupUI() {
	doc := js.Global().Get("document")
	
	// Get elements
	convertBtn := doc.Call("getElementById", "convert-btn")
	tempInput := doc.Call("getElementById", "temperature-input")
	conversionType := doc.Call("getElementById", "conversion-type")
	resultDiv := doc.Call("getElementById", "result")
	
	// Set initial state
	resultDiv.Set("innerText", "")
	resultDiv.Get("classList").Call("remove", "error")
	
	// Add click handler
	convertBtn.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		handleConversion(tempInput, conversionType, resultDiv)
		return nil
	}))
	
	// Add enter key handler
	tempInput.Call("addEventListener", "keypress", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if args[0].Get("key").String() == "Enter" {
			handleConversion(tempInput, conversionType, resultDiv)
		}
		return nil
	}))
}

func handleConversion(tempInput, conversionType, resultDiv js.Value) {
	inputVal := tempInput.Get("value").String()
	convType := conversionType.Get("value").String()
	
	// Clear error state
	resultDiv.Get("classList").Call("remove", "error")
	
	// Validate input
	if inputVal == "" {
		resultDiv.Set("innerText", "⚠️ Please enter a temperature")
		resultDiv.Get("classList").Call("add", "error")
		return
	}
	
	temp, err := strconv.ParseFloat(inputVal, 64)
	if err != nil {
		resultDiv.Set("innerText", "❌ Invalid number")
		resultDiv.Get("classList").Call("add", "error")
		return
	}
	
	var result float64
	var resultText string
	
	if convType == "c-to-f" {
		result = celsiusToFahrenheit(temp)
		resultText = fmt.Sprintf("%.1f°C = %.1f°F", temp, result)
	} else {
		result = fahrenheitToCelsius(temp)
		resultText = fmt.Sprintf("%.1f°F = %.1f°C", temp, result)
	}
	
	resultDiv.Set("innerText", resultText)
}

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}

func fahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}
