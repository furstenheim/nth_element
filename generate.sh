#!/usr/bin/env bash
sed -E 's/Buckets/IntBuckets/g; s/Select/IntSelect/g; s/sort.Interface/nthElementUtils.IntSorter/g; s/_ "github.com\/furstenheim\/nth_element\/utils"/"github.com\/furstenheim\/nth_element\/utils"/g' FloydRivest/sort.go > FloydRivest/intSort.go
sed -i '/sort/d' FloydRivest/intSort.go

sed -E 's/Buckets/IntBuckets/g; s/partition/Intpartition/g; s/Select\(/IntSelect\(/g; s/sort.Interface/nthElementUtils.IntSorter/g; s/_ "github.com\/furstenheim\/nth_element\/utils"/"github.com\/furstenheim\/nth_element\/utils"/g' QuickSelect/sort.go > QuickSelect/intSort.go
sed -i '/sort/d' QuickSelect/intSort.go


sed -E 's/Buckets/IntBuckets/g;
s/Select\(/IntSelect\(/g;
s/sort.Interface/nthElementUtils.IntSorter/g;
s/_select/int__select/g;
s/pivotFunc/int_pivotFunc/g;
s/insertionSortMedian/int_insertionSortMedian/g;
s/partition/int_partition/g;
s/_ "github.com\/furstenheim\/nth_element\/utils"/"github.com\/furstenheim\/nth_element\/utils"/g' MediansOfMedians/sort.go > MediansOfMedians/intSort.go
sed -i '/sort/d' MediansOfMedians/intSort.go