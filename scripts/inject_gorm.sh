#!/bin/bash

for file in ./src/model/*
do
    ./scripts/gorm.sh "$file"
done