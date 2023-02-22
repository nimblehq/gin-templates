#!/bin/sh

# go tool cover -func coverage/coverage.out results in the multiline text telling coverage percentage on each function in the following format
#   github.com/nimblehq/gulf-approval-web/helpers/config.go:8:                      GetConfigPrefix    100.0%
#   github.com/nimblehq/gulf-approval-web/helpers/config.go:28:                     GetStringConfig    100.0%
#   github.com/nimblehq/gulf-approval-web/lib/api/v1/controllers/health.go:13:      HealthStatus       100.0%
#   total:                                                                          (statements)       100.0%
# grep total to get the line start with `total` which contain the overall coverage percentage
# awk '{print substr($3, 1, length($3)-1)}' with the built-in variable `$3` to grab the 3rd part after the line is splited with space and substr to remove the %
# which will result in 100.0
coverage=$(go tool cover -func coverage/coverage.out | grep total | awk '{print substr($3, 1, length($3)-1)}')
expected=100

if [ "$(echo $coverage '<' $expected | bc -l)" -eq 1 ]; then\
  echo "coverage percentage is too low ${coverage}%, the expected percentage is ${expected}%"
  exit 1
else\
  echo "coverage percentage meet expectation ${coverage}%"
  exit 0
fi;
