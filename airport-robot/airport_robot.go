package airportrobot

import "fmt"

type Greeter interface {
    LanguageName() string
    Greet(name string) string
}

func SayHello(name string, g Greeter) string {

}



