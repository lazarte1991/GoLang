# GoLang

TPE GoLang -

Integrantes: 
    - Lazarte Leandro
    - Poppe Rocio
    - Acosta Torrissi Marcos

Para este trabajo, dentro del archivo entities.go, se crea la estructrura Result como se dió en el enunciado, con los tres atributos Type, Value and Length. La función NewResult genera una instancia de la estructura mencionada. 
Luego, dentro del archivo main, se crea la función parser que recibe como paŕametro una cadena de caracteres, toma la misma y la divide de manera tal que setea los valores de los atributos de la estructura Result y la devuelve. 
En la función stringToInt, se le pasan los caracteres que estan en la segunda y tercera posición de la cadena como parámetros, para convertirlos a un entero.
Dentro del archivo main_test.go se realizan los tests correspondientes para probar el correcto funcionamiento de las funciones creadas. Se generaron el archivo out y el out.html para chequear qué parte del código está cubierta por el test con el comando go test ./... -count=1 -coverprofile=out.test y go tool cover -html out -o out.html respectivamente.
Con el comando go test ./... -count=1 -cover se ejecuta el test y se puede apreciar que se esta cubriendo el 80 % del código generado. 


