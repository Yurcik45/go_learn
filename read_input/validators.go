package main

import "strings"

func ValidateName(s string) bool {
  return len(s) > 2
}

func ValidateEmail(s string) bool {
  condition := strings.Contains(s, "@") &&
               strings.Contains(s, ".")

  return condition
}

func ValidateAge(i int) bool {
  return i >= 18
}

