#!/usr/bin/env bash

sed -E 's/FloydRivestBuckets/IntFloydRivestBuckets/g; s/FloydRivestSelect/IntFloydRivestSelect/g; s/sort.Interface/IntSorter/g' floydRivest.go > intFloydRivest.go
sed -i '/sort/d' intFloydRivest.go