#!/bin/bash
printf "\nRegenerating gqlgen files\n"

while [[ "$#" -gt 0 ]]; do case $1 in
  -r|--resolvers)
    rm -f internal/gql/resolvers/generated/resolver.go
  ;;
  *) echo "Unknown parameter passed: $1"; exit 1;;
esac; shift; done

time go run -v github.com/99designs/gqlgen
printf "\nDone.\n\n"