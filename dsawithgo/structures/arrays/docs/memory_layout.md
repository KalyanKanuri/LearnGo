# Memory Layout of arrays

## Address Index

An Integer array in the memory looks like this.
P.S. the memory addresses are simulated not real

```txt
Address         Value
-------        -------
1000    --->     10
1008    --->     20
1016    --->     30
1024    --->     40
1032    --->     50
```

if we observe clearly we kept a difference of 8 for each memory address because a integer occupies 8 bytes of space in the memory.

## Contiguous Memory

contiguous simply means storing next to each other without any gaps

e.g., see below, the data in the memory is stored as below

```txt
+----+----+----+----+----+
|10  |20  |30  |40  |50  |
+----+----+----+----+----+
```

not distributed as shown below

```txt
10

          20

    30

                 40

50
```

## Finding an element in an index

suppose we want to access 3rd element in the array, that is element at index 3, if we consider the first memory address layout of the array the address of the first element is 1000 i.e., Base

system follows below formulae to get the value of an array at an index

Address = Base + (index * size)

=> Address = 1000 + (3 * 8)
   Address = 1000 + 24 = 1024

i.e., the memory address of 3rd element is 1024.
