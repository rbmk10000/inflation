package main
import "fmt"
var inflation_rate [15] float32

func main() {
	inflation_rate[0] = 1.6 // 2010	
	inflation_rate[1] = 3.2 // 2011 	
	inflation_rate[2] = 2.1 // 2012	
	inflation_rate[3] = 1.5 // 2013	
	inflation_rate[4] = 1.6 // 2014	
	inflation_rate[5] = 0.1 // 2015	
	inflation_rate[6] = 1.3 // 2016	
	inflation_rate[7] = 2.1 // 2017	
	inflation_rate[8] = 2.4 // 2018	
	inflation_rate[9] = 1.8 // 2019	
	inflation_rate[10] = 1.2// 2020
	inflation_rate[11] = 4.7// 2021
	inflation_rate[12] = 8.0// 2022
	inflation_rate[13] = 4.1// 2023
	inflation_rate[14] = 2.9// 2024
	var start_salary float32 = 2000
	var corrected_salary float32  = start_salary
	var real_salary float32  = start_salary
	var infl_rt float32 = 0
	//var start_year int = 2010
	for i :=0; i < 14; i++ {
		infl_rt = inflation_rate[i]/100
		corrected_salary = corrected_salary + corrected_salary * infl_rt 
		real_salary =  real_salary - start_salary * infl_rt
		fmt.Println(corrected_salary, ":", real_salary)
	}
}
