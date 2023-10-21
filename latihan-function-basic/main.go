package main
import "fmt"

func notReturn(car map[string]string)  {
		fmt.Println("Mobil",car["name"],"Berwarna",car["color"] )
	
}

func Return(car map[string]string) string {
	var kata string
	kata = "Mobil" + " "+car["name"] +" " + "Berwarna" +" "+car["color"] 

	return kata

}

func main(){
    var car = map[string]string{}
    car["name"] = "BWM"
    car["color"] = "Black"

	notReturn(car)
	fmt.Println(Return(car))
}

