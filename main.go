package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

func main () {
	var cnt = 0
	var icnt = 0
	var espaco string
	//Opening CSV
	file, _ := os.Open("problems.csv")
	defer file.Close()
	reader := csv.NewReader(file)
	data,_ := reader.ReadAll()
	//Initializing the timer
	fmt.Printf("Digite t para iniciar o timer\n")
	fmt.Scanf("%s", &espaco)
	if espaco == "t" {
		
		for _, value := range data {
			icnt++
			fmt.Printf("%s = ", value[0])
			//Setting the the timer to init alongside the loop
			time.AfterFunc(30 * time.Second, func (){
				fmt.Println()
				fmt.Printf("Voce acertou %d de %d", cnt, icnt )
				log.Fatal("Tempo limite excedido")
			})
			var res string
			fmt.Scanf("%s", &res)
			if res == value[1]{
				fmt.Printf("correto\n")
				cnt++
			} else {
				fmt.Printf("Incorreto\n")
			}
		}
		fmt.Printf("Voce acertou %d de %d", cnt, icnt )
	}
}