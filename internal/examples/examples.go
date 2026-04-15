package examples

import (
	"fmt"
	example"Library/internal"
)

func RunExamples() {
    // Examples of functions from the numbers package
    fmt.Println("Numbers package examples:")
    sum, err := example.Add(1, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Sum:", sum)
    }

    sub, err := example.Sub(5, 3)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Subtraction:", sub)
    }

    mult, err := example.Mult(2, 3)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Multiplication:", mult)
    }

    div, err := example.Div(4, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Division:", div)
    }

    avg, err := example.Avg(4, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Average:", avg)
    }

    min, err := example.Min(4, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Minimum:", min)
    }

    max, err := example.Max(4, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Maximum:", max)
    }

    abs, err := example.Abs(-4)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Absolute value:", abs)
    }

    fact, err := example.Fact(5)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Factorial:", fact)
    }

    isEven, err := example.IsEven(4)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Is even:", isEven)
    }

    isOdd, err := example.IsOdd(3)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Is odd:", isOdd)
    }

    // Examples of functions from the strings package
    fmt.Println("\nStrings package examples:")
    concat, err := example.Concat("Hello, ", "World!")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Concatenated string:", concat)
    }

    contains, err := example.Contains("Hello, World!", "World")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Contains substring:", contains)
    }

    count, err := example.Count("Hello, World!", "o")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Count of substring:", count)
    }

    reversed, err := example.Reverse("Hello, World!")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Reversed string:", reversed)
    }

    upper, err := example.ToUpper("Hello, World!")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Uppercase string:", upper)
    }

    lower, err := example.ToLower("Hello, World!")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Lowercase string:", lower)
    }

    isPalindrome, err := example.IsPalindrome("madam")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Is palindrome:", isPalindrome)
    }

    isPalindrome, err = example.IsPalindrome("hello")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Is palindrome:", isPalindrome)
    }

    // Examples of functions from the slices package
    fmt.Println("\nSlices package examples:")
    intSlice := []int{5, 2, 9, 1, 5, 6}
    fmt.Println("Original int slice:", intSlice)

    example.BubbleSort(intSlice)
    fmt.Println("Bubble sorted int slice:", intSlice)

    intSlice = []int{5, 2, 9, 1, 5, 6}
    example.InsertionSort(intSlice)
    fmt.Println("Insertion sorted int slice:", intSlice)

    intSlice = []int{5, 2, 9, 1, 5, 6}
    example.SelectionSort(intSlice)
    fmt.Println("Selection sorted int slice:", intSlice)

    intSlice = []int{5, 2, 9, 1, 5, 6}
    sortedIntSlice := example.MergeSort(intSlice)
    fmt.Println("Merge sorted int slice:", sortedIntSlice)

    sumSlice, err := example.SumSlice(intSlice)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Sum of int slice:", sumSlice)
    }

    minSlice, err := example.MinSlice(intSlice)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Min of int slice:", minSlice)
    }

    maxSlice, err := example.MaxSlice(intSlice)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Max of int slice:", maxSlice)
    }

    avgSlice, err := example.AvgSlice(intSlice)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Avg of int slice:", avgSlice)
    }
}