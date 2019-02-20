#!/usr/bin/env bash
sed -E 's/Buckets/IntBuckets/g; s/Select/IntSelect/g; s/sort.Interface/nthElementUtils.IntSorter/g; s/_ "github.com\/furstenheim\/nth_element\/utils"/"github.com\/furstenheim\/nth_element\/utils"/g' FloydRivest/sort.go > FloydRivest/intSort.go
sed -i '/sort/d' FloydRivest/intSort.go