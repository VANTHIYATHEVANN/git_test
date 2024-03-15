package main
import "fmt"
import "errors"
type area interface{
	Area() int
}
type perimter interface{
	Perimeter() int
}
type Rectangle struct{
	length int
	breadth int
}
type Square struct{
	side int
}
func (rect Rectangle)Area() int{
	return rect.length*rect.breadth
}
func (square Square)Area() int{
	return square.side*square.side
}
func (square Square)Perimeter() int{
	return 4*square.side
}
func NewSquare(side int)(Square,error){
	if(side>0){
		square:=Square{side}
		return square,nil
	}else if(side==0){
		square := Square{side}
		msg:=fmt.Sprintf("Zero not allowed")
		err:=errors.New(msg)
		return square,err
	}else{
		square := Square{side}
		msg:=fmt.Sprintf("Negative number not allowed")
		err:=errors.New(msg)
		return square,err
	}
}
func NewRectangle(length ,breadth int) (Rectangle,error){
		if(length>0 && breadth>0){
			rect := Rectangle {length,breadth}
			return rect,nil
		}else if(length==0 || breadth==0){
			rect := Rectangle {length,breadth}
			msg:=fmt.Sprintf("Zero not allowed")
			err:=errors.New(msg)
			return rect,err
		}else{
			rect := Rectangle {length,breadth}
			msg:=fmt.Sprintf("Negative values not allowed")
			err:=errors.New(msg)
			return rect,err
		}
}
func areaRectangle(rect *Rectangle) (int,error){
	if(rect.length>0 && rect.breadth>0){
		return rect.length*rect.breadth,nil
	}else if(rect.length==0 || rect.breadth==0){
		msg:=fmt.Sprintf("Zero not allowed")
		err:=errors.New(msg)
		return 0,err
	}else{
		msg:=fmt.Sprintf("Negative values not allowed")
		err:=errors.New(msg)
		return 0,err
	}
}

func perimeterRectangle(rect *Rectangle)(int, error){
	if(rect.length>0 && rect.breadth>0){
		return 2*(rect.length+rect.breadth),nil
	}else if(rect.length==0 || rect.breadth==0){
		msg:=fmt.Sprintf("Zero not allowed")
		err:=errors.New(msg)
		return 0,err
	}else{
		msg:=fmt.Sprintf("Negative values not allowed")
		err:=errors.New(msg)
		return 0,err
	}
}
func main(){
	rect,_:=NewRectangle(5,4)
	fmt.Println(rect.Area())
}