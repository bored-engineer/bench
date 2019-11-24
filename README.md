# bench
Benchmarks of various Golang stdlib functions 

## String Trimming
Tested on 11/24/2019 on `darwin`/`amd64` with `go1.13.1`:
```
BenchmarkBytesTrimLeftFuncSpace/many-12         	39149709	      32.5 ns/op
BenchmarkBytesTrimLeftSpace/many-12              	9234144 	      129 ns/op
BenchmarkStringsTrimLeftFuncSpace/many-12        	45523309	      27.2 ns/op
BenchmarkStringsTrimLeftSpace/many-12            	9048939 	      130 ns/op

BenchmarkBytesTrimLeftFuncSpace/none-12          	182077730	      6.38 ns/op
BenchmarkBytesTrimLeftSpace/none-12              	14094928	      86.7 ns/op
BenchmarkStringsTrimLeftFuncSpace/none-12        	208282578	      5.47 ns/op
BenchmarkStringsTrimLeftSpace/none-12            	16269506	      79.6 ns/op

BenchmarkBytesTrimLeftFuncSpace/one-12           	100000000	      10.4 ns/op
BenchmarkBytesTrimLeftSpace/one-12               	13244439	      93.0 ns/op
BenchmarkStringsTrimLeftFuncSpace/one-12         	140431153	      8.49 ns/op
BenchmarkStringsTrimLeftSpace/one-12             	13601450	      86.0 ns/op

BenchmarkBytesTrimLeftFuncSpace/reallybig-12     	365420  	      3804 ns/op
BenchmarkBytesTrimLeftSpace/reallybig-12         	201292  	      6414 ns/op
BenchmarkStringsTrimLeftFuncSpace/reallybig-12   	374078  	      2976 ns/op
BenchmarkStringsTrimLeftSpace/reallybig-12       	155467  	      7592 ns/op
```
