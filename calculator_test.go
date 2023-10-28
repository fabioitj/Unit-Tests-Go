package main;

import (
	"testing"
);

func TestShouldSumCorrectly(t *testing.T) {
	sum := Sum(1, 2, 3) // arrange
	result := 6 // act
	if sum != result { // assert
		t.Error("Valor Esperado:", result, "Valor Recebido:", sum);
	}
}

func TestShouldSumIncorrectly(t *testing.T) {
	sum := Sum(1, 2, 3) // arrange
	result := 7 // act
	if sum != result { // assert
		t.Error("Valor Esperado:", result, "Valor Recebido:", sum);
	}
}

func TestShouldSubtractCorrectly(t *testing.T) {
	subtract := Subtract(1, 2, 3) // arrange
	result := -6 // act
	if subtract != result { // assert
		t.Error("Valor Esperado:", result, "Valor Recebido:", subtract);
	}
}

func TestShouldSubtractIncorrectly(t *testing.T) {
	subtract := Subtract(1, 2, 3) // arrange
	result := -7 // act
	if subtract != result { // assert
		t.Error("Valor Esperado:", result, "Valor Recebido:", subtract);
	}
}

func TestShouldMultiplyCorrectly(t *testing.T) {
	multiply := Multiply(1, 2, 3) // arrange
	result := 6 // act
	if multiply != result { // assert
		t.Error("Valor Esperado:", result, "Valor Recebido:", multiply);
	}
}

func TestShouldMultiplyIncorrectly(t *testing.T) {
	multiply := Multiply(1, 2, 3) // arrange
	result := 8 // act
	if multiply != result { // assert
		t.Error("Valor Esperado:", result, "Valor Recebido:", multiply);
	}
}

func TestShouldDivideCorrectly(t *testing.T) {
	divide := Divide(10, 2) // arrange
	result := 5.0 // act
	if divide != result { // assert
		t.Error("Valor Esperado:", result, "Valor Recebido:", divide);
	}
}

func TestShouldDivideIncorrectly(t *testing.T) {
	divide := Divide(10, 2) // arrange
	result := 6.0 // act
	if divide != result { // assert
		t.Error("Valor Esperado:", result, "Valor Recebido:", divide);
	}
}