module Main (main) where
--llama a la funcion miembro para comprobar si existe el string solicitado en una cadena de texto
ejercicio1 palabra lista =
    filter contienenP lista
    where
        contienenP p = 
            if miembro palabra (split p) == True 
                then True 
            else False
          
split "" = []
split xs = ys : (split . drop 1) zs
   where (ys, zs) = span (/=' ') xs
-- Manda como primera lista, la lista que tiene los elementos que se desean comprobar y la segunda lista
-- es la que alberga todos los elementos
ejercicio2 [] _ = True
ejercicio2 l1 l2 =
    miembro (head l1) l2 && ejercicio2 (tail l1) l2
miembro _ [] = False
miembro e (x:xs) 
    | e == x = True
    | otherwise = miembro e xs
-- queda muy simple con el concat ya que no encontre una funcion para comparar si un elemento dentro
-- de una lista es tambien una lista
-- el ejercicio3_1 funciona para lista de lista la cual tiene una lista de numeros dentro [[[a]]]
-- el ejercicio3_2 funciona para una lista que tiene una lista de numeros, es mas simple [[a]]
ejercicio3_1 l = (concat . concat) l
ejercicio3_2 l  = concat l

-- La 3 y la 4 vienen siendo practicamente el mismo ejercicio, sin embargo para no repetir el mismo codigo en el
-- tercer ejercicio se implementa de diferente manera
ejercicio4 lista =
    concat(map juntar_sub lista)
    where 
        juntar_sub sublista = juntar sublista
                   
juntar lista =
    concat (map pegar lista)
    where
        pegar e = [e]

main :: IO ()
main = do
    print("Resultados del ejercicio 1: ")
    print (ejercicio1 "la" ["hola la casa", "Pepe la lanza", "Cama"])
    print(ejercicio1 "la" ["hola la casa", "Pepe la lanza", "Cama", "mininmo la taza"])
    print("Resultados del ejercicio 2: ")
    print(ejercicio2 [1,2,3,40] [4,1,6,7,0,2,5,3,20,4])
    print(ejercicio2 [] [4,1,6,7,0,2,5,3,20,4])
    print(ejercicio2 [1,2,3] [4,1,6,7,0,2,5,3,20,4])
    print("Resultados del ejercicio 3: ")
    print(ejercicio3_1 [[[1,2],[3]],[[4],[1,4,5]]])
    print (ejercicio3_2 [[1,2,3],[4,5,6]])
    print("Resultados del ejercicio 4: ")
    print (ejercicio4 [[1,2,3],[4,5,6],[7,8,9,10,11,12,13]])