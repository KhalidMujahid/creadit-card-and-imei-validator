export function creditCardValidator(input: string): boolean {
  const arr: number[] = [];

  // parse the input to integer then push it to arr (array)
  for (let i = 0; i < input.length; i++) {
    arr.push(parseInt(input[i]));
  }

  // the logic Luhn Algorithm
  for (let i = arr.length - 2; i >= 0; i = i - 2) {
    let temp: number = arr[i];

    temp = temp * 2;

    if (temp > 9) {
      temp = (temp % 10) + 1;
    }

    arr[i] = temp;
  }

  // sum
  let count = 0;

  for (let i = 0; i < arr.length; i++) {
    count += arr[i];
  }

  return count % 10 === 0;
}
