Logical operators:

! (logical NOT)
|| (logical OR) a > b || a > c// only able to check the second condition if the first condition is false. If the first condition is true, it does not move on to the second condition. Only 1 condition needs to be true for ||
&& (logical AND) a > b && a > c // again, both conditions need to be satisfied, so if the first condition is false, the compiler will not check the second condition. Both have to be true for &&.

Bitwise Operators:

& (bitwise AND)
| (bitwise OR)
^ (bitwise XOR)
&^(bit clear, AND NOT)
<< (left shift)
>> (right shift)

Comparison Operators:

== (equal)
!= (not equal)
< (less than)
<= (less than or equal to)
> (greater than)
>= (greater than or equal to)