# Taller de programación imperativa en Go

**Curso:** Paradigmas de Programación, ITCR, Ingeniería en Computadores
**Lenguaje:** Go
**Integrantes:**
- Mendez Rodriguez Eythan Yohel
- Mora Barrantes Franklin Santiago
- Venegas Arias Josue Gregorio

## Cómo ejecutarlo

1. Instalar Go desde https://go.dev/dl y comprobar con `go version`.
2. Abrir una terminal en la carpeta del proyecto (donde está `main.go`).
3. Correr:

```
go run main.go
```

No hace falta `go.mod` para correr un solo archivo con `go run main.go`. Todas las pruebas están definidas dentro de `main`, no hay entrada interactiva.

## Resultados de las pruebas

Cada prueba imprime `OK` o `FALLO` comparando contra el resultado esperado del enunciado.

´
Entrada: [4 -2 0 7 -5]

== Ejercicio 1 ==
antes:  a = 8 b = -3
despues: a = -3 b = 8

== Ejercicio 2 ==
Suma: 4
Minimo: -5
Maximo: 7
Negativos: 2
Recorrido inverso: -5 7 0 -2 4 
Entrada despues: [4 -2 0 7 -5]

== Ejercicio 3 ==
original: [4 -2 0 7 -5]
copia:    [99 -2 0 7 -5]

== Ejercicio 4 ==
ordenado: [-5 -2 0 4 7]
original: [4 -2 0 7 -5]

== Ejercicio 5 ==
findPair [-4 -1 0 2 5] objetivo 1
Par: (-4, 5)
 -> true OK
findPair [-4 -1 0 2 5] objetivo 8
No existe
 -> false OK
findPair [0 0] objetivo 0
Par: (0, 0)
 -> true OK
findPair [3] objetivo 6
No existe
 -> false OK
findPair [] objetivo 0
No existe
 -> false OK

== Ejercicio 6 ==
threeSum [-1 0 1 2 -1 -4]
(-1, -1, 2)
(-1, 0, 1)
 -> cantidad: 2 OK
 -> entrada intacta: [-1 0 1 2 -1 -4]
threeSum [0 1 1]
Ninguna
 -> cantidad: 0 OK
 -> entrada intacta: [0 1 1]
threeSum [0 0 0]
(0, 0, 0)
 -> cantidad: 1 OK
 -> entrada intacta: [0 0 0]
threeSum [0 0 0 0]
(0, 0, 0)
 -> cantidad: 1 OK
 -> entrada intacta: [0 0 0 0]
threeSum [-2 0 0 2 2]
(-2, 0, 2)
 -> cantidad: 1 OK
 -> entrada intacta: [-2 0 0 2 2]
threeSum [1 -1]
Ninguna
 -> cantidad: 0 OK
 -> entrada intacta: [1 -1]

== Ejercicio 7 ==
height:  [3 0 2 0 4]
leftMax:  [3 3 3 3 4]
rightMax: [4 4 4 4 4]

== Ejercicio 8 ==
agua [5] -> 0 OK
agua [0 1 2 3] -> 0 OK
agua [3 2 1 0] -> 0 OK
agua [2 2 2] -> 0 OK
agua [3 0 3] -> 3 OK
agua [3 0 2 0 4] -> 7 OK
agua [4 2 0 3 2 5] -> 9 OK
agua [0 1 0 2 1 0 1 3 2 1 2 1] -> 6 OK
agua [5 1 3 1 4] -> 7 OK

Detalle por posicion de [4 2 0 3 2 5]:
  pos 0: altura=4 leftMax=4 rightMax=5 nivel=4 agua=0
  pos 1: altura=2 leftMax=4 rightMax=5 nivel=4 agua=2
  pos 2: altura=0 leftMax=4 rightMax=5 nivel=4 agua=4
  pos 3: altura=3 leftMax=4 rightMax=5 nivel=4 agua=1
  pos 4: altura=2 leftMax=4 rightMax=5 nivel=4 agua=2
  pos 5: altura=5 leftMax=5 rightMax=5 nivel=5 agua=0
  total: 9
´

Resumen de lo que se espera:

| Ejercicio | Resultado esperado |
|---|---|
| 1 | a = -3, b = 8 |
| 2 | Suma 4, mínimo -5, máximo 7, negativos 2, inverso `-5 7 0 -2 4` |
| 3 | original `[4 -2 0 7 -5]`, copia `[99 -2 0 7 -5]` |
| 4 | `[-5 -2 0 4 7]` |
| 5 | `(-4, 5)`, no existe, `(0, 0)`, no existe |
| 6 | 2 ternas, 0, 1, 1, 1 |
| 7 | leftMax `[3 3 3 3 4]`, rightMax `[4 4 4 4 4]` |
| 8 | 0, 0, 0, 0, 3, 7, 9, 6 |

Casos con tamaño insuficiente (resultado vacío): `findPair([3], 6)`, `findPair([], 0)` y `threeSum([1, -1])`.

