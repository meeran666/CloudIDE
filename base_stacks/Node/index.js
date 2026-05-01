// Utility: generate random integers
function randomInt(min, max) {
  return Math.floor(Math.random() * (max - min + 1)) + min;
}

// Generate an array of random numbers
function generateArray(size, min = 1, max = 100) {
  return Array.from({ length: size }, () => randomInt(min, max));
}

// Sum of array
function sum(arr) {
  return arr.reduce((acc, val) => acc + val, 0);
}

// Average
function average(arr) {
  return sum(arr) / arr.length;
}

// Factorial (recursive)
function factorial(n) {
  if (n <= 1) return 1;
  return n * factorial(n - 1);
}

// Fibonacci (iterative - efficient)
function fibonacci(n) {
  if (n <= 1) return n;
  let a = 0, b = 1;
  for (let i = 2; i <= n; i++) {
    [a, b] = [b, a + b];
  }
  return b;
}

// Prime check
function isPrime(n) {
  if (n < 2) return false;
  for (let i = 2; i * i <= n; i++) {
    if (n % i === 0) return false;
  }
  return true;
}

// Filter primes from array
function getPrimes(arr) {
  return arr.filter(isPrime);
}

// Sort (quick sort implementation)
function quickSort(arr) {
  if (arr.length <= 1) return arr;

  const pivot = arr[arr.length - 1];
  const left = [];
  const right = [];

  for (let i = 0; i < arr.length - 1; i++) {
    arr[i] < pivot ? left.push(arr[i]) : right.push(arr[i]);
  }

  return [...quickSort(left), pivot, ...quickSort(right)];
}

// Main execution
function main() {
  const numbers = generateArray(10, 1, 50);

  console.log("Generated numbers:", numbers);
  console.log("Sorted:", quickSort(numbers));
  console.log("Sum:", sum(numbers));
  console.log("Average:", average(numbers));

  const primes = getPrimes(numbers);
  console.log("Prime numbers:", primes);

  console.log("Factorial of 5:", factorial(5));
  console.log("Fibonacci of 10:", fibonacci(10));
}

// Run program
main();