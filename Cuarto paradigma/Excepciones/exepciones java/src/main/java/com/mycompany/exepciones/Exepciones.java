/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Project/Maven2/JavaApp/src/main/java/${packagePath}/${mainClassName}.java to edit this template
 */

package com.mycompany.exepciones;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Scanner;
import java.util.Set;

/**
 *
 * @author dinar
 */
public class Exepciones {
    static ArrayList<Object> listaObjetos = new ArrayList<Object>();
    static ArrayList<String> listaOperaciones = new ArrayList<String>();
    
    public static boolean isNumeric(String cadena){
        try{
            Integer.parseInt(cadena);
            return true;
        }catch(NumberFormatException nfe){
            return false;
        }
    }
    
    public static void mostrarListaObjetos(){
        for (int i = 0; i < listaObjetos.size(); i ++){
            if(listaObjetos.get(i) instanceof Num == true){
                System.out.println( ((Num)listaObjetos.get(i)).getNumero_var());
            }
            else{
                System.out.println( ((Op)listaObjetos.get(i)).getOperacion_var());
            }
        }
    }
    
    public static void descomponer ()throws ErrorCF{
        ArrayList<String> listaCaracteres = new ArrayList<String>();
        Scanner datos = new Scanner(System.in);
        System.out.print("Ingrese la operacion: ");
        
        String operacionNueva = datos.nextLine();
        System.out.println("");
        System.out.println(operacionNueva);
        String []listaDatos = operacionNueva.split("");
        for(String letra: listaDatos){
            listaCaracteres.add(letra);
        }
        
        //Try
        boolean validarExpresion = true;
        boolean validarCaracter = true;
        int i = 0;
        while (i < listaCaracteres.size()){
            if (i == 0){
                if (isNumeric(listaCaracteres.get(i))){
                    validarExpresion = true;
                    Num n = new Num(Integer.parseInt(listaCaracteres.get(i)));
                    listaObjetos.add(n);
                }
                else{
                    validarExpresion = false;
                    if (isNumeric(listaCaracteres.get(i)) || listaOperaciones.contains(listaCaracteres.get(i))){
                        validarCaracter = true;
                    }
                    else{
                        validarExpresion = false;
                    }
                    break;
                }
            }
            else if (i == 1){
                if (listaOperaciones.contains(listaCaracteres.get(i))){
                    validarExpresion = true;
                    Op o = new Op(listaCaracteres.get(i));
                    listaObjetos.add(o);
                }
                else{
                    validarExpresion = false;
                    if (isNumeric(listaCaracteres.get(i)) || listaOperaciones.contains(listaCaracteres.get(i))){
                        validarCaracter = true;
                    }
                    else{
                        validarExpresion = false;
                    }
                    break;
                }
            }
            else if ((i % 2) == 0){
                if (isNumeric(listaCaracteres.get(i))){
                    validarExpresion = true;
                    Num n = new Num(Integer.parseInt(listaCaracteres.get(i)));
                    listaObjetos.add(n);
                }
                else{
                    validarExpresion = false;
                    if (isNumeric(listaCaracteres.get(i)) || listaOperaciones.contains(listaCaracteres.get(i))){
                        validarCaracter = true;
                    }
                    else{
                        validarExpresion = false;
                    }
                    break;
                
                }
            }
            else{
                if (listaOperaciones.contains(listaCaracteres.get(i))){
                    validarExpresion = true;
                    Op o = new Op(listaCaracteres.get(i));
                    listaObjetos.add(o);
                    
                }
                else{
                    validarExpresion = false;
                    if (isNumeric(listaCaracteres.get(i)) || listaOperaciones.contains(listaCaracteres.get(i))){
                        validarCaracter = true;
                    }
                    else{
                        validarExpresion = false;
                    }
                    break;
                }
            }
            i = i + 1;
            
        }
        
        if (validarCaracter && validarExpresion){
        System.out.println("Lista con objetos llena!");
        mostrarListaObjetos();

        }
        if  (!validarCaracter){
            throw new ErrorCF("La expresion presenta un error de caracteres.");
        }
        if  (!validarExpresion){
            throw new ErrorCF("La expresion presenta un error de expresion.");
        }
        
        
    }
    public static void calcular(){
        int num1 = 0;
        int num2 = 0;
        String operacion = "";
        int resultado = 0;
        while (listaObjetos.size() >= 3){
            num1 = ((Num)listaObjetos.get(0)).getNumero_var();
            operacion = ((Op)listaObjetos.get(1)).getOperacion_var();
            num2 = ((Num)listaObjetos.get(2)).getNumero_var();
            
            if (operacion.equals("+")){
                resultado = num1 + num2;
            }
            if (operacion.equals("-")){
                
                resultado = num1 - num2;
            }
            if (operacion.equals("*")){
                
                resultado = num1 * num2;
            }
            if (operacion.equals("/")){
                
                resultado = num1 / num2;
            }
            Num nuevoN = new Num(resultado);
            listaObjetos.remove(0);
            listaObjetos.remove(0);
            listaObjetos.remove(0);
            listaObjetos.add(0, nuevoN);
            System.out.println("Lista nueva");
            mostrarListaObjetos();
        }
    }
    
    public static void main(String[] args) {
        listaOperaciones.add("+");
        listaOperaciones.add("*");
        listaOperaciones.add("-");
        listaOperaciones.add("/");
        try{
            descomponer();
            calcular();
        }catch(ErrorCF e){
            System.out.println("Error capturado: " + e.getMessage());
            
        }
        
        //String text= new String("Hello, My name is Sachin");  
        /* Splits the sentence by the delimeter passed as an argument */  
        //String[] sentences = text.split("");  
        //System.out.println(Arrays.toString(sentences));  
        
    }
}
