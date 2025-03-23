package main
import("fmt" 
"errors"
"unicode/utf8") 
const Value int = 2344
func main(){
	fmt.Println("hello from goland")
	var First int32 = 8
	var  Second float32 = 12.3
	var Total float32 =  float32(First) + Second
	var Greating  = "goHello" +" "+ "world"
	fmt.Println(Total,Greating)
	fmt.Println(utf8.RuneCountInString("Y"))
	fmt.Println(len("γ"))
	var1 ,var2 := 1,2
	fmt.Println(var1+var2)
    myVar,myvar2 := "strg", "string"
	fmt.Println(myVar, myvar2)
	var print = "Hello \n world"
	fmt.Println(Value)
	printme(print)
	result,divide,err := Division(2,2)
	if err!= nil{
		fmt.Println(err)
	}else if(result ==0){
		fmt.Printf("remainder is zero (%v) ",result)
	}else{
		fmt.Printf("Value of Division is %v and It remainder is %v",divide,result )
	}

	
}

func printme(printvalue string){
	fmt.Println(printvalue)
}
func Division (Numo int , Deno int)(int,float64,error){
	if Deno == 0{
		err:= errors.New("denominator is equal to zero so result is undifined")
		return  0,0,err
	}else{
		result := Numo%Deno
		divide := float64(Numo/Deno)
		return result,divide,nil
	}
	
}