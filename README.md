P.S. My first Project made and will be only maintained

Library Project
Version
Current project version: v1.0.1

Description
This project is a set of libraries for working with numbers, strings, and slices in the Go programming language. The project implements various functions for performing mathematical operations, sorting, and working with strings.

Project Structure
main.go: The main file that runs examples of using the libraries.

Installation
 
Clone the repository:
git clone github.com/AlexeiSvt/library

Navigate to the project directory
cd library

Usage
Run the examples by executing the command:
go run main.go

# Package Descriptions

## numbers

- `Add`: Adds two numbers.
- `Sub`: Subtracts one number from another.
- `Mult`: Multiplies two numbers.
- `Div`: Divides one number by another.
- `Avg`: Calculates the average of two numbers.
- `Min`: Finds the minimum of two numbers.
- `Max`: Finds the maximum of two numbers.
- `Abs`: Calculates the absolute value of a number.
- `Fact`: Calculates the factorial of a number.
- `IsEven`: Checks if a number is even.
- `IsOdd`: Checks if a number is odd.

## strings

- `Concat`: Concatenates two strings.
- `Contains`: Checks if a string contains a substring.
- `Count`: Counts the number of occurrences of a substring in a string.
- `Reverse`: Reverses a string.
- `ToUpper`: Converts a string to uppercase.
- `ToLower`: Converts a string to lowercase.
- `IsPalindrome`: Checks if a string is a palindrome.

## slices

- `SumSlice`: Calculates the sum of all elements in a slice.
- `MinSlice`: Finds the minimum element in a slice.
- `MaxSlice`: Finds the maximum element in a slice.
- `AvgSlice`: Calculates the average of all elements in a slice.

- `BubbleSort`: Sorts a slice using the bubble sort algorithm.
- `InsertionSort`: Sorts a slice using the insertion sort algorithm.
- `SelectionSort`: Sorts a slice using the selection sort algorithm.
- `MergeSort`: Sorts a slice using the merge sort algorithm.

## interfaces

The `interfaces` package contains interfaces and type constraints:

- `Number`: A type constraint that allows any numeric type.
- `IntUint`: A type constraint that allows any integer or unsigned integer type.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
