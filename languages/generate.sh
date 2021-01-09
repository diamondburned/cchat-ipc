#!/usr/bin/env bash
declare -A languages
languages=(
	[go]="go"
	[ts]="ts"

	# Currently segfaults.
	# [rs]="rust"
)

for dir in "${!languages[@]}"; {
	rm -rf   "./$dir"
	flatc -o "./$dir" "--${languages[$dir]}" ../schema/*.fbs || {
		echo "Failed to generate $dir; removing."
		rm -rf "./$dir"
	}
}
