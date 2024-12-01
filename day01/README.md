## Solution

### Part 1

1. Create two lists `left` and `right`.
2. For each line read, split it around `   `. Add the first element to `left` and the second element to `right`.
3. Sort both lists.
4. Initialize a variable `sum = 0`. Iterate through both lists, find the distance between corresponding elements and add them to `sum += | left[i] - right[i] |`.

### Part 2

1. Find the frequency of elements in the right list using a map.
2. Retrieve the frequency of each element in the left list using the map and calculate `total similarity score = sum(val \* freq(val))`
