#!/bin/sh
# scripts/completions.sh
set -e
rm -rf completions
mkdir completions

for sh in bash zsh fish; do
	./bin/gmd completion "$sh" >"completions/gmd.$sh"
done