package validator

func Is_welcome(age uint) string {

	var msg string
	if age > 0 && age < 21 {
		//fmt.Println("Eres menor de edad y no puedes entrar al bar de moe")
		msg = "Eres menor de edad y no puedes entrar al bar de moe"
		return msg
	} else if age >= 21 {
		//fmt.Println("Sos bienvenido al bar de moe, la primera cerveza es gratis")
		msg = "Sos bienvenido al bar de moe, la primera cerveza es gratis"
		return msg
	} else {
		//fmt.Println("El numero no esta dentro del rango")
		msg = "El numero no esta dentro del rango"
		return msg
	}
}
