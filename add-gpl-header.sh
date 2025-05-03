#!/bin/bash

# Directory to start searching from (default to current directory)
DIR="${1:-.}"

# Define the GPLv3 header content
GPL_HEADER="Macrame is a program that enables the user to create keyboard macros and button panels. 
The macros are saved as simple JSON files and can be linked to the button panels. The panels can 
be created with HTML and CSS.

Copyright (C) 2025 Jesse Malotaux

This program is free software: you can redistribute it and/or modify 
it under the terms of the GNU General Public License as published by 
the Free Software Foundation, either version 3 of the License, or 
(at your option) any later version.

This program is distributed in the hope that it will be useful, 
but WITHOUT ANY WARRANTY; without even the implied warranty of 
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the 
GNU General Public License for more details.

You should have received a copy of the GNU General Public License 
along with this program. If not, see <https://www.gnu.org/licenses/>."

# Loop through all files in the directory (and subdirectories), excluding node_modules
find "$DIR" \( -iname \*.go -o \( -path "$DIR/fe/src/*.js" -o -path "$DIR/fe/src/*.vue" -o -path "$DIR/fe/src/*.css" \) \) ! -path "*/node_modules/*" | while read file
do
    # Check if the file already has a GPL header
    if ! grep -q "Copyright (C) 2025 Jesse Malotaux" "$file"; then
        # Check if it's a Vue file and handle it carefully
        if [[ "$file" == *.vue ]]; then
            echo "Adding GPL header to: $file"
            # Prepend the GPL header to Vue files as raw text
            # Make sure we don't add comment marks that break Vue syntax
            echo -e "<!--\n$GPL_HEADER\n-->\n\n$(cat "$file")" > "$file"
        else
            echo "Adding GPL header to: $file"
            # Prepend the GPL header to other files (go, js, ts, etc.) as comments
            echo -e "/*\n$GPL_HEADER\n*/\n\n$(cat "$file")" > "$file"
        fi
    else
        echo "GPL header already present in: $file"
    fi
done

echo "GPL header addition complete!"
