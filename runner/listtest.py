def calculate_sum(numbers):
    result = 0
    for num in numbers:
        if num % 2 == 0:  # เลขคู่
            result += num
        else:  # เลขคี่
            result -= num
    return result

   
print(" *** Sum even / Subtract odd ***")
input_str = input("Enter numbers : ")
numbers = [int(num) for num in input_str.split()]

result = calculate_sum(numbers)
print(f"Sum is {result}")

