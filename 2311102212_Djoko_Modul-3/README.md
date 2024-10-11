Guided1
program main()
kamus:
    a, b : integer
algoritma: 
    input (a,b)
    if a >= b then
        output(permutasi(a, b))
    else then
        output(permutasi(b, a))
end program

function faktorial(n : integer) integer 
kamus: 
    hasil : integer
    i : integer
    i <- 1
    hasil <- 1
algoritma:
    while i <= n do 
        hasil <- hasil * 1
    return hasil
end function

Guided2
function permutasi(n, r : integer) integer 
algoritma
    return faktorial(n) / faktorial(n-r)

program volumeLuasBalok
kamus:
    p,l,t real
algoritma:
    input(p,l,t)
    output("Volume dari balok terebut adalah %.2f",volumeBalok(p,l,t))
    output("Luas dari balok terebut adalah %.2f",luasBalok(p,l,t))
end program

function volumeBalok(p,l,t : real) real
algoritma:
    return(p * l * t)
end function

function luasBalok(p,l,t : real) real 
algoritma:
    return 2 * ((p*l) + (p*t) + (l*t))
end function

Unguided1
program menghitungKombinasiPermutasi
kamus:
    a1,a2,b1,b2 : integer
    p1,p2,c1,c2 : integer
algoritma:
    input(a1,b1,a2,b2)
    p1 <- permutasian(a1, a2)
    p2 <- permutasian(b1, b2)
    c1 <- kombinasi(a1, a2)
    c2 <- kombinasi(b1, b2)
    output(p1, c1, p2, c2)
end program

function faktorial(n : integer) integer 
kamus: 
    hasil : integer
    i : integer
    i <- 1
    hasil <- 1
algoritma:
    while i <= n do 
        hasil <- hasil * 1
    return hasil
end function

function permutasian(n, r : integer) integer 
kamus:
    n <- faktorial(n)
    r <- faktorial(n-r)
    hasil : integer
algoritma:
    hasil <- n/r
    return hasil
end function

function kombinasi(n, r:integer) integer
kamus:
    n <- faktorial(n)
    hr <- faktorial(n-r)
    hasil : integer
algoritma:
    hasil <- n/(r*hr)
    return hasil
end function

Unguided_2
program fungsiMatematika
kamus:
    a,b,c : integer
algoritma:
    input(a, b, c)
    output(fogoh(a))
    output(gohof(b))
    output(hofog(c))
end program

function fogoh(a : integer) integer 
kamus:
    f, g, h, hasil : integer
algoritma:
    h <- a+1
    g <- h-2
    f <- h * h
    hasil <- f
    return hasil
end function

function gohof(a : integer) integer 
kamus:
    f, g, h, hasil : integer
algoritma:
    f <- a * a
    h <- f+1
    g <- h-2
    hasil <- g
    return hasil
end function

function hofog(a : integer) integer 
kamus:
    f, g, h, hasil : integer
algoritma:
    g <- a-2
    f <- g * g
    h <- f+1
    hasil <- g
    return hasil
end function

program menentukanPosisiTitikSembarang
kamus:
    cx1, cy1, r1, cx2, cy2, r2, x, y : real
    hasil : string
    d1, d2 : string
algoritma:
    input(cx1, cy1, r1)
    input(cx2, cy2, r2)
    input(x, y)
    d1 <- mencariTitikSembarang(cx1, cy1 ,r1, x, y)
    d2 <- mencariTitikSembarang(cx2, cy2 ,r2, x, y)
    if d1 == dl && d2 == dl then
        output ("Titik di dalam lingkaran 1 dan 2)
    else if d1 == ll && d2 == ll then
        output ("Titik di luar lingkaran 1 dan 2)
    else if d1 == dl && d2 == ll then
        output ("Titik di dalam lingkaran 1)
    else if d1 == ll && d2 == dl then
        output ("Titik di dalam lingkaran 2)
    else then
        output ("Titik berada di tepi lingkaran")    
    
end program

function kuadrat(a : integer) integer 
algoritma
    return a * a
end function

function mencariTitikSembarang(cx, cy, r, x, y : integer) string
kamus:
    hasil : string
algoritma:
    hasil : math.sqrt((kuadrat(x-cx)) - kuadrat(y-cy))
    if hasil < r then
        hasil <- dl
        return hasil
    else if hasil = r then
        hasil <- tl
        return hasil
    else if hasil > r then
        hasil <- ll
        return hasil
end function