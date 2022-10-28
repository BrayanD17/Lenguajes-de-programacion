

from cmath import inf
from mimetypes import init

class Num:
    def __init__(self, numero):
        self.n = numero

class Op:
    def __init__(self, operacion):
        self.operacion = operacion

class ErrorCaracter_Expresion(Exception):
    def __init__(self, mensaje, informacion):
        super().__init__()
        self.mensaje = mensaje
        self.informacion = informacion




listaObjetos = []
listaOperaciones = ['+','-','/','*']
def descomponer():
    listaDatos = []
    datos = input("Ingrese las operaciones y numeros en infijo: ")
    for letra in datos:
        listaDatos.append(letra)
    print(listaDatos)

    try:
        i = 0
        validarExpresion = True
        validarCaracter = True
        while i < len(listaDatos):
            if i == 0:
                if listaDatos[i].isdigit():
                    validarExpresion = True
                    num = Num(int(listaDatos[i]))
                    listaObjetos.append(num)
                else:
                    validarExpresion = False
                    if listaDatos[i].isdigit() or listaDatos[i] in listaOperaciones:
                        validarCaracter = True
                    else:
                        validarCaracter = False
                    break
            
            elif i == 1:
                if listaDatos[i] in listaOperaciones:
                    validarExpresion = True
                    op = Op(listaDatos[i])
                    listaObjetos.append(op)
                else:
                    validarExpresion = False
                    if listaDatos[i].isdigit() or listaDatos[i] in listaOperaciones:
                        validarCaracter = True
                    else:
                        validarCaracter = False
                    break
            
            elif (i % 2 == 0):
                if listaDatos[i].isdigit():
                    validarExpresion = True
                    num = Num(int(listaDatos[i]))
                    listaObjetos.append(num)
                else:
                    validarExpresion = False
                    if listaDatos[i].isdigit() or listaDatos[i] in listaOperaciones:
                        validarCaracter = True
                    else:
                        validarCaracter = False
                    break
            
            else:
                if listaDatos[i] in listaOperaciones:
                    validarExpresion = True
                    op = Op(listaDatos[i])
                    listaObjetos.append(op)
                else:
                    validarExpresion = False
                    if listaDatos[i].isdigit() or listaDatos[i] in listaOperaciones:
                        validarCaracter = True
                    else:
                        validarCaracter = False
                    break

            i = i + 1
        if validarExpresion and validarCaracter:
            print("Lista con los objetos esta lista para trabajar!")
            verListaObjetos()
        if not(validarCaracter):
            raise ErrorCaracter_Expresion("La expresion presenta un error en caracteres: ", datos)
        if not(validarExpresion):
            raise ErrorCaracter_Expresion("La expresion presenta no se cuenta en infijo: ", datos)

    except ErrorCaracter_Expresion as e:
        print(e.mensaje, e.informacion)
def ifverListaObjetos():
    for ob in listaObjetos:
        if type(ob) is Num:
            print(ob.n)
        if type(ob) is Op:
            print(ob.operacion)
def calcular():
    numero1 = 0
    numero2 = 0
    resultado = 0

    while len(listaObjetos) >= 3:
        numero1 = listaObjetos[0].n
        numero2 = listaObjetos[2].n
        resultado = 0
        operar = listaObjetos[1].operacion
        if operar == '+':
            resultado = numero1 + numero2
        if operar == '-':
            resultado = numero1 - numero2

        if operar == '*':
            resultado = numero1 * numero2

        if operar == '/':
            resultado = numero1 / numero2

        newnum = Num(resultado)
        listaObjetos.pop(0)
        listaObjetos.pop(0)
        listaObjetos.pop(0)
        listaObjetos.insert(0,newnum)

    verListaObjetos()

#5*2/4+3
descomponer()
calcular()