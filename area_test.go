package main
import "testing"

func TestRectangle(t *testing.T){
	got,err:=NewRectangle(5,4)
	want:=Rectangle{5,4}
	if(err!=nil){
		t.Errorf("Error")
	}else if got!=want{
		t.Errorf("got %d want %d",got,want)
	}
}
func TestAreaRectangle(t *testing.T){
	obj,err:=NewRectangle(5,4)
	got_area,_:=areaRectangle(&obj)
	want_area:=obj.length*obj.breadth
	if(err!=nil){
		t.Error(err)
	}else if got_area!=want_area{
		t.Errorf("got %d want %d",got_area,want_area)
	}
}

func TestAreaRectangleZero(t *testing.T){
	_,err:=NewRectangle(0,4)
	if(err==nil){
		t.Errorf("Area can't be zero")
	}
}
func TestAreaRectangleNegValues(t *testing.T){
	_,err:=NewRectangle(-5,4)
	if(err==nil){
		t.Errorf("Area can't be negative")
	}
}

func TestPerimeterRectangle(t *testing.T){
	obj,err:=NewRectangle(5,4)
	got_perimeter,_:=perimeterRectangle(&obj)
	want_peri:=2*(obj.length+obj.breadth)
	if(err!=nil){
		t.Error(err)
	}else if got_perimeter!=want_peri{
		t.Errorf("got %d want %d",got_perimeter,want_peri)
	}

}
func TestPerimeterRectangleZero(t *testing.T){
	_,err:=NewRectangle(5,0)
	if(err==nil){
		t.Error("Cant be zero")
	}

}
func TestPerimeterRectangleNegValues(t *testing.T){
	_,err:=NewRectangle(-5,4)
	if(err==nil){
		t.Errorf("negative values not possible")
	}
}

func TestAreaSquare(t *testing.T){
	s,err:=NewSquare(10)
	if(err!=nil){
		t.Error(err)
	}else{
		got:=s.Area()
		want:=100
		if(got!=want){
			t.Errorf("got %d want %d",got,want)
		}
	}
}

func TestAreaSquareNeg(t *testing.T){
	_,err:=NewSquare(-10)
	if(err==nil){
		t.Error("Negative values not allowed")
	}
}

func TestAreaSquareZero(t *testing.T){
	_,err:=NewSquare(0)
	if(err==nil){
		t.Error("zero not possible")
	}
}

func TestPerimeterSquare(t *testing.T){
	s,err:=NewSquare(10)
	if(err!=nil){
		t.Error(err)
	}else{
		got:=s.Perimeter()
		want:=40
		if(got!=want){
			t.Errorf("got %d want %d",got,want)
		}
	}
}

func TestPerimeterSquareNeg(t *testing.T){
	_,err:=NewSquare(-10)
	if(err==nil){
		t.Errorf("negative values")
	}
}

func TestPerimeterSquareZero(t *testing.T){
	_,err:=NewSquare(0)
	if(err==nil){
		t.Error("Zero error")
	}
}