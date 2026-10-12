package main

import (
	"flag"
	"fmt"

	"github.com/fatih/color"
	"github.com/miloter/go-greetctl/internal/logx"

	// Usamos un alias <gr> para el paquete <greet>
	gr "github.com/miloter/go-greetctl/pkg/greet"
)

func main() {
	name := flag.String("name", "mundo", "-name mundo")
	flag.Parse()
	msg := gr.Hello(*name)
	fmt.Printf("cmd > %s!\n", msg)
	color.Green("Hola %s, saludos\n", *name)
	color.Blue("Hola %s, saludos\n", *name)
	logx.Info("FIN todo OK")
}