### Caso propio de agua atrapada

`[5, 1, 3, 1, 4]` debe dar **7**. El nivel del agua lo marca el menor de los dos máximos a cada lado. En las posiciones 1, 2 y 3 el nivel es 4 (el 5 de la izquierda es más alto que el 4 de la derecha), así que hay 3 + 1 + 3 = 7 unidades.

### Detalle por posición de `[4, 2, 0, 3, 2, 5]`

| Pos | Altura | leftMax | rightMax | Nivel | Agua |
|---|---|---|---|---|---|
| 0 | 4 | 4 | 5 | 4 | 0 |
| 1 | 2 | 4 | 5 | 4 | 2 |
| 2 | 0 | 4 | 5 | 4 | 4 |
| 3 | 3 | 4 | 5 | 4 | 1 |
| 4 | 2 | 4 | 5 | 4 | 2 |
| 5 | 5 | 5 | 5 | 5 | 0 |

Total: 9.

## Explicaciones

### Ejercicio 1: por qué `swap` recibe punteros

Go pasa los argumentos por valor, o sea que la función recibe copias. Si `swap` recibiera dos `int`, cambiaría las copias y las variables de `main` quedarían igual. Con punteros recibe las direcciones y puede escribir en las variables originales.

Dentro de la función, modificar `a` cambia solo el puntero local (a dónde apunta), mientras que modificar `*a` cambia el valor al que apunta, que es el de `main`.

La variable `aux` es local: existe solo dentro de `swap` y desaparece al terminar la función.

### Ejercicio 3: por qué `other := original` no copia

Un slice es un descriptor con un puntero al arreglo subyacente, una longitud y una capacidad. `other := original` copia ese descriptor, no los elementos, así que ambos slices apuntan al mismo arreglo y un cambio por uno se ve por el otro. Para tener datos independientes hay que crear un slice nuevo con `make` y copiar elemento por elemento, que es lo que hace `copyArray`.

### Ejercicio 5: por qué se puede descartar un extremo

Con el slice ordenado, si `values[left] + values[right]` es menor que el objetivo, ese `values[left]` no puede formar el par con ningún valor de la zona que queda, porque el mayor posible ya es `values[right]`. Entonces se descarta con `left++`. Si la suma es mayor, pasa lo mismo con `values[right]` y se hace `right--`. Cada paso elimina un candidato, así que se recorre una sola vez.

### Ejercicio 6: cómo se evitan ternas duplicadas en 3Sum

Se ordena una copia de la entrada y se fija un primer elemento `i`. Para los otros dos se usan dos índices (`left` y `right`). Los duplicados se evitan en dos puntos:

- Si `sorted[i]` es igual a `sorted[i-1]`, se salta ese `i`, porque ya se probaron todas las ternas que empiezan con ese valor.
- Cuando se encuentra una terna, se avanza `left` y `right` y se siguen moviendo mientras el valor sea igual al anterior, para no repetir la misma pareja.

Como al ordenar los repetidos quedan juntos, esas dos reglas bastan. Los valores repetidos sí se permiten si hay suficientes apariciones, porque cada elemento viene de una posición distinta. La búsqueda es O(n²) después de ordenar.

### Ejercicio 8: por qué el agua depende del menor de los dos máximos

El agua sobre una columna puede subir hasta que se derrama por el lado más bajo. Si a la izquierda hay una pared de 5 y a la derecha una de 4, el agua llega a 4 y ahí se sale por la derecha. Por eso el nivel en cada posición es `min(leftMax, rightMax)` y el agua es ese nivel menos la altura de la columna. Como `leftMax` y `rightMax` incluyen la altura actual, el resultado nunca es negativo.

El total se acumula en `int64`.

## Ejemplo de alcance de variables

En `swap`:

```go
func swap(a *int, b *int) {
	aux := *a
	*a = *b
	*b = aux
}
```

`aux` solo existe dentro de `swap`. Fuera de la función no se puede usar, y cada llamada crea la suya.

Otro caso es el `for` de `findPair`: `s` se declara dentro del ciclo y su alcance es solo ese bloque. Si se usara `:=` con un nombre que ya existe afuera, crearía una variable nueva que oculta la de afuera (shadowing).

## Manejo de memoria: Go vs C

| | C | Go |
|---|---|---|
| Reservar | `malloc` | `make` (slices) |
| Fallo al reservar | hay que comprobar `NULL` | no se maneja en este taller |
| Liberar | `free` manual | automático, lo hace el recolector de basura |
| Errores típicos | fugas, doble `free`, uso tras liberar | no hay `free` que olvidar |
| Arreglos | puntero + tamaño aparte | slice: puntero, longitud y capacidad juntos |

En C el programador es dueño de la memoria y decide cuándo liberarla. En Go la memoria se recupera cuando deja de ser alcanzable, lo que quita varios errores pero cuesta algo de control y rendimiento.
