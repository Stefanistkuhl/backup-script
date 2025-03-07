#!/bin/bash
set -a
source .env
set +a

cd $SRC_DIR

for dir in */; do
  if [[ -d "$dir" ]]; then
    echo "Compressing directory: ${dir%/}"
    tar -czvf "$COMPRESSED_DIRS_DIR/${dir%/}.tar.gz" "$dir"
  fi
done

rsync -avz -e "ssh -p $SSH_PORT" $COMPRESSED_DIRS_DIR/* $SSH_USER@$SSH_HOST:$DST_DIR
