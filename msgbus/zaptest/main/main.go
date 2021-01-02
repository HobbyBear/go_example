package main

type Foods struct {
	peopleList []People `json:"people_list"`
}

func (f *Foods)AddObserver(p People)  {
	f.peopleList = append(f.peopleList, p)
}

func (f *Foods)Produced()  {
	for _,p := range f.peopleList{
		p.Eat(f)
	}
}

type People interface {
	Eat(f *Foods)
}




func main() {

}
