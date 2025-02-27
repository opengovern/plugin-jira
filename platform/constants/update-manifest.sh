#!/bin/bash

file_path="manifest.yaml"

# Append multiple lines to the file
cat <<EOF >> "$file_path"

DescriberTag: local-$TAG
UpdateDate: $(date +%Y-%m-%d) $(date +%H:%M:%S)

EOF

echo "Data has been appended to $file_path"