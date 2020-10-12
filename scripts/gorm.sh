#!/bin/bash

g () {
  sed "s/json:\"$1,omitempty\"/json:\"$1,omitempty\" gorm:\"$2\"/"
}

cat $1 \
| g "id" "primary_key" \
| g "name" "varchar(100)" \
> $1.tmp && mv $1{.tmp,}
