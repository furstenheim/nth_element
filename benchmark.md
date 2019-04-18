=== RUN   TestSelect_BigArray
--- PASS: TestSelect_BigArray (2.31s)
=== RUN   TestSelect_RandomArray
--- FAIL: TestSelect_RandomArray (1.02s)
    select_test.go:70: 
        	Error Trace:	select_test.go:70
        	Error:      	"%!s(int=5099083273840999716)" is not less or equal than "%!s(int=5098757367429803214)"
        	Test:       	TestSelect_RandomArray
    select_test.go:70: 
        	Error Trace:	select_test.go:70
        	Error:      	"%!s(int=5099083273840999716)" is not less or equal than "%!s(int=5098771971608903699)"
        	Test:       	TestSelect_RandomArray
    select_test.go:70: 
        	Error Trace:	select_test.go:70
        	Error:      	"%!s(int=5099083273840999716)" is not less or equal than "%!s(int=4979504428492036838)"
        	Test:       	TestSelect_RandomArray
    select_test.go:70: 
        	Error Trace:	select_test.go:70
        	Error:      	"%!s(int=5099083273840999716)" is not less or equal than "%!s(int=5039543713335519498)"
        	Test:       	TestSelect_RandomArray
=== RUN   TestBucketsSize1
--- PASS: TestBucketsSize1 (0.00s)
=== RUN   TestBucketsBig
--- PASS: TestBucketsBig (0.52s)
=== RUN   TestSelectKnownArray
--- PASS: TestSelectKnownArray (0.00s)
=== RUN   TestSelectKnownArray2
--- PASS: TestSelectKnownArray2 (0.00s)
=== RUN   TestSelectVariousIndices
--- PASS: TestSelectVariousIndices (0.00s)
FAIL
exit status 1
FAIL	_/home/gabi/Gabi/apps/FloydRivest	3.859s
