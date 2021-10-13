# GoLang

TPE GoLang -

Integrantes: 
    - Acosta Torrissi Marcos
    - Lazarte Leandro
    - Poppe Rocio

Para este trabajo, dentro del archivo entities.go, se crea la estructrura Result como se dió en el enunciado, con los tres atributos Type, Value and Length. La función NewResult genera una instancia de la estructura mencionada. 
Luego, dentro del archivo main, se crea la función parser que recibe como paŕametro una cadena de caracteres, toma la misma y la divide de manera tal que, antes de instanciar un nuevo Result, va chequeando bloque a bloque si los datos enviados por parametro son compatibles con la estructura a instanciar. Esto se resuelve con tres if anidados; en el primero chequea que el formato se "NN" o "TX", en el segundo chequea que el tercer y cuarto caracter de la cadena sean numericos (se llama a la función stringToInt, se le pasan los estos caracteres como parámetros, para convertirlos a un entero, de no ser posible, devuelve 0 y no cumple con la condicion del if) y en el ultimo if se asegura que la cantidad de caracteres destinados al atributo Value, coincida con el numero entero antes devuelto por la funcion stringToInt. En todos los casos que no entre al if, no instancia el Result y setea un mensaje de error acorde al chequeo realizado. Si todo es correcto, se hace el new de la estructura y el error queda como nil. En ambos casos retorna los resultados obtenidos de los chequeos.

Dentro del archivo main_test.go se realizan los tests correspondientes para probar el correcto funcionamiento de las funciones creadas. Se generaron el archivo out y el out.html para chequear qué parte del código está cubierta por el test con el comando go test ./... -count=1 -coverprofile=out.test y go tool cover -html out -o out.html respectivamente.
Con el comando go test ./... -count=1 -cover se ejecuta el test y se puede apreciar que se esta cubriendo el 83.3 % del código generado. 


