#!/usr/bin/env bash

sed -E 's/Buckets/IntBuckets/g; s/Select/IntSelect/g; s/sort.Interface/IntSorter/g' sorter.go > sortInt.go
sed -i '/sort/d' sortInt.go