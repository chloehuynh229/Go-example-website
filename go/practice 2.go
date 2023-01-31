// 1 . Viết app giải phương trình bậc 2. Hint: ax^2 + bx + c =0. Nhập a, b, c cho ra kết quả. (IF)
package main

import (
	"fmt"
)

func main() {

	// Println function is used to
	// display output in the next line
	fmt.Println("Input a: ")
	var a float64  // var then variable name then variable type
	fmt.Scanln(&a) // Taking input from user

	fmt.Println("Input b: ")
	var b float64
	fmt.Scanln(&b)

	fmt.Println("Input c: ")
	var c float64
	fmt.Scanln(&c)

	if a == 0 && b == 0 && c == 0 {
		fmt.Println("PT vo so nghiem.")
	
	} else if a,b == 0 && c != 0 {
		fmt.Println("PT vo nghiem")

	} else {
		if a,b != 0 && c == 0 {
		fmt.Println("PT co 1 nghiem x = 0"
		} else if a == 0 && b,c != 0{
		fmt.Println("x =", -b/a)
	}
}

// 2. Nhập số a , check số a có phải prime number(số nguyên tố) không?

// 3. Cho a = 7, b = 777.
// Step 1: Swap 2 số này dùng function (a = 777 và b =7).
// Step 2: Swap 2 số này không dùng function.
// Step 3: Swap 2 số này không dùng function không dùng biến thứ 3.

// 4. Gọi 123 là perfect number vì 1+2+3=1*2*3. Tìm tất cả các perfect number nhỏ hơn 1000, 1000000, 10 triệu.

// 5. Nhập matrix có dạng 5x5( 5 hàng 5 cột). Check matrix này có đối xứng không? (Loop)
